package inference_synthesis

import (
	"fmt"

	alloraMath "github.com/allora-network/allora-chain/math"

	emissions "github.com/allora-network/allora-chain/x/emissions/types"
)

type RunningWeightedLoss struct {
	UnnormalizedWeightedLoss Loss
	SumWeight                Weight
}

// Update the running information needed to calculate weighted loss per worker
func RunningWeightedAvgUpdate(
	runningWeightedAvg *RunningWeightedLoss,
	nextWeight Weight,
	nextValue Weight,
) (RunningWeightedLoss, error) {
	nextValTimesWeight, err := nextValue.Mul(nextWeight)
	if err != nil {
		return RunningWeightedLoss{}, err
	}
	newUnnormalizedWeightedLoss, err := runningWeightedAvg.UnnormalizedWeightedLoss.Add(nextValTimesWeight)
	if err != nil {
		return RunningWeightedLoss{}, err
	}
	newSumWeight, err := runningWeightedAvg.SumWeight.Add(nextWeight)
	if err != nil {
		return RunningWeightedLoss{}, err
	}
	return RunningWeightedLoss{
		UnnormalizedWeightedLoss: newUnnormalizedWeightedLoss,
		SumWeight:                newSumWeight,
	}, nil
}

// Convert the running weighted average objects to WorkerAttributedValues
func convertMapOfRunningWeightedLossesToWorkerAttributedValue[T emissions.WorkerAttributedValue | emissions.WithheldWorkerAttributedValue](
	runningWeightedLosses map[Worker]*RunningWeightedLoss,
	sortedWorkers []Worker,
	epsilon alloraMath.Dec,
) []*T {
	weightedLosses := make([]*T, 0)
	for _, worker := range sortedWorkers {
		runningLoss, ok := runningWeightedLosses[worker]
		if !ok {
			continue
		}
		if runningLoss.SumWeight.Lt(epsilon) {
			fmt.Println("Sum weight for worker is 0", worker)
			continue
		}
		normalizedWeightedLoss, err := runningLoss.UnnormalizedWeightedLoss.Quo(runningLoss.SumWeight)
		if err != nil {
			fmt.Println("Error normalizing weighted loss: ", err)
			continue
		}
		weightedLosses = append(weightedLosses, &T{
			Worker: worker,
			Value:  normalizedWeightedLoss,
		})
	}
	return weightedLosses
}

