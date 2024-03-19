package module_test

import (
	"time"

	cosmosMath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allora-network/allora-chain/x/emissions/module"
	"github.com/allora-network/allora-chain/x/emissions/types"
)

func (s *ModuleTestSuite) TestGetWorkerScoreInferenceTask() {
	timeNow := uint64(time.Now().UTC().Unix())

	// Create a topic
	topicIds, err := mockCreateTopics(s, 1)
	s.NoError(err, "Error creating topic")
	topicId := topicIds[0]

	// Create and register 2 reputers in topic
	reputers, err := mockSomeReputers(s, topicId)
	s.NoError(err, "Error creating reputers")

	// Create and register 2 workers in topic
	workers, err := mockSomeWorkers(s, topicId)
	s.NoError(err, "Error creating workers")

	// Add a lossBundle for each reputer
	var reputersLossBundles []*types.LossBundle
	reputer1LossBundle := types.LossBundle{
		TopicId: topicId,
		Reputer: reputers[0].String(),
		InfererLosses: []*types.WorkerAttributedLoss{
			{
				Worker: workers[0].String(),
				Value: cosmosMath.NewUint(90),
			},
			{
				Worker: workers[1].String(),
				Value: cosmosMath.NewUint(100),
			},
		},
		// Increased loss when removing for worker 1
		OneOutLosses: []*types.WorkerAttributedLoss{
			{
				Worker: workers[0].String(),
				Value: cosmosMath.NewUint(115),
			},
			{
				Worker: workers[1].String(),
				Value: cosmosMath.NewUint(100),
			},
		},
	}
	reputersLossBundles = append(reputersLossBundles, &reputer1LossBundle)
	reputer2LossBundle := types.LossBundle{
		TopicId: topicId,
		Reputer: reputers[1].String(),
		InfererLosses: []*types.WorkerAttributedLoss{
			{
				Worker: workers[0].String(),
				Value: cosmosMath.NewUint(85),
			},
			{
				Worker: workers[1].String(),
				Value: cosmosMath.NewUint(100),
			},
		},
		// Increased loss when removing for worker 1
		OneOutLosses: []*types.WorkerAttributedLoss{
			{
				Worker: workers[0].String(),
				Value: cosmosMath.NewUint(120),
			},
			{
				Worker: workers[1].String(),
				Value: cosmosMath.NewUint(100),
			},
		},
	}
	reputersLossBundles = append(reputersLossBundles, &reputer2LossBundle)

	err = s.emissionsKeeper.InsertLossBudles(s.ctx, topicId, timeNow, types.LossBundles{LossBundles: reputersLossBundles})
	s.NoError(err, "Error adding lossBundle for worker")

	// Get LossBundles
	lossBundles, err := s.emissionsKeeper.GetLossBundles(s.ctx, topicId, timeNow)
	s.NoError(err, "Error getting lossBundles")

	// Get reputers stakes and reported losses for each worker
	var reputersStakes []float64
	var reputersWorker1ReportedLosses []float64
	var reputersWorker2ReportedLosses []float64
	var reputersWorker1ReportedOneOutLosses []float64
	var reputersWorker2ReportedOneOutLosses []float64

	for _, lossBundle := range lossBundles.LossBundles {
		reputerAddr, err := sdk.AccAddressFromBech32(lossBundle.Reputer)
		s.NoError(err, "Error getting reputerAddr")

		reputerStake, err := s.emissionsKeeper.GetStakeOnTopicFromReputer(s.ctx, topicId, reputerAddr)
		s.NoError(err, "Error getting reputerStake")

		reputerStakeFloat := float64(reputerStake.BigInt().Int64())
		reputersStakes = append(reputersStakes, reputerStakeFloat)
		for _, workerLoss := range lossBundle.InfererLosses {
			if workerLoss.Worker == workers[0].String() {
				reputersWorker1ReportedLosses = append(reputersWorker1ReportedLosses, float64(workerLoss.Value.BigInt().Int64()))
			} else if workerLoss.Worker == workers[1].String() {
				reputersWorker2ReportedLosses = append(reputersWorker2ReportedLosses, float64(workerLoss.Value.BigInt().Int64()))
			}
		}

		// Add OneOutLosses
		for _, workerLoss := range lossBundle.OneOutLosses {
			if workerLoss.Worker == workers[0].String() {
				reputersWorker1ReportedOneOutLosses = append(reputersWorker1ReportedOneOutLosses, float64(workerLoss.Value.BigInt().Int64()))
			} else if workerLoss.Worker == workers[1].String() {
				reputersWorker2ReportedOneOutLosses = append(reputersWorker2ReportedOneOutLosses, float64(workerLoss.Value.BigInt().Int64()))
			}
		}
	}

	// Get Stake Weighted Loss - Inference Loss
	worker1StakeWeightedLoss, err := module.GetStakeWeightedLoss(reputersStakes, reputersWorker1ReportedLosses)
	s.NoError(err, "Error getting stakeWeightedLoss")
	s.NotEqual(0, worker1StakeWeightedLoss, "Expected worker1StakeWeightedLoss to be non-zero")

	worker2StakeWeightedLoss, err := module.GetStakeWeightedLoss(reputersStakes, reputersWorker2ReportedLosses)
	s.NoError(err, "Error getting stakeWeightedLoss")
	s.NotEqual(0, worker2StakeWeightedLoss, "Expected worker2StakeWeightedLoss to be non-zero")

	// Get Stake Weighted Loss - OneOut Loss
	worker1StakeWeightedOneOutLoss, err := module.GetStakeWeightedLoss(reputersStakes, reputersWorker1ReportedOneOutLosses)
	s.NoError(err, "Error getting stakeWeightedLoss")
	s.NotEqual(0, worker1StakeWeightedOneOutLoss, "Expected worker1StakeWeightedOneOutLoss to be non-zero")

	worker2StakeWeightedOneOutLoss, err := module.GetStakeWeightedLoss(reputersStakes, reputersWorker2ReportedOneOutLosses)
	s.NoError(err, "Error getting stakeWeightedLoss")
	s.NotEqual(0, worker2StakeWeightedOneOutLoss, "Expected worker2StakeWeightedOneOutLoss to be non-zero")

	// Get Worker Score
	worker1Score := module.GetWorkerScore(worker1StakeWeightedLoss, worker1StakeWeightedOneOutLoss)
	s.NotEqual(0, worker1Score, "Expected worker1Score to be non-zero")

	worker2Score := module.GetWorkerScore(worker2StakeWeightedLoss, worker2StakeWeightedOneOutLoss)
	s.NotEqual(0, worker2Score, "Expected worker2Score to be non-zero")
}

