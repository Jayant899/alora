package module

import (
	"fmt"
	"math"

	errors "cosmossdk.io/errors"
	emissions "github.com/allora-network/allora-chain/x/emissions/types"
)

// GetfUniqueAgg calculates the unique value or impact of each forecaster.
// f^+
func GetfUniqueAgg(numForecasters float64) float64 {
	return 1.0 / math.Pow(2.0, (numForecasters-1.0))
}

// GetFinalWorkerScoreForecastTask calculates the worker score in forecast task.
// T_ik
func GetFinalWorkerScoreForecastTask(scoreOneIn, scoreOneOut, fUniqueAgg float64) float64 {
	return fUniqueAgg*scoreOneIn + (1-fUniqueAgg)*scoreOneOut
}

// GetWorkerScore calculates the worker score based on the losses and lossesCut.
// Consider the staked weighted inference loss and one-out loss to calculate the worker score.
// T_ij / T^-_ik / T^+_ik
func GetWorkerScore(losses, lossesOneOut float64) float64 {
	deltaLogLoss := math.Log10(lossesOneOut) - math.Log10(losses)
	return deltaLogLoss
}

// GetStakeWeightedLoss calculates the stake-weighted average loss.
// Consider the losses and the stake of each reputer to calculate the stake-weighted loss.
// The stake weighted loss is used to calculate the network-wide losses.
// L_i / L_ij / L_ik / L^-_i / L^-_il / L^+_ik
func GetStakeWeightedLoss(reputersStakes, reputersReportedLosses []float64) (float64, error) {
	if len(reputersStakes) != len(reputersReportedLosses) {
		return 0, fmt.Errorf("slices must have the same length")
	}

	totalStake := 0.0
	for _, stake := range reputersStakes {
		totalStake += stake
	}

	if totalStake == 0 {
		return 0, fmt.Errorf("total stake cannot be zero")
	}

	var stakeWeightedLoss float64 = 0
	for i, loss := range reputersReportedLosses {
		if loss <= 0 {
			return 0, fmt.Errorf("loss values must be greater than zero")
		}
		weightedLoss := (reputersStakes[i] * math.Log10(loss)) / totalStake
		stakeWeightedLoss += weightedLoss
	}

	return stakeWeightedLoss, nil
}

// Implements the potential function phi for the module
// this is equation 6 from the litepaper:
// ϕ_p(x) = (ln(1 + e^x))^p
//
// error handling:
// float Inf can be generated for values greater than 1.7976931348623157e+308
// e^x can create Inf
// ln(blah)^p can create Inf for sufficiently large ln result
// NaN is impossible as 1+e^x is always positive no matter the value of x
// and pow only produces NaN for NaN input
// therefore we only return one type of error and that is if phi overflows.
func phi(p float64, x float64) (float64, error) {
	if math.IsNaN(p) || math.IsInf(p, 0) || math.IsNaN(x) || math.IsInf(x, 0) {
		return 0, emissions.ErrPhiInvalidInput
	}
	eToTheX := math.Exp(x)
	onePlusEToTheX := 1 + eToTheX
	if math.IsInf(onePlusEToTheX, 0) {
		return 0, emissions.ErrEToTheXExponentiationIsInfinity
	}
	naturalLog := math.Log(onePlusEToTheX)
	result := math.Pow(naturalLog, p)
	if math.IsInf(result, 0) {
		return 0, emissions.ErrLnToThePExponentiationIsInfinity
	}
	// should theoretically never be possible with the above checks
	if math.IsNaN(result) {
		return 0, emissions.ErrPhiResultIsNaN
	}
	return result, nil
}

