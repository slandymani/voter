// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package qvault

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

// QVaultBalanceDetails is an auto generated low-level Go binding around an user-defined struct.
type QVaultBalanceDetails struct {
	CurrentBalance           *big.Int
	NormalizedBalance        *big.Int
	CompoundRate             *big.Int
	LastUpdateOfCompoundRate *big.Int
	InterestRate             *big.Int
}

// QVaultValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type QVaultValidatorInfo struct {
	Validator                  common.Address
	ActualStake                *big.Int
	NormalizedStake            *big.Int
	CompoundRate               *big.Int
	LatestUpdateOfCompoundRate *big.Int
	IdealStake                 *big.Int
	ClaimableReward            *big.Int
}

// TimeLockEntry is an auto generated low-level Go binding around an user-defined struct.
type TimeLockEntry struct {
	Amount       *big.Int
	ReleaseStart *big.Int
	ReleaseEnd   *big.Int
}

// VotingLockInfo is an auto generated low-level Go binding around an user-defined struct.
type VotingLockInfo struct {
	LockedAmount        *big.Int
	LockedUntil         *big.Int
	PendingUnlockAmount *big.Int
	PendingUnlockTime   *big.Int
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardClaimed\",\"inputs\":[{\"type\":\"address\",\"name\":\"_delegator\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_totalReward\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_newBalance\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTimeLock\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"tuple\",\"name\":\"_timelock\",\"internalType\":\"structTimeLockEntry\",\"indexed\":false,\"components\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseStart\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseEnd\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Purged\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"tuple\",\"name\":\"_timelock\",\"internalType\":\"structTimeLockEntry\",\"indexed\":false,\"components\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseStart\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseEnd\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDeposited\",\"inputs\":[{\"type\":\"address\",\"name\":\"_user\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_newDepositAmount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_newBalance\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserWithdrawn\",\"inputs\":[{\"type\":\"address\",\"name\":\"_user\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_withdrawnAmount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_newBalance\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"aggregatedNormalizedBalance\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"allowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"tokenOwner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"announceUnlock\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"_userAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"claimStakeDelegatorReward\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractCompoundRateKeeper\"}],\"name\":\"compoundRateKeeper\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"decimals\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"decreaseAllowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"subtractedValue\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"delegateStake\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"_delegatedTo\",\"internalType\":\"address[]\"},{\"type\":\"uint256[]\",\"name\":\"_stakes\",\"internalType\":\"uint256[]\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"deposit\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"depositFromPool\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"depositOnBehalfOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_releaseStart\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_releaseEnd\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"depositTo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structQVault.BalanceDetails\",\"components\":[{\"type\":\"uint256\",\"name\":\"currentBalance\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"normalizedBalance\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"compoundRate\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"lastUpdateOfCompoundRate\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"interestRate\",\"internalType\":\"uint256\"}]}],\"name\":\"getBalanceDetails\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"\",\"internalType\":\"structQVault.ValidatorInfo[]\",\"components\":[{\"type\":\"address\",\"name\":\"validator\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"actualStake\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"normalizedStake\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"compoundRate\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"latestUpdateOfCompoundRate\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"idealStake\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"claimableReward\",\"internalType\":\"uint256\"}]}],\"name\":\"getDelegationsList\",\"inputs\":[{\"type\":\"address\",\"name\":\"_delegator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structVotingLockInfo\",\"components\":[{\"type\":\"uint256\",\"name\":\"lockedAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"lockedUntil\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pendingUnlockAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pendingUnlockTime\",\"internalType\":\"uint256\"}]}],\"name\":\"getLockInfo\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getMinimumBalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getNormalizedBalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"_userAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"\",\"internalType\":\"structTimeLockEntry[]\",\"components\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseStart\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseEnd\",\"internalType\":\"uint256\"}]}],\"name\":\"getTimeLocks\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getUserBalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"_userAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"increaseAllowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"addedValue\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_registry\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"lock\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"purgeTimeLocks\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"receiver\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transferAndCall\",\"inputs\":[{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"receiver\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unlock\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"updateCompoundRate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"withdrawTo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}]}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// AggregatedNormalizedBalance is a free data retrieval call binding the contract method 0x6d00d28a.
//
// Solidity: function aggregatedNormalizedBalance() view returns(uint256)
func (_Contracts *ContractsCaller) AggregatedNormalizedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "aggregatedNormalizedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AggregatedNormalizedBalance is a free data retrieval call binding the contract method 0x6d00d28a.
//
// Solidity: function aggregatedNormalizedBalance() view returns(uint256)
func (_Contracts *ContractsSession) AggregatedNormalizedBalance() (*big.Int, error) {
	return _Contracts.Contract.AggregatedNormalizedBalance(&_Contracts.CallOpts)
}

