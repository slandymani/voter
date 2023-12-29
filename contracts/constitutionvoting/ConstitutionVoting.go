// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package constitutionvoting

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BaseVotingWeightInfo is an auto generated low-level Go binding around an user-defined struct.
type BaseVotingWeightInfo struct {
	OwnWeight        *big.Int
	VotingAgent      common.Address
	DelegationStatus uint8
	LockedUntil      *big.Int
}

// IParametersVotingParameterInfo is an auto generated low-level Go binding around an user-defined struct.
type IParametersVotingParameterInfo struct {
	ParamKey     string
	ParamType    uint8
	AddrValue    common.Address
	BoolValue    bool
	Bytes32Value [32]byte
	StrValue     string
	UintValue    *big.Int
}

// IQthVotingVotingWeightInfo is an auto generated low-level Go binding around an user-defined struct.
type IQthVotingVotingWeightInfo struct {
	HasAlreadyVoted bool
	CanVote         bool
	Base            BaseVotingWeightInfo
}

// IVotingBaseProposal is an auto generated low-level Go binding around an user-defined struct.
type IVotingBaseProposal struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}

// IVotingVotingCounters is an auto generated low-level Go binding around an user-defined struct.
type IVotingVotingCounters struct {
	WeightFor     *big.Int
	WeightAgainst *big.Int
	VetosCount    *big.Int
}

// IVotingVotingParams is an auto generated low-level Go binding around an user-defined struct.
type IVotingVotingParams struct {
	VotingStartTime    *big.Int
	VotingEndTime      *big.Int
	VetoEndTime        *big.Int
	ProposalExecutionP *big.Int
	RequiredQuorum     *big.Int
	RequiredMajority   *big.Int
	RequiredSMajority  *big.Int
	RequiredSQuorum    *big.Int
}

// IVotingVotingStats is an auto generated low-level Go binding around an user-defined struct.
type IVotingVotingStats struct {
	RequiredQuorum        *big.Int
	CurrentQuorum         *big.Int
	RequiredMajority      *big.Int
	CurrentMajority       *big.Int
	CurrentVetoPercentage *big.Int
}

