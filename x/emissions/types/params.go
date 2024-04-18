package types

import (
	cosmosMath "cosmossdk.io/math"
	alloraMath "github.com/allora-network/allora-chain/math"
)

type BlockHeight = int64

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		Version:                         "0.0.3",                                   // version of the protocol should be in lockstep with github release tag version
		RewardCadence:                   int64(600),                                // length of an "epoch" for rewards payouts in blocks; coupled with TopicRewardAlpha
		MinTopicUnmetDemand:             cosmosMath.NewUint(100),                   // total unmet demand for a topic < this => don't run inference solicatation or loss update
		MaxTopicsPerBlock:               uint64(2048),                              // max number of topics to run cadence for per block
		MinRequestUnmetDemand:           cosmosMath.NewUint(1),                     // delete requests if they have below this demand remaining
		MaxMissingInferencePercent:      alloraMath.MustNewDecFromString("0.2"),    // if a worker has this percentage of inferences missing, they are penalized
		RequiredMinimumStake:            cosmosMath.NewUint(100),                   // minimum stake required to be a worker or reputer
		RemoveStakeDelayWindow:          int64(60 * 60 * 24),                       // 1 day in seconds
		MinEpochLength:                  1,                                         // 1 block
		MaxInferenceRequestValidity:     int64(6 * 60 * 24 * 7 * 52),               // approximately 1 year in number of blocks
		MaxRequestCadence:               int64(6 * 60 * 24 * 7 * 52),               // approximately 1 year in number of blocks
		Sharpness:                       alloraMath.NewDecFromInt64(20),            // controls going from stake-weighted consensus at low values to majority vote of above-average stake holders at high values
		BetaEntropy:                     alloraMath.MustNewDecFromString("0.25"),   // controls resilience of reward payouts against copycat workers
		LearningRate:                    alloraMath.MustNewDecFromString("0.01"),   // speed of gradient descent
		MaxGradientThreshold:            alloraMath.MustNewDecFromString("0.001"),  // gradient descent stops when gradient falls below this
		MinStakeFraction:                alloraMath.MustNewDecFromString("0.5"),    // minimum fraction of stake that should be listened to when setting consensus listening coefficients
		MaxWorkersPerTopicRequest:       uint64(20),                                // maximum number of workers that can be assigned to a single inference request
		MaxReputersPerTopicRequest:      uint64(20),                                // maximum number of reputers that can be assigned to a single loss request
		Epsilon:                         alloraMath.MustNewDecFromString("0.0001"), // 0 threshold to prevent div by 0 and 0-approximation errors
		PInferenceSynthesis:             alloraMath.NewDecFromInt64(2),             // free parameter used in the gradient function phi' for inference synthesis
		PRewardSpread:                   alloraMath.NewDecFromInt64(1),             // fiducial value = 1; Exponent for W_i total reward allocated to reputers per timestep
		AlphaRegret:                     alloraMath.MustNewDecFromString("0.1"),    // how much to weight the most recent log-loss differences in regret EMA update
		MaxUnfulfilledWorkerRequests:    uint64(100),                               // maximum number of outstanding nonces for worker requests from the chain
		MaxUnfulfilledReputerRequests:   uint64(100),                               // maximum number of outstanding nonces for reputer requests from the chain
		NumberExpectedInferenceSybils:   uint64(10),                                // global number for calculation tax of worker reward
		SybilTaxExponent:                uint64(3),                                 // global parameter for calculation tax of worker reward
		TopicRewardStakeImportance:      alloraMath.MustNewDecFromString("0.5"),    // importance of stake in determining rewards for a topic
		TopicRewardFeeRevenueImportance: alloraMath.MustNewDecFromString("0.5"),    // importance of fee revenue in determining rewards for a topic
		TopicRewardAlpha:                alloraMath.MustNewDecFromString("0.5"),    // alpha for topic reward calculation; coupled with RewardCadence
		TaskRewardAlpha:                 alloraMath.MustNewDecFromString("0.1"),    // alpha for task reward calculation used to calculate  ~U_ij, ~V_ik, ~W_im
		ValidatorsVsAlloraPercentReward: alloraMath.MustNewDecFromString("0.25"),   // 25% rewards go to cosmos network validators
		MaxSamplesToScaleScores:         uint64(10),                                // maximum number of previous scores to store and use for standard deviation calculation
		MaxTopWorkersToReward:           uint64(10),                                // max this many top workers by score are rewarded for a topic
		MaxTopReputersToReward:          uint64(10),                                // max this many top reputers by score are rewarded for a topic
		CreateTopicFee:                  cosmosMath.NewInt(10),                     // topic registration fee
		SigmoidA:                        alloraMath.NewDecFromInt64(8),             // sigmoid function parameter, a = 8
		SigmoidB:                        alloraMath.MustNewDecFromString("0.5"),    // sigmoid function parameter, b = 0.5
		MaxRetriesToFulfilNoncesWorker:  int64(3),                                  // max throttle of simultaneous unfulfilled worker requests
		MaxRetriesToFulfilNoncesReputer: int64(3),                                  // max throttle of simultaneous unfulfilled reputer requests
	}
}