// Adjusted stake for calculating consensus S hat
// ^S_im = 1 - ϕ_1^−1(η) * ϕ1[ −η * (((N_r * a_im * S_im) / (Σ_m(a_im * S_im))) − 1 )]
// we use eta = 20 as the fiducial value decided in the paper
// phi_1 refers to the phi function with p = 1
// INPUTS:
// This function expects that allStakes
// and allListeningCoefficients are slices of the same length
// and the index to each slice corresponds to the same reputer
func adjustedStake(
	stake float64,
	allStakes []float64,
	listeningCoefficient float64,
	allListeningCoefficients []float64,
	numReputers float64,
) (float64, error) {
	if len(allStakes) != len(allListeningCoefficients) ||
		len(allStakes) == 0 ||
		len(allListeningCoefficients) == 0 {
		return 0, emissions.ErrAdjustedStakeInvalidSliceLength
	}
	// renaming variables just to be more legible with the formula
	S_im := stake
	a_im := listeningCoefficient
	N_r := numReputers

	denominator := 0.0
	for i, s := range allStakes {
		a := allListeningCoefficients[i]
		denominator += (a * s)
	}
	numerator := N_r * a_im * S_im
	stakeFraction := numerator / denominator
	stakeFraction = stakeFraction - 1
	stakeFraction = stakeFraction * -20 // eta = 20

	phi_1_stakeFraction, err := phi(1, stakeFraction)
	if err != nil {
		return 0, err
	}
	phi_1_Eta, err := phi(1, 20)
	if err != nil {
		return 0, err
	}
	// phi_1_Eta is taken to the -1 power
	// and then multiplied by phi_1_stakeFraction
	// so we can just treat it as phi_1_stakeFraction / phi_1_Eta
	phiVal := phi_1_stakeFraction / phi_1_Eta
	ret := 1 - phiVal

	if math.IsInf(ret, 0) {
		return 0, errors.Wrapf(emissions.ErrAdjustedStakeIsInfinity, "stake: %f", stake)
	}
	if math.IsNaN(ret) {
		return 0, errors.Wrapf(emissions.ErrAdjustedStakeIsNaN, "stake: %f", stake)
	}
	return ret, nil
}

// Used by Rewards fraction functions,
// all the exponential moving average functions take the form
// x_average=α*x_current + (1-α)*x_previous
//
// this covers the equations
// Uij = αUij + (1 − α)Ui−1,j
// ̃Vik = αVik + (1 − α)Vi−1,k
// ̃Wim = αWim + (1 − α)Wi−1,m
func exponentialMovingAverage(alpha float64, current float64, previous float64) (float64, error) {
	if math.IsNaN(alpha) || math.IsInf(alpha, 0) {
		return 0, errors.Wrapf(emissions.ErrExponentialMovingAverageInvalidInput, "alpha: %f", alpha)
	}
	if math.IsNaN(current) || math.IsInf(current, 0) {
		return 0, errors.Wrapf(emissions.ErrExponentialMovingAverageInvalidInput, "current: %f", current)
	}
	if math.IsNaN(previous) || math.IsInf(previous, 0) {
		return 0, errors.Wrapf(emissions.ErrExponentialMovingAverageInvalidInput, "previous: %f", previous)
	}

	// THE ONLY LINE OF CODE IN THIS FUNCTION
	// THAT ISN'T ERROR CHECKING IS HERE
	ret := alpha*current + (1-alpha)*previous

	if math.IsInf(ret, 0) {
		return 0, emissions.ErrExponentialMovingAverageIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrExponentialMovingAverageIsNaN
	}
	return ret, nil
}

// f_ij, f_ik, and f_im are all reward fractions
// that require computing the ratio of one participant to all participants
// yes this is extremely simple math
// yes we write a separate function for it anyway. The compiler can inline it if necessary
// normalizeToArray = value / sum(allValues)
// this covers equations
// f_ij =  (̃U_ij) / ∑_j(̃Uij)
// f_ik = (̃Vik) / ∑_k(̃Vik)
// fim =  (̃Wim) / ∑_m(̃Wim)
func normalizeAgainstSlice(value float64, allValues []float64) (float64, error) {
	if len(allValues) == 0 {
		return 0, emissions.ErrFractionInvalidSliceLength
	}
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, errors.Wrapf(emissions.ErrFractionInvalidInput, "value: %f", value)
	}
	sumValues := 0.0
	for i, v := range allValues {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return 0, errors.Wrapf(emissions.ErrFractionInvalidInput, "allValues[%d]: %f", i, v)
		}
		sumValues += v
	}
	if sumValues == 0 {
		return 0, emissions.ErrFractionDivideByZero
	}
	ret := value / sumValues
	if math.IsInf(ret, 0) {
		return 0, emissions.ErrFractionIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrFractionIsNaN
	}

	return ret, nil
}