func (s *ModuleTestSuite) TestGetStakeWeightedLoss() {
	timeNow := uint64(time.Now().UTC().Unix())

	// Create a topic
	topicIds, err := mockCreateTopics(s, 1)
	s.NoError(err, "Error creating topic")
	topicId := topicIds[0]

	// Create and register 2 reputers in topic
	reputers, err := mockSomeReputers(s, topicId)
	s.NoError(err, "Error creating reputers")

	// Add a lossBundle for each reputer
	losses := []cosmosMath.Uint{cosmosMath.NewUint(150), cosmosMath.NewUint(250)}

	var newLossBundles []*types.LossBundle
	for i, reputer := range reputers {
		lossBundle := types.LossBundle{
			Reputer:      reputer.String(),
			CombinedLoss: losses[i],
		}
		newLossBundles = append(newLossBundles, &lossBundle)
	}

	err = s.emissionsKeeper.InsertLossBudles(s.ctx, topicId, timeNow, types.LossBundles{LossBundles: newLossBundles})
	s.NoError(err, "Error adding lossBundle for reputer")

	var reputersStakes []float64
	var reputersReportedLosses []float64

	// Get LossBundles
	lossBundles, err := s.emissionsKeeper.GetLossBundles(s.ctx, topicId, timeNow)
	s.NoError(err, "Error getting lossBundles")

	// Get stakes and reported losses
	for _, lossBundle := range lossBundles.LossBundles {
		reputerAddr, err := sdk.AccAddressFromBech32(lossBundle.Reputer)
		s.NoError(err, "Error getting reputerAddr")

		reputerStake, err := s.emissionsKeeper.GetStakeOnTopicFromReputer(s.ctx, topicId, reputerAddr)
		s.NoError(err, "Error getting reputerStake")

		reputerStakeFloat, _ := reputerStake.BigInt().Float64()
		reputersStakes = append(reputersStakes, reputerStakeFloat)
		reputersReportedLosses = append(reputersReportedLosses, float64(lossBundle.CombinedLoss.BigInt().Int64()))
	}

	expectedStakeWeightedLoss, err := module.GetStakeWeightedLoss(reputersStakes, reputersReportedLosses)
	s.NoError(err, "Error getting stakeWeightedLoss")
	s.NotEqual(0, expectedStakeWeightedLoss, "Expected stakeWeightedLoss to be non-zero")
}