// ConstitutionvotingMetaData contains all meta data concerning the Constitutionvoting contract.
var ConstitutionvotingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"ProposalCreated\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_id\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"tuple\",\"name\":\"_proposal\",\"internalType\":\"structIVoting.BaseProposal\",\"indexed\":false,\"components\":[{\"type\":\"string\",\"name\":\"remark\",\"internalType\":\"string\"},{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIVoting.VotingParams\",\"components\":[{\"type\":\"uint256\",\"name\":\"votingStartTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"votingEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"vetoEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"proposalExecutionP\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredQuorum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredMajority\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredSMajority\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredSQuorum\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"counters\",\"internalType\":\"structIVoting.VotingCounters\",\"components\":[{\"type\":\"uint256\",\"name\":\"weightFor\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"weightAgainst\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"vetosCount\",\"internalType\":\"uint256\"}]},{\"type\":\"bool\",\"name\":\"executed\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalExecuted\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"_constitutionHash\",\"internalType\":\"bytes32\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumReached\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserVoted\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint8\",\"name\":\"_votingOption\",\"internalType\":\"enumIVoting.VotingOption\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VetoOccurred\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"constitutionHash\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"createAddrProposal\",\"inputs\":[{\"type\":\"string\",\"name\":\"_remark\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"_proposalType\",\"internalType\":\"enumConstitutionVoting.Classification\"},{\"type\":\"bytes32\",\"name\":\"_hashValue\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"_paramKey\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"_paramValue\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"createBoolProposal\",\"inputs\":[{\"type\":\"string\",\"name\":\"_remark\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"_proposalType\",\"internalType\":\"enumConstitutionVoting.Classification\"},{\"type\":\"bytes32\",\"name\":\"_hashValue\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"_paramKey\",\"internalType\":\"string\"},{\"type\":\"bool\",\"name\":\"_paramValue\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"createBytesProposal\",\"inputs\":[{\"type\":\"string\",\"name\":\"_remark\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"_proposalType\",\"internalType\":\"enumConstitutionVoting.Classification\"},{\"type\":\"bytes32\",\"name\":\"_hashValue\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"_paramKey\",\"internalType\":\"string\"},{\"type\":\"bytes32\",\"name\":\"_paramValue\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"createProposal\",\"inputs\":[{\"type\":\"string\",\"name\":\"_remark\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"_proposalType\",\"internalType\":\"enumConstitutionVoting.Classification\"},{\"type\":\"bytes32\",\"name\":\"_hashValue\",\"internalType\":\"bytes32\"},{\"type\":\"tuple[]\",\"name\":\"_parametersArr\",\"internalType\":\"structIParametersVoting.ParameterInfo[]\",\"components\":[{\"type\":\"string\",\"name\":\"paramKey\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"paramType\",\"internalType\":\"enumIParametersVoting.ParameterType\"},{\"type\":\"address\",\"name\":\"addrValue\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"boolValue\",\"internalType\":\"bool\"},{\"type\":\"bytes32\",\"name\":\"bytes32Value\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"strValue\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"uintValue\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"createStrProposal\",\"inputs\":[{\"type\":\"string\",\"name\":\"_remark\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"_proposalType\",\"internalType\":\"enumConstitutionVoting.Classification\"},{\"type\":\"bytes32\",\"name\":\"_hashValue\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"_paramKey\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_paramValue\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"createUintProposal\",\"inputs\":[{\"type\":\"string\",\"name\":\"_remark\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"_proposalType\",\"internalType\":\"enumConstitutionVoting.Classification\"},{\"type\":\"bytes32\",\"name\":\"_hashValue\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"_paramKey\",\"internalType\":\"string\"},{\"type\":\"uint32\",\"name\":\"_paramValue\",\"internalType\":\"uint32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"execute\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"getConstitutionHash\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"\",\"internalType\":\"structIParametersVoting.ParameterInfo[]\",\"components\":[{\"type\":\"string\",\"name\":\"paramKey\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"paramType\",\"internalType\":\"enumIParametersVoting.ParameterType\"},{\"type\":\"address\",\"name\":\"addrValue\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"boolValue\",\"internalType\":\"bool\"},{\"type\":\"bytes32\",\"name\":\"bytes32Value\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"strValue\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"uintValue\",\"internalType\":\"uint256\"}]}],\"name\":\"getParametersArr\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structIVoting.BaseProposal\",\"components\":[{\"type\":\"string\",\"name\":\"remark\",\"internalType\":\"string\"},{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIVoting.VotingParams\",\"components\":[{\"type\":\"uint256\",\"name\":\"votingStartTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"votingEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"vetoEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"proposalExecutionP\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredQuorum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredMajority\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredSMajority\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredSQuorum\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"counters\",\"internalType\":\"structIVoting.VotingCounters\",\"components\":[{\"type\":\"uint256\",\"name\":\"weightFor\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"weightAgainst\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"vetosCount\",\"internalType\":\"uint256\"}]},{\"type\":\"bool\",\"name\":\"executed\",\"internalType\":\"bool\"}]}],\"name\":\"getProposal\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structIVoting.VotingStats\",\"components\":[{\"type\":\"uint256\",\"name\":\"requiredQuorum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"currentQuorum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredMajority\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"currentMajority\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"currentVetoPercentage\",\"internalType\":\"uint256\"}]}],\"name\":\"getProposalStats\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"enumIVoting.ProposalStatus\"}],\"name\":\"getStatus\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getVetosNumber\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getVetosPercentage\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structIQthVoting.VotingWeightInfo\",\"components\":[{\"type\":\"bool\",\"name\":\"hasAlreadyVoted\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"canVote\",\"internalType\":\"bool\"},{\"type\":\"tuple\",\"name\":\"base\",\"internalType\":\"structBaseVotingWeightInfo\",\"components\":[{\"type\":\"uint256\",\"name\":\"ownWeight\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"votingAgent\",\"internalType\":\"address\"},{\"type\":\"uint8\",\"name\":\"delegationStatus\",\"internalType\":\"enumDelegationStatus\"},{\"type\":\"uint256\",\"name\":\"lockedUntil\",\"internalType\":\"uint256\"}]}]}],\"name\":\"getVotingWeightInfo\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"hasRootVetoed\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"hasUserVoted\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_constitutionHash\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"_registry\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"proposalCounter\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"base\",\"internalType\":\"structIVoting.BaseProposal\",\"components\":[{\"type\":\"string\",\"name\":\"remark\",\"internalType\":\"string\"},{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIVoting.VotingParams\",\"components\":[{\"type\":\"uint256\",\"name\":\"votingStartTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"votingEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"vetoEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"proposalExecutionP\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredQuorum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredMajority\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredSMajority\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"requiredSQuorum\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"counters\",\"internalType\":\"structIVoting.VotingCounters\",\"components\":[{\"type\":\"uint256\",\"name\":\"weightFor\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"weightAgainst\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"vetosCount\",\"internalType\":\"uint256\"}]},{\"type\":\"bool\",\"name\":\"executed\",\"internalType\":\"bool\"}]},{\"type\":\"uint256\",\"name\":\"parametersSize\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"classification\",\"internalType\":\"enumConstitutionVoting.Classification\"},{\"type\":\"bytes32\",\"name\":\"newConstitutionHash\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"currentConstitutionHash\",\"internalType\":\"bytes32\"}],\"name\":\"proposals\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"veto\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"voteAgainst\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"voteFor\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"}]}]",
}

// ConstitutionvotingABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstitutionvotingMetaData.ABI instead.
var ConstitutionvotingABI = ConstitutionvotingMetaData.ABI

// Constitutionvoting is an auto generated Go binding around an Ethereum contract.
type Constitutionvoting struct {
	ConstitutionvotingCaller     // Read-only binding to the contract
	ConstitutionvotingTransactor // Write-only binding to the contract
	ConstitutionvotingFilterer   // Log filterer for contract events
}

// ConstitutionvotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstitutionvotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstitutionvotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstitutionvotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstitutionvotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstitutionvotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstitutionvotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstitutionvotingSession struct {
	Contract     *Constitutionvoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ConstitutionvotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstitutionvotingCallerSession struct {
	Contract *ConstitutionvotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ConstitutionvotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstitutionvotingTransactorSession struct {
	Contract     *ConstitutionvotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ConstitutionvotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstitutionvotingRaw struct {
	Contract *Constitutionvoting // Generic contract binding to access the raw methods on
}

// ConstitutionvotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstitutionvotingCallerRaw struct {
	Contract *ConstitutionvotingCaller // Generic read-only contract binding to access the raw methods on
}

// ConstitutionvotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstitutionvotingTransactorRaw struct {
	Contract *ConstitutionvotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstitutionvoting creates a new instance of Constitutionvoting, bound to a specific deployed contract.
func NewConstitutionvoting(address common.Address, backend bind.ContractBackend) (*Constitutionvoting, error) {
	contract, err := bindConstitutionvoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constitutionvoting{ConstitutionvotingCaller: ConstitutionvotingCaller{contract: contract}, ConstitutionvotingTransactor: ConstitutionvotingTransactor{contract: contract}, ConstitutionvotingFilterer: ConstitutionvotingFilterer{contract: contract}}, nil
}

// NewConstitutionvotingCaller creates a new read-only instance of Constitutionvoting, bound to a specific deployed contract.
func NewConstitutionvotingCaller(address common.Address, caller bind.ContractCaller) (*ConstitutionvotingCaller, error) {
	contract, err := bindConstitutionvoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstitutionvotingCaller{contract: contract}, nil
}

// NewConstitutionvotingTransactor creates a new write-only instance of Constitutionvoting, bound to a specific deployed contract.
func NewConstitutionvotingTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstitutionvotingTransactor, error) {
	contract, err := bindConstitutionvoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstitutionvotingTransactor{contract: contract}, nil
}

// NewConstitutionvotingFilterer creates a new log filterer instance of Constitutionvoting, bound to a specific deployed contract.
func NewConstitutionvotingFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstitutionvotingFilterer, error) {
	contract, err := bindConstitutionvoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstitutionvotingFilterer{contract: contract}, nil
}

// bindConstitutionvoting binds a generic wrapper to an already deployed contract.
func bindConstitutionvoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConstitutionvotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constitutionvoting *ConstitutionvotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constitutionvoting.Contract.ConstitutionvotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constitutionvoting *ConstitutionvotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.ConstitutionvotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constitutionvoting *ConstitutionvotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.ConstitutionvotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constitutionvoting *ConstitutionvotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constitutionvoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constitutionvoting *ConstitutionvotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constitutionvoting *ConstitutionvotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.contract.Transact(opts, method, params...)
}

// ConstitutionHash is a free data retrieval call binding the contract method 0xc7d93fd4.
//
// Solidity: function constitutionHash() view returns(bytes32)
func (_Constitutionvoting *ConstitutionvotingCaller) ConstitutionHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "constitutionHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConstitutionHash is a free data retrieval call binding the contract method 0xc7d93fd4.
//
// Solidity: function constitutionHash() view returns(bytes32)
func (_Constitutionvoting *ConstitutionvotingSession) ConstitutionHash() ([32]byte, error) {
	return _Constitutionvoting.Contract.ConstitutionHash(&_Constitutionvoting.CallOpts)
}

// ConstitutionHash is a free data retrieval call binding the contract method 0xc7d93fd4.
//
// Solidity: function constitutionHash() view returns(bytes32)
func (_Constitutionvoting *ConstitutionvotingCallerSession) ConstitutionHash() ([32]byte, error) {
	return _Constitutionvoting.Contract.ConstitutionHash(&_Constitutionvoting.CallOpts)
}

// GetConstitutionHash is a free data retrieval call binding the contract method 0xd29b3992.
//
// Solidity: function getConstitutionHash(uint256 _id) view returns(bytes32)
func (_Constitutionvoting *ConstitutionvotingCaller) GetConstitutionHash(opts *bind.CallOpts, _id *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "getConstitutionHash", _id)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConstitutionHash is a free data retrieval call binding the contract method 0xd29b3992.
//
// Solidity: function getConstitutionHash(uint256 _id) view returns(bytes32)
func (_Constitutionvoting *ConstitutionvotingSession) GetConstitutionHash(_id *big.Int) ([32]byte, error) {
	return _Constitutionvoting.Contract.GetConstitutionHash(&_Constitutionvoting.CallOpts, _id)
}