// We define a modified entropy for each class
// ({F_i, G_i, H_i} for the inference, forecasting, and reputer tasks, respectively
// Fi = - ∑_j( f_ij * ln(f_ij) * (N_{i,eff} / N_i)^β )
// Gi = - ∑_k( f_ik * ln(f_ik) * (N_{f,eff} / N_f)^β )
// Hi = - ∑_m( f_im * ln(f_im) * (N_{r,eff} / N_r)^β )
// we use beta = 0.25 as a fiducial value
func entropy(allFs []float64, N_eff float64, numParticipants float64, beta float64) (float64, error) {
	if math.IsInf(N_eff, 0) ||
		math.IsNaN(N_eff) ||
		math.IsInf(numParticipants, 0) ||
		math.IsNaN(numParticipants) ||
		math.IsInf(beta, 0) ||
		math.IsNaN(beta) {
		return 0, errors.Wrapf(
			emissions.ErrEntropyInvalidInput,
			"N_eff: %f, numParticipants: %f, beta: %f",
			N_eff,
			numParticipants,
			beta,
		)
	}
	// simple variable rename to look more like the equations,
	// hopefully compiler is smart enough to inline it
	N := numParticipants

	multiplier := N_eff / N
	multiplier = math.Pow(multiplier, beta)

	sum := 0.0
	for i, f := range allFs {
		if math.IsInf(f, 0) || math.IsNaN(f) {
			return 0, errors.Wrapf(emissions.ErrEntropyInvalidInput, "allFs[%d]: %f", i, f)
		}
		sum += f * math.Log(f)
	}

	ret := -1 * sum * multiplier
	if math.IsInf(ret, 0) {
		return 0, errors.Wrapf(
			emissions.ErrEntropyIsInfinity,
			"sum of f: %f, multiplier: %f",
			sum,
			multiplier,
		)
	}
	if math.IsNaN(ret) {
		return 0, errors.Wrapf(
			emissions.ErrEntropyIsNaN,
			"sum of f: %f, multiplier: %f",
			sum,
			multiplier,
		)
	}
	return ret, nil
}

// The number ratio term captures the number of participants in the network
// to prevent sybil attacks in the rewards distribution
// This function captures
// N_{i,eff} = 1 / ∑_j( f_ij^2 )
// N_{f,eff} = 1 / ∑_k( f_ik^2 )
// N_{r,eff} = 1 / ∑_m( f_im^2 )
func numberRatio(rewardFractions []float64) (float64, error) {
	if len(rewardFractions) == 0 {
		return 0, emissions.ErrNumberRatioInvalidSliceLength
	}
	sum := 0.0
	for i, f := range rewardFractions {
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return 0, errors.Wrapf(emissions.ErrNumberRatioInvalidInput, "rewardFractions[%d]: %f", i, f)
		}
		sum += f * f
	}
	if sum == 0 {
		return 0, emissions.ErrNumberRatioDivideByZero
	}
	ret := 1 / sum
	if math.IsInf(ret, 0) {
		return 0, emissions.ErrNumberRatioIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrNumberRatioIsNaN
	}
	return ret, nil
}