func DefaultParamsVersion() string {
	return DefaultParams().Version
}

func DefaultParamsEpochLength() BlockHeight {
	return DefaultParams().RewardCadence
}

func DefaultParamsMinTopicUnmetDemand() cosmosMath.Uint {
	return DefaultParams().MinTopicUnmetDemand
}

func DefaultParamsMaxTopicsPerBlock() uint64 {
	return DefaultParams().MaxTopicsPerBlock
}

func DefaultParamsMinRequestUnmetDemand() cosmosMath.Uint {
	return DefaultParams().MinRequestUnmetDemand
}

func DefaultParamsMaxMissingInferencePercent() alloraMath.Dec {
	return DefaultParams().MaxMissingInferencePercent
}

func DefaultParamsRequiredMinimumStake() cosmosMath.Uint {
	return DefaultParams().RequiredMinimumStake
}

func DefaultParamsRemoveStakeDelayWindow() BlockHeight {
	return DefaultParams().RemoveStakeDelayWindow
}

func DefaultParamsMinEpochLength() BlockHeight {
	return DefaultParams().MinEpochLength
}

func DefaultParamsMaxInferenceRequestValidity() BlockHeight {
	return DefaultParams().MaxInferenceRequestValidity
}

func DefaultParamsMaxRequestCadence() BlockHeight {
	return DefaultParams().MaxRequestCadence
}

func DefaultParamsSharpness() alloraMath.Dec {
	return DefaultParams().Sharpness
}

func DefaultParamsBetaEntropy() alloraMath.Dec {
	return DefaultParams().BetaEntropy
}

func DefaultParamsLearningRate() alloraMath.Dec {
	return DefaultParams().LearningRate
}

func DefaultParamsMaxGradientThreshold() alloraMath.Dec {
	return DefaultParams().MaxGradientThreshold
}

func DefaultParamsMinStakeFraction() alloraMath.Dec {
	return DefaultParams().MinStakeFraction
}

func DefaultParamsMaxWorkersPerTopicRequest() uint64 {
	return DefaultParams().MaxWorkersPerTopicRequest
}

func DefaultParamsMaxReputersPerTopicRequest() uint64 {
	return DefaultParams().MaxReputersPerTopicRequest
}

func DefaultParamsEpsilon() alloraMath.Dec {
	return DefaultParams().Epsilon
}

func DefaultParamsPInferenceSynthesis() alloraMath.Dec {
	return DefaultParams().PInferenceSynthesis
}

func DefaultParamsPRewardSpread() alloraMath.Dec {
	return DefaultParams().PRewardSpread
}

func DefaultParamsAlphaRegret() alloraMath.Dec {
	return DefaultParams().AlphaRegret
}

func DefaultParamsMaxUnfulfilledWorkerRequestNonces() uint64 {
	return DefaultParams().MaxUnfulfilledWorkerRequests
}

func DefaultParamsMaxUnfulfilledReputerRequestNonces() uint64 {
	return DefaultParams().MaxUnfulfilledReputerRequests
}

func DefaultParamsNumberExpectedInfernceSybils() uint64 {
	return DefaultParams().NumberExpectedInferenceSybils
}

func DefaultParamsSybilTaxExponent() uint64 {
	return DefaultParams().SybilTaxExponent
}

func DefaultParamsTopicRewardStakeImportance() alloraMath.Dec {
	return DefaultParams().TopicRewardStakeImportance
}

func DefaultParamsTopicRewardFeeRevenueImportance() alloraMath.Dec {
	return DefaultParams().TopicRewardFeeRevenueImportance
}

func DefaultParamsTopicRewardAlpha() alloraMath.Dec {
	return DefaultParams().TopicRewardAlpha
}

func DefaultParamsValidatorsVsAlloraPercentReward() alloraMath.Dec {
	return DefaultParams().ValidatorsVsAlloraPercentReward
}

func DefaultParamsMaxSamplesToScaleScores() uint64 {
	return DefaultParams().MaxSamplesToScaleScores
}

func DefaultParamsCreateTopicFee() cosmosMath.Int {
	return DefaultParams().CreateTopicFee
}

func DefaultParamsMaxTopWorkersToReward() uint64 {
	return DefaultParams().MaxTopWorkersToReward
}

func DefaultParamsMaxTopReputersToReward() uint64 {
	return DefaultParams().MaxTopReputersToReward
}

func DefaultParamsSigmoidA() alloraMath.Dec {
	return DefaultParams().SigmoidA
}

func DefaultParamsSigmoidB() alloraMath.Dec {
	return DefaultParams().SigmoidB
}

func DefaultParamsMaxRetriesToFulfilNoncesWorker() int64 {
	return DefaultParams().MaxRetriesToFulfilNoncesWorker
}

func DefaultParamsMaxRetriesToFulfilNoncesReputer() int64 {
	return DefaultParams().MaxRetriesToFulfilNoncesReputer
}

// Validate does the sanity check on the params.
func (p Params) Validate() error {
	// Sanity check goes here.
	return nil
}
