// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/distribution/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/distribution/tags
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/distribution/types
package distribution

import (
	"github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

const (
	DefaultParamspace                = keeper.DefaultParamspace
	DefaultCodespace                 = types.DefaultCodespace
	CodeInvalidInput                 = types.CodeInvalidInput
	CodeNoDistributionInfo           = types.CodeNoDistributionInfo
	CodeNoValidatorCommission        = types.CodeNoValidatorCommission
	CodeSetWithdrawAddrDisabled      = types.CodeSetWithdrawAddrDisabled
	FeeCollectorName                 = types.FeeCollectorName
	ModuleName                       = types.ModuleName
	StoreKey                         = types.StoreKey
	TStoreKey                        = types.TStoreKey
	RouterKey                        = types.RouterKey
	QuerierRoute                     = types.QuerierRoute
	ProposalTypeCommunityPoolSpend   = types.ProposalTypeCommunityPoolSpend
	QueryParams                      = types.QueryParams
	QueryValidatorOutstandingRewards = types.QueryValidatorOutstandingRewards
	QueryValidatorCommission         = types.QueryValidatorCommission
	QueryValidatorSlashes            = types.QueryValidatorSlashes
	QueryDelegationRewards           = types.QueryDelegationRewards
	QueryDelegatorTotalRewards       = types.QueryDelegatorTotalRewards
	QueryDelegatorValidators         = types.QueryDelegatorValidators
	QueryWithdrawAddr                = types.QueryWithdrawAddr
	QueryCommunityPool               = types.QueryCommunityPool
	ParamCommunityTax                = types.ParamCommunityTax
	ParamBaseProposerReward          = types.ParamBaseProposerReward
	ParamBonusProposerReward         = types.ParamBonusProposerReward
	ParamWithdrawAddrEnabled         = types.ParamWithdrawAddrEnabled
)

var (
	// functions aliases
	RegisterInvariants                         = keeper.RegisterInvariants
	AllInvariants                              = keeper.AllInvariants
	NonNegativeOutstandingInvariant            = keeper.NonNegativeOutstandingInvariant
	CanWithdrawInvariant                       = keeper.CanWithdrawInvariant
	ReferenceCountInvariant                    = keeper.ReferenceCountInvariant
	ModuleAccountInvariant                     = keeper.ModuleAccountInvariant
	NewKeeper                                  = keeper.NewKeeper
	GetValidatorOutstandingRewardsAddress      = keeper.GetValidatorOutstandingRewardsAddress
	GetDelegatorWithdrawInfoAddress            = keeper.GetDelegatorWithdrawInfoAddress
	GetDelegatorStartingInfoAddresses          = keeper.GetDelegatorStartingInfoAddresses
	GetValidatorHistoricalRewardsAddressPeriod = keeper.GetValidatorHistoricalRewardsAddressPeriod
	GetValidatorCurrentRewardsAddress          = keeper.GetValidatorCurrentRewardsAddress
	GetValidatorAccumulatedCommissionAddress   = keeper.GetValidatorAccumulatedCommissionAddress
	GetValidatorSlashEventAddressHeight        = keeper.GetValidatorSlashEventAddressHeight
	GetValidatorOutstandingRewardsKey          = keeper.GetValidatorOutstandingRewardsKey
	GetDelegatorWithdrawAddrKey                = keeper.GetDelegatorWithdrawAddrKey
	GetDelegatorStartingInfoKey                = keeper.GetDelegatorStartingInfoKey
	GetValidatorHistoricalRewardsPrefix        = keeper.GetValidatorHistoricalRewardsPrefix
	GetValidatorHistoricalRewardsKey           = keeper.GetValidatorHistoricalRewardsKey
	GetValidatorCurrentRewardsKey              = keeper.GetValidatorCurrentRewardsKey
	GetValidatorAccumulatedCommissionKey       = keeper.GetValidatorAccumulatedCommissionKey
	GetValidatorSlashEventPrefix               = keeper.GetValidatorSlashEventPrefix
	GetValidatorSlashEventKey                  = keeper.GetValidatorSlashEventKey
	ParamKeyTable                              = keeper.ParamKeyTable
	HandleCommunityPoolSpendProposal           = keeper.HandleCommunityPoolSpendProposal
	NewQuerier                                 = keeper.NewQuerier
	MakeTestCodec                              = keeper.MakeTestCodec
	CreateTestInputDefault                     = keeper.CreateTestInputDefault
	CreateTestInputAdvanced                    = keeper.CreateTestInputAdvanced
	RegisterCodec                              = types.RegisterCodec
	NewDelegatorStartingInfo                   = types.NewDelegatorStartingInfo
	ErrNilDelegatorAddr                        = types.ErrNilDelegatorAddr
	ErrNilWithdrawAddr                         = types.ErrNilWithdrawAddr
	ErrNilValidatorAddr                        = types.ErrNilValidatorAddr
	ErrNoDelegationDistInfo                    = types.ErrNoDelegationDistInfo
	ErrNoValidatorDistInfo                     = types.ErrNoValidatorDistInfo
	ErrNoValidatorCommission                   = types.ErrNoValidatorCommission
	ErrSetWithdrawAddrDisabled                 = types.ErrSetWithdrawAddrDisabled
	ErrBadDistribution                         = types.ErrBadDistribution
	ErrInvalidProposalAmount                   = types.ErrInvalidProposalAmount
	ErrEmptyProposalRecipient                  = types.ErrEmptyProposalRecipient
	InitialFeePool                             = types.InitialFeePool
	NewGenesisState                            = types.NewGenesisState
	DefaultGenesisState                        = types.DefaultGenesisState
	ValidateGenesis                            = types.ValidateGenesis
	NewMsgSetWithdrawAddress                   = types.NewMsgSetWithdrawAddress
	NewMsgWithdrawDelegatorReward              = types.NewMsgWithdrawDelegatorReward
	NewMsgWithdrawValidatorCommission          = types.NewMsgWithdrawValidatorCommission
	NewCommunityPoolSpendProposal              = types.NewCommunityPoolSpendProposal
	NewQueryValidatorOutstandingRewardsParams  = types.NewQueryValidatorOutstandingRewardsParams
	NewQueryValidatorCommissionParams          = types.NewQueryValidatorCommissionParams
	NewQueryValidatorSlashesParams             = types.NewQueryValidatorSlashesParams
	NewQueryDelegationRewardsParams            = types.NewQueryDelegationRewardsParams
	NewQueryDelegatorParams                    = types.NewQueryDelegatorParams
	NewQueryDelegatorWithdrawAddrParams        = types.NewQueryDelegatorWithdrawAddrParams
	NewQueryDelegatorTotalRewardsResponse      = types.NewQueryDelegatorTotalRewardsResponse
	NewDelegationDelegatorReward               = types.NewDelegationDelegatorReward
	NewValidatorHistoricalRewards              = types.NewValidatorHistoricalRewards
	NewValidatorCurrentRewards                 = types.NewValidatorCurrentRewards
	InitialValidatorAccumulatedCommission      = types.InitialValidatorAccumulatedCommission
	NewValidatorSlashEvent                     = types.NewValidatorSlashEvent

	// variable aliases
	FeePoolKey                           = keeper.FeePoolKey
	ProposerKey                          = keeper.ProposerKey
	ValidatorOutstandingRewardsPrefix    = keeper.ValidatorOutstandingRewardsPrefix
	DelegatorWithdrawAddrPrefix          = keeper.DelegatorWithdrawAddrPrefix
	DelegatorStartingInfoPrefix          = keeper.DelegatorStartingInfoPrefix
	ValidatorHistoricalRewardsPrefix     = keeper.ValidatorHistoricalRewardsPrefix
	ValidatorCurrentRewardsPrefix        = keeper.ValidatorCurrentRewardsPrefix
	ValidatorAccumulatedCommissionPrefix = keeper.ValidatorAccumulatedCommissionPrefix
	ValidatorSlashEventPrefix            = keeper.ValidatorSlashEventPrefix
	ParamStoreKeyCommunityTax            = keeper.ParamStoreKeyCommunityTax
	ParamStoreKeyBaseProposerReward      = keeper.ParamStoreKeyBaseProposerReward
	ParamStoreKeyBonusProposerReward     = keeper.ParamStoreKeyBonusProposerReward
	ParamStoreKeyWithdrawAddrEnabled     = keeper.ParamStoreKeyWithdrawAddrEnabled
	TestAddrs                            = keeper.TestAddrs
	EventTypeRewards                     = types.EventTypeRewards
	EventTypeCommission                  = types.EventTypeCommission
	AttributeValueCategory               = types.AttributeValueCategory
	AttributeKeyValidator                = types.AttributeKeyValidator
	ModuleCdc                            = types.ModuleCdc
)

type (
	Hooks                                  = keeper.Hooks
	Keeper                                 = keeper.Keeper
	DelegatorStartingInfo                  = types.DelegatorStartingInfo
	CodeType                               = types.CodeType
	FeePool                                = types.FeePool
	DelegatorWithdrawInfo                  = types.DelegatorWithdrawInfo
	ValidatorOutstandingRewardsRecord      = types.ValidatorOutstandingRewardsRecord
	ValidatorAccumulatedCommissionRecord   = types.ValidatorAccumulatedCommissionRecord
	ValidatorHistoricalRewardsRecord       = types.ValidatorHistoricalRewardsRecord
	ValidatorCurrentRewardsRecord          = types.ValidatorCurrentRewardsRecord
	DelegatorStartingInfoRecord            = types.DelegatorStartingInfoRecord
	ValidatorSlashEventRecord              = types.ValidatorSlashEventRecord
	GenesisState                           = types.GenesisState
	MsgSetWithdrawAddress                  = types.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward             = types.MsgWithdrawDelegatorReward
	MsgWithdrawValidatorCommission         = types.MsgWithdrawValidatorCommission
	CommunityPoolSpendProposal             = types.CommunityPoolSpendProposal
	QueryValidatorOutstandingRewardsParams = types.QueryValidatorOutstandingRewardsParams
	QueryValidatorCommissionParams         = types.QueryValidatorCommissionParams
	QueryValidatorSlashesParams            = types.QueryValidatorSlashesParams
	QueryDelegationRewardsParams           = types.QueryDelegationRewardsParams
	QueryDelegatorParams                   = types.QueryDelegatorParams
	QueryDelegatorWithdrawAddrParams       = types.QueryDelegatorWithdrawAddrParams
	QueryDelegatorTotalRewardsResponse     = types.QueryDelegatorTotalRewardsResponse
	DelegationDelegatorReward              = types.DelegationDelegatorReward
	ValidatorHistoricalRewards             = types.ValidatorHistoricalRewards
	ValidatorCurrentRewards                = types.ValidatorCurrentRewards
	ValidatorAccumulatedCommission         = types.ValidatorAccumulatedCommission
	ValidatorSlashEvent                    = types.ValidatorSlashEvent
	ValidatorSlashEvents                   = types.ValidatorSlashEvents
	ValidatorOutstandingRewards            = types.ValidatorOutstandingRewards
)