// inference rewards calculation
// U_i = ((1 - χ) * γ * F_i * E_i ) / (F_i + G_i + H_i)
func inferenceRewards(
	chi float64,
	gamma float64,
	entropyInference float64,
	entropyForecasting float64,
	entropyReputer float64,
	timeStep float64,
) (float64, error) {
	if math.IsNaN(chi) || math.IsInf(chi, 0) ||
		math.IsNaN(gamma) || math.IsInf(gamma, 0) ||
		math.IsNaN(entropyInference) || math.IsInf(entropyInference, 0) ||
		math.IsNaN(entropyForecasting) || math.IsInf(entropyForecasting, 0) ||
		math.IsNaN(entropyReputer) || math.IsInf(entropyReputer, 0) ||
		math.IsNaN(timeStep) || math.IsInf(timeStep, 0) {
		return 0, errors.Wrapf(
			emissions.ErrInferenceRewardsInvalidInput,
			"chi: %f, gamma: %f, entropyInference: %f, entropyForecasting: %f, entropyReputer: %f, timeStep: %f",
			chi,
			gamma,
			entropyInference,
			entropyForecasting,
			entropyReputer,
			timeStep,
		)
	}
	ret := ((1 - chi) * gamma * entropyInference * timeStep) / (entropyInference + entropyForecasting + entropyReputer)
	if math.IsInf(ret, 0) {
		return 0, emissions.ErrInferenceRewardsIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrInferenceRewardsIsNaN
	}
	return ret, nil
}

// forecaster rewards calculation
// V_i = (χ * γ * G_i * E_i) / (F_i + G_i + H_i)
func forecastingRewards(
	chi float64,
	gamma float64,
	entropyInference float64,
	entropyForecasting float64,
	entropyReputer float64,
	timeStep float64,
) (float64, error) {
	if math.IsNaN(chi) || math.IsInf(chi, 0) ||
		math.IsNaN(gamma) || math.IsInf(gamma, 0) ||
		math.IsNaN(entropyInference) || math.IsInf(entropyInference, 0) ||
		math.IsNaN(entropyForecasting) || math.IsInf(entropyForecasting, 0) ||
		math.IsNaN(entropyReputer) || math.IsInf(entropyReputer, 0) ||
		math.IsNaN(timeStep) || math.IsInf(timeStep, 0) {
		return 0, errors.Wrapf(
			emissions.ErrForecastingRewardsInvalidInput,
			"chi: %f, gamma: %f, entropyInference: %f, entropyForecasting: %f, entropyReputer: %f, timeStep: %f",
			chi,
			gamma,
			entropyInference,
			entropyForecasting,
			entropyReputer,
			timeStep,
		)
	}
	ret := (chi * gamma * entropyForecasting * timeStep) / (entropyInference + entropyForecasting + entropyReputer)
	if math.IsInf(ret, 0) {
		return 0, emissions.ErrForecastingRewardsIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrForecastingRewardsIsNaN
	}
	return ret, nil
}

// reputer rewards calculation
// W_i = (H_i * E_i) / (F_i + G_i + H_i)
func reputerRewards(
	entropyInference float64,
	entropyForecasting float64,
	entropyReputer float64,
	timeStep float64,
) (float64, error) {
	if math.IsNaN(entropyInference) || math.IsInf(entropyInference, 0) ||
		math.IsNaN(entropyForecasting) || math.IsInf(entropyForecasting, 0) ||
		math.IsNaN(entropyReputer) || math.IsInf(entropyReputer, 0) ||
		math.IsNaN(timeStep) || math.IsInf(timeStep, 0) {
		return 0, errors.Wrapf(
			emissions.ErrReputerRewardsInvalidInput,
			"entropyInference: %f, entropyForecasting: %f, entropyReputer: %f, timeStep: %f",
			entropyInference,
			entropyForecasting,
			entropyReputer,
			timeStep,
		)
	}
	ret := (entropyReputer * timeStep) / (entropyInference + entropyForecasting + entropyReputer)
	if math.IsInf(ret, 0) {
		return 0, emissions.ErrReputerRewardsIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrReputerRewardsIsNaN
	}
	return ret, nil
}

// The performance score of the entire forecasting task T_i
// is positive if the removal of the forecasting task would
// increase the network loss, and is negative if its removal
// would decrease the network loss
// We subtract the log-loss of the complete network inference
// (L_i) from that of the naive network (L_i^-), which is
// obtained by omitting all forecast-implied inferences
// T_i = log L_i^- - log L_i
func forecastingPerformanceScore(
	naiveNetworkInferenceLoss float64,
	networkInferenceLoss float64,
) (float64, error) {
	if math.IsNaN(networkInferenceLoss) || math.IsInf(networkInferenceLoss, 0) ||
		math.IsNaN(naiveNetworkInferenceLoss) || math.IsInf(naiveNetworkInferenceLoss, 0) {
		return 0, errors.Wrapf(
			emissions.ErrForecastingPerformanceScoreInvalidInput,
			"networkInferenceLoss: %f, naiveNetworkInferenceLoss: %f",
			networkInferenceLoss,
			naiveNetworkInferenceLoss,
		)
	}
	ret := math.Log10(naiveNetworkInferenceLoss) - math.Log10(networkInferenceLoss)

	if math.IsInf(ret, 0) {
		return 0, emissions.ErrForecastingPerformanceScoreIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrForecastingPerformanceScoreIsNaN
	}
	return ret, nil
}