// AggregatedNormalizedBalance is a free data retrieval call binding the contract method 0x6d00d28a.
//
// Solidity: function aggregatedNormalizedBalance() view returns(uint256)
func (_Contracts *ContractsCallerSession) AggregatedNormalizedBalance() (*big.Int, error) {
	return _Contracts.Contract.AggregatedNormalizedBalance(&_Contracts.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256)
func (_Contracts *ContractsCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allowance", tokenOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256)
func (_Contracts *ContractsSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256)
func (_Contracts *ContractsCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Contracts *ContractsSession) BalanceOf(_userAddress common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, _userAddress)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceOf(_userAddress common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, _userAddress)
}

// CompoundRateKeeper is a free data retrieval call binding the contract method 0x3d4d5d4d.
//
// Solidity: function compoundRateKeeper() view returns(address)
func (_Contracts *ContractsCaller) CompoundRateKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "compoundRateKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CompoundRateKeeper is a free data retrieval call binding the contract method 0x3d4d5d4d.
//
// Solidity: function compoundRateKeeper() view returns(address)
func (_Contracts *ContractsSession) CompoundRateKeeper() (common.Address, error) {
	return _Contracts.Contract.CompoundRateKeeper(&_Contracts.CallOpts)
}

// CompoundRateKeeper is a free data retrieval call binding the contract method 0x3d4d5d4d.
//
// Solidity: function compoundRateKeeper() view returns(address)
func (_Contracts *ContractsCallerSession) CompoundRateKeeper() (common.Address, error) {
	return _Contracts.Contract.CompoundRateKeeper(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Contracts *ContractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Contracts *ContractsSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Contracts *ContractsCallerSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Contracts *ContractsCaller) GetBalanceDetails(opts *bind.CallOpts) (QVaultBalanceDetails, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBalanceDetails")

	if err != nil {
		return *new(QVaultBalanceDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(QVaultBalanceDetails)).(*QVaultBalanceDetails)

	return out0, err

}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Contracts *ContractsSession) GetBalanceDetails() (QVaultBalanceDetails, error) {
	return _Contracts.Contract.GetBalanceDetails(&_Contracts.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Contracts *ContractsCallerSession) GetBalanceDetails() (QVaultBalanceDetails, error) {
	return _Contracts.Contract.GetBalanceDetails(&_Contracts.CallOpts)
}

// GetDelegationsList is a free data retrieval call binding the contract method 0x2f3e0784.
//
// Solidity: function getDelegationsList(address _delegator) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Contracts *ContractsCaller) GetDelegationsList(opts *bind.CallOpts, _delegator common.Address) ([]QVaultValidatorInfo, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getDelegationsList", _delegator)

	if err != nil {
		return *new([]QVaultValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]QVaultValidatorInfo)).(*[]QVaultValidatorInfo)

	return out0, err

}

// GetDelegationsList is a free data retrieval call binding the contract method 0x2f3e0784.
//
// Solidity: function getDelegationsList(address _delegator) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Contracts *ContractsSession) GetDelegationsList(_delegator common.Address) ([]QVaultValidatorInfo, error) {
	return _Contracts.Contract.GetDelegationsList(&_Contracts.CallOpts, _delegator)
}

