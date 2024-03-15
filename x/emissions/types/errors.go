package types

import "cosmossdk.io/errors"

var (
	ErrTopicReputerStakeDoesNotExist                     = errors.Register(ModuleName, 1, "topic reputer stake does not exist")
	ErrIntegerUnderflowTopicReputerStake                 = errors.Register(ModuleName, 2, "integer underflow for topic reputer stake")
	ErrIntegerUnderflowTarget                            = errors.Register(ModuleName, 3, "integer underflow for target")
	ErrIntegerUnderflowTopicStake                        = errors.Register(ModuleName, 4, "integer underflow for topic stake")
	ErrIntegerUnderflowAllTopicStakeSum                  = errors.Register(ModuleName, 5, "integer underflow for all topic stake sum")
	ErrIntegerUnderflowTotalStake                        = errors.Register(ModuleName, 6, "integer underflow for total stake")
	ErrIterationLengthDoesNotMatch                       = errors.Register(ModuleName, 7, "iteration length does not match")
	ErrInvalidTopicId                                    = errors.Register(ModuleName, 8, "invalid topic ID")
	ErrReputerAlreadyRegisteredInTopic                   = errors.Register(ModuleName, 9, "reputer already registered in topic")
	ErrWorkerAlreadyRegisteredInTopic                    = errors.Register(ModuleName, 10, "worker already registered in topic")
	ErrAddressAlreadyRegisteredInATopic                  = errors.Register(ModuleName, 11, "address already registered in a topic")
	ErrAddressIsNotRegisteredInAnyTopic                  = errors.Register(ModuleName, 12, "address is not registered in any topic")
	ErrAddressIsNotRegisteredInThisTopic                 = errors.Register(ModuleName, 13, "address is not registered in this topic")
	ErrInsufficientStakeToRegister                       = errors.Register(ModuleName, 14, "insufficient stake to register")
	ErrLibP2PKeyRequired                                 = errors.Register(ModuleName, 15, "libp2p key required")
	ErrAddressNotRegistered                              = errors.Register(ModuleName, 16, "address not registered")
	ErrStakeTargetNotRegistered                          = errors.Register(ModuleName, 17, "stake target not registered")
	ErrTopicIdOfStakerAndTargetDoNotMatch                = errors.Register(ModuleName, 18, "topic ID of staker and target do not match")
	ErrInsufficientStakeToRemove                         = errors.Register(ModuleName, 19, "insufficient stake to remove")
	ErrNoStakeToRemove                                   = errors.Register(ModuleName, 20, "no stake to remove")
	ErrDoNotSetMapValueToZero                            = errors.Register(ModuleName, 21, "do not set map value to zero")
	ErrBlockHeightNegative                               = errors.Register(ModuleName, 22, "block height negative")
	ErrBlockHeightLessThanPrevious                       = errors.Register(ModuleName, 23, "block height less than previous")
	ErrModifyStakeBeforeBondLessThanAmountModified       = errors.Register(ModuleName, 24, "modify stake before bond less than amount modified")
	ErrModifyStakeBeforeSumGreaterThanSenderStake        = errors.Register(ModuleName, 25, "modify stake before sum greater than sender stake")
	ErrModifyStakeSumBeforeNotEqualToSumAfter            = errors.Register(ModuleName, 26, "modify stake sum before not equal to sum after")
	ErrConfirmRemoveStakeNoRemovalStarted                = errors.Register(ModuleName, 27, "confirm remove stake no removal started")
	ErrConfirmRemoveStakeTooEarly                        = errors.Register(ModuleName, 28, "confirm remove stake too early")
	ErrConfirmRemoveStakeTooLate                         = errors.Register(ModuleName, 29, "confirm remove stake too late")
	ErrScalarMultiplyNegative                            = errors.Register(ModuleName, 30, "scalar multiply negative")
	ErrDivideMapValuesByZero                             = errors.Register(ModuleName, 31, "divide map values by zero")
	ErrTopicIdListValueDecodeInvalidLength               = errors.Register(ModuleName, 32, "topic ID list value decode invalid length")
	ErrTopicIdListValueDecodeJsonInvalidLength           = errors.Register(ModuleName, 33, "topic ID list value decode JSON invalid length")
	ErrTopicIdListValueDecodeJsonInvalidFormat           = errors.Register(ModuleName, 34, "topic ID list value decode JSON invalid format")
	ErrTopicDoesNotExist                                 = errors.Register(ModuleName, 35, "topic does not exist")
	ErrCannotRemoveMoreStakeThanStakedInTopic            = errors.Register(ModuleName, 36, "cannot remove more stake than staked in topic")
	ErrInferenceRequestAlreadyInMempool                  = errors.Register(ModuleName, 37, "inference request already in mempool")
	ErrInferenceRequestBidAmountLessThanPrice            = errors.Register(ModuleName, 38, "inference request bid amount less than price")
	ErrInferenceRequestTimestampValidUntilInPast         = errors.Register(ModuleName, 39, "inference request timestamp valid until in past")
	ErrInferenceRequestTimestampValidUntilTooFarInFuture = errors.Register(ModuleName, 40, "inference request timestamp valid until too far in future")
	ErrInferenceRequestCadenceTooFast                    = errors.Register(ModuleName, 41, "inference request cadence too fast")
	ErrInferenceRequestCadenceTooSlow                    = errors.Register(ModuleName, 42, "inference request cadence too slow")
	ErrInferenceRequestWillNeverBeScheduled              = errors.Register(ModuleName, 43, "inference request will never be scheduled")
	ErrOwnerCannotBeEmpty                                = errors.Register(ModuleName, 44, "owner cannot be empty")
	ErrInsufficientStakeAfterRemoval                     = errors.Register(ModuleName, 45, "insufficient stake after removal")
	ErrInferenceRequestBidAmountTooLow                   = errors.Register(ModuleName, 46, "inference request bid amount too low")
	ErrIntegerUnderflowUnmetDemand                       = errors.Register(ModuleName, 47, "integer underflow for unmet demand")
	ErrInferenceCadenceBelowMinimum                      = errors.Register(ModuleName, 48, "inference cadence must be at least 60 seconds (1 minute)")
	ErrLossCadenceBelowMinimum                           = errors.Register(ModuleName, 49, "loss cadence must be at least 10800 seconds (3 hours)")
	ErrNotWhitelistAdmin                                 = errors.Register(ModuleName, 50, "not whitelist admin")
	ErrNotInTopicCreationWhitelist                       = errors.Register(ModuleName, 51, "not in topic creation whitelist")
	ErrNotInReputerWhitelist                             = errors.Register(ModuleName, 52, "not in reputer whitelist")
	ErrTopicNotEnoughDemand                              = errors.Register(ModuleName, 53, "topic not enough demand")
	ErrInvalidRequestId                                  = errors.Register(ModuleName, 54, "invalid request ID")
	ErrInferenceRequestNotInMempool                      = errors.Register(ModuleName, 55, "inference request not in mempool")
	ErrIntegerUnderflowStakeFromDelegator                = errors.Register(ModuleName, 56, "integer underflow for stake from delegator")
	ErrIntegerUnderflowDelegatedStakePlacement           = errors.Register(ModuleName, 57, "integer underflow for delegated stake placement")
	ErrIntegerUnderflowDelegatedStakeUponReputer         = errors.Register(ModuleName, 58, "integer underflow for delegated stake upon reputer")
)