// sigmoid function
// σ(x) = 1/(1+e^{-x}) = e^x/(1+e^x)
func sigmoid(x float64) (float64, error) {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return 0, emissions.ErrSigmoidInvalidInput
	}
	ret := math.Exp(x) / (1 + math.Exp(x))
	if math.IsInf(ret, 0) {
		return 0, emissions.ErrSigmoidIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrSigmoidIsNaN
	}
	return ret, nil
}

// we apply a utility function to the forecasting performance score
// to let the forecasting task utility range from the interval [0.1, 0.5]
// χ = 0.1 + 0.4σ(a*T_i − b)
// sigma is the sigmoid function
// a has fiduciary value of 8
// b has fiduciary value of 0.5
func forecastingUtility(forecastingPerformanceScore float64, a float64, b float64) (float64, error) {
	if math.IsNaN(forecastingPerformanceScore) || math.IsInf(forecastingPerformanceScore, 0) ||
		math.IsNaN(a) || math.IsInf(a, 0) ||
		math.IsNaN(b) || math.IsInf(b, 0) {
		return 0, emissions.ErrForecastingUtilityInvalidInput
	}
	ret, err := sigmoid(a*forecastingPerformanceScore - b)
	if err != nil {
		return 0, err
	}
	ret = 0.1 + 0.4*ret
	if math.IsInf(ret, 0) {
		return 0, emissions.ErrForecastingUtilityIsInfinity
	}
	if math.IsNaN(ret) {
		return 0, emissions.ErrForecastingUtilityIsNaN
	}
	return ret, nil
}

// renormalize with a factor γ to ensure that the
// total reward allocated to workers (Ui + Vi)
// remains constant (otherwise, this would go at the expense of reputers)
// γ = (F_i + G_i) / ( (1 − χ)*F_i + χ*G_i)
func normalizationFactor(
	entropyInference float64,
	entropyForecasting float64,
	forecastingUtility float64,
) (float64, error) {
	if math.IsNaN(entropyInference) || math.IsInf(entropyInference, 0) ||
		math.IsNaN(entropyForecasting) || math.IsInf(entropyForecasting, 0) ||
		math.IsNaN(forecastingUtility) || math.IsInf(forecastingUtility, 0) {
		return 0, errors.Wrapf(
			emissions.ErrNormalizationFactorInvalidInput,
			"entropyInference: %f, entropyForecasting: %f, forecastingUtility: %f",
			entropyInference,
			entropyForecasting,
			forecastingUtility,
		)
	}
	numerator := entropyInference + entropyForecasting
	denominator := (1-forecastingUtility)*entropyInference + forecastingUtility*entropyForecasting
	ret := numerator / denominator
	if math.IsInf(ret, 0) {
		return 0, errors.Wrapf(
			emissions.ErrNormalizationFactorIsInfinity,
			"numerator: %f, denominator: %f entropyInference: %f, entropyForecasting: %f, forecastingUtility: %f",
			numerator,
			denominator,
			entropyInference,
			entropyForecasting,
			forecastingUtility,
		)
	}
	if math.IsNaN(ret) {
		return 0, errors.Wrapf(
			emissions.ErrNormalizationFactorIsNaN,
			"numerator: %f, denominator: %f entropyInference: %f, entropyForecasting: %f, forecastingUtility: %f",
			numerator,
			denominator,
			entropyInference,
			entropyForecasting,
			forecastingUtility,
		)
	}

	return ret, nil
}