// GetConstitutionHash is a free data retrieval call binding the contract method 0xd29b3992.
//
// Solidity: function getConstitutionHash(uint256 _id) view returns(bytes32)
func (_Constitutionvoting *ConstitutionvotingCallerSession) GetConstitutionHash(_id *big.Int) ([32]byte, error) {
	return _Constitutionvoting.Contract.GetConstitutionHash(&_Constitutionvoting.CallOpts, _id)
}

// GetParametersArr is a free data retrieval call binding the contract method 0xefde1c43.
//
// Solidity: function getParametersArr(uint256 _proposalId) view returns((string,uint8,address,bool,bytes32,string,uint256)[])
func (_Constitutionvoting *ConstitutionvotingCaller) GetParametersArr(opts *bind.CallOpts, _proposalId *big.Int) ([]IParametersVotingParameterInfo, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "getParametersArr", _proposalId)

	if err != nil {
		return *new([]IParametersVotingParameterInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IParametersVotingParameterInfo)).(*[]IParametersVotingParameterInfo)

	return out0, err

}

// GetParametersArr is a free data retrieval call binding the contract method 0xefde1c43.
//
// Solidity: function getParametersArr(uint256 _proposalId) view returns((string,uint8,address,bool,bytes32,string,uint256)[])
func (_Constitutionvoting *ConstitutionvotingSession) GetParametersArr(_proposalId *big.Int) ([]IParametersVotingParameterInfo, error) {
	return _Constitutionvoting.Contract.GetParametersArr(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetParametersArr is a free data retrieval call binding the contract method 0xefde1c43.
//
// Solidity: function getParametersArr(uint256 _proposalId) view returns((string,uint8,address,bool,bytes32,string,uint256)[])
func (_Constitutionvoting *ConstitutionvotingCallerSession) GetParametersArr(_proposalId *big.Int) ([]IParametersVotingParameterInfo, error) {
	return _Constitutionvoting.Contract.GetParametersArr(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_Constitutionvoting *ConstitutionvotingCaller) GetProposal(opts *bind.CallOpts, _proposalId *big.Int) (IVotingBaseProposal, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "getProposal", _proposalId)

	if err != nil {
		return *new(IVotingBaseProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingBaseProposal)).(*IVotingBaseProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_Constitutionvoting *ConstitutionvotingSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _Constitutionvoting.Contract.GetProposal(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_Constitutionvoting *ConstitutionvotingCallerSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _Constitutionvoting.Contract.GetProposal(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Constitutionvoting *ConstitutionvotingCaller) GetProposalStats(opts *bind.CallOpts, _proposalId *big.Int) (IVotingVotingStats, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "getProposalStats", _proposalId)

	if err != nil {
		return *new(IVotingVotingStats), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingVotingStats)).(*IVotingVotingStats)

	return out0, err

}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Constitutionvoting *ConstitutionvotingSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _Constitutionvoting.Contract.GetProposalStats(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Constitutionvoting *ConstitutionvotingCallerSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _Constitutionvoting.Contract.GetProposalStats(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_Constitutionvoting *ConstitutionvotingCaller) GetStatus(opts *bind.CallOpts, _id *big.Int) (uint8, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "getStatus", _id)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_Constitutionvoting *ConstitutionvotingSession) GetStatus(_id *big.Int) (uint8, error) {
	return _Constitutionvoting.Contract.GetStatus(&_Constitutionvoting.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_Constitutionvoting *ConstitutionvotingCallerSession) GetStatus(_id *big.Int) (uint8, error) {
	return _Constitutionvoting.Contract.GetStatus(&_Constitutionvoting.CallOpts, _id)
}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingCaller) GetVetosNumber(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "getVetosNumber", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) GetVetosNumber(_proposalId *big.Int) (*big.Int, error) {
	return _Constitutionvoting.Contract.GetVetosNumber(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingCallerSession) GetVetosNumber(_proposalId *big.Int) (*big.Int, error) {
	return _Constitutionvoting.Contract.GetVetosNumber(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingCaller) GetVetosPercentage(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "getVetosPercentage", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) GetVetosPercentage(_proposalId *big.Int) (*big.Int, error) {
	return _Constitutionvoting.Contract.GetVetosPercentage(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingCallerSession) GetVetosPercentage(_proposalId *big.Int) (*big.Int, error) {
	return _Constitutionvoting.Contract.GetVetosPercentage(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_Constitutionvoting *ConstitutionvotingCaller) GetVotingWeightInfo(opts *bind.CallOpts, _proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "getVotingWeightInfo", _proposalId)

	if err != nil {
		return *new(IQthVotingVotingWeightInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IQthVotingVotingWeightInfo)).(*IQthVotingVotingWeightInfo)

	return out0, err

}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_Constitutionvoting *ConstitutionvotingSession) GetVotingWeightInfo(_proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	return _Constitutionvoting.Contract.GetVotingWeightInfo(&_Constitutionvoting.CallOpts, _proposalId)
}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_Constitutionvoting *ConstitutionvotingCallerSession) GetVotingWeightInfo(_proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	return _Constitutionvoting.Contract.GetVotingWeightInfo(&_Constitutionvoting.CallOpts, _proposalId)
}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_Constitutionvoting *ConstitutionvotingCaller) HasRootVetoed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "hasRootVetoed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_Constitutionvoting *ConstitutionvotingSession) HasRootVetoed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Constitutionvoting.Contract.HasRootVetoed(&_Constitutionvoting.CallOpts, arg0, arg1)
}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_Constitutionvoting *ConstitutionvotingCallerSession) HasRootVetoed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Constitutionvoting.Contract.HasRootVetoed(&_Constitutionvoting.CallOpts, arg0, arg1)
}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_Constitutionvoting *ConstitutionvotingCaller) HasUserVoted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "hasUserVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_Constitutionvoting *ConstitutionvotingSession) HasUserVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Constitutionvoting.Contract.HasUserVoted(&_Constitutionvoting.CallOpts, arg0, arg1)
}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_Constitutionvoting *ConstitutionvotingCallerSession) HasUserVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Constitutionvoting.Contract.HasUserVoted(&_Constitutionvoting.CallOpts, arg0, arg1)
}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingCaller) ProposalCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "proposalCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) ProposalCounter() (*big.Int, error) {
	return _Constitutionvoting.Contract.ProposalCounter(&_Constitutionvoting.CallOpts)
}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_Constitutionvoting *ConstitutionvotingCallerSession) ProposalCounter() (*big.Int, error) {
	return _Constitutionvoting.Contract.ProposalCounter(&_Constitutionvoting.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) base, uint256 parametersSize, uint8 classification, bytes32 newConstitutionHash, bytes32 currentConstitutionHash)