func CalcNetworkLosses(
	stakesByReputer map[Worker]Stake,
	reputerReportedLosses emissions.ReputerValueBundles,
	epsilon alloraMath.Dec,
) (emissions.ValueBundle, error) {
	// Make map from inferer to their running weighted-average loss
	runningWeightedCombinedLoss := RunningWeightedLoss{alloraMath.ZeroDec(), alloraMath.ZeroDec()}
	runningWeightedInfererLosses := make(map[Worker]*RunningWeightedLoss)
	runningWeightedForecasterLosses := make(map[Worker]*RunningWeightedLoss)
	runningWeightedNaiveLoss := RunningWeightedLoss{alloraMath.ZeroDec(), alloraMath.ZeroDec()}
	runningWeightedOneOutInfererLosses := make(map[Worker]*RunningWeightedLoss) // Withheld worker -> Forecaster -> Loss
	runningWeightedOneOutForecasterLosses := make(map[Worker]*RunningWeightedLoss)
	runningWeightedOneInForecasterLosses := make(map[Worker]*RunningWeightedLoss)

	for _, report := range reputerReportedLosses.ReputerValueBundles {
		if report.ValueBundle != nil {
			stakeAmount, err := alloraMath.NewDecFromSdkUint(stakesByReputer[report.ValueBundle.Reputer])
			if err != nil {
				return emissions.ValueBundle{}, err
			}

			// Update combined loss with reputer reported loss and stake
			runningWeightedCombinedLoss, err = RunningWeightedAvgUpdate(&runningWeightedCombinedLoss, stakeAmount, report.ValueBundle.CombinedValue)
			if err != nil {
				fmt.Println("Error updating running weighted average for next combined loss: ", err)
				return emissions.ValueBundle{}, err
			}

			// Not all reputers may have reported losses on the same set of inferers => important that the code below doesn't assume that!
			// Update inferer losses
			for _, loss := range report.ValueBundle.InfererValues {
				if runningWeightedInfererLosses[loss.Worker] == nil {
					runningWeightedInfererLosses[loss.Worker] = &RunningWeightedLoss{
						UnnormalizedWeightedLoss: alloraMath.ZeroDec(),
						SumWeight:                alloraMath.ZeroDec(),
					}
				}

				nextAvg, err := RunningWeightedAvgUpdate(runningWeightedInfererLosses[loss.Worker], stakeAmount, loss.Value)
				if err != nil {
					fmt.Println("Error updating running weighted average for inferer: ", err)
					return emissions.ValueBundle{}, err
				}
				runningWeightedInfererLosses[loss.Worker] = &nextAvg
			}

			// Update forecaster losses
			for _, loss := range report.ValueBundle.ForecasterValues {
				if runningWeightedForecasterLosses[loss.Worker] == nil {
					runningWeightedForecasterLosses[loss.Worker] = &RunningWeightedLoss{
						UnnormalizedWeightedLoss: alloraMath.ZeroDec(),
						SumWeight:                alloraMath.ZeroDec(),
					}
				}

				nextAvg, err := RunningWeightedAvgUpdate(runningWeightedForecasterLosses[loss.Worker], stakeAmount, loss.Value)
				if err != nil {
					fmt.Println("Error updating running weighted average for forecaster: ", err)
					return emissions.ValueBundle{}, err
				}
				runningWeightedForecasterLosses[loss.Worker] = &nextAvg
			}

			// Update naive loss
			runningWeightedNaiveLoss, err = RunningWeightedAvgUpdate(&runningWeightedNaiveLoss, stakeAmount, report.ValueBundle.NaiveValue)
			if err != nil {
				fmt.Println("Error updating running weighted average for naive loss: ", err)
				return emissions.ValueBundle{}, err
			}

			// Update one-out inferer losses
			for _, loss := range report.ValueBundle.OneOutInfererValues {
				if runningWeightedOneOutInfererLosses[loss.Worker] == nil {
					runningWeightedOneOutInfererLosses[loss.Worker] = &RunningWeightedLoss{
						UnnormalizedWeightedLoss: alloraMath.ZeroDec(),
						SumWeight:                alloraMath.ZeroDec(),
					}
				}
				nextAvg, err := RunningWeightedAvgUpdate(runningWeightedOneOutInfererLosses[loss.Worker], stakeAmount, loss.Value)
				if err != nil {
					fmt.Println("Error updating running weighted average for one-out inferer: ", err)
					return emissions.ValueBundle{}, err
				}
				runningWeightedOneOutInfererLosses[loss.Worker] = &nextAvg
			}

			// Update one-out forecaster losses
			for _, loss := range report.ValueBundle.OneOutForecasterValues {
				if runningWeightedOneOutForecasterLosses[loss.Worker] == nil {
					runningWeightedOneOutForecasterLosses[loss.Worker] = &RunningWeightedLoss{
						UnnormalizedWeightedLoss: alloraMath.ZeroDec(),
						SumWeight:                alloraMath.ZeroDec(),
					}
				}

				nextAvg, err := RunningWeightedAvgUpdate(runningWeightedOneOutForecasterLosses[loss.Worker], stakeAmount, loss.Value)
				if err != nil {
					fmt.Println("Error updating running weighted average for one-out forecaster: ", err)
					return emissions.ValueBundle{}, err
				}
				runningWeightedOneOutForecasterLosses[loss.Worker] = &nextAvg
			}

			// Update one-in forecaster losses
			for _, loss := range report.ValueBundle.OneInForecasterValues {
				if runningWeightedOneInForecasterLosses[loss.Worker] == nil {
					runningWeightedOneInForecasterLosses[loss.Worker] = &RunningWeightedLoss{
						UnnormalizedWeightedLoss: alloraMath.ZeroDec(),
						SumWeight:                alloraMath.ZeroDec(),
					}
				}

				nextAvg, err := RunningWeightedAvgUpdate(runningWeightedOneInForecasterLosses[loss.Worker], stakeAmount, loss.Value)
				if err != nil {
					fmt.Println("Error updating running weighted average for one-in forecaster: ", err)
					return emissions.ValueBundle{}, err
				}
				runningWeightedOneInForecasterLosses[loss.Worker] = &nextAvg
			}
		}
	}

	sortedInferers := GetSortedStringKeys(runningWeightedInfererLosses)
	sortedForecasters := GetSortedStringKeys(runningWeightedForecasterLosses)
	if runningWeightedCombinedLoss.SumWeight.Lt(epsilon) {
		fmt.Println("Sum weight for combined and naive loss is 0")
		return emissions.ValueBundle{}, emissions.ErrFractionDivideByZero
	}
	combinedValue, err := runningWeightedCombinedLoss.UnnormalizedWeightedLoss.Quo(runningWeightedCombinedLoss.SumWeight)
	if err != nil {
		fmt.Println("Error normalizing combined loss: ", err)
		return emissions.ValueBundle{}, err
	}
	// Convert the running weighted averages to WorkerAttributedValue/WithheldWorkerAttributedValue for inferers and forecasters
	infererLosses := convertMapOfRunningWeightedLossesToWorkerAttributedValue[emissions.WorkerAttributedValue](runningWeightedInfererLosses, sortedInferers, epsilon)
	forecasterLosses := convertMapOfRunningWeightedLossesToWorkerAttributedValue[emissions.WorkerAttributedValue](runningWeightedForecasterLosses, sortedForecasters, epsilon)
	naiveValue, err := runningWeightedNaiveLoss.UnnormalizedWeightedLoss.Quo(runningWeightedNaiveLoss.SumWeight)
	if err != nil {
		fmt.Println("Error normalizing naive loss: ", err)
		return emissions.ValueBundle{}, err
	}
	oneOutInfererLosses := convertMapOfRunningWeightedLossesToWorkerAttributedValue[emissions.WithheldWorkerAttributedValue](runningWeightedOneOutInfererLosses, sortedInferers, epsilon)
	oneOutForecasterLosses := convertMapOfRunningWeightedLossesToWorkerAttributedValue[emissions.WithheldWorkerAttributedValue](runningWeightedOneOutForecasterLosses, sortedForecasters, epsilon)
	oneInForecasterLosses := convertMapOfRunningWeightedLossesToWorkerAttributedValue[emissions.WorkerAttributedValue](runningWeightedOneInForecasterLosses, sortedForecasters, epsilon)

	output := emissions.ValueBundle{
		CombinedValue:          combinedValue,
		InfererValues:          infererLosses,
		ForecasterValues:       forecasterLosses,
		NaiveValue:             naiveValue,
		OneOutInfererValues:    oneOutInfererLosses,
		OneOutForecasterValues: oneOutForecasterLosses,
		OneInForecasterValues:  oneInForecasterLosses,
	}

	return output, nil
}