// GetDelegationsList is a free data retrieval call binding the contract method 0x2f3e0784.
//
// Solidity: function getDelegationsList(address _delegator) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Contracts *ContractsCallerSession) GetDelegationsList(_delegator common.Address) ([]QVaultValidatorInfo, error) {
	return _Contracts.Contract.GetDelegationsList(&_Contracts.CallOpts, _delegator)
}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Contracts *ContractsCaller) GetLockInfo(opts *bind.CallOpts) (VotingLockInfo, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLockInfo")

	if err != nil {
		return *new(VotingLockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(VotingLockInfo)).(*VotingLockInfo)

	return out0, err

}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Contracts *ContractsSession) GetLockInfo() (VotingLockInfo, error) {
	return _Contracts.Contract.GetLockInfo(&_Contracts.CallOpts)
}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Contracts *ContractsCallerSession) GetLockInfo() (VotingLockInfo, error) {
	return _Contracts.Contract.GetLockInfo(&_Contracts.CallOpts)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Contracts *ContractsCaller) GetMinimumBalance(opts *bind.CallOpts, _account common.Address, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMinimumBalance", _account, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Contracts *ContractsSession) GetMinimumBalance(_account common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMinimumBalance(&_Contracts.CallOpts, _account, _timestamp)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMinimumBalance(_account common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMinimumBalance(&_Contracts.CallOpts, _account, _timestamp)
}

// GetNormalizedBalance is a free data retrieval call binding the contract method 0x5d247049.
//
// Solidity: function getNormalizedBalance(address _userAddress) view returns(uint256)
func (_Contracts *ContractsCaller) GetNormalizedBalance(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getNormalizedBalance", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNormalizedBalance is a free data retrieval call binding the contract method 0x5d247049.
//
// Solidity: function getNormalizedBalance(address _userAddress) view returns(uint256)
func (_Contracts *ContractsSession) GetNormalizedBalance(_userAddress common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetNormalizedBalance(&_Contracts.CallOpts, _userAddress)
}

// GetNormalizedBalance is a free data retrieval call binding the contract method 0x5d247049.
//
// Solidity: function getNormalizedBalance(address _userAddress) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetNormalizedBalance(_userAddress common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetNormalizedBalance(&_Contracts.CallOpts, _userAddress)
}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Contracts *ContractsCaller) GetTimeLocks(opts *bind.CallOpts, _account common.Address) ([]TimeLockEntry, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getTimeLocks", _account)

	if err != nil {
		return *new([]TimeLockEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]TimeLockEntry)).(*[]TimeLockEntry)

	return out0, err

}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Contracts *ContractsSession) GetTimeLocks(_account common.Address) ([]TimeLockEntry, error) {
	return _Contracts.Contract.GetTimeLocks(&_Contracts.CallOpts, _account)
}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Contracts *ContractsCallerSession) GetTimeLocks(_account common.Address) ([]TimeLockEntry, error) {
	return _Contracts.Contract.GetTimeLocks(&_Contracts.CallOpts, _account)
}

// GetUserBalance is a free data retrieval call binding the contract method 0x47734892.
//
// Solidity: function getUserBalance(address _userAddress) view returns(uint256)
func (_Contracts *ContractsCaller) GetUserBalance(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getUserBalance", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserBalance is a free data retrieval call binding the contract method 0x47734892.
//
// Solidity: function getUserBalance(address _userAddress) view returns(uint256)
func (_Contracts *ContractsSession) GetUserBalance(_userAddress common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetUserBalance(&_Contracts.CallOpts, _userAddress)
}

// GetUserBalance is a free data retrieval call binding the contract method 0x47734892.
//
// Solidity: function getUserBalance(address _userAddress) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetUserBalance(_userAddress common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetUserBalance(&_Contracts.CallOpts, _userAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Contracts *ContractsTransactor) AnnounceUnlock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "announceUnlock", _amount)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Contracts *ContractsSession) AnnounceUnlock(_amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnnounceUnlock(&_Contracts.TransactOpts, _amount)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) AnnounceUnlock(_amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnnounceUnlock(&_Contracts.TransactOpts, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, amount)
}