func (_Constitutionvoting *ConstitutionvotingCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Base                    IVotingBaseProposal
	ParametersSize          *big.Int
	Classification          uint8
	NewConstitutionHash     [32]byte
	CurrentConstitutionHash [32]byte
}, error) {
	var out []interface{}
	err := _Constitutionvoting.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Base                    IVotingBaseProposal
		ParametersSize          *big.Int
		Classification          uint8
		NewConstitutionHash     [32]byte
		CurrentConstitutionHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Base = *abi.ConvertType(out[0], new(IVotingBaseProposal)).(*IVotingBaseProposal)
	outstruct.ParametersSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Classification = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.NewConstitutionHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.CurrentConstitutionHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) base, uint256 parametersSize, uint8 classification, bytes32 newConstitutionHash, bytes32 currentConstitutionHash)
func (_Constitutionvoting *ConstitutionvotingSession) Proposals(arg0 *big.Int) (struct {
	Base                    IVotingBaseProposal
	ParametersSize          *big.Int
	Classification          uint8
	NewConstitutionHash     [32]byte
	CurrentConstitutionHash [32]byte
}, error) {
	return _Constitutionvoting.Contract.Proposals(&_Constitutionvoting.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) base, uint256 parametersSize, uint8 classification, bytes32 newConstitutionHash, bytes32 currentConstitutionHash)
func (_Constitutionvoting *ConstitutionvotingCallerSession) Proposals(arg0 *big.Int) (struct {
	Base                    IVotingBaseProposal
	ParametersSize          *big.Int
	Classification          uint8
	NewConstitutionHash     [32]byte
	CurrentConstitutionHash [32]byte
}, error) {
	return _Constitutionvoting.Contract.Proposals(&_Constitutionvoting.CallOpts, arg0)
}

