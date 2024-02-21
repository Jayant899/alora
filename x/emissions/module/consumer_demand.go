package module

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	cosmosMath "cosmossdk.io/math"
	state "github.com/allora-network/allora-chain/x/emissions"
	"github.com/allora-network/allora-chain/x/emissions/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// A structure to hold the original value and a random tiebreaker
type SortableItem[T any] struct {
	Value      T
	Weight     uint64
	Tiebreaker float64
}

type RequestId = string

type TopicId = uint64

type PriceAndReturn = struct {
	Price  cosmosMath.Uint
	Return cosmosMath.Uint
}

// Sorts the given slice of topics in descending order according to their corresponding return, using randomness as tiebreaker
// e.g. ([]uint64{1, 2, 3}, map[uint64]uint64{1: 2, 2: 2, 3: 3}, 0) -> [3, 1, 2] or [3, 2, 1]
func SortTopicsByReturnDescWithRandomTiebreaker(valsToSort []state.Topic, weights map[TopicId]PriceAndReturn, randSeed uint64) []state.Topic {
	// Convert the slice of Ts to a slice of SortableItems, each with a random tiebreaker
	r := rand.New(rand.NewSource(int64(randSeed)))
	items := make([]SortableItem[state.Topic], len(valsToSort))
	for i, topic := range valsToSort {
		items[i] = SortableItem[state.Topic]{topic, weights[topic.Id].Price.Uint64(), r.Float64()}
	}

	// Sort the slice of SortableItems
	// If the values are equal, the tiebreaker will decide their order
	sort.Slice(items, func(i, j int) bool {
		if items[i].Value == items[j].Value {
			return items[i].Tiebreaker > items[j].Tiebreaker
		}
		return items[i].Weight > items[j].Weight
	})

	// Extract and print the sorted values to demonstrate the sorting
	sortedValues := make([]state.Topic, len(valsToSort))
	for i, item := range items {
		sortedValues[i] = item.Value
	}
	return sortedValues
}

// Check that a request:
//
//	should be checked in this timestep
//	AND didn't expire
//	AND would have enough funds to cover the potential next price
//	AND admits and acceptable max price per inference
func IsValidAtPrice(ctx sdk.Context, am AppModule, req state.InferenceRequest, price cosmosMath.Uint, currentTime uint64) (*bool, error) {
	reqId, err := req.GetRequestId()
	if err != nil {
		fmt.Println("Error getting request id: ", err)
		return nil, err
	}
	reqUnmetDemand, err := am.keeper.GetRequestDemand(ctx, reqId)
	if err != nil {
		fmt.Println("Error getting request demand: ", err)
		return nil, err
	}
	res := req.LastChecked+req.Cadence <= currentTime && req.TimestampValidUntil > currentTime && reqUnmetDemand.GTE(price) && req.MaxPricePerInference.LTE(price)
	return &res, nil
}