// Same as CalcNetworkLosses() but just returns the combined loss
func CalcCombinedNetworkLoss(
	stakesByReputer map[Worker]Stake,
	reputerReportedLosses *emissions.ReputerValueBundles,
	epsilon alloraMath.Dec,
) (Loss, error) {
	// Make map from inferer to their running weighted-average loss
	runningWeightedCombinedLoss := RunningWeightedLoss{alloraMath.ZeroDec(), alloraMath.ZeroDec()}

	for _, report := range reputerReportedLosses.ReputerValueBundles {
		if report.ValueBundle != nil {
			stakeAmount, err := alloraMath.NewDecFromSdkUint(stakesByReputer[report.ValueBundle.Reputer])
			if err != nil {
				fmt.Println("Error converting stake to Dec: ", err)
				return Loss{}, err
			}

			// Update combined loss with reputer reported loss and stake
			nextCombinedLoss, err := RunningWeightedAvgUpdate(
				&runningWeightedCombinedLoss,
				stakeAmount,
				report.ValueBundle.CombinedValue,
			)
			if err != nil {
				fmt.Println("Error updating running weighted average for combined loss: ", err)
				return alloraMath.ZeroDec(), err
			}
			runningWeightedCombinedLoss = nextCombinedLoss
		}
	}

	if runningWeightedCombinedLoss.SumWeight.Lt(epsilon) {
		fmt.Println("Sum weight for combined loss is 0")
		return Loss{}, emissions.ErrFractionDivideByZero
	}
	combinedValue, err := runningWeightedCombinedLoss.UnnormalizedWeightedLoss.Quo(runningWeightedCombinedLoss.SumWeight)
	if err != nil {
		fmt.Println("Error normalizing combined loss: ", err)
		return Loss{}, err
	}

	return combinedValue, nil
}