// CreateAddrProposal is a paid mutator transaction binding the contract method 0xf0fd6cd6.
//
// Solidity: function createAddrProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, address _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactor) CreateAddrProposal(opts *bind.TransactOpts, _remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue common.Address) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "createAddrProposal", _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateAddrProposal is a paid mutator transaction binding the contract method 0xf0fd6cd6.
//
// Solidity: function createAddrProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, address _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) CreateAddrProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue common.Address) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateAddrProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateAddrProposal is a paid mutator transaction binding the contract method 0xf0fd6cd6.
//
// Solidity: function createAddrProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, address _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactorSession) CreateAddrProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue common.Address) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateAddrProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateBoolProposal is a paid mutator transaction binding the contract method 0xf5709743.
//
// Solidity: function createBoolProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, bool _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactor) CreateBoolProposal(opts *bind.TransactOpts, _remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue bool) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "createBoolProposal", _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateBoolProposal is a paid mutator transaction binding the contract method 0xf5709743.
//
// Solidity: function createBoolProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, bool _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) CreateBoolProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue bool) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateBoolProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateBoolProposal is a paid mutator transaction binding the contract method 0xf5709743.
//
// Solidity: function createBoolProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, bool _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactorSession) CreateBoolProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue bool) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateBoolProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateBytesProposal is a paid mutator transaction binding the contract method 0x9eb5c74c.
//
// Solidity: function createBytesProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, bytes32 _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactor) CreateBytesProposal(opts *bind.TransactOpts, _remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue [32]byte) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "createBytesProposal", _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateBytesProposal is a paid mutator transaction binding the contract method 0x9eb5c74c.
//
// Solidity: function createBytesProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, bytes32 _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) CreateBytesProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue [32]byte) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateBytesProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateBytesProposal is a paid mutator transaction binding the contract method 0x9eb5c74c.
//
// Solidity: function createBytesProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, bytes32 _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactorSession) CreateBytesProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue [32]byte) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateBytesProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x0abe4c2e.
//
// Solidity: function createProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, (string,uint8,address,bool,bytes32,string,uint256)[] _parametersArr) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactor) CreateProposal(opts *bind.TransactOpts, _remark string, _proposalType uint8, _hashValue [32]byte, _parametersArr []IParametersVotingParameterInfo) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "createProposal", _remark, _proposalType, _hashValue, _parametersArr)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x0abe4c2e.
//
// Solidity: function createProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, (string,uint8,address,bool,bytes32,string,uint256)[] _parametersArr) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) CreateProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _parametersArr []IParametersVotingParameterInfo) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _parametersArr)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x0abe4c2e.
//
// Solidity: function createProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, (string,uint8,address,bool,bytes32,string,uint256)[] _parametersArr) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactorSession) CreateProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _parametersArr []IParametersVotingParameterInfo) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _parametersArr)
}