// Big Principle:
//
//	The price of inference is determined by the demand for inference vs the scarcity of capacity to run topic-level shit
//
// We process all inference requests, draw demand from them, possibly remove them from the mempool
// in order to determine which topics we should calculate, the current price per inference, and total
// demand to meet with supply (how much demand to draw).
//
// Find the demand per epoch of the topic by counting inference requests per topic that have not yet been filtered out
// Order topics by "demand" (num unfiltered requests)
// Take top keeper.MAX_TOPICS_PER_BLOCK number of topics with the highest demand (sort desc by num requests)
// If number of topics > keeper.TARGET_CAPACITY_PER_BLOCK, price := price * (1 + keeper.PRICE_CHANGE_PERCENT)
// If number of topics < keeper.TARGET_CAPACITY_PER_BLOCK, price := max(keeper.MIN_PRICE_PER_EPOCH, price * (1 - keeper.PRICE_CHANGE_PERCENT))
// Deduct from all unfiltered requests `price` amount among those top topics
// Delete requests with negligible demand leftover
func ChurnAndDrawFromRequestsToGetTopActiveTopicsAndMetDemand(ctx sdk.Context, am AppModule, currentTime uint64) (*[]state.Topic, *cosmosMath.Uint, error) {
	// Inactivate topics with below keeper.MIN_TOPIC_DEMAND demand)
	topicsActive, err := am.keeper.GetActiveTopics(ctx)
	if err != nil {
		fmt.Println("Error getting active topics: ", err)
		return nil, nil, err
	}
	for _, topic := range topicsActive {
		topicUnmetDemand, err := am.keeper.GetUnmetDemand(ctx, topic.Id)
		if err != nil {
			fmt.Println("Error getting unmet demand: ", err)
			return nil, nil, err
		}
		if topicUnmetDemand.LT(cosmosMath.NewUint(keeper.MIN_TOPIC_DEMAND)) {
			fmt.Printf("Inactivating topic due to no demand: %v metadata: %s", topic.Id, topic.Metadata)
			am.keeper.InactivateTopic(ctx, topic.Id)
		}
	}

	topicsActive, err = am.keeper.GetActiveTopics(ctx)
	if err != nil {
		fmt.Println("Error getting active topics: ", err)
		return nil, nil, err
	}
	fmt.Println("Active topics: ", len(topicsActive))

	topicsActiveWithDemand := make([]state.Topic, 0)
	_IsValidAtPrice := func(req state.InferenceRequest, price cosmosMath.Uint) (*bool, error) {
		return IsValidAtPrice(ctx, am, req, price, currentTime)
	}
	topicBestPrices := make(map[TopicId]PriceAndReturn)
	requestsToDrawDemandFrom := make(map[TopicId]map[string]state.InferenceRequest, 0)
	for _, topic := range topicsActive {
		inferenceRequests, err := am.keeper.GetMempoolInferenceRequestsForTopic(ctx, topic.Id)
		if err != nil {
			fmt.Println("Error getting mempool inference requests: ", err)
			return nil, nil, err
		}

		// Initialize a map of request price to array of valid requests
		demandCurve := make(map[cosmosMath.Uint]map[RequestId]state.InferenceRequest)

		// Loop through inference requests and then loop again (nested) checking validity of all other inferences at the first inference's max price
		for _, req := range inferenceRequests {
			// Check validity of current request at its own price
			isValidAtPrice, err := _IsValidAtPrice(req, req.MaxPricePerInference)
			if err != nil {
				fmt.Println("Error checking if request is valid at price: ", err)
				return nil, nil, err
			}
			if *isValidAtPrice {
				reqId, err := req.GetRequestId()
				if err != nil {
					fmt.Println("Error getting request id: ", err)
					return nil, nil, err
				}

				// Initialize map of requests within demand curve. These nested map of requests are valid for the wrapping key price
				potentialNextPrice := req.MaxPricePerInference
				if demandCurve[potentialNextPrice] == nil {
					demandCurve[potentialNextPrice] = make(map[RequestId]state.InferenceRequest)
				}
				// We already checked the validity of this request at its own price above, so we can add it to the demand curve
				demandCurve[potentialNextPrice][reqId] = req

				for _, req2 := range inferenceRequests {
					reqId2, err := req2.GetRequestId()
					if err != nil {
						fmt.Println("Error getting request id: ", err)
						return nil, nil, err
					}
					_, ok := demandCurve[potentialNextPrice][reqId2]
					if !ok {
						isValidAtPrice, err := _IsValidAtPrice(req2, potentialNextPrice)
						if err != nil {
							fmt.Println("Error checking if request is valid at price: ", err)
							return nil, nil, err
						}
						if *isValidAtPrice {
							demandCurve[potentialNextPrice][reqId2] = req2
						}
					}
				}
			}
		}

		// Loop through demand curve and find the price that maximizes the total return from the demand curve
		maxReturn := cosmosMath.NewUint(0)
		priceOfMaxReturn := cosmosMath.NewUint(0)
		for price, requests := range demandCurve {
			if cosmosMath.NewUint(uint64(len(requests))).Mul(price).GT(maxReturn) {
				maxReturn = cosmosMath.NewUint(uint64(len(requests))).Mul(price)
				priceOfMaxReturn = price
			}
		}
		topicBestPrices[topic.Id] = PriceAndReturn{priceOfMaxReturn, maxReturn}
		requestsToDrawDemandFrom[topic.Id] = demandCurve[priceOfMaxReturn]
	}

	// Sort topics by topicBestPrices
	sortedTopics := SortTopicsByReturnDescWithRandomTiebreaker(topicsActiveWithDemand, topicBestPrices, currentTime)
	// Take top am.keeper.MAX_TOPICS_PER_BLOCK number of topics with the highest demand
	topTopicsByReturn := sortedTopics[:uint(math.Min(float64(len(sortedTopics)), keeper.MAX_TOPICS_PER_BLOCK))]

	// Determine how many funds to draw from demand and Remove depleted/insufficiently funded requests
	totalFundsToDrawFromDemand := cosmosMath.NewUint(0)
	for _, topic := range topTopicsByReturn {
		bestPrice := topicBestPrices[topic.Id].Price
		numRequestsServed := 0
		for _, req := range requestsToDrawDemandFrom[topic.Id] {
			reqId, err := req.GetRequestId()
			if err != nil {
				fmt.Println("Error getting request demand: ", err)
				return nil, nil, err
			}
			reqDemand, err := am.keeper.GetRequestDemand(ctx, reqId)
			if err != nil {
				fmt.Println("Error getting request demand: ", err)
				return nil, nil, err
			}
			// all the previous conditionals were already applied to the requests in the previous loop
			// => should never be negative
			newReqDemand := reqDemand.Sub(bestPrice)
			am.keeper.SetRequestDemand(ctx, reqId, newReqDemand)
			if newReqDemand.LT(cosmosMath.NewUint(keeper.MIN_UNMET_DEMAND)) {
				// Should convey to users to not surprise them. This helps prevent spamming the mempool with requests that are not worth serving
				// The effectively burned dust is 1-time "cost" the consumer incurs when they create "subscriptions" they don't ever refill nor fill enough
				// This encourages consumers to maximize how much they fund any single request, discouraging a pattern of many less-funded requests
				am.keeper.RemoveFromMempool(ctx, req)
			}
			numRequestsServed++
		}
		totalFundsToDrawFromDemand = totalFundsToDrawFromDemand.Add(bestPrice.Mul(cosmosMath.NewUint(uint64(numRequestsServed))))
	}

	return &topTopicsByReturn, &totalFundsToDrawFromDemand, nil
}