// ClaimStakeDelegatorReward is a paid mutator transaction binding the contract method 0xe892302d.
//
// Solidity: function claimStakeDelegatorReward() returns(bool)
func (_Contracts *ContractsTransactor) ClaimStakeDelegatorReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "claimStakeDelegatorReward")
}

// ClaimStakeDelegatorReward is a paid mutator transaction binding the contract method 0xe892302d.
//
// Solidity: function claimStakeDelegatorReward() returns(bool)
func (_Contracts *ContractsSession) ClaimStakeDelegatorReward() (*types.Transaction, error) {
	return _Contracts.Contract.ClaimStakeDelegatorReward(&_Contracts.TransactOpts)
}

// ClaimStakeDelegatorReward is a paid mutator transaction binding the contract method 0xe892302d.
//
// Solidity: function claimStakeDelegatorReward() returns(bool)
func (_Contracts *ContractsTransactorSession) ClaimStakeDelegatorReward() (*types.Transaction, error) {
	return _Contracts.Contract.ClaimStakeDelegatorReward(&_Contracts.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contracts *ContractsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contracts *ContractsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DecreaseAllowance(&_Contracts.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contracts *ContractsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DecreaseAllowance(&_Contracts.TransactOpts, spender, subtractedValue)
}

// DelegateStake is a paid mutator transaction binding the contract method 0x80563e69.
//
// Solidity: function delegateStake(address[] _delegatedTo, uint256[] _stakes) returns()
func (_Contracts *ContractsTransactor) DelegateStake(opts *bind.TransactOpts, _delegatedTo []common.Address, _stakes []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "delegateStake", _delegatedTo, _stakes)
}

// DelegateStake is a paid mutator transaction binding the contract method 0x80563e69.
//
// Solidity: function delegateStake(address[] _delegatedTo, uint256[] _stakes) returns()
func (_Contracts *ContractsSession) DelegateStake(_delegatedTo []common.Address, _stakes []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DelegateStake(&_Contracts.TransactOpts, _delegatedTo, _stakes)
}

// DelegateStake is a paid mutator transaction binding the contract method 0x80563e69.
//
// Solidity: function delegateStake(address[] _delegatedTo, uint256[] _stakes) returns()
func (_Contracts *ContractsTransactorSession) DelegateStake(_delegatedTo []common.Address, _stakes []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DelegateStake(&_Contracts.TransactOpts, _delegatedTo, _stakes)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(bool)
func (_Contracts *ContractsTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(bool)
func (_Contracts *ContractsSession) Deposit() (*types.Transaction, error) {
	return _Contracts.Contract.Deposit(&_Contracts.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(bool)
func (_Contracts *ContractsTransactorSession) Deposit() (*types.Transaction, error) {
	return _Contracts.Contract.Deposit(&_Contracts.TransactOpts)
}

// DepositFromPool is a paid mutator transaction binding the contract method 0x31989b58.
//
// Solidity: function depositFromPool() payable returns()
func (_Contracts *ContractsTransactor) DepositFromPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositFromPool")
}

// DepositFromPool is a paid mutator transaction binding the contract method 0x31989b58.
//
// Solidity: function depositFromPool() payable returns()
func (_Contracts *ContractsSession) DepositFromPool() (*types.Transaction, error) {
	return _Contracts.Contract.DepositFromPool(&_Contracts.TransactOpts)
}

// DepositFromPool is a paid mutator transaction binding the contract method 0x31989b58.
//
// Solidity: function depositFromPool() payable returns()
func (_Contracts *ContractsTransactorSession) DepositFromPool() (*types.Transaction, error) {
	return _Contracts.Contract.DepositFromPool(&_Contracts.TransactOpts)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Contracts *ContractsTransactor) DepositOnBehalfOf(opts *bind.TransactOpts, _account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositOnBehalfOf", _account, _releaseStart, _releaseEnd)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Contracts *ContractsSession) DepositOnBehalfOf(_account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositOnBehalfOf(&_Contracts.TransactOpts, _account, _releaseStart, _releaseEnd)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Contracts *ContractsTransactorSession) DepositOnBehalfOf(_account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositOnBehalfOf(&_Contracts.TransactOpts, _account, _releaseStart, _releaseEnd)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address _recipient) payable returns(bool)
func (_Contracts *ContractsTransactor) DepositTo(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositTo", _recipient)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address _recipient) payable returns(bool)
func (_Contracts *ContractsSession) DepositTo(_recipient common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DepositTo(&_Contracts.TransactOpts, _recipient)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address _recipient) payable returns(bool)
func (_Contracts *ContractsTransactorSession) DepositTo(_recipient common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DepositTo(&_Contracts.TransactOpts, _recipient)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contracts *ContractsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contracts *ContractsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.IncreaseAllowance(&_Contracts.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contracts *ContractsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.IncreaseAllowance(&_Contracts.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Contracts *ContractsSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Contracts *ContractsTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, _registry)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Contracts *ContractsTransactor) Lock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "lock", _amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Contracts *ContractsSession) Lock(_amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Lock(&_Contracts.TransactOpts, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) Lock(_amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Lock(&_Contracts.TransactOpts, _amount)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Contracts *ContractsTransactor) PurgeTimeLocks(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "purgeTimeLocks", _account)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Contracts *ContractsSession) PurgeTimeLocks(_account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PurgeTimeLocks(&_Contracts.TransactOpts, _account)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Contracts *ContractsTransactorSession) PurgeTimeLocks(_account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PurgeTimeLocks(&_Contracts.TransactOpts, _account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, receiver, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_Contracts *ContractsTransactor) TransferAndCall(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferAndCall", _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_Contracts *ContractsSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.TransferAndCall(&_Contracts.TransactOpts, _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.TransferAndCall(&_Contracts.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address receiver, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", owner, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address receiver, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) TransferFrom(owner common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, owner, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address receiver, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferFrom(owner common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, owner, receiver, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Contracts *ContractsTransactor) Unlock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlock", _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Contracts *ContractsSession) Unlock(_amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Unlock(&_Contracts.TransactOpts, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) Unlock(_amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Unlock(&_Contracts.TransactOpts, _amount)
}

// UpdateCompoundRate is a paid mutator transaction binding the contract method 0xc0ab9cbc.
//
// Solidity: function updateCompoundRate() returns(uint256)
func (_Contracts *ContractsTransactor) UpdateCompoundRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateCompoundRate")
}

// UpdateCompoundRate is a paid mutator transaction binding the contract method 0xc0ab9cbc.
//
// Solidity: function updateCompoundRate() returns(uint256)
func (_Contracts *ContractsSession) UpdateCompoundRate() (*types.Transaction, error) {
	return _Contracts.Contract.UpdateCompoundRate(&_Contracts.TransactOpts)
}

// UpdateCompoundRate is a paid mutator transaction binding the contract method 0xc0ab9cbc.
//
// Solidity: function updateCompoundRate() returns(uint256)
func (_Contracts *ContractsTransactorSession) UpdateCompoundRate() (*types.Transaction, error) {
	return _Contracts.Contract.UpdateCompoundRate(&_Contracts.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Contracts *ContractsSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Contracts *ContractsTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, _amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address _recipient, uint256 _amount) returns(bool)
func (_Contracts *ContractsTransactor) WithdrawTo(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawTo", _recipient, _amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address _recipient, uint256 _amount) returns(bool)
func (_Contracts *ContractsSession) WithdrawTo(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawTo(&_Contracts.TransactOpts, _recipient, _amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address _recipient, uint256 _amount) returns(bool)
func (_Contracts *ContractsTransactorSession) WithdrawTo(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawTo(&_Contracts.TransactOpts, _recipient, _amount)
}

// ContractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contracts contract.
type ContractsApprovalIterator struct {
	Event *ContractsApproval // Event containing the contract specifics and raw log

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
func (it *ContractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApproval)
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
		it.Event = new(ContractsApproval)
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
func (it *ContractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApproval represents a Approval event raised by the Contracts contract.
type ContractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalIterator{contract: _Contracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApproval)
				if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) ParseApproval(log types.Log) (*ContractsApproval, error) {
	event := new(ContractsApproval)
	if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDelegatorRewardClaimedIterator is returned from FilterDelegatorRewardClaimed and is used to iterate over the raw logs and unpacked data for DelegatorRewardClaimed events raised by the Contracts contract.
type ContractsDelegatorRewardClaimedIterator struct {
	Event *ContractsDelegatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ContractsDelegatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDelegatorRewardClaimed)
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
		it.Event = new(ContractsDelegatorRewardClaimed)
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
func (it *ContractsDelegatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDelegatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDelegatorRewardClaimed represents a DelegatorRewardClaimed event raised by the Contracts contract.
type ContractsDelegatorRewardClaimed struct {
	Delegator   common.Address
	TotalReward *big.Int
	NewBalance  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardClaimed is a free log retrieval operation binding the contract event 0xdf2cbec787902946d1147a360aed309764340def25e5912a6a885b67ee2dcdb8.
//
// Solidity: event DelegatorRewardClaimed(address indexed _delegator, uint256 _totalReward, uint256 _newBalance)
func (_Contracts *ContractsFilterer) FilterDelegatorRewardClaimed(opts *bind.FilterOpts, _delegator []common.Address) (*ContractsDelegatorRewardClaimedIterator, error) {

	var _delegatorRule []interface{}
	for _, _delegatorItem := range _delegator {
		_delegatorRule = append(_delegatorRule, _delegatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DelegatorRewardClaimed", _delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDelegatorRewardClaimedIterator{contract: _Contracts.contract, event: "DelegatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardClaimed is a free log subscription operation binding the contract event 0xdf2cbec787902946d1147a360aed309764340def25e5912a6a885b67ee2dcdb8.
//
// Solidity: event DelegatorRewardClaimed(address indexed _delegator, uint256 _totalReward, uint256 _newBalance)
func (_Contracts *ContractsFilterer) WatchDelegatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *ContractsDelegatorRewardClaimed, _delegator []common.Address) (event.Subscription, error) {

	var _delegatorRule []interface{}
	for _, _delegatorItem := range _delegator {
		_delegatorRule = append(_delegatorRule, _delegatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DelegatorRewardClaimed", _delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDelegatorRewardClaimed)
				if err := _Contracts.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
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

// ParseDelegatorRewardClaimed is a log parse operation binding the contract event 0xdf2cbec787902946d1147a360aed309764340def25e5912a6a885b67ee2dcdb8.
//
// Solidity: event DelegatorRewardClaimed(address indexed _delegator, uint256 _totalReward, uint256 _newBalance)
func (_Contracts *ContractsFilterer) ParseDelegatorRewardClaimed(log types.Log) (*ContractsDelegatorRewardClaimed, error) {
	event := new(ContractsDelegatorRewardClaimed)
	if err := _Contracts.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsNewTimeLockIterator is returned from FilterNewTimeLock and is used to iterate over the raw logs and unpacked data for NewTimeLock events raised by the Contracts contract.
type ContractsNewTimeLockIterator struct {
	Event *ContractsNewTimeLock // Event containing the contract specifics and raw log

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
func (it *ContractsNewTimeLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsNewTimeLock)
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
		it.Event = new(ContractsNewTimeLock)
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
func (it *ContractsNewTimeLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsNewTimeLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsNewTimeLock represents a NewTimeLock event raised by the Contracts contract.
type ContractsNewTimeLock struct {
	Account  common.Address
	Timelock TimeLockEntry
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewTimeLock is a free log retrieval operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Contracts *ContractsFilterer) FilterNewTimeLock(opts *bind.FilterOpts, _account []common.Address) (*ContractsNewTimeLockIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "NewTimeLock", _accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsNewTimeLockIterator{contract: _Contracts.contract, event: "NewTimeLock", logs: logs, sub: sub}, nil
}

// WatchNewTimeLock is a free log subscription operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Contracts *ContractsFilterer) WatchNewTimeLock(opts *bind.WatchOpts, sink chan<- *ContractsNewTimeLock, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "NewTimeLock", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsNewTimeLock)
				if err := _Contracts.contract.UnpackLog(event, "NewTimeLock", log); err != nil {
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

// ParseNewTimeLock is a log parse operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Contracts *ContractsFilterer) ParseNewTimeLock(log types.Log) (*ContractsNewTimeLock, error) {
	event := new(ContractsNewTimeLock)
	if err := _Contracts.contract.UnpackLog(event, "NewTimeLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPurgedIterator is returned from FilterPurged and is used to iterate over the raw logs and unpacked data for Purged events raised by the Contracts contract.
type ContractsPurgedIterator struct {
	Event *ContractsPurged // Event containing the contract specifics and raw log

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
func (it *ContractsPurgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPurged)
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
		it.Event = new(ContractsPurged)
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
func (it *ContractsPurgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPurgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPurged represents a Purged event raised by the Contracts contract.
type ContractsPurged struct {
	Account  common.Address
	Timelock TimeLockEntry
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurged is a free log retrieval operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Contracts *ContractsFilterer) FilterPurged(opts *bind.FilterOpts, _account []common.Address) (*ContractsPurgedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Purged", _accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPurgedIterator{contract: _Contracts.contract, event: "Purged", logs: logs, sub: sub}, nil
}

// WatchPurged is a free log subscription operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Contracts *ContractsFilterer) WatchPurged(opts *bind.WatchOpts, sink chan<- *ContractsPurged, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Purged", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPurged)
				if err := _Contracts.contract.UnpackLog(event, "Purged", log); err != nil {
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

// ParsePurged is a log parse operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Contracts *ContractsFilterer) ParsePurged(log types.Log) (*ContractsPurged, error) {
	event := new(ContractsPurged)
	if err := _Contracts.contract.UnpackLog(event, "Purged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

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
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
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
		it.Event = new(ContractsTransfer)
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
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_Contracts *ContractsFilterer) ParseTransfer(log types.Log) (*ContractsTransfer, error) {
	event := new(ContractsTransfer)
	if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransfer0Iterator is returned from FilterTransfer0 and is used to iterate over the raw logs and unpacked data for Transfer0 events raised by the Contracts contract.
type ContractsTransfer0Iterator struct {
	Event *ContractsTransfer0 // Event containing the contract specifics and raw log

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
func (it *ContractsTransfer0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer0)
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
		it.Event = new(ContractsTransfer0)
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
func (it *ContractsTransfer0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer0 represents a Transfer0 event raised by the Contracts contract.
type ContractsTransfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer0 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsTransfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransfer0Iterator{contract: _Contracts.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

// WatchTransfer0 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *ContractsTransfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer0)
				if err := _Contracts.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

// ParseTransfer0 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) ParseTransfer0(log types.Log) (*ContractsTransfer0, error) {
	event := new(ContractsTransfer0)
	if err := _Contracts.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUserDepositedIterator is returned from FilterUserDeposited and is used to iterate over the raw logs and unpacked data for UserDeposited events raised by the Contracts contract.
type ContractsUserDepositedIterator struct {
	Event *ContractsUserDeposited // Event containing the contract specifics and raw log

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
func (it *ContractsUserDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUserDeposited)
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
		it.Event = new(ContractsUserDeposited)
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
func (it *ContractsUserDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUserDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUserDeposited represents a UserDeposited event raised by the Contracts contract.
type ContractsUserDeposited struct {
	User             common.Address
	NewDepositAmount *big.Int
	NewBalance       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUserDeposited is a free log retrieval operation binding the contract event 0x9fa0303acb2d73e4e86bf54dd3fba8208feacfb4f5a79175dc1f36b674da0dbb.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _newDepositAmount, uint256 _newBalance)
func (_Contracts *ContractsFilterer) FilterUserDeposited(opts *bind.FilterOpts, _user []common.Address) (*ContractsUserDepositedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "UserDeposited", _userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUserDepositedIterator{contract: _Contracts.contract, event: "UserDeposited", logs: logs, sub: sub}, nil
}

// WatchUserDeposited is a free log subscription operation binding the contract event 0x9fa0303acb2d73e4e86bf54dd3fba8208feacfb4f5a79175dc1f36b674da0dbb.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _newDepositAmount, uint256 _newBalance)
func (_Contracts *ContractsFilterer) WatchUserDeposited(opts *bind.WatchOpts, sink chan<- *ContractsUserDeposited, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "UserDeposited", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUserDeposited)
				if err := _Contracts.contract.UnpackLog(event, "UserDeposited", log); err != nil {
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

// ParseUserDeposited is a log parse operation binding the contract event 0x9fa0303acb2d73e4e86bf54dd3fba8208feacfb4f5a79175dc1f36b674da0dbb.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _newDepositAmount, uint256 _newBalance)
func (_Contracts *ContractsFilterer) ParseUserDeposited(log types.Log) (*ContractsUserDeposited, error) {
	event := new(ContractsUserDeposited)
	if err := _Contracts.contract.UnpackLog(event, "UserDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUserWithdrawnIterator is returned from FilterUserWithdrawn and is used to iterate over the raw logs and unpacked data for UserWithdrawn events raised by the Contracts contract.
type ContractsUserWithdrawnIterator struct {
	Event *ContractsUserWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractsUserWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUserWithdrawn)
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
		it.Event = new(ContractsUserWithdrawn)
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
func (it *ContractsUserWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUserWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUserWithdrawn represents a UserWithdrawn event raised by the Contracts contract.
type ContractsUserWithdrawn struct {
	User            common.Address
	WithdrawnAmount *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUserWithdrawn is a free log retrieval operation binding the contract event 0x43389e74a5f67d287aa20ee5677bf6eaea427b4b436414723562a06c75debdc1.
//
// Solidity: event UserWithdrawn(address indexed _user, uint256 _withdrawnAmount, uint256 _newBalance)
func (_Contracts *ContractsFilterer) FilterUserWithdrawn(opts *bind.FilterOpts, _user []common.Address) (*ContractsUserWithdrawnIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "UserWithdrawn", _userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUserWithdrawnIterator{contract: _Contracts.contract, event: "UserWithdrawn", logs: logs, sub: sub}, nil
}

// WatchUserWithdrawn is a free log subscription operation binding the contract event 0x43389e74a5f67d287aa20ee5677bf6eaea427b4b436414723562a06c75debdc1.
//
// Solidity: event UserWithdrawn(address indexed _user, uint256 _withdrawnAmount, uint256 _newBalance)
func (_Contracts *ContractsFilterer) WatchUserWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsUserWithdrawn, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "UserWithdrawn", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUserWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "UserWithdrawn", log); err != nil {
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

// ParseUserWithdrawn is a log parse operation binding the contract event 0x43389e74a5f67d287aa20ee5677bf6eaea427b4b436414723562a06c75debdc1.
//
// Solidity: event UserWithdrawn(address indexed _user, uint256 _withdrawnAmount, uint256 _newBalance)
func (_Contracts *ContractsFilterer) ParseUserWithdrawn(log types.Log) (*ContractsUserWithdrawn, error) {
	event := new(ContractsUserWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "UserWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