// CreateStrProposal is a paid mutator transaction binding the contract method 0x32a4922e.
//
// Solidity: function createStrProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, string _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactor) CreateStrProposal(opts *bind.TransactOpts, _remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue string) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "createStrProposal", _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateStrProposal is a paid mutator transaction binding the contract method 0x32a4922e.
//
// Solidity: function createStrProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, string _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) CreateStrProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue string) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateStrProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateStrProposal is a paid mutator transaction binding the contract method 0x32a4922e.
//
// Solidity: function createStrProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, string _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactorSession) CreateStrProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue string) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateStrProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateUintProposal is a paid mutator transaction binding the contract method 0x51af0c70.
//
// Solidity: function createUintProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, uint32 _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactor) CreateUintProposal(opts *bind.TransactOpts, _remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue uint32) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "createUintProposal", _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateUintProposal is a paid mutator transaction binding the contract method 0x51af0c70.
//
// Solidity: function createUintProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, uint32 _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingSession) CreateUintProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue uint32) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateUintProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// CreateUintProposal is a paid mutator transaction binding the contract method 0x51af0c70.
//
// Solidity: function createUintProposal(string _remark, uint8 _proposalType, bytes32 _hashValue, string _paramKey, uint32 _paramValue) returns(uint256)
func (_Constitutionvoting *ConstitutionvotingTransactorSession) CreateUintProposal(_remark string, _proposalType uint8, _hashValue [32]byte, _paramKey string, _paramValue uint32) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.CreateUintProposal(&_Constitutionvoting.TransactOpts, _remark, _proposalType, _hashValue, _paramKey, _paramValue)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingTransactor) Execute(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "execute", _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.Execute(&_Constitutionvoting.TransactOpts, _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingTransactorSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.Execute(&_Constitutionvoting.TransactOpts, _proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0x6910e334.
//
// Solidity: function initialize(bytes32 _constitutionHash, address _registry) returns()
func (_Constitutionvoting *ConstitutionvotingTransactor) Initialize(opts *bind.TransactOpts, _constitutionHash [32]byte, _registry common.Address) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "initialize", _constitutionHash, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x6910e334.
//
// Solidity: function initialize(bytes32 _constitutionHash, address _registry) returns()
func (_Constitutionvoting *ConstitutionvotingSession) Initialize(_constitutionHash [32]byte, _registry common.Address) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.Initialize(&_Constitutionvoting.TransactOpts, _constitutionHash, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x6910e334.
//
// Solidity: function initialize(bytes32 _constitutionHash, address _registry) returns()
func (_Constitutionvoting *ConstitutionvotingTransactorSession) Initialize(_constitutionHash [32]byte, _registry common.Address) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.Initialize(&_Constitutionvoting.TransactOpts, _constitutionHash, _registry)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _id) returns()
func (_Constitutionvoting *ConstitutionvotingTransactor) Veto(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "veto", _id)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _id) returns()
func (_Constitutionvoting *ConstitutionvotingSession) Veto(_id *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.Veto(&_Constitutionvoting.TransactOpts, _id)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _id) returns()
func (_Constitutionvoting *ConstitutionvotingTransactorSession) Veto(_id *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.Veto(&_Constitutionvoting.TransactOpts, _id)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingTransactor) VoteAgainst(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "voteAgainst", _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.VoteAgainst(&_Constitutionvoting.TransactOpts, _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingTransactorSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.VoteAgainst(&_Constitutionvoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingTransactor) VoteFor(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.contract.Transact(opts, "voteFor", _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.VoteFor(&_Constitutionvoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_Constitutionvoting *ConstitutionvotingTransactorSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _Constitutionvoting.Contract.VoteFor(&_Constitutionvoting.TransactOpts, _proposalId)
}

// ConstitutionvotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Constitutionvoting contract.
type ConstitutionvotingProposalCreatedIterator struct {
	Event *ConstitutionvotingProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstitutionvotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstitutionvotingProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstitutionvotingProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstitutionvotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstitutionvotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstitutionvotingProposalCreated represents a ProposalCreated event raised by the Constitutionvoting contract.
type ConstitutionvotingProposalCreated struct {
	Id       *big.Int
	Proposal IVotingBaseProposal
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_Constitutionvoting *ConstitutionvotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*ConstitutionvotingProposalCreatedIterator, error) {

	logs, sub, err := _Constitutionvoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &ConstitutionvotingProposalCreatedIterator{contract: _Constitutionvoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_Constitutionvoting *ConstitutionvotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *ConstitutionvotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Constitutionvoting.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstitutionvotingProposalCreated)
				if err := _Constitutionvoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_Constitutionvoting *ConstitutionvotingFilterer) ParseProposalCreated(log types.Log) (*ConstitutionvotingProposalCreated, error) {
	event := new(ConstitutionvotingProposalCreated)
	if err := _Constitutionvoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstitutionvotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Constitutionvoting contract.
type ConstitutionvotingProposalExecutedIterator struct {
	Event *ConstitutionvotingProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstitutionvotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstitutionvotingProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstitutionvotingProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstitutionvotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstitutionvotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstitutionvotingProposalExecuted represents a ProposalExecuted event raised by the Constitutionvoting contract.
type ConstitutionvotingProposalExecuted struct {
	ProposalId       *big.Int
	ConstitutionHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x56a007d3eea04bd347e571f3451382cb2a33ef5fd102b9a63846ff8d787f43cf.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId, bytes32 _constitutionHash)
func (_Constitutionvoting *ConstitutionvotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts, _proposalId []*big.Int) (*ConstitutionvotingProposalExecutedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _Constitutionvoting.contract.FilterLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &ConstitutionvotingProposalExecutedIterator{contract: _Constitutionvoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x56a007d3eea04bd347e571f3451382cb2a33ef5fd102b9a63846ff8d787f43cf.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId, bytes32 _constitutionHash)
func (_Constitutionvoting *ConstitutionvotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *ConstitutionvotingProposalExecuted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _Constitutionvoting.contract.WatchLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstitutionvotingProposalExecuted)
				if err := _Constitutionvoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0x56a007d3eea04bd347e571f3451382cb2a33ef5fd102b9a63846ff8d787f43cf.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId, bytes32 _constitutionHash)
func (_Constitutionvoting *ConstitutionvotingFilterer) ParseProposalExecuted(log types.Log) (*ConstitutionvotingProposalExecuted, error) {
	event := new(ConstitutionvotingProposalExecuted)
	if err := _Constitutionvoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstitutionvotingQuorumReachedIterator is returned from FilterQuorumReached and is used to iterate over the raw logs and unpacked data for QuorumReached events raised by the Constitutionvoting contract.
type ConstitutionvotingQuorumReachedIterator struct {
	Event *ConstitutionvotingQuorumReached // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstitutionvotingQuorumReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstitutionvotingQuorumReached)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstitutionvotingQuorumReached)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstitutionvotingQuorumReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstitutionvotingQuorumReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstitutionvotingQuorumReached represents a QuorumReached event raised by the Constitutionvoting contract.
type ConstitutionvotingQuorumReached struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterQuorumReached is a free log retrieval operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_Constitutionvoting *ConstitutionvotingFilterer) FilterQuorumReached(opts *bind.FilterOpts) (*ConstitutionvotingQuorumReachedIterator, error) {

	logs, sub, err := _Constitutionvoting.contract.FilterLogs(opts, "QuorumReached")
	if err != nil {
		return nil, err
	}
	return &ConstitutionvotingQuorumReachedIterator{contract: _Constitutionvoting.contract, event: "QuorumReached", logs: logs, sub: sub}, nil
}

// WatchQuorumReached is a free log subscription operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_Constitutionvoting *ConstitutionvotingFilterer) WatchQuorumReached(opts *bind.WatchOpts, sink chan<- *ConstitutionvotingQuorumReached) (event.Subscription, error) {

	logs, sub, err := _Constitutionvoting.contract.WatchLogs(opts, "QuorumReached")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstitutionvotingQuorumReached)
				if err := _Constitutionvoting.contract.UnpackLog(event, "QuorumReached", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQuorumReached is a log parse operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_Constitutionvoting *ConstitutionvotingFilterer) ParseQuorumReached(log types.Log) (*ConstitutionvotingQuorumReached, error) {
	event := new(ConstitutionvotingQuorumReached)
	if err := _Constitutionvoting.contract.UnpackLog(event, "QuorumReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstitutionvotingUserVotedIterator is returned from FilterUserVoted and is used to iterate over the raw logs and unpacked data for UserVoted events raised by the Constitutionvoting contract.
type ConstitutionvotingUserVotedIterator struct {
	Event *ConstitutionvotingUserVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstitutionvotingUserVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstitutionvotingUserVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstitutionvotingUserVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstitutionvotingUserVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstitutionvotingUserVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstitutionvotingUserVoted represents a UserVoted event raised by the Constitutionvoting contract.
type ConstitutionvotingUserVoted struct {
	ProposalId   *big.Int
	VotingOption uint8
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserVoted is a free log retrieval operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_Constitutionvoting *ConstitutionvotingFilterer) FilterUserVoted(opts *bind.FilterOpts, _proposalId []*big.Int) (*ConstitutionvotingUserVotedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _Constitutionvoting.contract.FilterLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &ConstitutionvotingUserVotedIterator{contract: _Constitutionvoting.contract, event: "UserVoted", logs: logs, sub: sub}, nil
}

// WatchUserVoted is a free log subscription operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_Constitutionvoting *ConstitutionvotingFilterer) WatchUserVoted(opts *bind.WatchOpts, sink chan<- *ConstitutionvotingUserVoted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _Constitutionvoting.contract.WatchLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstitutionvotingUserVoted)
				if err := _Constitutionvoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserVoted is a log parse operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_Constitutionvoting *ConstitutionvotingFilterer) ParseUserVoted(log types.Log) (*ConstitutionvotingUserVoted, error) {
	event := new(ConstitutionvotingUserVoted)
	if err := _Constitutionvoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstitutionvotingVetoOccurredIterator is returned from FilterVetoOccurred and is used to iterate over the raw logs and unpacked data for VetoOccurred events raised by the Constitutionvoting contract.
type ConstitutionvotingVetoOccurredIterator struct {
	Event *ConstitutionvotingVetoOccurred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstitutionvotingVetoOccurredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstitutionvotingVetoOccurred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstitutionvotingVetoOccurred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstitutionvotingVetoOccurredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstitutionvotingVetoOccurredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstitutionvotingVetoOccurred represents a VetoOccurred event raised by the Constitutionvoting contract.
type ConstitutionvotingVetoOccurred struct {
	Id     *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVetoOccurred is a free log retrieval operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_Constitutionvoting *ConstitutionvotingFilterer) FilterVetoOccurred(opts *bind.FilterOpts, id []*big.Int, sender []common.Address) (*ConstitutionvotingVetoOccurredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Constitutionvoting.contract.FilterLogs(opts, "VetoOccurred", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ConstitutionvotingVetoOccurredIterator{contract: _Constitutionvoting.contract, event: "VetoOccurred", logs: logs, sub: sub}, nil
}

// WatchVetoOccurred is a free log subscription operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_Constitutionvoting *ConstitutionvotingFilterer) WatchVetoOccurred(opts *bind.WatchOpts, sink chan<- *ConstitutionvotingVetoOccurred, id []*big.Int, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Constitutionvoting.contract.WatchLogs(opts, "VetoOccurred", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstitutionvotingVetoOccurred)
				if err := _Constitutionvoting.contract.UnpackLog(event, "VetoOccurred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVetoOccurred is a log parse operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_Constitutionvoting *ConstitutionvotingFilterer) ParseVetoOccurred(log types.Log) (*ConstitutionvotingVetoOccurred, error) {
	event := new(ConstitutionvotingVetoOccurred)
	if err := _Constitutionvoting.contract.UnpackLog(event, "VetoOccurred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
