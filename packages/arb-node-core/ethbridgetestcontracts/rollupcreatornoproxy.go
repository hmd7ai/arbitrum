// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbsRollupABI is the input ABI used to generate the binding from.
const AbsRollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterSendAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"NodeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentNodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterInboxBatchEndCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterInboxBatchAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"indexed\":false,\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"NodeRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startNode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endNode\",\"type\":\"uint256\"}],\"name\":\"NodesDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNode\",\"type\":\"uint256\"}],\"name\":\"StakerReassigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeSendAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"executionHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proposedTimes\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"maxMessageCounts\",\"type\":\"uint256[2]\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedBridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraChallengeTimeBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[6]\",\"name\":\"connectedContracts\",\"type\":\"address[6]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"requireUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireUnresolvedExists\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstUnresolvedNodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestNodeCreated\",\"type\":\"uint256\"}],\"name\":\"requiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerBridge\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expectedNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"},{\"internalType\":\"uint256\",\"name\":\"beforeProposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxMaxCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sequencerBatchProof\",\"type\":\"bytes\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbsRollup is an auto generated Go binding around an Ethereum contract.
type AbsRollup struct {
	AbsRollupCaller     // Read-only binding to the contract
	AbsRollupTransactor // Write-only binding to the contract
	AbsRollupFilterer   // Log filterer for contract events
}

// AbsRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsRollupSession struct {
	Contract     *AbsRollup        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsRollupCallerSession struct {
	Contract *AbsRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AbsRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsRollupTransactorSession struct {
	Contract     *AbsRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AbsRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsRollupRaw struct {
	Contract *AbsRollup // Generic contract binding to access the raw methods on
}

// AbsRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsRollupCallerRaw struct {
	Contract *AbsRollupCaller // Generic read-only contract binding to access the raw methods on
}

// AbsRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsRollupTransactorRaw struct {
	Contract *AbsRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsRollup creates a new instance of AbsRollup, bound to a specific deployed contract.
func NewAbsRollup(address common.Address, backend bind.ContractBackend) (*AbsRollup, error) {
	contract, err := bindAbsRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbsRollup{AbsRollupCaller: AbsRollupCaller{contract: contract}, AbsRollupTransactor: AbsRollupTransactor{contract: contract}, AbsRollupFilterer: AbsRollupFilterer{contract: contract}}, nil
}

// NewAbsRollupCaller creates a new read-only instance of AbsRollup, bound to a specific deployed contract.
func NewAbsRollupCaller(address common.Address, caller bind.ContractCaller) (*AbsRollupCaller, error) {
	contract, err := bindAbsRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsRollupCaller{contract: contract}, nil
}

// NewAbsRollupTransactor creates a new write-only instance of AbsRollup, bound to a specific deployed contract.
func NewAbsRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsRollupTransactor, error) {
	contract, err := bindAbsRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsRollupTransactor{contract: contract}, nil
}

// NewAbsRollupFilterer creates a new log filterer instance of AbsRollup, bound to a specific deployed contract.
func NewAbsRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsRollupFilterer, error) {
	contract, err := bindAbsRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsRollupFilterer{contract: contract}, nil
}

// bindAbsRollup binds a generic wrapper to an already deployed contract.
func bindAbsRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsRollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbsRollup *AbsRollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbsRollup.Contract.AbsRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbsRollup *AbsRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsRollup.Contract.AbsRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbsRollup *AbsRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbsRollup.Contract.AbsRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbsRollup *AbsRollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbsRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbsRollup *AbsRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbsRollup *AbsRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbsRollup.Contract.contract.Transact(opts, method, params...)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_AbsRollup *AbsRollupCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "_stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_AbsRollup *AbsRollupSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _AbsRollup.Contract.StakerMap(&_AbsRollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_AbsRollup *AbsRollupCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _AbsRollup.Contract.StakerMap(&_AbsRollup.CallOpts, arg0)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.AmountStaked(&_AbsRollup.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.AmountStaked(&_AbsRollup.CallOpts, staker)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ArbGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "arbGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_AbsRollup *AbsRollupSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _AbsRollup.Contract.ArbGasSpeedLimitPerBlock(&_AbsRollup.CallOpts)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _AbsRollup.Contract.ArbGasSpeedLimitPerBlock(&_AbsRollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) BaseStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "baseStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_AbsRollup *AbsRollupSession) BaseStake() (*big.Int, error) {
	return _AbsRollup.Contract.BaseStake(&_AbsRollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) BaseStake() (*big.Int, error) {
	return _AbsRollup.Contract.BaseStake(&_AbsRollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_AbsRollup *AbsRollupCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_AbsRollup *AbsRollupSession) ChallengeFactory() (common.Address, error) {
	return _AbsRollup.Contract.ChallengeFactory(&_AbsRollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) ChallengeFactory() (common.Address, error) {
	return _AbsRollup.Contract.ChallengeFactory(&_AbsRollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ConfirmPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "confirmPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _AbsRollup.Contract.ConfirmPeriodBlocks(&_AbsRollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _AbsRollup.Contract.ConfirmPeriodBlocks(&_AbsRollup.CallOpts)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) CountStakedZombies(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "countStakedZombies", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_AbsRollup *AbsRollupSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.CountStakedZombies(&_AbsRollup.CallOpts, node)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.CountStakedZombies(&_AbsRollup.CallOpts, node)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_AbsRollup *AbsRollupCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_AbsRollup *AbsRollupSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _AbsRollup.Contract.CurrentChallenge(&_AbsRollup.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_AbsRollup *AbsRollupCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _AbsRollup.Contract.CurrentChallenge(&_AbsRollup.CallOpts, staker)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_AbsRollup *AbsRollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _AbsRollup.Contract.CurrentRequiredStake(&_AbsRollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _AbsRollup.Contract.CurrentRequiredStake(&_AbsRollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_AbsRollup *AbsRollupCaller) DelayedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "delayedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_AbsRollup *AbsRollupSession) DelayedBridge() (common.Address, error) {
	return _AbsRollup.Contract.DelayedBridge(&_AbsRollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) DelayedBridge() (common.Address, error) {
	return _AbsRollup.Contract.DelayedBridge(&_AbsRollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ExtraChallengeTimeBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "extraChallengeTimeBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _AbsRollup.Contract.ExtraChallengeTimeBlocks(&_AbsRollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _AbsRollup.Contract.ExtraChallengeTimeBlocks(&_AbsRollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_AbsRollup *AbsRollupSession) FirstUnresolvedNode() (*big.Int, error) {
	return _AbsRollup.Contract.FirstUnresolvedNode(&_AbsRollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _AbsRollup.Contract.FirstUnresolvedNode(&_AbsRollup.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_AbsRollup *AbsRollupCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_AbsRollup *AbsRollupSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.GetNode(&_AbsRollup.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_AbsRollup *AbsRollupCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.GetNode(&_AbsRollup.CallOpts, nodeNum)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_AbsRollup *AbsRollupCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_AbsRollup *AbsRollupSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _AbsRollup.Contract.GetNodeHash(&_AbsRollup.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_AbsRollup *AbsRollupCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _AbsRollup.Contract.GetNodeHash(&_AbsRollup.CallOpts, index)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_AbsRollup *AbsRollupCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_AbsRollup *AbsRollupSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.GetStakerAddress(&_AbsRollup.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_AbsRollup *AbsRollupCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.GetStakerAddress(&_AbsRollup.CallOpts, stakerNum)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_AbsRollup *AbsRollupCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_AbsRollup *AbsRollupSession) IsMaster() (bool, error) {
	return _AbsRollup.Contract.IsMaster(&_AbsRollup.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_AbsRollup *AbsRollupCallerSession) IsMaster() (bool, error) {
	return _AbsRollup.Contract.IsMaster(&_AbsRollup.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_AbsRollup *AbsRollupCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_AbsRollup *AbsRollupSession) IsStaked(staker common.Address) (bool, error) {
	return _AbsRollup.Contract.IsStaked(&_AbsRollup.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_AbsRollup *AbsRollupCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _AbsRollup.Contract.IsStaked(&_AbsRollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_AbsRollup *AbsRollupCaller) IsZombie(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "isZombie", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_AbsRollup *AbsRollupSession) IsZombie(staker common.Address) (bool, error) {
	return _AbsRollup.Contract.IsZombie(&_AbsRollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_AbsRollup *AbsRollupCallerSession) IsZombie(staker common.Address) (bool, error) {
	return _AbsRollup.Contract.IsZombie(&_AbsRollup.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_AbsRollup *AbsRollupSession) LastStakeBlock() (*big.Int, error) {
	return _AbsRollup.Contract.LastStakeBlock(&_AbsRollup.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) LastStakeBlock() (*big.Int, error) {
	return _AbsRollup.Contract.LastStakeBlock(&_AbsRollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_AbsRollup *AbsRollupSession) LatestConfirmed() (*big.Int, error) {
	return _AbsRollup.Contract.LatestConfirmed(&_AbsRollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) LatestConfirmed() (*big.Int, error) {
	return _AbsRollup.Contract.LatestConfirmed(&_AbsRollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_AbsRollup *AbsRollupSession) LatestNodeCreated() (*big.Int, error) {
	return _AbsRollup.Contract.LatestNodeCreated(&_AbsRollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _AbsRollup.Contract.LatestNodeCreated(&_AbsRollup.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.LatestStakedNode(&_AbsRollup.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.LatestStakedNode(&_AbsRollup.CallOpts, staker)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_AbsRollup *AbsRollupSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _AbsRollup.Contract.MinimumAssertionPeriod(&_AbsRollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _AbsRollup.Contract.MinimumAssertionPeriod(&_AbsRollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_AbsRollup *AbsRollupCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_AbsRollup *AbsRollupSession) NodeFactory() (common.Address, error) {
	return _AbsRollup.Contract.NodeFactory(&_AbsRollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) NodeFactory() (common.Address, error) {
	return _AbsRollup.Contract.NodeFactory(&_AbsRollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_AbsRollup *AbsRollupCaller) Outbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "outbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_AbsRollup *AbsRollupSession) Outbox() (common.Address, error) {
	return _AbsRollup.Contract.Outbox(&_AbsRollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) Outbox() (common.Address, error) {
	return _AbsRollup.Contract.Outbox(&_AbsRollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbsRollup *AbsRollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbsRollup *AbsRollupSession) Owner() (common.Address, error) {
	return _AbsRollup.Contract.Owner(&_AbsRollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) Owner() (common.Address, error) {
	return _AbsRollup.Contract.Owner(&_AbsRollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AbsRollup *AbsRollupCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AbsRollup *AbsRollupSession) Paused() (bool, error) {
	return _AbsRollup.Contract.Paused(&_AbsRollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AbsRollup *AbsRollupCallerSession) Paused() (bool, error) {
	return _AbsRollup.Contract.Paused(&_AbsRollup.CallOpts)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_AbsRollup *AbsRollupCaller) RequireUnresolved(opts *bind.CallOpts, nodeNum *big.Int) error {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "requireUnresolved", nodeNum)

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_AbsRollup *AbsRollupSession) RequireUnresolved(nodeNum *big.Int) error {
	return _AbsRollup.Contract.RequireUnresolved(&_AbsRollup.CallOpts, nodeNum)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_AbsRollup *AbsRollupCallerSession) RequireUnresolved(nodeNum *big.Int) error {
	return _AbsRollup.Contract.RequireUnresolved(&_AbsRollup.CallOpts, nodeNum)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_AbsRollup *AbsRollupCaller) RequireUnresolvedExists(opts *bind.CallOpts) error {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "requireUnresolvedExists")

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_AbsRollup *AbsRollupSession) RequireUnresolvedExists() error {
	return _AbsRollup.Contract.RequireUnresolvedExists(&_AbsRollup.CallOpts)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_AbsRollup *AbsRollupCallerSession) RequireUnresolvedExists() error {
	return _AbsRollup.Contract.RequireUnresolvedExists(&_AbsRollup.CallOpts)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) RequiredStake(opts *bind.CallOpts, blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "requiredStake", blockNumber, firstUnresolvedNodeNum, latestNodeCreated)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_AbsRollup *AbsRollupSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _AbsRollup.Contract.RequiredStake(&_AbsRollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _AbsRollup.Contract.RequiredStake(&_AbsRollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_AbsRollup *AbsRollupCaller) RollupEventBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "rollupEventBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_AbsRollup *AbsRollupSession) RollupEventBridge() (common.Address, error) {
	return _AbsRollup.Contract.RollupEventBridge(&_AbsRollup.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) RollupEventBridge() (common.Address, error) {
	return _AbsRollup.Contract.RollupEventBridge(&_AbsRollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_AbsRollup *AbsRollupCaller) SequencerBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "sequencerBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_AbsRollup *AbsRollupSession) SequencerBridge() (common.Address, error) {
	return _AbsRollup.Contract.SequencerBridge(&_AbsRollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) SequencerBridge() (common.Address, error) {
	return _AbsRollup.Contract.SequencerBridge(&_AbsRollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_AbsRollup *AbsRollupSession) StakerCount() (*big.Int, error) {
	return _AbsRollup.Contract.StakerCount(&_AbsRollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) StakerCount() (*big.Int, error) {
	return _AbsRollup.Contract.StakerCount(&_AbsRollup.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_AbsRollup *AbsRollupSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.WithdrawableFunds(&_AbsRollup.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.WithdrawableFunds(&_AbsRollup.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_AbsRollup *AbsRollupCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_AbsRollup *AbsRollupSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.ZombieAddress(&_AbsRollup.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_AbsRollup *AbsRollupCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.ZombieAddress(&_AbsRollup.CallOpts, zombieNum)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_AbsRollup *AbsRollupSession) ZombieCount() (*big.Int, error) {
	return _AbsRollup.Contract.ZombieCount(&_AbsRollup.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ZombieCount() (*big.Int, error) {
	return _AbsRollup.Contract.ZombieCount(&_AbsRollup.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_AbsRollup *AbsRollupSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _AbsRollup.Contract.ZombieLatestStakedNode(&_AbsRollup.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _AbsRollup.Contract.ZombieLatestStakedNode(&_AbsRollup.CallOpts, zombieNum)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_AbsRollup *AbsRollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "completeChallenge", winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_AbsRollup *AbsRollupSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.CompleteChallenge(&_AbsRollup.TransactOpts, winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_AbsRollup *AbsRollupTransactorSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.CompleteChallenge(&_AbsRollup.TransactOpts, winningStaker, losingStaker)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_AbsRollup *AbsRollupTransactor) ConfirmNextNode(opts *bind.TransactOpts, beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "confirmNextNode", beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_AbsRollup *AbsRollupSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.ConfirmNextNode(&_AbsRollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_AbsRollup *AbsRollupTransactorSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.ConfirmNextNode(&_AbsRollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_AbsRollup *AbsRollupTransactor) CreateChallenge(opts *bind.TransactOpts, stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "createChallenge", stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_AbsRollup *AbsRollupSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.CreateChallenge(&_AbsRollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_AbsRollup *AbsRollupTransactorSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.CreateChallenge(&_AbsRollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_AbsRollup *AbsRollupTransactor) Initialize(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "initialize", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_AbsRollup *AbsRollupSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.Initialize(&_AbsRollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_AbsRollup *AbsRollupTransactorSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.Initialize(&_AbsRollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AbsRollup *AbsRollupTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AbsRollup *AbsRollupSession) Pause() (*types.Transaction, error) {
	return _AbsRollup.Contract.Pause(&_AbsRollup.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AbsRollup *AbsRollupTransactorSession) Pause() (*types.Transaction, error) {
	return _AbsRollup.Contract.Pause(&_AbsRollup.TransactOpts)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_AbsRollup *AbsRollupTransactor) ReduceDeposit(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "reduceDeposit", target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_AbsRollup *AbsRollupSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.ReduceDeposit(&_AbsRollup.TransactOpts, target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_AbsRollup *AbsRollupTransactorSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.ReduceDeposit(&_AbsRollup.TransactOpts, target)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_AbsRollup *AbsRollupTransactor) RejectNextNode(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "rejectNextNode", stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_AbsRollup *AbsRollupSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.RejectNextNode(&_AbsRollup.TransactOpts, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_AbsRollup *AbsRollupTransactorSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.RejectNextNode(&_AbsRollup.TransactOpts, stakerAddress)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupTransactor) RemoveOldOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "removeOldOutbox", _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveOldOutbox(&_AbsRollup.TransactOpts, _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupTransactorSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveOldOutbox(&_AbsRollup.TransactOpts, _outbox)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_AbsRollup *AbsRollupTransactor) RemoveOldZombies(opts *bind.TransactOpts, startIndex *big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "removeOldZombies", startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_AbsRollup *AbsRollupSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveOldZombies(&_AbsRollup.TransactOpts, startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_AbsRollup *AbsRollupTransactorSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveOldZombies(&_AbsRollup.TransactOpts, startIndex)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_AbsRollup *AbsRollupTransactor) RemoveZombie(opts *bind.TransactOpts, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "removeZombie", zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_AbsRollup *AbsRollupSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveZombie(&_AbsRollup.TransactOpts, zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_AbsRollup *AbsRollupTransactorSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveZombie(&_AbsRollup.TransactOpts, zombieNum, maxNodes)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AbsRollup *AbsRollupTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AbsRollup *AbsRollupSession) Resume() (*types.Transaction, error) {
	return _AbsRollup.Contract.Resume(&_AbsRollup.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AbsRollup *AbsRollupTransactorSession) Resume() (*types.Transaction, error) {
	return _AbsRollup.Contract.Resume(&_AbsRollup.TransactOpts)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_AbsRollup *AbsRollupTransactor) ReturnOldDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "returnOldDeposit", stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_AbsRollup *AbsRollupSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.ReturnOldDeposit(&_AbsRollup.TransactOpts, stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_AbsRollup *AbsRollupTransactorSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.ReturnOldDeposit(&_AbsRollup.TransactOpts, stakerAddress)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_AbsRollup *AbsRollupTransactor) SetInbox(opts *bind.TransactOpts, _inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "setInbox", _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_AbsRollup *AbsRollupSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _AbsRollup.Contract.SetInbox(&_AbsRollup.TransactOpts, _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_AbsRollup *AbsRollupTransactorSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _AbsRollup.Contract.SetInbox(&_AbsRollup.TransactOpts, _inbox, _enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupTransactor) SetOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "setOutbox", _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.SetOutbox(&_AbsRollup.TransactOpts, _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupTransactorSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.SetOutbox(&_AbsRollup.TransactOpts, _outbox)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_AbsRollup *AbsRollupTransactor) StakeOnExistingNode(opts *bind.TransactOpts, nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "stakeOnExistingNode", nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_AbsRollup *AbsRollupSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.StakeOnExistingNode(&_AbsRollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_AbsRollup *AbsRollupTransactorSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.StakeOnExistingNode(&_AbsRollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_AbsRollup *AbsRollupTransactor) StakeOnNewNode(opts *bind.TransactOpts, expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "stakeOnNewNode", expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_AbsRollup *AbsRollupSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.StakeOnNewNode(&_AbsRollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_AbsRollup *AbsRollupTransactorSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.StakeOnNewNode(&_AbsRollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_AbsRollup *AbsRollupTransactor) WithdrawStakerFunds(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "withdrawStakerFunds", destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_AbsRollup *AbsRollupSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.WithdrawStakerFunds(&_AbsRollup.TransactOpts, destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_AbsRollup *AbsRollupTransactorSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.WithdrawStakerFunds(&_AbsRollup.TransactOpts, destination)
}

// AbsRollupNodeConfirmedIterator is returned from FilterNodeConfirmed and is used to iterate over the raw logs and unpacked data for NodeConfirmed events raised by the AbsRollup contract.
type AbsRollupNodeConfirmedIterator struct {
	Event *AbsRollupNodeConfirmed // Event containing the contract specifics and raw log

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
func (it *AbsRollupNodeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupNodeConfirmed)
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
		it.Event = new(AbsRollupNodeConfirmed)
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
func (it *AbsRollupNodeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupNodeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupNodeConfirmed represents a NodeConfirmed event raised by the AbsRollup contract.
type AbsRollupNodeConfirmed struct {
	NodeNum        *big.Int
	AfterSendAcc   [32]byte
	AfterSendCount *big.Int
	AfterLogAcc    [32]byte
	AfterLogCount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeConfirmed is a free log retrieval operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_AbsRollup *AbsRollupFilterer) FilterNodeConfirmed(opts *bind.FilterOpts, nodeNum []*big.Int) (*AbsRollupNodeConfirmedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupNodeConfirmedIterator{contract: _AbsRollup.contract, event: "NodeConfirmed", logs: logs, sub: sub}, nil
}

// WatchNodeConfirmed is a free log subscription operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_AbsRollup *AbsRollupFilterer) WatchNodeConfirmed(opts *bind.WatchOpts, sink chan<- *AbsRollupNodeConfirmed, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupNodeConfirmed)
				if err := _AbsRollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
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

// ParseNodeConfirmed is a log parse operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_AbsRollup *AbsRollupFilterer) ParseNodeConfirmed(log types.Log) (*AbsRollupNodeConfirmed, error) {
	event := new(AbsRollupNodeConfirmed)
	if err := _AbsRollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the AbsRollup contract.
type AbsRollupNodeCreatedIterator struct {
	Event *AbsRollupNodeCreated // Event containing the contract specifics and raw log

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
func (it *AbsRollupNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupNodeCreated)
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
		it.Event = new(AbsRollupNodeCreated)
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
func (it *AbsRollupNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupNodeCreated represents a NodeCreated event raised by the AbsRollup contract.
type AbsRollupNodeCreated struct {
	NodeNum                 *big.Int
	ParentNodeHash          [32]byte
	NodeHash                [32]byte
	ExecutionHash           [32]byte
	InboxMaxCount           *big.Int
	AfterInboxBatchEndCount *big.Int
	AfterInboxBatchAcc      [32]byte
	AssertionBytes32Fields  [2][3][32]byte
	AssertionIntFields      [2][4]*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_AbsRollup *AbsRollupFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int, parentNodeHash [][32]byte) (*AbsRollupNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupNodeCreatedIterator{contract: _AbsRollup.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_AbsRollup *AbsRollupFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *AbsRollupNodeCreated, nodeNum []*big.Int, parentNodeHash [][32]byte) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupNodeCreated)
				if err := _AbsRollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
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

// ParseNodeCreated is a log parse operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_AbsRollup *AbsRollupFilterer) ParseNodeCreated(log types.Log) (*AbsRollupNodeCreated, error) {
	event := new(AbsRollupNodeCreated)
	if err := _AbsRollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupNodeRejectedIterator is returned from FilterNodeRejected and is used to iterate over the raw logs and unpacked data for NodeRejected events raised by the AbsRollup contract.
type AbsRollupNodeRejectedIterator struct {
	Event *AbsRollupNodeRejected // Event containing the contract specifics and raw log

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
func (it *AbsRollupNodeRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupNodeRejected)
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
		it.Event = new(AbsRollupNodeRejected)
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
func (it *AbsRollupNodeRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupNodeRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupNodeRejected represents a NodeRejected event raised by the AbsRollup contract.
type AbsRollupNodeRejected struct {
	NodeNum *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeRejected is a free log retrieval operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_AbsRollup *AbsRollupFilterer) FilterNodeRejected(opts *bind.FilterOpts, nodeNum []*big.Int) (*AbsRollupNodeRejectedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupNodeRejectedIterator{contract: _AbsRollup.contract, event: "NodeRejected", logs: logs, sub: sub}, nil
}

// WatchNodeRejected is a free log subscription operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_AbsRollup *AbsRollupFilterer) WatchNodeRejected(opts *bind.WatchOpts, sink chan<- *AbsRollupNodeRejected, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupNodeRejected)
				if err := _AbsRollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
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

// ParseNodeRejected is a log parse operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_AbsRollup *AbsRollupFilterer) ParseNodeRejected(log types.Log) (*AbsRollupNodeRejected, error) {
	event := new(AbsRollupNodeRejected)
	if err := _AbsRollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupNodesDestroyedIterator is returned from FilterNodesDestroyed and is used to iterate over the raw logs and unpacked data for NodesDestroyed events raised by the AbsRollup contract.
type AbsRollupNodesDestroyedIterator struct {
	Event *AbsRollupNodesDestroyed // Event containing the contract specifics and raw log

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
func (it *AbsRollupNodesDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupNodesDestroyed)
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
		it.Event = new(AbsRollupNodesDestroyed)
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
func (it *AbsRollupNodesDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupNodesDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupNodesDestroyed represents a NodesDestroyed event raised by the AbsRollup contract.
type AbsRollupNodesDestroyed struct {
	StartNode *big.Int
	EndNode   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodesDestroyed is a free log retrieval operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_AbsRollup *AbsRollupFilterer) FilterNodesDestroyed(opts *bind.FilterOpts, startNode []*big.Int, endNode []*big.Int) (*AbsRollupNodesDestroyedIterator, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupNodesDestroyedIterator{contract: _AbsRollup.contract, event: "NodesDestroyed", logs: logs, sub: sub}, nil
}

// WatchNodesDestroyed is a free log subscription operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_AbsRollup *AbsRollupFilterer) WatchNodesDestroyed(opts *bind.WatchOpts, sink chan<- *AbsRollupNodesDestroyed, startNode []*big.Int, endNode []*big.Int) (event.Subscription, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupNodesDestroyed)
				if err := _AbsRollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
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

// ParseNodesDestroyed is a log parse operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_AbsRollup *AbsRollupFilterer) ParseNodesDestroyed(log types.Log) (*AbsRollupNodesDestroyed, error) {
	event := new(AbsRollupNodesDestroyed)
	if err := _AbsRollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the AbsRollup contract.
type AbsRollupOwnerFunctionCalledIterator struct {
	Event *AbsRollupOwnerFunctionCalled // Event containing the contract specifics and raw log

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
func (it *AbsRollupOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupOwnerFunctionCalled)
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
		it.Event = new(AbsRollupOwnerFunctionCalled)
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
func (it *AbsRollupOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the AbsRollup contract.
type AbsRollupOwnerFunctionCalled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0x6eb4661c504ed343cbf8e1e4f8024569117770d3cd0062bfe94b454cf0c48f20.
//
// Solidity: event OwnerFunctionCalled()
func (_AbsRollup *AbsRollupFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts) (*AbsRollupOwnerFunctionCalledIterator, error) {

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "OwnerFunctionCalled")
	if err != nil {
		return nil, err
	}
	return &AbsRollupOwnerFunctionCalledIterator{contract: _AbsRollup.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0x6eb4661c504ed343cbf8e1e4f8024569117770d3cd0062bfe94b454cf0c48f20.
//
// Solidity: event OwnerFunctionCalled()
func (_AbsRollup *AbsRollupFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *AbsRollupOwnerFunctionCalled) (event.Subscription, error) {

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "OwnerFunctionCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupOwnerFunctionCalled)
				if err := _AbsRollup.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
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

// ParseOwnerFunctionCalled is a log parse operation binding the contract event 0x6eb4661c504ed343cbf8e1e4f8024569117770d3cd0062bfe94b454cf0c48f20.
//
// Solidity: event OwnerFunctionCalled()
func (_AbsRollup *AbsRollupFilterer) ParseOwnerFunctionCalled(log types.Log) (*AbsRollupOwnerFunctionCalled, error) {
	event := new(AbsRollupOwnerFunctionCalled)
	if err := _AbsRollup.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AbsRollup contract.
type AbsRollupPausedIterator struct {
	Event *AbsRollupPaused // Event containing the contract specifics and raw log

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
func (it *AbsRollupPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupPaused)
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
		it.Event = new(AbsRollupPaused)
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
func (it *AbsRollupPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupPaused represents a Paused event raised by the AbsRollup contract.
type AbsRollupPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AbsRollup *AbsRollupFilterer) FilterPaused(opts *bind.FilterOpts) (*AbsRollupPausedIterator, error) {

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AbsRollupPausedIterator{contract: _AbsRollup.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AbsRollup *AbsRollupFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AbsRollupPaused) (event.Subscription, error) {

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupPaused)
				if err := _AbsRollup.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AbsRollup *AbsRollupFilterer) ParsePaused(log types.Log) (*AbsRollupPaused, error) {
	event := new(AbsRollupPaused)
	if err := _AbsRollup.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the AbsRollup contract.
type AbsRollupRollupChallengeStartedIterator struct {
	Event *AbsRollupRollupChallengeStarted // Event containing the contract specifics and raw log

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
func (it *AbsRollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupRollupChallengeStarted)
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
		it.Event = new(AbsRollupRollupChallengeStarted)
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
func (it *AbsRollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the AbsRollup contract.
type AbsRollupRollupChallengeStarted struct {
	ChallengeContract common.Address
	Asserter          common.Address
	Challenger        common.Address
	ChallengedNode    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_AbsRollup *AbsRollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts, challengeContract []common.Address) (*AbsRollupRollupChallengeStartedIterator, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupRollupChallengeStartedIterator{contract: _AbsRollup.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_AbsRollup *AbsRollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *AbsRollupRollupChallengeStarted, challengeContract []common.Address) (event.Subscription, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupRollupChallengeStarted)
				if err := _AbsRollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
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

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_AbsRollup *AbsRollupFilterer) ParseRollupChallengeStarted(log types.Log) (*AbsRollupRollupChallengeStarted, error) {
	event := new(AbsRollupRollupChallengeStarted)
	if err := _AbsRollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the AbsRollup contract.
type AbsRollupRollupCreatedIterator struct {
	Event *AbsRollupRollupCreated // Event containing the contract specifics and raw log

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
func (it *AbsRollupRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupRollupCreated)
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
		it.Event = new(AbsRollupRollupCreated)
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
func (it *AbsRollupRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupRollupCreated represents a RollupCreated event raised by the AbsRollup contract.
type AbsRollupRollupCreated struct {
	MachineHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_AbsRollup *AbsRollupFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*AbsRollupRollupCreatedIterator, error) {

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &AbsRollupRollupCreatedIterator{contract: _AbsRollup.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_AbsRollup *AbsRollupFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *AbsRollupRollupCreated) (event.Subscription, error) {

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupRollupCreated)
				if err := _AbsRollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_AbsRollup *AbsRollupFilterer) ParseRollupCreated(log types.Log) (*AbsRollupRollupCreated, error) {
	event := new(AbsRollupRollupCreated)
	if err := _AbsRollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupStakerReassignedIterator is returned from FilterStakerReassigned and is used to iterate over the raw logs and unpacked data for StakerReassigned events raised by the AbsRollup contract.
type AbsRollupStakerReassignedIterator struct {
	Event *AbsRollupStakerReassigned // Event containing the contract specifics and raw log

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
func (it *AbsRollupStakerReassignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupStakerReassigned)
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
		it.Event = new(AbsRollupStakerReassigned)
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
func (it *AbsRollupStakerReassignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupStakerReassignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupStakerReassigned represents a StakerReassigned event raised by the AbsRollup contract.
type AbsRollupStakerReassigned struct {
	Staker  common.Address
	NewNode *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakerReassigned is a free log retrieval operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_AbsRollup *AbsRollupFilterer) FilterStakerReassigned(opts *bind.FilterOpts, staker []common.Address) (*AbsRollupStakerReassignedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupStakerReassignedIterator{contract: _AbsRollup.contract, event: "StakerReassigned", logs: logs, sub: sub}, nil
}

// WatchStakerReassigned is a free log subscription operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_AbsRollup *AbsRollupFilterer) WatchStakerReassigned(opts *bind.WatchOpts, sink chan<- *AbsRollupStakerReassigned, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupStakerReassigned)
				if err := _AbsRollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
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

// ParseStakerReassigned is a log parse operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_AbsRollup *AbsRollupFilterer) ParseStakerReassigned(log types.Log) (*AbsRollupStakerReassigned, error) {
	event := new(AbsRollupStakerReassigned)
	if err := _AbsRollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AbsRollup contract.
type AbsRollupUnpausedIterator struct {
	Event *AbsRollupUnpaused // Event containing the contract specifics and raw log

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
func (it *AbsRollupUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupUnpaused)
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
		it.Event = new(AbsRollupUnpaused)
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
func (it *AbsRollupUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupUnpaused represents a Unpaused event raised by the AbsRollup contract.
type AbsRollupUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AbsRollup *AbsRollupFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AbsRollupUnpausedIterator, error) {

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AbsRollupUnpausedIterator{contract: _AbsRollup.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AbsRollup *AbsRollupFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AbsRollupUnpaused) (event.Subscription, error) {

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupUnpaused)
				if err := _AbsRollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AbsRollup *AbsRollupFilterer) ParseUnpaused(log types.Log) (*AbsRollupUnpaused, error) {
	event := new(AbsRollupUnpaused)
	if err := _AbsRollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"deliverMessageToInbox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b610e948061007d6000396000f3fe6080604052600436106100ad5760003560e01c806302bbfad1146100b25780633dbcc8d1146100f9578063413b35bd1461010e578063715018a6146101555780637ee943291461016c5780638da5cb5b146101b2578063945e1147146101c75780639e5d4c49146101f1578063ab5d894314610302578063c29372de14610317578063cee3d7281461034a578063d9dd67ab14610385578063e45b7ce6146103af578063f2fde38b146103ea575b600080fd5b6100e7600480360360608110156100c857600080fd5b5060ff813516906001600160a01b03602082013516906040013561041d565b60408051918252519081900360200190f35b34801561010557600080fd5b506100e7610533565b34801561011a57600080fd5b506101416004803603602081101561013157600080fd5b50356001600160a01b0316610539565b604080519115158252519081900360200190f35b34801561016157600080fd5b5061016a61055a565b005b34801561017857600080fd5b506101966004803603602081101561018f57600080fd5b50356105f4565b604080516001600160a01b039092168252519081900360200190f35b3480156101be57600080fd5b5061019661061b565b3480156101d357600080fd5b50610196600480360360208110156101ea57600080fd5b503561062a565b3480156101fd57600080fd5b506102816004803603606081101561021457600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561024357600080fd5b82018360208201111561025557600080fd5b803590602001918460018302840111600160201b8311171561027657600080fd5b509092509050610637565b60405180831515815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102c65781810151838201526020016102ae565b50505050905090810190601f1680156102f35780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561030e57600080fd5b5061019661079b565b34801561032357600080fd5b506101416004803603602081101561033a57600080fd5b50356001600160a01b03166107aa565b34801561035657600080fd5b5061016a6004803603604081101561036d57600080fd5b506001600160a01b03813516906020013515156107cc565b34801561039157600080fd5b506100e7600480360360208110156103a857600080fd5b5035610a08565b3480156103bb57600080fd5b5061016a600480360360408110156103d257600080fd5b506001600160a01b0381351690602001351515610a26565b3480156103f657600080fd5b5061016a6004803603602081101561040d57600080fd5b50356001600160a01b0316610c61565b3360009081526001602081905260408220015460ff16610475576040805162461bcd60e51b815260206004820152600e60248201526d09c9ea8be8ca49e9abe929c849eb60931b604482015290519081900360640190fd5b600654600061048986864342863a8a610d51565b9050600082156104b157600660018403815481106104a357fe5b906000526020600020015490505b60066104bd8284610dc2565b8154600181018355600092835260209283902001556040805133815260ff8a16928101929092526001600160a01b038816828201526060820187905251829185917f23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf79181900360800190a3509095945050505050565b60065490565b6001600160a01b031660009081526002602052604090206001015460ff1690565b610562610dee565b6001600160a01b031661057361061b565b6001600160a01b0316146105bc576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1f833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610e3f833981519152908390a3600080546001600160a01b0319169055565b6003818154811061060157fe5b6000918252602090912001546001600160a01b0316905081565b6000546001600160a01b031690565b6004818154811061060157fe5b3360009081526002602052604081206001015460609060ff16610693576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b82156106ee576106ab866001600160a01b0316610df2565b6106ee576040805162461bcd60e51b815260206004820152600f60248201526e1393d7d0d3d11157d05517d11154d5608a1b604482015290519081900360640190fd5b600580546001600160a01b0319811633179091556040516001600160a01b0391821691881690879087908790808383808284376040519201945060009350909150508083038185875af1925050503d8060008114610768576040519150601f19603f3d011682016040523d82523d6000602084013e61076d565b606091505b50600580546001600160a01b0319166001600160a01b03949094169390931790925597909650945050505050565b6005546001600160a01b031681565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b6107d4610dee565b6001600160a01b03166107e561061b565b6001600160a01b03161461082e576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1f833981519152604482015290519081900360640190fd5b6001600160a01b0382166000908152600260205260409020600181015460ff168080156108585750825b8061086a57508015801561086a575082155b15610876575050610a04565b821561090557604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319169091179055610a01565b60048054600019810190811061091757fe5b6000918252602090912001548254600480546001600160a01b0390931692909190811061094057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000015460026000600485600001548154811061098857fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205560048054806109b857fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff191690555b50505b5050565b60068181548110610a1557fe5b600091825260209091200154905081565b610a2e610dee565b6001600160a01b0316610a3f61061b565b6001600160a01b031614610a88576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1f833981519152604482015290519081900360640190fd5b6001600160a01b03821660009081526001602081905260409091209081015460ff16808015610ab45750825b80610ac6575080158015610ac6575082155b15610ad2575050610a04565b8215610b6057604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319169091179055610a01565b600380546000198101908110610b7257fe5b6000918252602090912001548254600380546001600160a01b03909316929091908110610b9b57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600160006003856000015481548110610be357fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556003805480610c1357fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff1916905550505050565b610c69610dee565b6001600160a01b0316610c7a61061b565b6001600160a01b031614610cc3576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1f833981519152604482015290519081900360640190fd5b6001600160a01b038116610d085760405162461bcd60e51b8152600401808060200182810382526026815260200180610df96026913960400191505060405180910390fd5b600080546040516001600160a01b0380851693921691600080516020610e3f83398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6001600160601b0319166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b3390565b3b15159056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220916eba2c20193ad95ae845bfc395ed112a797faf67d09f6796784d3ebf065b8c64736f6c634300060c0033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeSession) ActiveOutbox() (common.Address, error) {
	return _Bridge.Contract.ActiveOutbox(&_Bridge.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeCallerSession) ActiveOutbox() (common.Address, error) {
	return _Bridge.Contract.ActiveOutbox(&_Bridge.CallOpts)
}

// AllowedInboxList is a free data retrieval call binding the contract method 0x7ee94329.
//
// Solidity: function allowedInboxList(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) AllowedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedInboxList is a free data retrieval call binding the contract method 0x7ee94329.
//
// Solidity: function allowedInboxList(uint256 ) view returns(address)
func (_Bridge *BridgeSession) AllowedInboxList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.AllowedInboxList(&_Bridge.CallOpts, arg0)
}

// AllowedInboxList is a free data retrieval call binding the contract method 0x7ee94329.
//
// Solidity: function allowedInboxList(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) AllowedInboxList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.AllowedInboxList(&_Bridge.CallOpts, arg0)
}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeCaller) AllowedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeSession) AllowedInboxes(inbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedInboxes(&_Bridge.CallOpts, inbox)
}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeCallerSession) AllowedInboxes(inbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedInboxes(&_Bridge.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_Bridge *BridgeSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.AllowedOutboxList(&_Bridge.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.AllowedOutboxList(&_Bridge.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedOutboxes(&_Bridge.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedOutboxes(&_Bridge.CallOpts, outbox)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeCaller) InboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "inboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Bridge.Contract.InboxAccs(&_Bridge.CallOpts, arg0)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeCallerSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Bridge.Contract.InboxAccs(&_Bridge.CallOpts, arg0)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Bridge *BridgeCaller) MessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "messageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Bridge *BridgeSession) MessageCount() (*big.Int, error) {
	return _Bridge.Contract.MessageCount(&_Bridge.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Bridge *BridgeCallerSession) MessageCount() (*big.Int, error) {
	return _Bridge.Contract.MessageCount(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeTransactor) DeliverMessageToInbox(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deliverMessageToInbox", kind, sender, messageDataHash)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeSession) DeliverMessageToInbox(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DeliverMessageToInbox(&_Bridge.TransactOpts, kind, sender, messageDataHash)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeTransactorSession) DeliverMessageToInbox(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DeliverMessageToInbox(&_Bridge.TransactOpts, kind, sender, messageDataHash)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeTransactor) ExecuteCall(opts *bind.TransactOpts, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeCall", destAddr, amount, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeSession) ExecuteCall(destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteCall(&_Bridge.TransactOpts, destAddr, amount, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeTransactorSession) ExecuteCall(destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteCall(&_Bridge.TransactOpts, destAddr, amount, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactor) SetInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setInbox", inbox, enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeSession) SetInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetInbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactorSession) SetInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetInbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeTransactor) SetOutbox(opts *bind.TransactOpts, outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setOutbox", outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOutbox(&_Bridge.TransactOpts, outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeTransactorSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOutbox(&_Bridge.TransactOpts, outbox, enabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// BridgeMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the Bridge contract.
type BridgeMessageDeliveredIterator struct {
	Event *BridgeMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMessageDelivered)
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
		it.Event = new(BridgeMessageDelivered)
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
func (it *BridgeMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMessageDelivered represents a MessageDelivered event raised by the Bridge contract.
type BridgeMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_Bridge *BridgeFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeMessageDeliveredIterator{contract: _Bridge.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_Bridge *BridgeFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMessageDelivered)
				if err := _Bridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_Bridge *BridgeFilterer) ParseMessageDelivered(log types.Log) (*BridgeMessageDelivered, error) {
	event := new(BridgeMessageDelivered)
	if err := _Bridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCreatorABI is the input ABI used to generate the binding from.
const BridgeCreatorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TemplatesUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adminProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sequencerDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sequencerDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"createBridge\",\"outputs\":[{\"internalType\":\"contractBridge\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractSequencerInbox\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractInbox\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delayedBridgeTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencerInboxTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inboxTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollupEventBridgeTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outboxTemplate\",\"type\":\"address\"}],\"name\":\"updateTemplates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeCreatorBin is the compiled bytecode used for deploying new contracts.
var BridgeCreatorBin = "0x608060405234801561001057600080fd5b50600061001b6101d7565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350604051610071906101db565b604051809103906000f08015801561008d573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b03929092169190911790556040516100ba906101e9565b604051809103906000f0801580156100d6573d6000803e3d6000fd5b50600280546001600160a01b0319166001600160a01b0392909216919091179055604051610103906101f7565b604051809103906000f08015801561011f573d6000803e3d6000fd5b50600380546001600160a01b0319166001600160a01b039290921691909117905560405161014c90610205565b604051809103906000f080158015610168573d6000803e3d6000fd5b50600480546001600160a01b0319166001600160a01b039290921691909117905560405161019590610213565b604051809103906000f0801580156101b1573d6000803e3d6000fd5b50600580546001600160a01b0319166001600160a01b0392909216919091179055610221565b3390565b610f1180620018e683390190565b61157a80620027f783390190565b610ce58062003d7183390190565b610ab48062004a5683390190565b611864806200550a83390190565b6116b580620002316000396000f3fe60806040523480156200001157600080fd5b50600436106200005e5760003560e01c80632147e58e1462000063578063715018a614620000b05780638da5cb5b14620000ba578063ed33557914620000e0578063f2fde38b1462000161575b600080fd5b620000ae600480360360a08110156200007b57600080fd5b506001600160a01b038135811691602081013582169160408201358116916060810135821691608090910135166200018a565b005b620000ae62000278565b620000c462000319565b604080516001600160a01b039092168252519081900360200190f35b62000125600480360360a0811015620000f857600080fd5b506001600160a01b0381358116916020810135821691604082013516906060810135906080013562000328565b604080516001600160a01b0396871681529486166020860152928516848401529084166060840152909216608082015290519081900360a00190f35b620000ae600480360360208110156200017957600080fd5b50356001600160a01b031662000865565b620001946200095e565b6001600160a01b0316620001a762000319565b6001600160a01b031614620001f2576040805162461bcd60e51b8152602060048201819052602482015260008051602062001640833981519152604482015290519081900360640190fd5b600180546001600160a01b03199081166001600160a01b0388811691909117909255600280548216878416179055600380548216868416179055600480548216858416179055600580549091169183169190911790556040517fc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b90600090a15050505050565b620002826200095e565b6001600160a01b03166200029562000319565b6001600160a01b031614620002e0576040805162461bcd60e51b8152602060048201819052602482015260008051602062001640833981519152604482015290519081900360640190fd5b600080546040516001600160a01b039091169060008051602062001660833981519152908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b60008060008060006200033a62000962565b6001546040516001600160a01b03909116908c90620003599062000997565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f0801580156200039d573d6000803e3d6000fd5b506001600160a01b0390811660208301526002546040519116908c90620003c49062000997565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f08015801562000408573d6000803e3d6000fd5b506001600160a01b0390811660408084019190915260035490519116908c90620004329062000997565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f08015801562000476573d6000803e3d6000fd5b506001600160a01b0390811660608301526004546040519116908c906200049d9062000997565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f080158015620004e1573d6000803e3d6000fd5b506001600160a01b0390811660808301526005546040519116908c90620005089062000997565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f0801580156200054c573d6000803e3d6000fd5b506001600160a01b0390811660a08301526040808301516020840151825163eb990c5960e01b815290841660048201528c84166024820152604481018c9052606481018b9052915192169163eb990c599160848082019260009290919082900301818387803b158015620005bf57600080fd5b505af1158015620005d4573d6000803e3d6000fd5b5050505080606001516001600160a01b031663c4d66de882602001516040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156200063057600080fd5b505af115801562000645573d6000803e3d6000fd5b5050505080608001516001600160a01b031663485cc95582602001518c6040518363ffffffff1660e01b815260040180836001600160a01b03168152602001826001600160a01b0316815260200192505050600060405180830381600087803b158015620006b257600080fd5b505af1158015620006c7573d6000803e3d6000fd5b505050508060a001516001600160a01b031663485cc9558b83602001516040518363ffffffff1660e01b815260040180836001600160a01b03168152602001826001600160a01b0316815260200192505050600060405180830381600087803b1580156200073457600080fd5b505af115801562000749573d6000803e3d6000fd5b5050505080602001516001600160a01b031663e45b7ce6826060015160016040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015620007b057600080fd5b505af1158015620007c5573d6000803e3d6000fd5b5050505080602001516001600160a01b031663f2fde38b8b6040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156200081d57600080fd5b505af115801562000832573d6000803e3d6000fd5b50505050602081015160408201516060830151608084015160a090940151929e919d509b50919950975095505050505050565b6200086f6200095e565b6001600160a01b03166200088262000319565b6001600160a01b031614620008cd576040805162461bcd60e51b8152602060048201819052602482015260008051602062001640833981519152604482015290519081900360640190fd5b6001600160a01b038116620009145760405162461bcd60e51b81526004018080602001828103825260268152602001806200161a6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216916000805160206200166083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b610c7480620009a68339019056fe608060405260405162000c7438038062000c74833981810160405260608110156200002957600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200005557600080fd5b9083019060208201858111156200006b57600080fd5b82516401000000008111828201881017156200008657600080fd5b82525081516020918201929091019080838360005b83811015620000b55781810151838201526020016200009b565b50505050905090810190601f168015620000e35780820380516001836020036101000a031916815260200191505b5060405250849150829050620000f98262000137565b8051156200011a57620001188282620001ae60201b620003821760201c565b505b50620001239050565b6200012e82620001dd565b505050620003bf565b6200014d816200020160201b620003ae1760201c565b6200018a5760405162461bcd60e51b815260040180806020018281038252603681526020018062000c186036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b6060620001d6838360405180606001604052806027815260200162000bf16027913962000207565b9392505050565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b3b151590565b6060620002148462000201565b620002515760405162461bcd60e51b815260040180806020018281038252602681526020018062000c4e6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b60208310620002915780518252601f19909201916020918201910162000270565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114620002f3576040519150601f19603f3d011682016040523d82523d6000602084013e620002f8565b606091505b5090925090506200030b82828662000315565b9695505050505050565b6060831562000326575081620001d6565b825115620003375782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200038357818101518382015260200162000369565b50505050905090810190601f168015620003b15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b61082280620003cf6000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100985780635c60da1b146101165780638f28397014610147578063f851a4401461017a5761005d565b3661005d5761005b61018f565b005b61005b61018f565b34801561007157600080fd5b5061005b6004803603602081101561008857600080fd5b50356001600160a01b03166101a9565b61005b600480360360408110156100ae57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156100d857600080fd5b8201836020820111156100ea57600080fd5b803590602001918460018302840111600160201b8311171561010b57600080fd5b5090925090506101e3565b34801561012257600080fd5b5061012b610260565b604080516001600160a01b039092168252519081900360200190f35b34801561015357600080fd5b5061005b6004803603602081101561016a57600080fd5b50356001600160a01b031661029d565b34801561018657600080fd5b5061012b610357565b6101976103b4565b6101a76101a2610414565b610427565b565b6101b161044b565b6001600160a01b0316336001600160a01b031614156101d8576101d38161045e565b6101e0565b6101e061018f565b50565b6101eb61044b565b6001600160a01b0316336001600160a01b031614156102535761020d8361045e565b61024d8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038292505050565b5061025b565b61025b61018f565b505050565b600061026a61044b565b6001600160a01b0316336001600160a01b031614156102925761028b610414565b905061029a565b61029a61018f565b90565b6102a561044b565b6001600160a01b0316336001600160a01b031614156101d8576001600160a01b0381166103035760405162461bcd60e51b815260040180806020018281038252603a8152602001806106ce603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61032c61044b565b604080516001600160a01b03928316815291841660208301528051918290030190a16101d38161049e565b600061036161044b565b6001600160a01b0316336001600160a01b031614156102925761028b61044b565b60606103a78383604051806060016040528060278152602001610728602791396104b0565b9392505050565b3b151590565b6103bc61044b565b6001600160a01b0316336001600160a01b0316141561040c5760405162461bcd60e51b81526004018080602001828103825260428152602001806107ab6042913960600191505060405180910390fd5b6101a76101a7565b6000805160206107088339815191525490565b3660008037600080366000845af43d6000803e808015610446573d6000f35b3d6000fd5b6000805160206106ae8339815191525490565b610467816105b3565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6000805160206106ae83398151915255565b60606104bb846103ae565b6104f65760405162461bcd60e51b81526004018080602001828103825260268152602001806107856026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106105345780518252601f199092019160209182019101610515565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610594576040519150601f19603f3d011682016040523d82523d6000602084013e610599565b606091505b50915091506105a9828286610609565b9695505050505050565b6105bc816103ae565b6105f75760405162461bcd60e51b815260040180806020018281038252603681526020018061074f6036913960400191505060405180910390fd5b60008051602061070883398151915255565b606083156106185750816103a7565b8251156106285782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561067257818101518382015260200161065a565b50505050905090810190601f16801561069f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfeb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f2061646472657373360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a264697066735822122081df10ca698e92e05c31196844f9e9c309ede61eb7e45463098c415a459a179564736f6c634300060c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163744f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122093b91a2c8e78d2f071f20c2beb658187ea0dddf500476c58c6578dfb8c8b93d864736f6c634300060c0033608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b610e948061007d6000396000f3fe6080604052600436106100ad5760003560e01c806302bbfad1146100b25780633dbcc8d1146100f9578063413b35bd1461010e578063715018a6146101555780637ee943291461016c5780638da5cb5b146101b2578063945e1147146101c75780639e5d4c49146101f1578063ab5d894314610302578063c29372de14610317578063cee3d7281461034a578063d9dd67ab14610385578063e45b7ce6146103af578063f2fde38b146103ea575b600080fd5b6100e7600480360360608110156100c857600080fd5b5060ff813516906001600160a01b03602082013516906040013561041d565b60408051918252519081900360200190f35b34801561010557600080fd5b506100e7610533565b34801561011a57600080fd5b506101416004803603602081101561013157600080fd5b50356001600160a01b0316610539565b604080519115158252519081900360200190f35b34801561016157600080fd5b5061016a61055a565b005b34801561017857600080fd5b506101966004803603602081101561018f57600080fd5b50356105f4565b604080516001600160a01b039092168252519081900360200190f35b3480156101be57600080fd5b5061019661061b565b3480156101d357600080fd5b50610196600480360360208110156101ea57600080fd5b503561062a565b3480156101fd57600080fd5b506102816004803603606081101561021457600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561024357600080fd5b82018360208201111561025557600080fd5b803590602001918460018302840111600160201b8311171561027657600080fd5b509092509050610637565b60405180831515815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102c65781810151838201526020016102ae565b50505050905090810190601f1680156102f35780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561030e57600080fd5b5061019661079b565b34801561032357600080fd5b506101416004803603602081101561033a57600080fd5b50356001600160a01b03166107aa565b34801561035657600080fd5b5061016a6004803603604081101561036d57600080fd5b506001600160a01b03813516906020013515156107cc565b34801561039157600080fd5b506100e7600480360360208110156103a857600080fd5b5035610a08565b3480156103bb57600080fd5b5061016a600480360360408110156103d257600080fd5b506001600160a01b0381351690602001351515610a26565b3480156103f657600080fd5b5061016a6004803603602081101561040d57600080fd5b50356001600160a01b0316610c61565b3360009081526001602081905260408220015460ff16610475576040805162461bcd60e51b815260206004820152600e60248201526d09c9ea8be8ca49e9abe929c849eb60931b604482015290519081900360640190fd5b600654600061048986864342863a8a610d51565b9050600082156104b157600660018403815481106104a357fe5b906000526020600020015490505b60066104bd8284610dc2565b8154600181018355600092835260209283902001556040805133815260ff8a16928101929092526001600160a01b038816828201526060820187905251829185917f23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf79181900360800190a3509095945050505050565b60065490565b6001600160a01b031660009081526002602052604090206001015460ff1690565b610562610dee565b6001600160a01b031661057361061b565b6001600160a01b0316146105bc576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1f833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610e3f833981519152908390a3600080546001600160a01b0319169055565b6003818154811061060157fe5b6000918252602090912001546001600160a01b0316905081565b6000546001600160a01b031690565b6004818154811061060157fe5b3360009081526002602052604081206001015460609060ff16610693576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b82156106ee576106ab866001600160a01b0316610df2565b6106ee576040805162461bcd60e51b815260206004820152600f60248201526e1393d7d0d3d11157d05517d11154d5608a1b604482015290519081900360640190fd5b600580546001600160a01b0319811633179091556040516001600160a01b0391821691881690879087908790808383808284376040519201945060009350909150508083038185875af1925050503d8060008114610768576040519150601f19603f3d011682016040523d82523d6000602084013e61076d565b606091505b50600580546001600160a01b0319166001600160a01b03949094169390931790925597909650945050505050565b6005546001600160a01b031681565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b6107d4610dee565b6001600160a01b03166107e561061b565b6001600160a01b03161461082e576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1f833981519152604482015290519081900360640190fd5b6001600160a01b0382166000908152600260205260409020600181015460ff168080156108585750825b8061086a57508015801561086a575082155b15610876575050610a04565b821561090557604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319169091179055610a01565b60048054600019810190811061091757fe5b6000918252602090912001548254600480546001600160a01b0390931692909190811061094057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000015460026000600485600001548154811061098857fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205560048054806109b857fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff191690555b50505b5050565b60068181548110610a1557fe5b600091825260209091200154905081565b610a2e610dee565b6001600160a01b0316610a3f61061b565b6001600160a01b031614610a88576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1f833981519152604482015290519081900360640190fd5b6001600160a01b03821660009081526001602081905260409091209081015460ff16808015610ab45750825b80610ac6575080158015610ac6575082155b15610ad2575050610a04565b8215610b6057604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319169091179055610a01565b600380546000198101908110610b7257fe5b6000918252602090912001548254600380546001600160a01b03909316929091908110610b9b57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600160006003856000015481548110610be357fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556003805480610c1357fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff1916905550505050565b610c69610dee565b6001600160a01b0316610c7a61061b565b6001600160a01b031614610cc3576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1f833981519152604482015290519081900360640190fd5b6001600160a01b038116610d085760405162461bcd60e51b8152600401808060200182810382526026815260200180610df96026913960400191505060405180910390fd5b600080546040516001600160a01b0380851693921691600080516020610e3f83398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6001600160601b0319166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b3390565b3b15159056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220916eba2c20193ad95ae845bfc395ed112a797faf67d09f6796784d3ebf065b8c64736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff1916600117905561154d8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c806306cc91b2146100a95780630a17a46414610130578063342025fa146102025780633dbcc8d11461021c5780635c1bba38146102245780636f791d29146102485780638af00545146102645780639afc500d146102b1578063b71939b114610381578063d9dd67ab14610389578063e367a2c1146103a6578063eb990c59146103ae575b600080fd5b610117600480360360408110156100bf57600080fd5b810190602081018135600160201b8111156100d957600080fd5b8201836020820111156100eb57600080fd5b803590602001918460018302840111600160201b8311171561010c57600080fd5b9193509150356103ea565b6040805192835260208301919091528051918290030190f35b610200600480360360c081101561014657600080fd5b810190602081018135600160201b81111561016057600080fd5b82018360208201111561017257600080fd5b803590602001918460018302840111600160201b8311171561019357600080fd5b919390929091602081019035600160201b8111156101b057600080fd5b8201836020820111156101c257600080fd5b803590602001918460208302840111600160201b831117156101e357600080fd5b91935091508035906020810135906040810135906060013561053c565b005b61020a610636565b60408051918252519081900360200190f35b61020a61063c565b61022c610642565b604080516001600160a01b039092168252519081900360200190f35b610250610651565b604080519115158252519081900360200190f35b610200600480360361010081101561027b57600080fd5b5080359060ff60208201351690604081019060808101359060a0810135906001600160a01b0360c0820135169060e0013561065a565b610200600480360360c08110156102c757600080fd5b810190602081018135600160201b8111156102e157600080fd5b8201836020820111156102f357600080fd5b803590602001918460018302840111600160201b8311171561031457600080fd5b919390929091602081019035600160201b81111561033157600080fd5b82018360208201111561034357600080fd5b803590602001918460208302840111600160201b8311171561036457600080fd5b9193509150803590602081013590604081013590606001356109c9565b61022c610a89565b61020a6004803603602081101561039f57600080fd5b5035610a98565b61020a610ab6565b610200600480360360808110156103c457600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135610abc565b600080826103fd57506000905080610534565b60008061043f87878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509250610b44915050565b90925090506000811561047a5761047188888560018087038154811061046157fe5b9060005260206000200154610bb2565b90935060010190505b60006001838154811061048957fe5b9060005260206000200154905060006104a48a8a8785610bb2565b90955090508288116104eb576040805162461bcd60e51b815260206004820152600b60248201526a10905510d217d4d510549560aa1b604482015290519081900360640190fd5b8088111561052c576040805162461bcd60e51b815260206004820152600960248201526810905510d217d1539160ba1b604482015290519081900360640190fd5b955093505050505b935093915050565b6002546000806105528b8b8b8b8b8b8b8b610d42565b9150915081837f43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063600254878f8f8f8f8f8f8f8c6001808054905003604051808c81526020018b8152602001806020018060200188815260200187815260200186815260200185815260200184815260200183810383528c8c82818152602001925080828437600083820152601f01601f191690910184810383528a8152602090810191508b908b0280828437600083820152604051601f909101601f19169092018290039f50909d5050505050505050505050505050a35050505050505050505050565b60075481565b60025481565b6005546001600160a01b031681565b60005460ff1690565b60035487116106a4576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b60006106ba8784883560208a01358989886111bf565b6006549091504387359091011061070b576040805162461bcd60e51b815260206004820152601060248201526f4d41585f44454c41595f424c4f434b5360801b604482015290519081900360640190fd5b6007544260208801359091011061075a576040805162461bcd60e51b815260206004820152600e60248201526d4d41585f44454c41595f54494d4560901b604482015290519081900360640190fd5b600060018911156107e457600480546040805163d9dd67ab60e01b81526001198d0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b1580156107b557600080fd5b505afa1580156107c9573d6000803e3d6000fd5b505050506040513d60208110156107df57600080fd5b505190505b6107ee8183611230565b600480546040805163d9dd67ab60e01b81526000198e0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561083e57600080fd5b505afa158015610852573d6000803e3d6000fd5b505050506040513d602081101561086857600080fd5b5051146108b2576040805162461bcd60e51b81526020600482015260136024820152722222a620aca2a22fa0a1a1aaa6aaa620aa27a960691b604482015290519081900360640190fd5b5050600254600154600090156108e2576001805460001981019081106108d457fe5b906000526020600020015490505b60008060006108f484868e434261125c565b92509250925060018290806001815401808255809150506001900390600052602060002001600090919091909150558060028190555083857f85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0838f60405180604001604052808881526020018981525060018080549050036040518085815260200184815260200183600260200280838360005b838110156109a0578181015183820152602001610988565b5050505090500182815260200194505050505060405180910390a3505050505050505050505050565b333214610a0b576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600254600080610a218b8b8b8b8b8b8b8b610d42565b6002546001546040805192835260208301899052828101849052600019909101606083015251929450909250839185917f90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812919081900360800190a35050505050505050505050565b6004546001600160a01b031681565b60018181548110610aa557fe5b600091825260209091200154905081565b60065481565b6004546001600160a01b031615610b09576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600480546001600160a01b039586166001600160a01b0319918216179091556005805494909516931692909217909255600691909155600755565b60008082845110158015610b5c575060208385510310155b610b99576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301610ba78585611496565b915091509250929050565b600080600080600080610bfc8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610b44915050565b8094508199505050610c458a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610b44915050565b8093508199505050610c8e8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610b44915050565b809250819950505082828260405160200180807129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b8152506012018481526020018381526020018281526020019350505050604051602081830303815290604052805190602001209250816001019350868314610d32576040805162461bcd60e51b815260206004820152600960248201526842415443485f41434360b81b604482015290519081900360640190fd5b5095989197509095505050505050565b60055460009081906001600160a01b03163314610d97576040805162461bcd60e51b815260206004820152600e60248201526d27a7262cafa9a2a8aaa2a721a2a960911b604482015290519081900360640190fd5b4360065487011015610de0576040805162461bcd60e51b815260206004820152600d60248201526c109313d0d2d7d513d3d7d3d311609a1b604482015290519081900360640190fd5b43861115610e25576040805162461bcd60e51b815260206004820152600d60248201526c424c4f434b5f544f4f5f4e455760981b604482015290519081900360640190fd5b4260075486011015610e6d576040805162461bcd60e51b815260206004820152600c60248201526b1512535157d513d3d7d3d31160a21b604482015290519081900360640190fd5b42851115610eb1576040805162461bcd60e51b815260206004820152600c60248201526b54494d455f544f4f5f4e455760a01b604482015290519081900360640190fd5b600354841015610efc576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b6001841015610f46576040805162461bcd60e51b8152602060048201526011602482015270135554d517d111531056515117d2539255607a1b604482015290519081900360640190fd5b6001600354101580610f56575086155b610fa1576040805162461bcd60e51b8152602060048201526017602482015276135554d517d111531056515117d253925517d4d5105495604a1b604482015290519081900360640190fd5b60015415610fc957600180546000198101908110610fbb57fe5b906000526020600020015491505b600254826000805b8a8110156110dd5760008e838f8f8f86818110610fea57fe5b90506020020135860192611000939291906114ef565b604051808383808284378083019250505092505050604051809103902090506000600390508d8d8481811061103157fe5b9050602002013560001415611044575060065b600061105682338f8f8b6000896111bf565b905085878260405160200180807129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b81525060120184815260200183815260200182815260200193505050506040516020818303038152906040528051906020012095508e8e858181106110ba57fe5b905060200201358501945086806001019750505050508080600101915050610fd1565b506110eb8284898c8c61125c565b600254929650945092508311611136576040805162461bcd60e51b815260206004820152600b60248201526a08a9aa0a8b2be8482a886960ab1b604482015290519081900360640190fd5b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60182905560028390558582146111af576040805162461bcd60e51b815260206004820152600960248201526841465445525f41434360b81b604482015290519081900360640190fd5b5050509850989650505050505050565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6001600160601b0319166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008060008060035487111561148c576004805460408051633dbcc8d160e01b815290516001600160a01b0390921692633dbcc8d1928282019260209290829003018186803b1580156112ae57600080fd5b505afa1580156112c2573d6000803e3d6000fd5b505050506040513d60208110156112d857600080fd5b5051871115611320576040805162461bcd60e51b815260206004820152600f60248201526e2222a620aca2a22faa27a7afa320a960891b604482015290519081900360640190fd5b600480546040805163d9dd67ab60e01b81526000198b0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561137057600080fd5b505afa158015611384573d6000803e3d6000fd5b505050506040513d602081101561139a57600080fd5b505160035460408051702232b630bcb2b21036b2b9b9b0b3b2b99d60791b602082810191909152603182019d909d52605181018c905260718101839052609181018b905260b18082018590528251808303909101815260d190910190915280519b019a909a2060608051608020919b8a03909a019991925060009061142b9060069083908b908b908f9084906111bf565b604080517129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b602080830191909152603282019e909e52605281018d905260728082019390935281518082039093018352609201905280519b019a909a2060038990559950506001909701965b9895505050505050565b600081602001835110156114e6576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b600080858511156114fe578182fd5b8386111561150a578182fd5b505082019391909203915056fea2646970667358221220aeba8c9d33854dd36c0e27f94d18d1ea3dbb28396ca1bc93ea42b2d2dacb9fc364736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff19166001179055610cb88061002d6000396000f3fe6080604052600436106100975760003560e01c80631fe927cf1461009c5780635075788b146101295780635e916758146101cc578063679b6ded1461025657806367ef3ab8146102ff5780636f791d291461038e5780638a631aa6146103b7578063ad9d4ba314610453578063b75436bb14610479578063b8c6a388146104f4578063c4d66de81461052c578063e78cea9214610561575b600080fd5b3480156100a857600080fd5b50610117600480360360208110156100bf57600080fd5b810190602081018135600160201b8111156100d957600080fd5b8201836020820111156100eb57600080fd5b803590602001918460018302840111600160201b8311171561010c57600080fd5b509092509050610592565b60408051918252519081900360200190f35b34801561013557600080fd5b50610117600480360360c081101561014c57600080fd5b8135916020810135916040820135916001600160a01b03606082013516916080820135919081019060c0810160a0820135600160201b81111561018e57600080fd5b8201836020820111156101a057600080fd5b803590602001918460018302840111600160201b831117156101c157600080fd5b509092509050610635565b610117600480360360808110156101e257600080fd5b8135916020810135916001600160a01b036040830135169190810190608081016060820135600160201b81111561021857600080fd5b82018360208201111561022a57600080fd5b803590602001918460018302840111600160201b8311171561024b57600080fd5b5090925090506106bd565b610117600480360361010081101561026d57600080fd5b6001600160a01b038235811692602081013592604082013592606083013581169260808101359091169160a08201359160c081013591810190610100810160e0820135600160201b8111156102c157600080fd5b8201836020820111156102d357600080fd5b803590602001918460018302840111600160201b831117156102f457600080fd5b50909250905061073b565b610117600480360360a081101561031557600080fd5b8135916020810135916040820135916001600160a01b036060820135169181019060a081016080820135600160201b81111561035057600080fd5b82018360208201111561036257600080fd5b803590602001918460018302840111600160201b8311171561038357600080fd5b5090925090506107f7565b34801561039a57600080fd5b506103a361087e565b604080519115158252519081900360200190f35b3480156103c357600080fd5b50610117600480360360a08110156103da57600080fd5b8135916020810135916001600160a01b036040830135169160608101359181019060a081016080820135600160201b81111561041557600080fd5b82018360208201111561042757600080fd5b803590602001918460018302840111600160201b8311171561044857600080fd5b509092509050610887565b6101176004803603602081101561046957600080fd5b50356001600160a01b03166108fb565b34801561048557600080fd5b506101176004803603602081101561049c57600080fd5b810190602081018135600160201b8111156104b657600080fd5b8201836020820111156104c857600080fd5b803590602001918460018302840111600160201b831117156104e957600080fd5b509092509050610958565b6101176004803603608081101561050a57600080fd5b506001600160a01b0381351690602081013590604081013590606001356109db565b34801561053857600080fd5b5061055f6004803603602081101561054f57600080fd5b50356001600160a01b0316610a92565b005b34801561056d57600080fd5b50610576610b0c565b604080516001600160a01b039092168252519081900360200190f35b60003332146105d6576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600061060060033386866040518083838082843760405192018290039091209350610b2092505050565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a29392505050565b60006106b160033360008b8b8b8b60601b60601c6001600160a01b03168b8b8b604051602001808960ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052610bbb565b98975050505050505050565b6000610731600733600189898960601b60601c6001600160a01b0316348a8a604051602001808860ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052610bbb565b9695505050505050565b60006107e96009338c60601b60601c6001600160a01b03168c348d8d60601b60601c6001600160a01b03168d60601b60601c6001600160a01b03168d8d8d8d90508e8e604051602001808c81526020018b81526020018a8152602001898152602001888152602001878152602001868152602001858152602001848152602001838380828437808301925050509b505050505050505050505050604051602081830303815290604052610bbb565b9a9950505050505050505050565b600061087360073360008a8a8a8a60601b60601c6001600160a01b0316348b8b604051602001808960ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052610bbb565b979650505050505050565b60005460ff1690565b600061087360033360018a8a8a60601b60601c6001600160a01b03168a8a8a604051602001808860ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052610bbb565b60408051600160f81b6020820152600060218201819052604182018190526001600160a01b0384166061830152346081808401919091528351808403909101815260a1909201909252610952906007908490610bbb565b92915050565b60008061098360033386866040518083838082843760405192018290039091209350610b2092505050565b905080600080516020610c63833981519152858560405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a29392505050565b6040805163679b6ded60e01b81526001600160a01b0386166004820152346024820152604481018590523360648201819052608482015260a4810184905260c4810183905261010060e4820152600061010482018190529151309163679b6ded9161014480830192602092919082900301818787803b158015610a5d57600080fd5b505af1158015610a71573d6000803e3d6000fd5b505050506040513d6020811015610a8757600080fd5b505195945050505050565b60005461010090046001600160a01b031615610ae4576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600080546001600160a01b0390921661010002610100600160a81b0319909216919091179055565b60005461010090046001600160a01b031681565b60008054604080516302bbfad160e01b815260ff871660048201526001600160a01b038681166024830152604482018690529151610100909304909116916302bbfad1913491606480830192602092919082900301818588803b158015610b8657600080fd5b505af1158015610b9a573d6000803e3d6000fd5b50505050506040513d6020811015610bb157600080fd5b5051949350505050565b600080610bd085858580519060200120610b20565b905080600080516020610c63833981519152846040518080602001828103825283818151815260200191508051906020019080838360005b83811015610c20578181015183820152602001610c08565b50505050905090810190601f168015610c4d5780820380516001836020036101000a031916815260200191505b509250505060405180910390a294935050505056feff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60ba2646970667358221220488558ba5fa286ea60ec52a8c54eed040a38bb3686c995260bf4769b82979aab64736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff19166001179055610a878061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c806316b9109b1461007d57806330a826b41461009c578063485cc955146100b957806364126c7c146100e75780636f791d29146101135780638b8ca1991461012f578063b0f2af2914610167578063f03c04a514610206575b600080fd5b61009a6004803603602081101561009357600080fd5b5035610232565b005b61009a600480360360208110156100b257600080fd5b50356102b3565b61009a600480360360408110156100cf57600080fd5b506001600160a01b0381358116916020013516610331565b61009a600480360360408110156100fd57600080fd5b50803590602001356001600160a01b03166103b6565b61011b6105dd565b604080519115158252519081900360200190f35b61009a6004803603608081101561014557600080fd5b50803590602081013590604081013590606001356001600160a01b03166105e6565b61009a600480360360e081101561017d57600080fd5b8135916020810135916040820135916060810135916001600160a01b03608083013581169260a08101359091169181019060e0810160c0820135600160201b8111156101c857600080fd5b8201836020820111156101da57600080fd5b803590602001918460018302840111600160201b831117156101fb57600080fd5b509092509050610687565b61009a6004803603604081101561021c57600080fd5b506001600160a01b038135169060200135610879565b6001546001600160a01b0316331461027f576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f81b602082015260218082018490528251808303909101815260419091019091526102b090610913565b50565b6001546001600160a01b03163314610300576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f91b602082015260218082018490528251808303909101815260419091019091526102b090610913565b6001546001600160a01b03161561037e576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b60008054610100600160a81b0319166101006001600160a01b0394851602179055600180546001600160a01b03191691909216179055565b6001546001600160a01b03163314610403576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60015460408051634f0f4aa960e01b81526004810185905290516001600160a01b03909216916000918391634f0f4aa991602480820192602092909190829003018186803b15801561045457600080fd5b505afa158015610468573d6000803e3d6000fd5b505050506040513d602081101561047e57600080fd5b5051604080516348b4573960e11b81526001600160a01b038681166004830152915192935090831691639168ae7291602480820192602092909190829003018186803b1580156104cd57600080fd5b505afa1580156104e1573d6000803e3d6000fd5b505050506040513d60208110156104f757600080fd5b5051610537576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b816001600160a01b0316632b2af0ab856040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561057b57600080fd5b505afa15801561058f573d6000803e3d6000fd5b505060408051600160fa1b6020820152602181018890526001600160a01b0387166041808301919091528251808303909101815260619091019091526105d792509050610913565b50505050565b60005460ff1690565b6001546001600160a01b03163314610633576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600060208201526021810186905260418101859052436061820152608181018490526001600160a01b03831660a1808301919091528251808303909101815260c19091019091526105d790610913565b6001546001600160a01b031633146106d4576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b6060888888888860601b60601c6001600160a01b03168860601b60601c6001600160a01b03168888604051602001808981526020018881526020018781526020018681526020018581526020018481526020018383808284376040805191909301818103601f190182528084526000805483516020808601919091206302bbfad160e01b855260048086015233602486015260448501529551939f50909d5061010090046001600160a01b03169b506302bbfad19a5060648082019a509398509096508690039091019350849250899150889050803b1580156107b657600080fd5b505af11580156107ca573d6000803e3d6000fd5b505050506040513d60208110156107e057600080fd5b505160408051602080825285518282015285519394508493600080516020610a32833981519152938793928392918301919085019080838360005b8381101561083357818101518382015260200161081b565b50505050905090810190601f1680156108605780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050505050505050565b6001546001600160a01b031633146108c6576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600360f81b60208201526001600160a01b0384166021820152604181018390524360618083019190915282518083039091018152608190910190915261090f90610913565b5050565b600080548251602080850191909120604080516302bbfad160e01b8152600860048201523360248201526044810192909252516101009093046001600160a01b0316936302bbfad193606480840194939192918390030190829087803b15801561097c57600080fd5b505af1158015610990573d6000803e3d6000fd5b505050506040513d60208110156109a657600080fd5b50516040805160208082528451828201528451600080516020610a32833981519152938693928392918301919085019080838360005b838110156109f45781810151838201526020016109dc565b50505050905090810190601f168015610a215780820380516001836020036101000a031916815260200191505b509250505060405180910390a25056feff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60ba26469706673582212208dc0296bd65615ead206d527d65af5d42dee111320c869c9002172ca197bed1164736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff191660011790556118378061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80627436d3146100a857806305d3efe6146101605780630c726847146101685780634654779014610228578063485cc955146102305780636d5161ec1461025e5780636f791d291461029757806380648b02146102b35780638515bc6a146102bb5780639c5cfe0b146102c35780639f0c04bf146103bf578063b0f305371461045e575b600080fd5b61014e600480360360608110156100be57600080fd5b810190602081018135600160201b8111156100d857600080fd5b8201836020820111156100ea57600080fd5b803590602001918460208302840111600160201b8311171561010b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505082359350505060200135610466565b60408051918252519081900360200190f35b61014e6104a1565b6102266004803603604081101561017e57600080fd5b810190602081018135600160201b81111561019857600080fd5b8201836020820111156101aa57600080fd5b803590602001918460018302840111600160201b831117156101cb57600080fd5b919390929091602081019035600160201b8111156101e857600080fd5b8201836020820111156101fa57600080fd5b803590602001918460208302840111600160201b8311171561021b57600080fd5b5090925090506104a7565b005b61014e610593565b6102266004803603604081101561024657600080fd5b506001600160a01b03813581169160200135166105a2565b61027b6004803603602081101561027457600080fd5b5035610680565b604080516001600160a01b039092168252519081900360200190f35b61029f6106a7565b604080519115158252519081900360200190f35b61027b6106b0565b61014e6106bf565b61022660048036036101408110156102da57600080fd5b81359190810190604081016020820135600160201b8111156102fb57600080fd5b82018360208201111561030d57600080fd5b803590602001918460208302840111600160201b8311171561032e57600080fd5b919390928235926001600160a01b03602082013581169360408301359091169260608301359260808101359260a08201359260c08301359261010081019060e00135600160201b81111561038157600080fd5b82018360208201111561039357600080fd5b803590602001918460018302840111600160201b831117156103b457600080fd5b5090925090506106d5565b61014e600480360360e08110156103d557600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561042057600080fd5b82018360208201111561043257600080fd5b803590602001918460018302840111600160201b8311171561045357600080fd5b509092509050610898565b61014e610935565b60006104998484846040516020018082815260200191505060405160208183030381529060405280519060200120610944565b949350505050565b60035490565b60005461010090046001600160a01b031633146104f9576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b806000805b8281101561058a5761056887838888888681811061051857fe5b9050602002013586019261052e93929190611304565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a1292505050565b84848281811061057457fe5b60200291909101359290920191506001016104fe565b50505050505050565b6005546001600160801b031690565b60005461010090046001600160a01b0316156105f4576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b60008054610100600160a81b0319166101006001600160a01b038581169190910291909117909155600180546001600160a01b03191691831691909117905560405161063f906112f7565b604051809103906000f08015801561065b573d6000803e3d6000fd5b50600280546001600160a01b0319166001600160a01b03929092169190911790555050565b6003818154811061068d57fe5b6000918252602090912001546001600160a01b0316905081565b60005460ff1690565b6004546001600160a01b031690565b600554600160801b90046001600160801b031690565b60006106e78989898989898989610898565b905061072a8d8d8d808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508f9250869150610bc39050565b8c896001600160a01b0316896001600160a01b03167f20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab189648d6040518082815260200191505060405180910390a46004805460058054600680546001600160a01b038f81166001600160a01b03198716179096556001600160801b038c8116600160801b9081028f83166001600160801b0319808816919091178416919091179096558c821695831695909517909255604080516020601f8b018190048102820181019092528981529690951695828416959490930482169391169161082c918e918b918b908b9081908401838280828437600092019190915250610e1c92505050565b600480546001600160a01b03959095166001600160a01b031990951694909417909355600580546001600160801b03928316600160801b029383166001600160801b03199182161783169390931790556006805491909316911617905550505050505050505050505050565b600060038960601b60601c6001600160a01b03168960601b60601c6001600160a01b0316898989898989604051602001808a60ff1660f81b815260010189815260200188815260200187815260200186815260200185815260200184815260200183838082843780830192505050995050505050505050505060405160208183030381529060405280519060200120905098975050505050505050565b6006546001600160801b031690565b825160009061010081111561095857600080fd5b8260005b82811015610a0857600286066109b55786818151811061097857fe5b60200260200101518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091506109fa565b818782815181106109c257fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b60028604955060010161095c565b5095945050505050565b805160009082908290610a2157fe5b01602001516001600160f81b0319161415610bc0578051606114610a79576040805162461bcd60e51b815260206004820152600a6024820152690848288be988a9c8ea8960b31b604482015290519081900360640190fd5b6000610a86826001611039565b90506000610a95836021611039565b90506000610aa4846041611039565b600254909150600090610abf906001600160a01b0316611092565b9050806001600160a01b0316635b36c66b83856040518363ffffffff1660e01b81526004018083815260200182815260200192505050600060405180830381600087803b158015610b0f57600080fd5b505af1158015610b23573d6000803e3d6000fd5b5050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810180546001600160a01b0386166001600160a01b0319909116179055604080518281526020810187905280820188905290519193508792507fe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131919081900360600190a250505050505b50565b61010083511115610c0c576040805162461bcd60e51b815260206004820152600e60248201526d50524f4f465f544f4f5f4c4f4e4760901b604482015290519081900360640190fd5b825160020a8210610c57576040805162461bcd60e51b815260206004820152601060248201526f1410551217d393d517d352539253505360821b604482015290519081900360640190fd5b6000610c64848484610466565b9050600060038681548110610c7557fe5b6000918252602090912001546001600160a01b0316905080610cca576040805162461bcd60e51b815260206004820152600960248201526809c9ebe9eaaa8849eb60bb1b604482015290519081900360640190fd5b600084865160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012090506000826001600160a01b03166357d61c0b85846040518363ffffffff1660e01b81526004018083815260200182815260200192505050602060405180830381600087803b158015610d4e57600080fd5b505af1158015610d62573d6000803e3d6000fd5b505050506040513d6020811015610d7857600080fd5b5051905080610e1257826001600160a01b03166383197ef06040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610dbc57600080fd5b505af1158015610dd0573d6000803e3d6000fd5b50505050600060038981548110610de357fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b5050505050505050565b600154604051639e5d4c4960e01b81526001600160a01b03858116600483019081526024830186905260606044840181815286516064860152865160009692959490921693639e5d4c49938a938a938a93909160849091019060208501908083838e5b83811015610e97578181015183820152602001610e7f565b50505050905090810190601f168015610ec45780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610ee557600080fd5b505af1158015610ef9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015610f2257600080fd5b815160208301805160405192949293830192919084600160201b821115610f4857600080fd5b908301906020820185811115610f5d57600080fd5b8251600160201b811182820188101715610f7657600080fd5b82525081516020918201929091019080838360005b83811015610fa3578181015183820152602001610f8b565b50505050905090810190601f168015610fd05780820380516001836020036101000a031916815260200191505b50604052505050915091508161103257805115610ff05780518082602001fd5b6040805162461bcd60e51b81526020600482015260126024820152711094925111d157d0d0531317d1905253115160721b604482015290519081900360640190fd5b5050505050565b60008160200183511015611089576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000816110a7816001600160a01b0316611250565b604051806040016040528060188152602001772727afa1a7a72a2920a1aa2fa1a627a722afa6a0a9aa22a960411b815250906111615760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561112657818101518382015260200161110e565b50505050905090810190601f1680156111535780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b15801561119b57600080fd5b505afa1580156111af573d6000803e3d6000fd5b505050506040513d60208110156111c557600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b6020820152906112365760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561112657818101518382015260200161110e565b50611249836001600160a01b031661125a565b9392505050565b803b15155b919050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116611255576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b6104d58061132d83390190565b60008085851115611313578182fd5b8386111561131f578182fd5b505082019391909203915056fe608060405234801561001057600080fd5b506000805460ff191660011790556104a88061002d6000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80635780e4e71461007257806357d61c0b1461008c5780635b36c66b146100af5780636f791d29146100d457806383197ef0146100f05780639db9af81146100f8578063ebf0c71714610115575b600080fd5b61007a61011d565b60408051918252519081900360200190f35b61007a600480360360408110156100a257600080fd5b5080359060200135610123565b6100d2600480360360408110156100c557600080fd5b508035906020013561023b565b005b6100dc610330565b604080519115158252519081900360200190f35b6100d2610339565b6100dc6004803603602081101561010e57600080fd5b503561039a565b61007a6103af565b60025481565b6000805461010090046001600160a01b0316331461017a576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b60008281526003602052604090205460ff16156101ce576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b600154831461020f576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b506000818152600360205260409020805460ff1916600117905560028054600019019081905592915050565b60005461010090046001600160a01b03161561028d576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600154156102d1576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b8161030e576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b60008054610100600160a81b0319163361010002179055600191909155600255565b60005460ff1690565b60005461010090046001600160a01b0316331461038f576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b610398336103b5565b565b60036020526000908152604090205460ff1681565b60015481565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff16156104655760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561042a578181015183820152602001610412565b50505050905090810190601f1680156104575780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316fffea26469706673582212205ecec9da6ebed4eccad9344ae0a522e3085e41b86c9eccc6fb729d9bcb7618e464736f6c634300060c0033a26469706673582212204ccec08ea791b7560390c0551e8cd2460e272073f7af02bb1876e4e89da7ad9964736f6c634300060c0033"

// DeployBridgeCreator deploys a new Ethereum contract, binding an instance of BridgeCreator to it.
func DeployBridgeCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeCreator, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeCreatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeCreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeCreator{BridgeCreatorCaller: BridgeCreatorCaller{contract: contract}, BridgeCreatorTransactor: BridgeCreatorTransactor{contract: contract}, BridgeCreatorFilterer: BridgeCreatorFilterer{contract: contract}}, nil
}

// BridgeCreator is an auto generated Go binding around an Ethereum contract.
type BridgeCreator struct {
	BridgeCreatorCaller     // Read-only binding to the contract
	BridgeCreatorTransactor // Write-only binding to the contract
	BridgeCreatorFilterer   // Log filterer for contract events
}

// BridgeCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeCreatorSession struct {
	Contract     *BridgeCreator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCreatorCallerSession struct {
	Contract *BridgeCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BridgeCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeCreatorTransactorSession struct {
	Contract     *BridgeCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BridgeCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeCreatorRaw struct {
	Contract *BridgeCreator // Generic contract binding to access the raw methods on
}

// BridgeCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCreatorCallerRaw struct {
	Contract *BridgeCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeCreatorTransactorRaw struct {
	Contract *BridgeCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeCreator creates a new instance of BridgeCreator, bound to a specific deployed contract.
func NewBridgeCreator(address common.Address, backend bind.ContractBackend) (*BridgeCreator, error) {
	contract, err := bindBridgeCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeCreator{BridgeCreatorCaller: BridgeCreatorCaller{contract: contract}, BridgeCreatorTransactor: BridgeCreatorTransactor{contract: contract}, BridgeCreatorFilterer: BridgeCreatorFilterer{contract: contract}}, nil
}

// NewBridgeCreatorCaller creates a new read-only instance of BridgeCreator, bound to a specific deployed contract.
func NewBridgeCreatorCaller(address common.Address, caller bind.ContractCaller) (*BridgeCreatorCaller, error) {
	contract, err := bindBridgeCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCreatorCaller{contract: contract}, nil
}

// NewBridgeCreatorTransactor creates a new write-only instance of BridgeCreator, bound to a specific deployed contract.
func NewBridgeCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeCreatorTransactor, error) {
	contract, err := bindBridgeCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCreatorTransactor{contract: contract}, nil
}

// NewBridgeCreatorFilterer creates a new log filterer instance of BridgeCreator, bound to a specific deployed contract.
func NewBridgeCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeCreatorFilterer, error) {
	contract, err := bindBridgeCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeCreatorFilterer{contract: contract}, nil
}

// bindBridgeCreator binds a generic wrapper to an already deployed contract.
func bindBridgeCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeCreator *BridgeCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeCreator.Contract.BridgeCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeCreator *BridgeCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCreator.Contract.BridgeCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeCreator *BridgeCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeCreator.Contract.BridgeCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeCreator *BridgeCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeCreator *BridgeCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeCreator *BridgeCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeCreator.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCreator *BridgeCreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCreator *BridgeCreatorSession) Owner() (common.Address, error) {
	return _BridgeCreator.Contract.Owner(&_BridgeCreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCreator *BridgeCreatorCallerSession) Owner() (common.Address, error) {
	return _BridgeCreator.Contract.Owner(&_BridgeCreator.CallOpts)
}

// CreateBridge is a paid mutator transaction binding the contract method 0xed335579.
//
// Solidity: function createBridge(address adminProxy, address rollup, address sequencer, uint256 sequencerDelayBlocks, uint256 sequencerDelaySeconds) returns(address, address, address, address, address)
func (_BridgeCreator *BridgeCreatorTransactor) CreateBridge(opts *bind.TransactOpts, adminProxy common.Address, rollup common.Address, sequencer common.Address, sequencerDelayBlocks *big.Int, sequencerDelaySeconds *big.Int) (*types.Transaction, error) {
	return _BridgeCreator.contract.Transact(opts, "createBridge", adminProxy, rollup, sequencer, sequencerDelayBlocks, sequencerDelaySeconds)
}

// CreateBridge is a paid mutator transaction binding the contract method 0xed335579.
//
// Solidity: function createBridge(address adminProxy, address rollup, address sequencer, uint256 sequencerDelayBlocks, uint256 sequencerDelaySeconds) returns(address, address, address, address, address)
func (_BridgeCreator *BridgeCreatorSession) CreateBridge(adminProxy common.Address, rollup common.Address, sequencer common.Address, sequencerDelayBlocks *big.Int, sequencerDelaySeconds *big.Int) (*types.Transaction, error) {
	return _BridgeCreator.Contract.CreateBridge(&_BridgeCreator.TransactOpts, adminProxy, rollup, sequencer, sequencerDelayBlocks, sequencerDelaySeconds)
}

// CreateBridge is a paid mutator transaction binding the contract method 0xed335579.
//
// Solidity: function createBridge(address adminProxy, address rollup, address sequencer, uint256 sequencerDelayBlocks, uint256 sequencerDelaySeconds) returns(address, address, address, address, address)
func (_BridgeCreator *BridgeCreatorTransactorSession) CreateBridge(adminProxy common.Address, rollup common.Address, sequencer common.Address, sequencerDelayBlocks *big.Int, sequencerDelaySeconds *big.Int) (*types.Transaction, error) {
	return _BridgeCreator.Contract.CreateBridge(&_BridgeCreator.TransactOpts, adminProxy, rollup, sequencer, sequencerDelayBlocks, sequencerDelaySeconds)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeCreator *BridgeCreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeCreator *BridgeCreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeCreator.Contract.RenounceOwnership(&_BridgeCreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeCreator *BridgeCreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeCreator.Contract.RenounceOwnership(&_BridgeCreator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeCreator *BridgeCreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeCreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeCreator *BridgeCreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeCreator.Contract.TransferOwnership(&_BridgeCreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeCreator *BridgeCreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeCreator.Contract.TransferOwnership(&_BridgeCreator.TransactOpts, newOwner)
}

// UpdateTemplates is a paid mutator transaction binding the contract method 0x2147e58e.
//
// Solidity: function updateTemplates(address _delayedBridgeTemplate, address _sequencerInboxTemplate, address _inboxTemplate, address _rollupEventBridgeTemplate, address _outboxTemplate) returns()
func (_BridgeCreator *BridgeCreatorTransactor) UpdateTemplates(opts *bind.TransactOpts, _delayedBridgeTemplate common.Address, _sequencerInboxTemplate common.Address, _inboxTemplate common.Address, _rollupEventBridgeTemplate common.Address, _outboxTemplate common.Address) (*types.Transaction, error) {
	return _BridgeCreator.contract.Transact(opts, "updateTemplates", _delayedBridgeTemplate, _sequencerInboxTemplate, _inboxTemplate, _rollupEventBridgeTemplate, _outboxTemplate)
}

// UpdateTemplates is a paid mutator transaction binding the contract method 0x2147e58e.
//
// Solidity: function updateTemplates(address _delayedBridgeTemplate, address _sequencerInboxTemplate, address _inboxTemplate, address _rollupEventBridgeTemplate, address _outboxTemplate) returns()
func (_BridgeCreator *BridgeCreatorSession) UpdateTemplates(_delayedBridgeTemplate common.Address, _sequencerInboxTemplate common.Address, _inboxTemplate common.Address, _rollupEventBridgeTemplate common.Address, _outboxTemplate common.Address) (*types.Transaction, error) {
	return _BridgeCreator.Contract.UpdateTemplates(&_BridgeCreator.TransactOpts, _delayedBridgeTemplate, _sequencerInboxTemplate, _inboxTemplate, _rollupEventBridgeTemplate, _outboxTemplate)
}

// UpdateTemplates is a paid mutator transaction binding the contract method 0x2147e58e.
//
// Solidity: function updateTemplates(address _delayedBridgeTemplate, address _sequencerInboxTemplate, address _inboxTemplate, address _rollupEventBridgeTemplate, address _outboxTemplate) returns()
func (_BridgeCreator *BridgeCreatorTransactorSession) UpdateTemplates(_delayedBridgeTemplate common.Address, _sequencerInboxTemplate common.Address, _inboxTemplate common.Address, _rollupEventBridgeTemplate common.Address, _outboxTemplate common.Address) (*types.Transaction, error) {
	return _BridgeCreator.Contract.UpdateTemplates(&_BridgeCreator.TransactOpts, _delayedBridgeTemplate, _sequencerInboxTemplate, _inboxTemplate, _rollupEventBridgeTemplate, _outboxTemplate)
}

// BridgeCreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeCreator contract.
type BridgeCreatorOwnershipTransferredIterator struct {
	Event *BridgeCreatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeCreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCreatorOwnershipTransferred)
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
		it.Event = new(BridgeCreatorOwnershipTransferred)
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
func (it *BridgeCreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCreatorOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeCreator contract.
type BridgeCreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeCreator *BridgeCreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeCreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeCreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCreatorOwnershipTransferredIterator{contract: _BridgeCreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeCreator *BridgeCreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeCreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeCreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCreatorOwnershipTransferred)
				if err := _BridgeCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeCreator *BridgeCreatorFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeCreatorOwnershipTransferred, error) {
	event := new(BridgeCreatorOwnershipTransferred)
	if err := _BridgeCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCreatorTemplatesUpdatedIterator is returned from FilterTemplatesUpdated and is used to iterate over the raw logs and unpacked data for TemplatesUpdated events raised by the BridgeCreator contract.
type BridgeCreatorTemplatesUpdatedIterator struct {
	Event *BridgeCreatorTemplatesUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeCreatorTemplatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCreatorTemplatesUpdated)
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
		it.Event = new(BridgeCreatorTemplatesUpdated)
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
func (it *BridgeCreatorTemplatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCreatorTemplatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCreatorTemplatesUpdated represents a TemplatesUpdated event raised by the BridgeCreator contract.
type BridgeCreatorTemplatesUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTemplatesUpdated is a free log retrieval operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_BridgeCreator *BridgeCreatorFilterer) FilterTemplatesUpdated(opts *bind.FilterOpts) (*BridgeCreatorTemplatesUpdatedIterator, error) {

	logs, sub, err := _BridgeCreator.contract.FilterLogs(opts, "TemplatesUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeCreatorTemplatesUpdatedIterator{contract: _BridgeCreator.contract, event: "TemplatesUpdated", logs: logs, sub: sub}, nil
}

// WatchTemplatesUpdated is a free log subscription operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_BridgeCreator *BridgeCreatorFilterer) WatchTemplatesUpdated(opts *bind.WatchOpts, sink chan<- *BridgeCreatorTemplatesUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeCreator.contract.WatchLogs(opts, "TemplatesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCreatorTemplatesUpdated)
				if err := _BridgeCreator.contract.UnpackLog(event, "TemplatesUpdated", log); err != nil {
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

// ParseTemplatesUpdated is a log parse operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_BridgeCreator *BridgeCreatorFilterer) ParseTemplatesUpdated(log types.Log) (*BridgeCreatorTemplatesUpdated, error) {
	event := new(BridgeCreatorTemplatesUpdated)
	if err := _BridgeCreator.contract.UnpackLog(event, "TemplatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC20RollupABI is the input ABI used to generate the binding from.
const ERC20RollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterSendAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"NodeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentNodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterInboxBatchEndCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterInboxBatchAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"indexed\":false,\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"NodeRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startNode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endNode\",\"type\":\"uint256\"}],\"name\":\"NodesDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNode\",\"type\":\"uint256\"}],\"name\":\"StakerReassigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"addToDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeSendAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"executionHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proposedTimes\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"maxMessageCounts\",\"type\":\"uint256[2]\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedBridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraChallengeTimeBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[6]\",\"name\":\"connectedContracts\",\"type\":\"address[6]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"newStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"requireUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireUnresolvedExists\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstUnresolvedNodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestNodeCreated\",\"type\":\"uint256\"}],\"name\":\"requiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerBridge\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expectedNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"},{\"internalType\":\"uint256\",\"name\":\"beforeProposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxMaxCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sequencerBatchProof\",\"type\":\"bytes\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC20RollupBin is the compiled bytecode used for deploying new contracts.
var ERC20RollupBin = "0x608060405234801561001057600080fd5b506000805460ff19908116600117909155600b80549091169055615536806100396000396000f3fe608060405234801561001057600080fd5b50600436106102935760003560e01c8063046f7da21461029857806304a28064146102a25780631e83d30f146102da5780632b2af0ab146102f75780632e7acfa6146103145780632f30cabd1461031c5780633e55c0c7146103425780633e96576e146103665780633fe386271461038c578063414f23fe1461041c57806345e38b641461043f578063488ed1a9146104475780634d26732d146104755780634f0f4aa91461047d57806351ed6a301461049a578063567ca41b146104a25780635c975abb146104c85780635dbaf68b146104e45780635e8ef106146104ec5780636177fd18146104f457806362a82d7d1461051a57806363721d6b1461053757806365f7f80d1461053f57806367425daf1461054757806369fd251c1461054f5780636b94c33b146105755780636f791d291461059b5780636f7d0026146105a35780637427be51146105cc57806376e7e23b146105f2578063771b2f97146105fa5780637ba9534a146106025780637e2d21551461060a57806381fbc98a1461062d5780638456cb59146106535780638640ce5f1461065b5780638da5cb5b1461066357806391c657e81461066b5780639e8a713f14610691578063ce11e6ab14610699578063d01e6602146106a1578063d735e21d146106be578063d93fe9c4146106c6578063dff69787146106ce578063e45b7ce6146106d6578063e8bd492214610704578063edfd03ed14610760578063ef40a6701461077d578063f31d863f146107a3578063f33e1fac14610874578063f3f0a03e14610891578063f51de41b146108bd578063f8d1f194146108c5578063fa7803e6146108e2578063fb64884e14610910578063fdaf57971461092d578063ff204f3b146109d4575b600080fd5b6102a06109fa565b005b6102c8600480360360208110156102b857600080fd5b50356001600160a01b0316610a67565b60408051918252519081900360200190f35b6102a0600480360360208110156102f057600080fd5b5035610b1f565b6102a06004803603602081101561030d57600080fd5b5035610b97565b6102c8610c33565b6102c86004803603602081101561033257600080fd5b50356001600160a01b0316610c39565b61034a610c54565b604080516001600160a01b039092168252519081900360200190f35b6102c86004803603602081101561037c57600080fd5b50356001600160a01b0316610c63565b6102a060048036036102408110156103a357600080fd5b813591602081019160e08201916101e08101359161020082013591908101906102408101610220820135600160201b8111156103de57600080fd5b8201836020820111156103f057600080fd5b803590602001918460018302840111600160201b8311171561041157600080fd5b509092509050610c81565b6102a06004803603604081101561043257600080fd5b5080359060200135611539565b6102c8611704565b6102a0600480360361014081101561045e57600080fd5b50604081016080820160c083016101008401611709565b6102c8611fb6565b61034a6004803603602081101561049357600080fd5b5035611fdb565b61034a611ff6565b6102a0600480360360208110156104b857600080fd5b50356001600160a01b0316612005565b6104d0612126565b604080519115158252519081900360200190f35b61034a61212f565b6102c861213e565b6104d06004803603602081101561050a57600080fd5b50356001600160a01b0316612144565b61034a6004803603602081101561053057600080fd5b503561216c565b6102c8612196565b6102c861219c565b6102a06121a2565b61034a6004803603602081101561056557600080fd5b50356001600160a01b031661220c565b6102a06004803603602081101561058b57600080fd5b50356001600160a01b031661222d565b6104d0612624565b6102c8600480360360608110156105b957600080fd5b508035906020810135906040013561262d565b6102a0600480360360208110156105e257600080fd5b50356001600160a01b0316612644565b6102c86126ef565b6102c86126f5565b6102c86126fb565b6102a06004803603604081101561062057600080fd5b5080359060200135612701565b6102c86004803603602081101561064357600080fd5b50356001600160a01b03166128e1565b6102a0612a03565b6102c8612a57565b61034a612a5d565b6104d06004803603602081101561068157600080fd5b50356001600160a01b0316612a6c565b61034a612ac6565b61034a612ad5565b61034a600480360360208110156106b757600080fd5b5035612ae4565b6102c8612b13565b61034a612b19565b6102c8612b28565b6102a0600480360360408110156106ec57600080fd5b506001600160a01b0381351690602001351515612b2e565b61072a6004803603602081101561071a57600080fd5b50356001600160a01b0316612c03565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b6102a06004803603602081101561077657600080fd5b5035612c3f565b6102c86004803603602081101561079357600080fd5b50356001600160a01b0316612ca4565b6102a0600480360360c08110156107b957600080fd5b81359190810190604081016020820135600160201b8111156107da57600080fd5b8201836020820111156107ec57600080fd5b803590602001918460018302840111600160201b8311171561080d57600080fd5b919390929091602081019035600160201b81111561082a57600080fd5b82018360208201111561083c57600080fd5b803590602001918460208302840111600160201b8311171561085d57600080fd5b919350915080359060208101359060400135612cc2565b6102c86004803603602081101561088a57600080fd5b5035613255565b6102a0600480360360408110156108a757600080fd5b506001600160a01b03813516906020013561327d565b61034a61339a565b6102c8600480360360208110156108db57600080fd5b50356133a9565b6102a0600480360360408110156108f857600080fd5b506001600160a01b03813581169160200135166133bb565b6102a06004803603602081101561092657600080fd5b5035613423565b6102a060048036036101c081101561094457600080fd5b8135916020810135916040820135916060810135916080820135916001600160a01b0360a082013581169260c08301359091169190810190610100810160e0820135600160201b81111561099757600080fd5b8201836020820111156109a957600080fd5b803590602001918460018302840111600160201b831117156109ca57600080fd5b919350915061353b565b6102a0600480360360208110156109ea57600080fd5b50356001600160a01b03166135c9565b6016546001600160a01b03163314610a46576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610a4e61367f565b6040516000805160206154e183398151915290600090a1565b600080610a72612196565b90506000805b82811015610b1557846001600160a01b0316639168ae72610a9883612ae4565b6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610ad557600080fd5b505afa158015610ae9573d6000803e3d6000fd5b505050506040513d6020811015610aff57600080fd5b505115610b0d576001909101905b600101610a78565b509150505b919050565b610b27612126565b15610b67576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b610b703361371f565b6000610b7a611fb6565b905080821015610b88578091505b610b9233836137b6565b505050565b610b9f612b13565b811015610be5576040805162461bcd60e51b815260206004820152600f60248201526e1053149150511657d11150d2511151608a1b604482015290519081900360640190fd5b610bed6126fb565b811115610c30576040805162461bcd60e51b815260206004820152600c60248201526b1113d154d39517d1561254d560a21b604482015290519081900360640190fd5b50565b600c5481565b6001600160a01b03166000908152600a602052604090205490565b6011546001600160a01b031681565b6001600160a01b031660009081526008602052604090206001015490565b610c89612126565b15610cc9576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b610cd233612144565b610d10576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6000610d1b33610c63565b9050610d256153d1565b6000610d3083611fdb565b90506000601160009054906101000a90046001600160a01b03166001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b158015610d8257600080fd5b505afa158015610d96573d6000803e3d6000fd5b505050506040513d6020811015610dac57600080fd5b50519050610db8615406565b60408051808201909152610e72908c60026000835b82821015610e0e5760408051606081810190925290808402860190600390839083908082843760009201919091525050508152600190910190602001610dcd565b50506040805180820190915291508d905060026000835b82821015610e665760408051608081810190925290808402860190600490839083908082843760009201919091525050508152600190910190602001610e25565b505050508b8b86613877565b90506000806000856001600160a01b031663701da98e6040518163ffffffff1660e01b815260040160206040518083038186803b158015610eb257600080fd5b505afa158015610ec6573d6000803e3d6000fd5b505050506040513d6020811015610edc57600080fd5b50518451610ee9906138c5565b14610f2d576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b6011546020850151604090810151815163036648d960e11b81526024810182905260048101928352604481018d90526001600160a01b03909316926306cc91b2928e928e929091908190606401858580828437600081840152601f19601f820116905080830192505050945050505050604080518083038186803b158015610fb457600080fd5b505afa158015610fc8573d6000803e3d6000fd5b505050506040513d6040811015610fde57600080fd5b508051602090910151855160e00151919450925060009061100090439061395a565b9050604b811015611045576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b84515160208601515160009161105b919061395a565b90508560000151610100015186602001516040015110158061108a5750600e546110869083906139b7565b8110155b806110b0575085516060908101516020880151909101516064916110ae919061395a565b145b6110ed576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b61110d6004611107600e54856139b790919063ffffffff16565b906139b7565b81111561114d576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b600e546000906111729061116c61116582600161395a565b8590613a10565b90613a68565b9050611200816111fa611190600c5443613a1090919063ffffffff16565b8c6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156111c957600080fd5b505afa1580156111dd573d6000803e3d6000fd5b505050506040513d60208110156111f357600080fd5b5051613acc565b90613a10565b93506000896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b15801561123d57600080fd5b505afa158015611251573d6000803e3d6000fd5b505050506040513d602081101561126757600080fd5b5051905080156112b9576112b68561127e83611fdb565b6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156111c957600080fd5b94505b5050868660200151604001511115611309576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b5050600080600080896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b15801561134a57600080fd5b505afa15801561135e573d6000803e3d6000fd5b505050506040513d602081101561137457600080fd5b505111905080156113f3576113ec896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b1580156113bb57600080fd5b505afa1580156113cf573d6000803e3d6000fd5b505050506040513d60208110156113e557600080fd5b50516133a9565b91506113ff565b6113fc8b6133a9565b91505b61140e878588888f8787613ae2565b9a50925050508f811461145f576040805162461bcd60e51b81526020600482015260146024820152730aa9c8ab0a08a86a88a88be9c9e888abe9082a6960631b604482015290519081900360640190fd5b6114733361146b6126fb565b600c54613d98565b50505050505050611483826133a9565b61148b6126fb565b7f8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a548b84608001518560400151866000015187602001518f8f6040518088815260200187815260200186815260200185815260200184815260200183600260600280828437600083820152601f01601f191690910190508261010080828437600083820152604051601f909101601f1916909201829003995090975050505050505050a3505050505050505050565b611541612126565b15611581576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b61158a33612144565b6115c8576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b806115d2836133a9565b14611611576040805162461bcd60e51b815260206004820152600a6024820152694e4f44455f52454f524760b01b604482015290519081900360640190fd5b611619612b13565b821015801561162f575061162b6126fb565b8211155b61163857600080fd5b600061164383611fdb565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561167e57600080fd5b505afa158015611692573d6000803e3d6000fd5b505050506040513d60208110156116a857600080fd5b50516116b333610c63565b146116f7576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b610b923384600c54613d98565b604b81565b611711612126565b15611751576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b6020840135843510611798576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b6117a06126fb565b602085013511156117e7576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b83356117f161219c565b10611837576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b600061184985825b6020020135611fdb565b9050600061185886600161183f565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561189357600080fd5b505afa1580156118a7573d6000803e3d6000fd5b505050506040513d60208110156118bd57600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b1580156118ff57600080fd5b505afa158015611913573d6000803e3d6000fd5b505050506040513d602081101561192957600080fd5b505114611969576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b6119838760005b60200201356001600160a01b031661371f565b61198e876001611970565b604080516348b4573960e11b81526001600160a01b03893581166004830152915191841691639168ae7291602480820192602092909190829003018186803b1580156119d957600080fd5b505afa1580156119ed573d6000803e3d6000fd5b505050506040513d6020811015611a0357600080fd5b5051611a4b576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b604080516348b4573960e11b81526001600160a01b0360208a81013582166004840152925190841692639168ae729260248082019391829003018186803b158015611a9557600080fd5b505afa158015611aa9573d6000803e3d6000fd5b505050506040513d6020811015611abf57600080fd5b5051611b07576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b611b1c853585358560005b6020020135613f2b565b826001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b158015611b5557600080fd5b505afa158015611b69573d6000803e3d6000fd5b505050506040513d6020811015611b7f57600080fd5b505114611bc0576040805162461bcd60e51b815260206004820152600a6024820152694348414c5f484153483160b01b604482015290519081900360640190fd5b611bd560208087013590860135856001611b12565b816001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b158015611c0e57600080fd5b505afa158015611c22573d6000803e3d6000fd5b505050506040513d6020811015611c3857600080fd5b505114611c79576040805162461bcd60e51b815260206004820152600a60248201526921a420a62fa420a9a41960b11b604482015290519081900360640190fd5b6000611dca611ceb846001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611cba57600080fd5b505afa158015611cce573d6000803e3d6000fd5b505050506040513d6020811015611ce457600080fd5b5051611fdb565b6001600160a01b031663d7ff5e356040518163ffffffff1660e01b815260040160206040518083038186803b158015611d2357600080fd5b505afa158015611d37573d6000803e3d6000fd5b505050506040513d6020811015611d4d57600080fd5b5051600d546111fa90818960006020020135886001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611d9857600080fd5b505afa158015611dac573d6000803e3d6000fd5b505050506040513d6020811015611dc257600080fd5b50519061395a565b90506020850135811015611e0457611dfc6001600160a01b0389351689600160200201356001600160a01b0316613f62565b505050611faf565b6014546000906001600160a01b0390811690638ecaab119030908a35908935908e35168e600160200201356001600160a01b0316611e5c8d600060028110611e4857fe5b60200201358a61395a90919063ffffffff16565b611e768e600160200201358b61395a90919063ffffffff16565b601154601054604080516001600160e01b031960e08d901b1681526001600160a01b039a8b166004820152602481019990995260448901979097529488166064880152928716608487015260a486019190915260c4850152841660e484015290921661010482015290516101248083019260209291908290030181600087803b158015611f0257600080fd5b505af1158015611f16573d6000803e3d6000fd5b505050506040513d6020811015611f2c57600080fd5b50519050611f556001600160a01b038a35168a600160200201356001600160a01b031683613fdd565b604080518a356001600160a01b0390811682526020808d01358216908301528a35828401529151918316917fa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda758799181900360600190a2505050505b5050505050565b600080611fc1612b13565b9050611fd54382611fd06126fb565b614027565b91505090565b6000908152600560205260409020546001600160a01b031690565b6017546001600160a01b031681565b6016546001600160a01b03163314612051576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6012546001600160a01b03828116911614156120a1576040805162461bcd60e51b815260206004820152600a602482015269086aaa4be9eaaa8849eb60b31b604482015290519081900360640190fd5b601054604080516319dc7ae560e31b81526001600160a01b038481166004830152600060248301819052925193169263cee3d7289260448084019391929182900301818387803b1580156120f457600080fd5b505af1158015612108573d6000803e3d6000fd5b50506040516000805160206154e1833981519152925060009150a150565b600b5460ff1690565b6014546001600160a01b031681565b600e5481565b6001600160a01b0316600090815260086020526040902060030154600160a01b900460ff1690565b60006007828154811061217b57fe5b6000918252602090912001546001600160a01b031692915050565b60095490565b60015490565b60006121ac612b13565b90506121b661219c565b811180156121cb57506121c76126fb565b8111155b610c30576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b6001600160a01b039081166000908152600860205260409020600301541690565b612235612126565b15612275576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b61227d6121a2565b600061228761219c565b90506000612293612b13565b905060006122a082611fdb565b905082816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156122dc57600080fd5b505afa1580156122f0573d6000803e3d6000fd5b505050506040513d602081101561230657600080fd5b505114156125865761231784612144565b612355576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b61236661236185610c63565b610b97565b806001600160a01b0316639168ae72856040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156123b357600080fd5b505afa1580156123c7573d6000803e3d6000fd5b505050506040513d60208110156123dd57600080fd5b505115612424576040805162461bcd60e51b815260206004820152601060248201526f14d51052d15117d3d397d5105491d15560821b604482015290519081900360640190fd5b806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b15801561245d57600080fd5b505afa158015612471573d6000803e3d6000fd5b5050505061247e83611fdb565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b1580156124b657600080fd5b505afa1580156124ca573d6000803e3d6000fd5b505050506124d86000612c3f565b6124e181610a67565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561251a57600080fd5b505afa15801561252e573d6000803e3d6000fd5b505050506040513d602081101561254457600080fd5b505114612586576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b61258e6142ba565b60135460408051630c2a09ad60e21b81526004810185905290516001600160a01b03909216916330a826b49160248082019260009290919082900301818387803b1580156125db57600080fd5b505af11580156125ef573d6000803e3d6000fd5b50506040518492507f9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a39150600090a250505050565b60005460ff1690565b600061263a848484614027565b90505b9392505050565b61264c612126565b1561268c576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b61269461219c565b61269d82610c63565b11156126dd576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b6126e68161371f565b610c30816142d0565b600f5481565b600d5481565b60035490565b612709612126565b15612749576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b612751612196565b821115612796576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b60006127a183612ae4565b905060006127ae84613255565b90506000806127bb612b13565b90505b8083101580156127cd57508482105b156128b95760006127dd84611fdb565b9050806001600160a01b03166396a9fdc0866040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561282e57600080fd5b505af1158015612842573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561287f57600080fd5b505afa158015612893573d6000803e3d6000fd5b505050506040513d60208110156128a957600080fd5b50519350506001909101906127be565b808310156128cf576128ca86614323565b6128d9565b6128d986846143bf565b505050505050565b60006128eb612126565b1561292b576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b6000612936336143e6565b6017546040805163a9059cbb60e01b81526001600160a01b03878116600483015260248201859052915193945091169163a9059cbb916044808201926020929091908290030181600087803b15801561298e57600080fd5b505af11580156129a2573d6000803e3d6000fd5b505050506040513d60208110156129b857600080fd5b50516129fd576040805162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b604482015290519081900360640190fd5b92915050565b6016546001600160a01b03163314612a4f576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610a4e614405565b60045490565b6016546001600160a01b031681565b6000805b600954811015612abd5760098181548110612a8757fe5b60009182526020909120600290910201546001600160a01b0384811691161415612ab5576001915050610b1a565b600101612a70565b50600092915050565b6013546001600160a01b031681565b6012546001600160a01b031681565b600060098281548110612af357fe5b60009182526020909120600290910201546001600160a01b031692915050565b60025490565b6015546001600160a01b031681565b60075490565b6016546001600160a01b03163314612b7a576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6010546040805163722dbe7360e11b81526001600160a01b03858116600483015284151560248301529151919092169163e45b7ce691604480830192600092919082900301818387803b158015612bd057600080fd5b505af1158015612be4573d6000803e3d6000fd5b50506040516000805160206154e1833981519152925060009150a15050565b6008602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6000612c49612196565b90506000612c55612b13565b9050825b82811015612c9e575b81612c6c82613255565b1015612c9657612c7b81614323565b60001990920191828110612c9157505050610c30565b612c62565b600101612c59565b50505050565b6001600160a01b031660009081526008602052604090206002015490565b612cca612126565b15612d0a576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b612d126121a2565b6000612d1c612b28565b11612d5b576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f5354414b45525360b01b604482015290519081900360640190fd5b6000612d65612b13565b90506000612d7282611fdb565b9050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b158015612dad57600080fd5b505afa158015612dc1573d6000803e3d6000fd5b50505050612dcd61219c565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015612e0657600080fd5b505afa158015612e1a573d6000803e3d6000fd5b505050506040513d6020811015612e3057600080fd5b505114612e73576040805162461bcd60e51b815260206004820152600c60248201526b24a72b20a624a22fa82922ab60a11b604482015290519081900360640190fd5b612e83612e7e61219c565b611fdb565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b158015612ebb57600080fd5b505afa158015612ecf573d6000803e3d6000fd5b50505050612edd6000612c3f565b612ef1612ee982610a67565b6111fa612b28565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015612f2a57600080fd5b505afa158015612f3e573d6000803e3d6000fd5b505050506040513d6020811015612f5457600080fd5b505114612f99576040805162461bcd60e51b815260206004820152600e60248201526d1393d517d0531317d4d51052d15160921b604482015290519081900360640190fd5b600061301a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c918291850190849080828437600081840152601f19601f820116905080830192505050505050508d614483565b90506130298b82878988614584565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b15801561306257600080fd5b505afa158015613076573d6000803e3d6000fd5b505050506040513d602081101561308c57600080fd5b5051146130cf576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b60125460408051630c72684760e01b815260048101918252604481018c90526001600160a01b0390921691630c726847918d918d918d918d919081906024810190606401878780828437600083820152601f01601f19169091018481038352858152602090810191508690860280828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b15801561317957600080fd5b505af115801561318d573d6000803e3d6000fd5b505050506131996145cb565b601354604080516316b9109b60e01b81526004810186905290516001600160a01b03909216916316b9109b9160248082019260009290919082900301818387803b1580156131e657600080fd5b505af11580156131fa573d6000803e3d6000fd5b505060408051848152602081018a90528082018990526060810188905290518693507f2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a292509081900360800190a25050505050505050505050565b60006009828154811061326457fe5b9060005260206000209060020201600101549050919050565b613285612126565b156132c5576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b6132cf82826145e4565b601754604080516323b872dd60e01b81523360048201523060248201526044810184905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b15801561332957600080fd5b505af115801561333d573d6000803e3d6000fd5b505050506040513d602081101561335357600080fd5b5051613396576040805162461bcd60e51b815260206004820152600d60248201526c1514905394d1915497d1905253609a1b604482015290519081900360640190fd5b5050565b6010546001600160a01b031681565b60009081526006602052604090205490565b6133c5828261463f565b6001600160a01b0316336001600160a01b031614613419576040805162461bcd60e51b815260206004820152600c60248201526b2ba927a723afa9a2a72222a960a11b604482015290519081900360640190fd5b6133968282613f62565b61342b612126565b1561346b576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b61347481614701565b601754604080516323b872dd60e01b81523360048201523060248201526044810184905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156134ce57600080fd5b505af11580156134e2573d6000803e3d6000fd5b505050506040513d60208110156134f857600080fd5b5051610c30576040805162461bcd60e51b815260206004820152600d60248201526c1514905394d1915497d1905253609a1b604482015290519081900360640190fd5b6017546001600160a01b031661358b576040805162461bcd60e51b815260206004820152601060248201526f2722a2a22fa9aa20a5a2afaa27a5a2a760811b604482015290519081900360640190fd5b61359d8a8a8a8a8a8a8a8a8a8a6148b0565b5050601780546001600160a01b0319166001600160a01b03949094169390931790925550505050505050565b6016546001600160a01b03163314613615576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601280546001600160a01b0319166001600160a01b03838116918217909255601054604080516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b1580156120f457600080fd5b613687612126565b6136cf576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b600b805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613702614c1a565b604080516001600160a01b039092168252519081900360200190a1565b61372881612144565b613766576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b60006137718261220c565b6001600160a01b031614610c30576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b038216600090815260086020526040812060028101548084111561381b576040805162461bcd60e51b815260206004820152601060248201526f544f4f5f4c4954544c455f5354414b4560801b604482015290519081900360640190fd5b6000613827828661395a565b600284018690556001600160a01b0387166000908152600a60205260409020549091506138549082613a10565b6001600160a01b0387166000908152600a60205260409020559250505092915050565b61387f615406565b6040805180820190915286518651829161389a918888614c1e565b81526020016138b9886001602002015188600160200201514387614c1e565b90529695505050505050565b6000816000015182602001518360400151846060015185608001518660a001518760c001518860e00151896101000151604051602001808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019950505050505050505050604051602081830303815290604052805190602001209050919050565b6000828211156139b1576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000826139c6575060006129fd565b828202828482816139d357fe5b041461263d5760405162461bcd60e51b81526004018080602001828103825260218152602001806154a06021913960400191505060405180910390fd5b60008282018381101561263d576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b6000808211613abb576040805162461bcd60e51b815260206004820152601a602482015279536166654d6174683a206469766973696f6e206279207a65726f60301b604482015290519081900360640190fd5b818381613ac457fe5b049392505050565b6000818311613adb578161263d565b5090919050565b6000613aec6153d1565b613af46153d1565b601160009054906101000a90046001600160a01b03166001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b158015613b4257600080fd5b505afa158015613b56573d6000803e3d6000fd5b505050506040513d6020811015613b6c57600080fd5b50516040820152613b7c86611fdb565b6001600160a01b031660a08201526000613b946126fb565b6001019050613ba28b614cbc565b60808301528882526020820188905260135460408051638b8ca19960e01b815260048101849052602481018a9052604481018d905233606482015290516001600160a01b0390921691638b8ca1999160848082019260009290919082900301818387803b158015613c1257600080fd5b505af1158015613c26573d6000803e3d6000fd5b505060155460208e01516001600160a01b03909116925063d45ab2b59150613c4d906138c5565b613c5c8e866080015143614ced565b613c658f614d02565b8b8f6040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b158015613cb957600080fd5b505af1158015613ccd573d6000803e3d6000fd5b505050506040513d6020811015613ce357600080fd5b50516001600160a01b031660608301525060808101516020820151600091613d0e9187918991614d32565b9050613d1e826060015182614d96565b8160a001516001600160a01b0316631bc09d0a613d396126fb565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015613d6f57600080fd5b505af1158015613d83573d6000803e3d6000fd5b50929d939c50929a5050505050505050505050565b6001600160a01b0380841660008181526008602090815260408083208784526005835281842054825163123334b760e11b815260048101969096529151909591909116938492632466696e9260248084019382900301818787803b158015613dff57600080fd5b505af1158015613e13573d6000803e3d6000fd5b505050506040513d6020811015613e2957600080fd5b505160018085018790559091508114156128d957600060056000846001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015613e7c57600080fd5b505afa158015613e90573d6000803e3d6000fd5b505050506040513d6020811015613ea657600080fd5b505181526020810191909152604001600020546001600160a01b0316905080636971dfe5613ed44388613a10565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015613f0a57600080fd5b505af1158015613f1e573d6000803e3d6000fd5b5050505050505050505050565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b6000613f6d82612ca4565b90506000613f7a84612ca4565b905080821115613f9b57613f98613f9184836137b6565b839061395a565b91505b60028204613fa98582614de0565b613fb3838261395a565b9250613fbe85614e11565b601654613fd4906001600160a01b031684614e3b565b611faf84614e7e565b6001600160a01b03928316600090815260086020526040808220600390810180549487166001600160a01b0319958616811790915594909516825290209092018054909216179055565b60008160018403141561403d5750600f5461263d565b600061404884611fdb565b6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561408057600080fd5b505afa158015614094573d6000803e3d6000fd5b505050506040513d60208110156140aa57600080fd5b50519050808510156140c0575050600f5461263d565b6140c861542b565b506040805161014081018252600181526201e05b60208201526201f7d191810191909152620138916060820152620329e160808201526201be4360a08201526204cb8c60c08201526201fbc460e082015262036d326101008201526202797361012082015261413561542b565b506040805161014081018252600181526201c03060208201526201b6999181019190915261fde26060820152620265c6608082015262013b8e60a0820152620329e160c08201526201389160e08201526201f7d16101008201526201537561012082015260006141a5888561395a565b905060006141c3600c5461116c600a856139b790919063ffffffff16565b905060ff6141d282600a613a68565b106141e6576000199550505050505061263d565b60006141f382600a613a68565b60020a9050600085600a8406600a811061420957fe5b602002015162ffffff168202905085600a8406600a811061422657fe5b602002015162ffffff1682828161423957fe5b04146142505760001997505050505050505061263d565b600061427586600a8606600a811061426457fe5b6020020151839062ffffff16613a68565b905080614280575060015b600f54808202908290828161429157fe5b04146142aa57600019995050505050505050505061263d565b9c9b505050505050505050505050565b6142c5600254614f2e565b600280546001019055565b6001600160a01b03811660009081526008602090815260408083206002810154600a90935292205461430191613a10565b6001600160a01b0383166000908152600a602052604090205561339682614fb0565b60098054600019810190811061433557fe5b90600052602060002090600202016009828154811061435057fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600980548061439357fe5b60008281526020812060026000199093019283020180546001600160a01b031916815560010155905550565b80600983815481106143cd57fe5b9060005260206000209060020201600101819055505050565b6001600160a01b03166000908152600a60205260408120805491905590565b61440d612126565b1561444d576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b600b805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613702614c1a565b81518351600091829184835b838110156145365760008882815181106144a557fe5b602002602001015190508381870111156144f5576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868b018101829020604080518084019690965285810191909152805180860382018152606090950190528351930192909220919094019360010161448f565b50818414614579576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b979650505050505050565b60408051602080820197909752808201959095526060850192909252608084019290925260a0808401929092528051808403909201825260c0909201909152805191012090565b6145d6600154614f2e565b600280546001818155019055565b6145ec612126565b1561462c576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b6146358261371f565b6133968282614de0565b6001600160a01b03808316600090815260086020526040808220848416835290822060038083015490820154939492939192908116911681146146b3576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b0381166146f8576040805162461bcd60e51b81526020600482015260076024820152661393d7d0d2105360ca1b604482015290519081900360640190fd5b95945050505050565b614709612126565b15614749576040805162461bcd60e51b815260206004820152601060248201526000805160206154c1833981519152604482015290519081900360640190fd5b61475233612144565b15614795576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b61479e33612a6c565b156147e3576040805162461bcd60e51b815260206004820152601060248201526f5354414b45525f49535f5a4f4d42494560801b604482015290519081900360640190fd5b6147eb611fb6565b811015614832576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b61483c33826150d6565b6013546001600160a01b031663f03c04a53361485661219c565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561489c57600080fd5b505af1158015611faf573d6000803e3d6000fd5b600c54156148f4576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b88614938576040805162461bcd60e51b815260206004820152600f60248201526e10905117d0d3d39197d411549253d1608a1b604482015290519081900360640190fd5b6010805482356001600160a01b039081166001600160a01b03199283161792839055601180546020860135831690841617905560128054909216604080860135831691821790935582516319dc7ae560e31b8152600481019190915260016024820152915192169163cee3d7289160448082019260009290919082900301818387803b1580156149c757600080fd5b505af11580156149db573d6000803e3d6000fd5b50505050806003600681106149ec57fe5b601380546001600160a01b0319166001600160a01b0360209390930293909301358216929092179091556010546040805163722dbe7360e11b8152606085013584166004820152600160248201529051919092169163e45b7ce691604480830192600092919082900301818387803b158015614a6757600080fd5b505af1158015614a7b573d6000803e3d6000fd5b505060135460405163b0f2af2960e01b8152600481018d8152602482018d9052604482018c9052606482018b90526001600160a01b038a8116608484015289811660a484015260e060c4840190815260e484018990529316945063b0f2af2993508d928d928d928d928d928d928d928d9261010401848480828437600081840152601f19601f8201169050808301925050509950505050505050505050600060405180830381600087803b158015614b3257600080fd5b505af1158015614b46573d6000803e3d6000fd5b5050505080600460068110614b5757fe5b601480546001600160a01b03199081166001600160a01b03602094909402949094013583169390931790556015805490921660a08401359091161790556000614b9f8b6151a1565b9050614baa81615298565b600c8a9055600d899055600e889055600f879055601680546001600160a01b0319166001600160a01b038716179055604080518c815290517f4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d916020908290030190a15050505050505050505050565b3390565b614c2661544a565b604080516101208101825285518152865160208201529081018560016020020151815260200185600260048110614c5957fe5b6020020151815260200185600360048110614c7057fe5b6020020151815260200186600160038110614c8757fe5b6020020151815260200186600260038110614c9e57fe5b60200201518152602001848152602001838152509050949350505050565b805180516020830151516000926129fd929182900390614cdb906152e7565b614ce886602001516152e7565b61531c565b600061263a8383866020015160400151613f2b565b805160a09081015160208301519182015160c083015160608401516080909401516000946129fd94939291614584565b60008085614d41576000614d44565b60015b905080858585604051602001808560ff1660f81b815260010184815260200183815260200182815260200194505050505060405160208183030381529060405280519060200120915050949350505050565b60038054600101808255600090815260056020908152604080832080546001600160a01b0319166001600160a01b0397909716969096179095559154815260069091529190912055565b6001600160a01b03821660009081526008602052604090206002810154614e079083613a10565b6002909101555050565b6001600160a01b0316600090815260086020526040902060030180546001600160a01b0319169055565b6001600160a01b0382166000908152600a6020526040902054614e5e9082613a10565b6001600160a01b039092166000908152600a602052604090209190915550565b6001600160a01b0381811660008181526008602090815260408083208151808301909252938152600180850154928201928352600980549182018155909352517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af600290930292830180546001600160a01b031916919095161790935591517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7b09092019190915561339682614fb0565b60008181526005602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b158015614f7a57600080fd5b505af1158015614f8e573d6000803e3d6000fd5b50505060009182525060056020526040902080546001600160a01b0319169055565b6001600160a01b03811660009081526008602052604090208054600780546000198101908110614fdc57fe5b600091825260209091200154600780546001600160a01b03909216918390811061500257fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600860006007848154811061504257fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600780548061507257fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03949094168152600890935250506040812081815560018101829055600281019190915560030180546001600160a81b0319169055565b6007805460018082019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688810180546001600160a01b039586166001600160a01b031991821681179092556040805160a08101825293845284546020858101918252858301978852600060608701818152608088018981529682526008909252929092209451855551948401949094559351600283015591516003909101805492511515600160a01b0260ff60a01b199290951692909316919091171691909117905543600455565b6000806151f7604051806101200160405280600081526020018581526020016000815260200160008152602001600081526020016000801b81526020016000801b815260200143815260200160018152506138c5565b6015546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905243608483015291519394506001600160a01b039092169263d45ab2b59260a4808201936020939283900390910190829087803b15801561526557600080fd5b505af1158015615279573d6000803e3d6000fd5b505050506040513d602081101561528f57600080fd5b50519392505050565b6000805260056020527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc80546001600160a01b0319166001600160a01b03929092169190911790556001600255565b60006129fd8260000151615317846040015185602001518660a0015187606001518860c00151896080015161535a565b6153a5565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b604051806040016040528061541961544a565b815260200161542661544a565b905290565b604051806101400160405280600a906020820280368337509192915050565b604051806101200160405280600081526020016000801916815260200160008152602001600081526020016000815260200160008019168152602001600080191681526020016000815260200160008152509056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775061757361626c653a20706175736564000000000000000000000000000000006eb4661c504ed343cbf8e1e4f8024569117770d3cd0062bfe94b454cf0c48f20a264697066735822122096801c2b2aeb0b78f581f8057d956b0791459bf414aa24934a96e705b472018d64736f6c634300060c0033"

// DeployERC20Rollup deploys a new Ethereum contract, binding an instance of ERC20Rollup to it.
func DeployERC20Rollup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Rollup, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20RollupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20RollupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Rollup{ERC20RollupCaller: ERC20RollupCaller{contract: contract}, ERC20RollupTransactor: ERC20RollupTransactor{contract: contract}, ERC20RollupFilterer: ERC20RollupFilterer{contract: contract}}, nil
}

// ERC20Rollup is an auto generated Go binding around an Ethereum contract.
type ERC20Rollup struct {
	ERC20RollupCaller     // Read-only binding to the contract
	ERC20RollupTransactor // Write-only binding to the contract
	ERC20RollupFilterer   // Log filterer for contract events
}

// ERC20RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20RollupSession struct {
	Contract     *ERC20Rollup      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20RollupCallerSession struct {
	Contract *ERC20RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20RollupTransactorSession struct {
	Contract     *ERC20RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20RollupRaw struct {
	Contract *ERC20Rollup // Generic contract binding to access the raw methods on
}

// ERC20RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20RollupCallerRaw struct {
	Contract *ERC20RollupCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20RollupTransactorRaw struct {
	Contract *ERC20RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Rollup creates a new instance of ERC20Rollup, bound to a specific deployed contract.
func NewERC20Rollup(address common.Address, backend bind.ContractBackend) (*ERC20Rollup, error) {
	contract, err := bindERC20Rollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Rollup{ERC20RollupCaller: ERC20RollupCaller{contract: contract}, ERC20RollupTransactor: ERC20RollupTransactor{contract: contract}, ERC20RollupFilterer: ERC20RollupFilterer{contract: contract}}, nil
}

// NewERC20RollupCaller creates a new read-only instance of ERC20Rollup, bound to a specific deployed contract.
func NewERC20RollupCaller(address common.Address, caller bind.ContractCaller) (*ERC20RollupCaller, error) {
	contract, err := bindERC20Rollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupCaller{contract: contract}, nil
}

// NewERC20RollupTransactor creates a new write-only instance of ERC20Rollup, bound to a specific deployed contract.
func NewERC20RollupTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20RollupTransactor, error) {
	contract, err := bindERC20Rollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupTransactor{contract: contract}, nil
}

// NewERC20RollupFilterer creates a new log filterer instance of ERC20Rollup, bound to a specific deployed contract.
func NewERC20RollupFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20RollupFilterer, error) {
	contract, err := bindERC20Rollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupFilterer{contract: contract}, nil
}

// bindERC20Rollup binds a generic wrapper to an already deployed contract.
func bindERC20Rollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Rollup *ERC20RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Rollup.Contract.ERC20RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Rollup *ERC20RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ERC20RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Rollup *ERC20RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ERC20RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Rollup *ERC20RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Rollup *ERC20RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Rollup *ERC20RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.contract.Transact(opts, method, params...)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_ERC20Rollup *ERC20RollupCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "_stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_ERC20Rollup *ERC20RollupSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _ERC20Rollup.Contract.StakerMap(&_ERC20Rollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_ERC20Rollup *ERC20RollupCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _ERC20Rollup.Contract.StakerMap(&_ERC20Rollup.CallOpts, arg0)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.AmountStaked(&_ERC20Rollup.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.AmountStaked(&_ERC20Rollup.CallOpts, staker)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ArbGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "arbGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _ERC20Rollup.Contract.ArbGasSpeedLimitPerBlock(&_ERC20Rollup.CallOpts)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _ERC20Rollup.Contract.ArbGasSpeedLimitPerBlock(&_ERC20Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) BaseStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "baseStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) BaseStake() (*big.Int, error) {
	return _ERC20Rollup.Contract.BaseStake(&_ERC20Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) BaseStake() (*big.Int, error) {
	return _ERC20Rollup.Contract.BaseStake(&_ERC20Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) ChallengeFactory() (common.Address, error) {
	return _ERC20Rollup.Contract.ChallengeFactory(&_ERC20Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) ChallengeFactory() (common.Address, error) {
	return _ERC20Rollup.Contract.ChallengeFactory(&_ERC20Rollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ConfirmPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "confirmPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _ERC20Rollup.Contract.ConfirmPeriodBlocks(&_ERC20Rollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _ERC20Rollup.Contract.ConfirmPeriodBlocks(&_ERC20Rollup.CallOpts)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) CountStakedZombies(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "countStakedZombies", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.CountStakedZombies(&_ERC20Rollup.CallOpts, node)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.CountStakedZombies(&_ERC20Rollup.CallOpts, node)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_ERC20Rollup *ERC20RollupSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _ERC20Rollup.Contract.CurrentChallenge(&_ERC20Rollup.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _ERC20Rollup.Contract.CurrentChallenge(&_ERC20Rollup.CallOpts, staker)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _ERC20Rollup.Contract.CurrentRequiredStake(&_ERC20Rollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _ERC20Rollup.Contract.CurrentRequiredStake(&_ERC20Rollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) DelayedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "delayedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) DelayedBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.DelayedBridge(&_ERC20Rollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) DelayedBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.DelayedBridge(&_ERC20Rollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ExtraChallengeTimeBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "extraChallengeTimeBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _ERC20Rollup.Contract.ExtraChallengeTimeBlocks(&_ERC20Rollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _ERC20Rollup.Contract.ExtraChallengeTimeBlocks(&_ERC20Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) FirstUnresolvedNode() (*big.Int, error) {
	return _ERC20Rollup.Contract.FirstUnresolvedNode(&_ERC20Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _ERC20Rollup.Contract.FirstUnresolvedNode(&_ERC20Rollup.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_ERC20Rollup *ERC20RollupSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.GetNode(&_ERC20Rollup.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.GetNode(&_ERC20Rollup.CallOpts, nodeNum)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_ERC20Rollup *ERC20RollupCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_ERC20Rollup *ERC20RollupSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _ERC20Rollup.Contract.GetNodeHash(&_ERC20Rollup.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_ERC20Rollup *ERC20RollupCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _ERC20Rollup.Contract.GetNodeHash(&_ERC20Rollup.CallOpts, index)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_ERC20Rollup *ERC20RollupSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.GetStakerAddress(&_ERC20Rollup.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.GetStakerAddress(&_ERC20Rollup.CallOpts, stakerNum)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ERC20Rollup *ERC20RollupCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ERC20Rollup *ERC20RollupSession) IsMaster() (bool, error) {
	return _ERC20Rollup.Contract.IsMaster(&_ERC20Rollup.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ERC20Rollup *ERC20RollupCallerSession) IsMaster() (bool, error) {
	return _ERC20Rollup.Contract.IsMaster(&_ERC20Rollup.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupSession) IsStaked(staker common.Address) (bool, error) {
	return _ERC20Rollup.Contract.IsStaked(&_ERC20Rollup.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _ERC20Rollup.Contract.IsStaked(&_ERC20Rollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupCaller) IsZombie(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "isZombie", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupSession) IsZombie(staker common.Address) (bool, error) {
	return _ERC20Rollup.Contract.IsZombie(&_ERC20Rollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupCallerSession) IsZombie(staker common.Address) (bool, error) {
	return _ERC20Rollup.Contract.IsZombie(&_ERC20Rollup.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) LastStakeBlock() (*big.Int, error) {
	return _ERC20Rollup.Contract.LastStakeBlock(&_ERC20Rollup.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) LastStakeBlock() (*big.Int, error) {
	return _ERC20Rollup.Contract.LastStakeBlock(&_ERC20Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) LatestConfirmed() (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestConfirmed(&_ERC20Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) LatestConfirmed() (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestConfirmed(&_ERC20Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) LatestNodeCreated() (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestNodeCreated(&_ERC20Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestNodeCreated(&_ERC20Rollup.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestStakedNode(&_ERC20Rollup.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestStakedNode(&_ERC20Rollup.CallOpts, staker)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _ERC20Rollup.Contract.MinimumAssertionPeriod(&_ERC20Rollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _ERC20Rollup.Contract.MinimumAssertionPeriod(&_ERC20Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) NodeFactory() (common.Address, error) {
	return _ERC20Rollup.Contract.NodeFactory(&_ERC20Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) NodeFactory() (common.Address, error) {
	return _ERC20Rollup.Contract.NodeFactory(&_ERC20Rollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) Outbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "outbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) Outbox() (common.Address, error) {
	return _ERC20Rollup.Contract.Outbox(&_ERC20Rollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) Outbox() (common.Address, error) {
	return _ERC20Rollup.Contract.Outbox(&_ERC20Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) Owner() (common.Address, error) {
	return _ERC20Rollup.Contract.Owner(&_ERC20Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) Owner() (common.Address, error) {
	return _ERC20Rollup.Contract.Owner(&_ERC20Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Rollup *ERC20RollupCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Rollup *ERC20RollupSession) Paused() (bool, error) {
	return _ERC20Rollup.Contract.Paused(&_ERC20Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Rollup *ERC20RollupCallerSession) Paused() (bool, error) {
	return _ERC20Rollup.Contract.Paused(&_ERC20Rollup.CallOpts)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_ERC20Rollup *ERC20RollupCaller) RequireUnresolved(opts *bind.CallOpts, nodeNum *big.Int) error {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "requireUnresolved", nodeNum)

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_ERC20Rollup *ERC20RollupSession) RequireUnresolved(nodeNum *big.Int) error {
	return _ERC20Rollup.Contract.RequireUnresolved(&_ERC20Rollup.CallOpts, nodeNum)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_ERC20Rollup *ERC20RollupCallerSession) RequireUnresolved(nodeNum *big.Int) error {
	return _ERC20Rollup.Contract.RequireUnresolved(&_ERC20Rollup.CallOpts, nodeNum)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_ERC20Rollup *ERC20RollupCaller) RequireUnresolvedExists(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "requireUnresolvedExists")

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_ERC20Rollup *ERC20RollupSession) RequireUnresolvedExists() error {
	return _ERC20Rollup.Contract.RequireUnresolvedExists(&_ERC20Rollup.CallOpts)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_ERC20Rollup *ERC20RollupCallerSession) RequireUnresolvedExists() error {
	return _ERC20Rollup.Contract.RequireUnresolvedExists(&_ERC20Rollup.CallOpts)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) RequiredStake(opts *bind.CallOpts, blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "requiredStake", blockNumber, firstUnresolvedNodeNum, latestNodeCreated)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _ERC20Rollup.Contract.RequiredStake(&_ERC20Rollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _ERC20Rollup.Contract.RequiredStake(&_ERC20Rollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) RollupEventBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "rollupEventBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) RollupEventBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.RollupEventBridge(&_ERC20Rollup.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) RollupEventBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.RollupEventBridge(&_ERC20Rollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) SequencerBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "sequencerBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) SequencerBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.SequencerBridge(&_ERC20Rollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) SequencerBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.SequencerBridge(&_ERC20Rollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) StakeToken() (common.Address, error) {
	return _ERC20Rollup.Contract.StakeToken(&_ERC20Rollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) StakeToken() (common.Address, error) {
	return _ERC20Rollup.Contract.StakeToken(&_ERC20Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) StakerCount() (*big.Int, error) {
	return _ERC20Rollup.Contract.StakerCount(&_ERC20Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) StakerCount() (*big.Int, error) {
	return _ERC20Rollup.Contract.StakerCount(&_ERC20Rollup.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.WithdrawableFunds(&_ERC20Rollup.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.WithdrawableFunds(&_ERC20Rollup.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_ERC20Rollup *ERC20RollupSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.ZombieAddress(&_ERC20Rollup.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.ZombieAddress(&_ERC20Rollup.CallOpts, zombieNum)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ZombieCount() (*big.Int, error) {
	return _ERC20Rollup.Contract.ZombieCount(&_ERC20Rollup.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ZombieCount() (*big.Int, error) {
	return _ERC20Rollup.Contract.ZombieCount(&_ERC20Rollup.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _ERC20Rollup.Contract.ZombieLatestStakedNode(&_ERC20Rollup.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _ERC20Rollup.Contract.ZombieLatestStakedNode(&_ERC20Rollup.CallOpts, zombieNum)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupTransactor) AddToDeposit(opts *bind.TransactOpts, stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "addToDeposit", stakerAddress, tokenAmount)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupSession) AddToDeposit(stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.AddToDeposit(&_ERC20Rollup.TransactOpts, stakerAddress, tokenAmount)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) AddToDeposit(stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.AddToDeposit(&_ERC20Rollup.TransactOpts, stakerAddress, tokenAmount)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_ERC20Rollup *ERC20RollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "completeChallenge", winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_ERC20Rollup *ERC20RollupSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.CompleteChallenge(&_ERC20Rollup.TransactOpts, winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.CompleteChallenge(&_ERC20Rollup.TransactOpts, winningStaker, losingStaker)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_ERC20Rollup *ERC20RollupTransactor) ConfirmNextNode(opts *bind.TransactOpts, beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "confirmNextNode", beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_ERC20Rollup *ERC20RollupSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ConfirmNextNode(&_ERC20Rollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ConfirmNextNode(&_ERC20Rollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_ERC20Rollup *ERC20RollupTransactor) CreateChallenge(opts *bind.TransactOpts, stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "createChallenge", stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_ERC20Rollup *ERC20RollupSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.CreateChallenge(&_ERC20Rollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.CreateChallenge(&_ERC20Rollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_ERC20Rollup *ERC20RollupTransactor) Initialize(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "initialize", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_ERC20Rollup *ERC20RollupSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Initialize(&_ERC20Rollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Initialize(&_ERC20Rollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupTransactor) NewStake(opts *bind.TransactOpts, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "newStake", tokenAmount)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupSession) NewStake(tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.NewStake(&_ERC20Rollup.TransactOpts, tokenAmount)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) NewStake(tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.NewStake(&_ERC20Rollup.TransactOpts, tokenAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Rollup *ERC20RollupTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Rollup *ERC20RollupSession) Pause() (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Pause(&_ERC20Rollup.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Pause(&_ERC20Rollup.TransactOpts)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_ERC20Rollup *ERC20RollupTransactor) ReduceDeposit(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "reduceDeposit", target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_ERC20Rollup *ERC20RollupSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ReduceDeposit(&_ERC20Rollup.TransactOpts, target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ReduceDeposit(&_ERC20Rollup.TransactOpts, target)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupTransactor) RejectNextNode(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "rejectNextNode", stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RejectNextNode(&_ERC20Rollup.TransactOpts, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RejectNextNode(&_ERC20Rollup.TransactOpts, stakerAddress)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupTransactor) RemoveOldOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "removeOldOutbox", _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveOldOutbox(&_ERC20Rollup.TransactOpts, _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveOldOutbox(&_ERC20Rollup.TransactOpts, _outbox)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_ERC20Rollup *ERC20RollupTransactor) RemoveOldZombies(opts *bind.TransactOpts, startIndex *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "removeOldZombies", startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_ERC20Rollup *ERC20RollupSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveOldZombies(&_ERC20Rollup.TransactOpts, startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveOldZombies(&_ERC20Rollup.TransactOpts, startIndex)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_ERC20Rollup *ERC20RollupTransactor) RemoveZombie(opts *bind.TransactOpts, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "removeZombie", zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_ERC20Rollup *ERC20RollupSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveZombie(&_ERC20Rollup.TransactOpts, zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveZombie(&_ERC20Rollup.TransactOpts, zombieNum, maxNodes)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ERC20Rollup *ERC20RollupTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ERC20Rollup *ERC20RollupSession) Resume() (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Resume(&_ERC20Rollup.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) Resume() (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Resume(&_ERC20Rollup.TransactOpts)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupTransactor) ReturnOldDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "returnOldDeposit", stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ReturnOldDeposit(&_ERC20Rollup.TransactOpts, stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ReturnOldDeposit(&_ERC20Rollup.TransactOpts, stakerAddress)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_ERC20Rollup *ERC20RollupTransactor) SetInbox(opts *bind.TransactOpts, _inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "setInbox", _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_ERC20Rollup *ERC20RollupSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.SetInbox(&_ERC20Rollup.TransactOpts, _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.SetInbox(&_ERC20Rollup.TransactOpts, _inbox, _enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupTransactor) SetOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "setOutbox", _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.SetOutbox(&_ERC20Rollup.TransactOpts, _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.SetOutbox(&_ERC20Rollup.TransactOpts, _outbox)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_ERC20Rollup *ERC20RollupTransactor) StakeOnExistingNode(opts *bind.TransactOpts, nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "stakeOnExistingNode", nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_ERC20Rollup *ERC20RollupSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.StakeOnExistingNode(&_ERC20Rollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.StakeOnExistingNode(&_ERC20Rollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_ERC20Rollup *ERC20RollupTransactor) StakeOnNewNode(opts *bind.TransactOpts, expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "stakeOnNewNode", expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_ERC20Rollup *ERC20RollupSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.StakeOnNewNode(&_ERC20Rollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.StakeOnNewNode(&_ERC20Rollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_ERC20Rollup *ERC20RollupTransactor) WithdrawStakerFunds(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "withdrawStakerFunds", destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.WithdrawStakerFunds(&_ERC20Rollup.TransactOpts, destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_ERC20Rollup *ERC20RollupTransactorSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.WithdrawStakerFunds(&_ERC20Rollup.TransactOpts, destination)
}

// ERC20RollupNodeConfirmedIterator is returned from FilterNodeConfirmed and is used to iterate over the raw logs and unpacked data for NodeConfirmed events raised by the ERC20Rollup contract.
type ERC20RollupNodeConfirmedIterator struct {
	Event *ERC20RollupNodeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20RollupNodeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupNodeConfirmed)
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
		it.Event = new(ERC20RollupNodeConfirmed)
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
func (it *ERC20RollupNodeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupNodeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupNodeConfirmed represents a NodeConfirmed event raised by the ERC20Rollup contract.
type ERC20RollupNodeConfirmed struct {
	NodeNum        *big.Int
	AfterSendAcc   [32]byte
	AfterSendCount *big.Int
	AfterLogAcc    [32]byte
	AfterLogCount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeConfirmed is a free log retrieval operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_ERC20Rollup *ERC20RollupFilterer) FilterNodeConfirmed(opts *bind.FilterOpts, nodeNum []*big.Int) (*ERC20RollupNodeConfirmedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupNodeConfirmedIterator{contract: _ERC20Rollup.contract, event: "NodeConfirmed", logs: logs, sub: sub}, nil
}

// WatchNodeConfirmed is a free log subscription operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_ERC20Rollup *ERC20RollupFilterer) WatchNodeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20RollupNodeConfirmed, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupNodeConfirmed)
				if err := _ERC20Rollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
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

// ParseNodeConfirmed is a log parse operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_ERC20Rollup *ERC20RollupFilterer) ParseNodeConfirmed(log types.Log) (*ERC20RollupNodeConfirmed, error) {
	event := new(ERC20RollupNodeConfirmed)
	if err := _ERC20Rollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the ERC20Rollup contract.
type ERC20RollupNodeCreatedIterator struct {
	Event *ERC20RollupNodeCreated // Event containing the contract specifics and raw log

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
func (it *ERC20RollupNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupNodeCreated)
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
		it.Event = new(ERC20RollupNodeCreated)
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
func (it *ERC20RollupNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupNodeCreated represents a NodeCreated event raised by the ERC20Rollup contract.
type ERC20RollupNodeCreated struct {
	NodeNum                 *big.Int
	ParentNodeHash          [32]byte
	NodeHash                [32]byte
	ExecutionHash           [32]byte
	InboxMaxCount           *big.Int
	AfterInboxBatchEndCount *big.Int
	AfterInboxBatchAcc      [32]byte
	AssertionBytes32Fields  [2][3][32]byte
	AssertionIntFields      [2][4]*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_ERC20Rollup *ERC20RollupFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int, parentNodeHash [][32]byte) (*ERC20RollupNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupNodeCreatedIterator{contract: _ERC20Rollup.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_ERC20Rollup *ERC20RollupFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *ERC20RollupNodeCreated, nodeNum []*big.Int, parentNodeHash [][32]byte) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupNodeCreated)
				if err := _ERC20Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
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

// ParseNodeCreated is a log parse operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_ERC20Rollup *ERC20RollupFilterer) ParseNodeCreated(log types.Log) (*ERC20RollupNodeCreated, error) {
	event := new(ERC20RollupNodeCreated)
	if err := _ERC20Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupNodeRejectedIterator is returned from FilterNodeRejected and is used to iterate over the raw logs and unpacked data for NodeRejected events raised by the ERC20Rollup contract.
type ERC20RollupNodeRejectedIterator struct {
	Event *ERC20RollupNodeRejected // Event containing the contract specifics and raw log

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
func (it *ERC20RollupNodeRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupNodeRejected)
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
		it.Event = new(ERC20RollupNodeRejected)
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
func (it *ERC20RollupNodeRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupNodeRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupNodeRejected represents a NodeRejected event raised by the ERC20Rollup contract.
type ERC20RollupNodeRejected struct {
	NodeNum *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeRejected is a free log retrieval operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_ERC20Rollup *ERC20RollupFilterer) FilterNodeRejected(opts *bind.FilterOpts, nodeNum []*big.Int) (*ERC20RollupNodeRejectedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupNodeRejectedIterator{contract: _ERC20Rollup.contract, event: "NodeRejected", logs: logs, sub: sub}, nil
}

// WatchNodeRejected is a free log subscription operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_ERC20Rollup *ERC20RollupFilterer) WatchNodeRejected(opts *bind.WatchOpts, sink chan<- *ERC20RollupNodeRejected, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupNodeRejected)
				if err := _ERC20Rollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
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

// ParseNodeRejected is a log parse operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_ERC20Rollup *ERC20RollupFilterer) ParseNodeRejected(log types.Log) (*ERC20RollupNodeRejected, error) {
	event := new(ERC20RollupNodeRejected)
	if err := _ERC20Rollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupNodesDestroyedIterator is returned from FilterNodesDestroyed and is used to iterate over the raw logs and unpacked data for NodesDestroyed events raised by the ERC20Rollup contract.
type ERC20RollupNodesDestroyedIterator struct {
	Event *ERC20RollupNodesDestroyed // Event containing the contract specifics and raw log

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
func (it *ERC20RollupNodesDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupNodesDestroyed)
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
		it.Event = new(ERC20RollupNodesDestroyed)
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
func (it *ERC20RollupNodesDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupNodesDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupNodesDestroyed represents a NodesDestroyed event raised by the ERC20Rollup contract.
type ERC20RollupNodesDestroyed struct {
	StartNode *big.Int
	EndNode   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodesDestroyed is a free log retrieval operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_ERC20Rollup *ERC20RollupFilterer) FilterNodesDestroyed(opts *bind.FilterOpts, startNode []*big.Int, endNode []*big.Int) (*ERC20RollupNodesDestroyedIterator, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupNodesDestroyedIterator{contract: _ERC20Rollup.contract, event: "NodesDestroyed", logs: logs, sub: sub}, nil
}

// WatchNodesDestroyed is a free log subscription operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_ERC20Rollup *ERC20RollupFilterer) WatchNodesDestroyed(opts *bind.WatchOpts, sink chan<- *ERC20RollupNodesDestroyed, startNode []*big.Int, endNode []*big.Int) (event.Subscription, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupNodesDestroyed)
				if err := _ERC20Rollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
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

// ParseNodesDestroyed is a log parse operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_ERC20Rollup *ERC20RollupFilterer) ParseNodesDestroyed(log types.Log) (*ERC20RollupNodesDestroyed, error) {
	event := new(ERC20RollupNodesDestroyed)
	if err := _ERC20Rollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the ERC20Rollup contract.
type ERC20RollupOwnerFunctionCalledIterator struct {
	Event *ERC20RollupOwnerFunctionCalled // Event containing the contract specifics and raw log

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
func (it *ERC20RollupOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupOwnerFunctionCalled)
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
		it.Event = new(ERC20RollupOwnerFunctionCalled)
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
func (it *ERC20RollupOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the ERC20Rollup contract.
type ERC20RollupOwnerFunctionCalled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0x6eb4661c504ed343cbf8e1e4f8024569117770d3cd0062bfe94b454cf0c48f20.
//
// Solidity: event OwnerFunctionCalled()
func (_ERC20Rollup *ERC20RollupFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts) (*ERC20RollupOwnerFunctionCalledIterator, error) {

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "OwnerFunctionCalled")
	if err != nil {
		return nil, err
	}
	return &ERC20RollupOwnerFunctionCalledIterator{contract: _ERC20Rollup.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0x6eb4661c504ed343cbf8e1e4f8024569117770d3cd0062bfe94b454cf0c48f20.
//
// Solidity: event OwnerFunctionCalled()
func (_ERC20Rollup *ERC20RollupFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *ERC20RollupOwnerFunctionCalled) (event.Subscription, error) {

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "OwnerFunctionCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupOwnerFunctionCalled)
				if err := _ERC20Rollup.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
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

// ParseOwnerFunctionCalled is a log parse operation binding the contract event 0x6eb4661c504ed343cbf8e1e4f8024569117770d3cd0062bfe94b454cf0c48f20.
//
// Solidity: event OwnerFunctionCalled()
func (_ERC20Rollup *ERC20RollupFilterer) ParseOwnerFunctionCalled(log types.Log) (*ERC20RollupOwnerFunctionCalled, error) {
	event := new(ERC20RollupOwnerFunctionCalled)
	if err := _ERC20Rollup.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20Rollup contract.
type ERC20RollupPausedIterator struct {
	Event *ERC20RollupPaused // Event containing the contract specifics and raw log

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
func (it *ERC20RollupPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupPaused)
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
		it.Event = new(ERC20RollupPaused)
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
func (it *ERC20RollupPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupPaused represents a Paused event raised by the ERC20Rollup contract.
type ERC20RollupPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20RollupPausedIterator, error) {

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20RollupPausedIterator{contract: _ERC20Rollup.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20RollupPaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupPaused)
				if err := _ERC20Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) ParsePaused(log types.Log) (*ERC20RollupPaused, error) {
	event := new(ERC20RollupPaused)
	if err := _ERC20Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the ERC20Rollup contract.
type ERC20RollupRollupChallengeStartedIterator struct {
	Event *ERC20RollupRollupChallengeStarted // Event containing the contract specifics and raw log

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
func (it *ERC20RollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupRollupChallengeStarted)
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
		it.Event = new(ERC20RollupRollupChallengeStarted)
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
func (it *ERC20RollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the ERC20Rollup contract.
type ERC20RollupRollupChallengeStarted struct {
	ChallengeContract common.Address
	Asserter          common.Address
	Challenger        common.Address
	ChallengedNode    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_ERC20Rollup *ERC20RollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts, challengeContract []common.Address) (*ERC20RollupRollupChallengeStartedIterator, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupRollupChallengeStartedIterator{contract: _ERC20Rollup.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_ERC20Rollup *ERC20RollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *ERC20RollupRollupChallengeStarted, challengeContract []common.Address) (event.Subscription, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupRollupChallengeStarted)
				if err := _ERC20Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
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

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_ERC20Rollup *ERC20RollupFilterer) ParseRollupChallengeStarted(log types.Log) (*ERC20RollupRollupChallengeStarted, error) {
	event := new(ERC20RollupRollupChallengeStarted)
	if err := _ERC20Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the ERC20Rollup contract.
type ERC20RollupRollupCreatedIterator struct {
	Event *ERC20RollupRollupCreated // Event containing the contract specifics and raw log

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
func (it *ERC20RollupRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupRollupCreated)
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
		it.Event = new(ERC20RollupRollupCreated)
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
func (it *ERC20RollupRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupRollupCreated represents a RollupCreated event raised by the ERC20Rollup contract.
type ERC20RollupRollupCreated struct {
	MachineHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_ERC20Rollup *ERC20RollupFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*ERC20RollupRollupCreatedIterator, error) {

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &ERC20RollupRollupCreatedIterator{contract: _ERC20Rollup.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_ERC20Rollup *ERC20RollupFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *ERC20RollupRollupCreated) (event.Subscription, error) {

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupRollupCreated)
				if err := _ERC20Rollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_ERC20Rollup *ERC20RollupFilterer) ParseRollupCreated(log types.Log) (*ERC20RollupRollupCreated, error) {
	event := new(ERC20RollupRollupCreated)
	if err := _ERC20Rollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupStakerReassignedIterator is returned from FilterStakerReassigned and is used to iterate over the raw logs and unpacked data for StakerReassigned events raised by the ERC20Rollup contract.
type ERC20RollupStakerReassignedIterator struct {
	Event *ERC20RollupStakerReassigned // Event containing the contract specifics and raw log

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
func (it *ERC20RollupStakerReassignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupStakerReassigned)
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
		it.Event = new(ERC20RollupStakerReassigned)
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
func (it *ERC20RollupStakerReassignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupStakerReassignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupStakerReassigned represents a StakerReassigned event raised by the ERC20Rollup contract.
type ERC20RollupStakerReassigned struct {
	Staker  common.Address
	NewNode *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakerReassigned is a free log retrieval operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_ERC20Rollup *ERC20RollupFilterer) FilterStakerReassigned(opts *bind.FilterOpts, staker []common.Address) (*ERC20RollupStakerReassignedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupStakerReassignedIterator{contract: _ERC20Rollup.contract, event: "StakerReassigned", logs: logs, sub: sub}, nil
}

// WatchStakerReassigned is a free log subscription operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_ERC20Rollup *ERC20RollupFilterer) WatchStakerReassigned(opts *bind.WatchOpts, sink chan<- *ERC20RollupStakerReassigned, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupStakerReassigned)
				if err := _ERC20Rollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
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

// ParseStakerReassigned is a log parse operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_ERC20Rollup *ERC20RollupFilterer) ParseStakerReassigned(log types.Log) (*ERC20RollupStakerReassigned, error) {
	event := new(ERC20RollupStakerReassigned)
	if err := _ERC20Rollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20Rollup contract.
type ERC20RollupUnpausedIterator struct {
	Event *ERC20RollupUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20RollupUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupUnpaused)
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
		it.Event = new(ERC20RollupUnpaused)
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
func (it *ERC20RollupUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupUnpaused represents a Unpaused event raised by the ERC20Rollup contract.
type ERC20RollupUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20RollupUnpausedIterator, error) {

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20RollupUnpausedIterator{contract: _ERC20Rollup.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20RollupUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupUnpaused)
				if err := _ERC20Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) ParseUnpaused(log types.Log) (*ERC20RollupUnpaused, error) {
	event := new(ERC20RollupUnpaused)
	if err := _ERC20Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInboxABI is the input ABI used to generate the binding from.
const IInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"arbTxCallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submissionRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"valueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxGasPrice\",\"type\":\"uint256\"}],\"name\":\"depositEthRetryable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IInbox is an auto generated Go binding around an Ethereum contract.
type IInbox struct {
	IInboxCaller     // Read-only binding to the contract
	IInboxTransactor // Write-only binding to the contract
	IInboxFilterer   // Log filterer for contract events
}

// IInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInboxSession struct {
	Contract     *IInbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInboxCallerSession struct {
	Contract *IInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInboxTransactorSession struct {
	Contract     *IInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInboxRaw struct {
	Contract *IInbox // Generic contract binding to access the raw methods on
}

// IInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInboxCallerRaw struct {
	Contract *IInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInboxTransactorRaw struct {
	Contract *IInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInbox creates a new instance of IInbox, bound to a specific deployed contract.
func NewIInbox(address common.Address, backend bind.ContractBackend) (*IInbox, error) {
	contract, err := bindIInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInbox{IInboxCaller: IInboxCaller{contract: contract}, IInboxTransactor: IInboxTransactor{contract: contract}, IInboxFilterer: IInboxFilterer{contract: contract}}, nil
}

// NewIInboxCaller creates a new read-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxCaller(address common.Address, caller bind.ContractCaller) (*IInboxCaller, error) {
	contract, err := bindIInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxCaller{contract: contract}, nil
}

// NewIInboxTransactor creates a new write-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IInboxTransactor, error) {
	contract, err := bindIInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxTransactor{contract: contract}, nil
}

// NewIInboxFilterer creates a new log filterer instance of IInbox, bound to a specific deployed contract.
func NewIInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IInboxFilterer, error) {
	contract, err := bindIInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInboxFilterer{contract: contract}, nil
}

// bindIInbox binds a generic wrapper to an already deployed contract.
func bindIInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.IInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IInbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxSession) Bridge() (common.Address, error) {
	return _IInbox.Contract.Bridge(&_IInbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxCallerSession) Bridge() (common.Address, error) {
	return _IInbox.Contract.Bridge(&_IInbox.CallOpts)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactor) CreateRetryableTicket(opts *bind.TransactOpts, destAddr common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "createRetryableTicket", destAddr, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_IInbox *IInboxSession) CreateRetryableTicket(destAddr common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.CreateRetryableTicket(&_IInbox.TransactOpts, destAddr, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) CreateRetryableTicket(destAddr common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.CreateRetryableTicket(&_IInbox.TransactOpts, destAddr, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, maxGas, gasPriceBid, data)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns(uint256)
func (_IInbox *IInboxTransactor) DepositEth(opts *bind.TransactOpts, destAddr common.Address) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "depositEth", destAddr)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns(uint256)
func (_IInbox *IInboxSession) DepositEth(destAddr common.Address) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEth(&_IInbox.TransactOpts, destAddr)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) DepositEth(destAddr common.Address) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEth(&_IInbox.TransactOpts, destAddr)
}

// DepositEthRetryable is a paid mutator transaction binding the contract method 0xb8c6a388.
//
// Solidity: function depositEthRetryable(address destAddr, uint256 maxSubmissionCost, uint256 maxGas, uint256 maxGasPrice) payable returns(uint256)
func (_IInbox *IInboxTransactor) DepositEthRetryable(opts *bind.TransactOpts, destAddr common.Address, maxSubmissionCost *big.Int, maxGas *big.Int, maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "depositEthRetryable", destAddr, maxSubmissionCost, maxGas, maxGasPrice)
}

// DepositEthRetryable is a paid mutator transaction binding the contract method 0xb8c6a388.
//
// Solidity: function depositEthRetryable(address destAddr, uint256 maxSubmissionCost, uint256 maxGas, uint256 maxGasPrice) payable returns(uint256)
func (_IInbox *IInboxSession) DepositEthRetryable(destAddr common.Address, maxSubmissionCost *big.Int, maxGas *big.Int, maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEthRetryable(&_IInbox.TransactOpts, destAddr, maxSubmissionCost, maxGas, maxGasPrice)
}

// DepositEthRetryable is a paid mutator transaction binding the contract method 0xb8c6a388.
//
// Solidity: function depositEthRetryable(address destAddr, uint256 maxSubmissionCost, uint256 maxGas, uint256 maxGasPrice) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) DepositEthRetryable(destAddr common.Address, maxSubmissionCost *big.Int, maxGas *big.Int, maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEthRetryable(&_IInbox.TransactOpts, destAddr, maxSubmissionCost, maxGas, maxGasPrice)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_IInbox *IInboxTransactor) SendContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendContractTransaction", maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_IInbox *IInboxSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendContractTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_IInbox *IInboxTransactorSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendContractTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL1FundedContractTransaction", maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_IInbox *IInboxSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedContractTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedContractTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL1FundedUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_IInbox *IInboxSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedUnsignedTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedUnsignedTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_IInbox *IInboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_IInbox *IInboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL2Message(&_IInbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_IInbox *IInboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL2Message(&_IInbox.TransactOpts, messageData)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_IInbox *IInboxTransactor) SendUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_IInbox *IInboxSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendUnsignedTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_IInbox *IInboxTransactorSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendUnsignedTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// IInboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the IInbox contract.
type IInboxInboxMessageDeliveredIterator struct {
	Event *IInboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IInboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxInboxMessageDelivered)
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
		it.Event = new(IInboxInboxMessageDelivered)
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
func (it *IInboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the IInbox contract.
type IInboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IInbox *IInboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*IInboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IInboxInboxMessageDeliveredIterator{contract: _IInbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IInbox *IInboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *IInboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxInboxMessageDelivered)
				if err := _IInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IInbox *IInboxFilterer) ParseInboxMessageDelivered(log types.Log) (*IInboxInboxMessageDelivered, error) {
	event := new(IInboxInboxMessageDelivered)
	if err := _IInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the IInbox contract.
type IInboxInboxMessageDeliveredFromOriginIterator struct {
	Event *IInboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(IInboxInboxMessageDeliveredFromOrigin)
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
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the IInbox contract.
type IInboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IInbox *IInboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*IInboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IInboxInboxMessageDeliveredFromOriginIterator{contract: _IInbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IInbox *IInboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *IInboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxInboxMessageDeliveredFromOrigin)
				if err := _IInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IInbox *IInboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*IInboxInboxMessageDeliveredFromOrigin, error) {
	event := new(IInboxInboxMessageDeliveredFromOrigin)
	if err := _IInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOutboxABI is the input ABI used to generate the binding from.
const IOutboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionIndex\",\"type\":\"uint256\"}],\"name\":\"OutBoxTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numInBatch\",\"type\":\"uint256\"}],\"name\":\"OutboxEntryCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1EthBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"processOutgoingMessages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IOutbox is an auto generated Go binding around an Ethereum contract.
type IOutbox struct {
	IOutboxCaller     // Read-only binding to the contract
	IOutboxTransactor // Write-only binding to the contract
	IOutboxFilterer   // Log filterer for contract events
}

// IOutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOutboxSession struct {
	Contract     *IOutbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOutboxCallerSession struct {
	Contract *IOutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IOutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOutboxTransactorSession struct {
	Contract     *IOutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IOutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOutboxRaw struct {
	Contract *IOutbox // Generic contract binding to access the raw methods on
}

// IOutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOutboxCallerRaw struct {
	Contract *IOutboxCaller // Generic read-only contract binding to access the raw methods on
}

// IOutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOutboxTransactorRaw struct {
	Contract *IOutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOutbox creates a new instance of IOutbox, bound to a specific deployed contract.
func NewIOutbox(address common.Address, backend bind.ContractBackend) (*IOutbox, error) {
	contract, err := bindIOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOutbox{IOutboxCaller: IOutboxCaller{contract: contract}, IOutboxTransactor: IOutboxTransactor{contract: contract}, IOutboxFilterer: IOutboxFilterer{contract: contract}}, nil
}

// NewIOutboxCaller creates a new read-only instance of IOutbox, bound to a specific deployed contract.
func NewIOutboxCaller(address common.Address, caller bind.ContractCaller) (*IOutboxCaller, error) {
	contract, err := bindIOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOutboxCaller{contract: contract}, nil
}

// NewIOutboxTransactor creates a new write-only instance of IOutbox, bound to a specific deployed contract.
func NewIOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IOutboxTransactor, error) {
	contract, err := bindIOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOutboxTransactor{contract: contract}, nil
}

// NewIOutboxFilterer creates a new log filterer instance of IOutbox, bound to a specific deployed contract.
func NewIOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IOutboxFilterer, error) {
	contract, err := bindIOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOutboxFilterer{contract: contract}, nil
}

// bindIOutbox binds a generic wrapper to an already deployed contract.
func bindIOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOutbox *IOutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOutbox.Contract.IOutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOutbox *IOutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOutbox.Contract.IOutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOutbox *IOutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOutbox.Contract.IOutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOutbox *IOutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOutbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOutbox *IOutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOutbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOutbox *IOutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOutbox.Contract.contract.Transact(opts, method, params...)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_IOutbox *IOutboxCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_IOutbox *IOutboxSession) L2ToL1Block() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1Block(&_IOutbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_IOutbox *IOutboxCallerSession) L2ToL1Block() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1Block(&_IOutbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_IOutbox *IOutboxCaller) L2ToL1EthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1EthBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_IOutbox *IOutboxSession) L2ToL1EthBlock() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1EthBlock(&_IOutbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_IOutbox *IOutboxCallerSession) L2ToL1EthBlock() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1EthBlock(&_IOutbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_IOutbox *IOutboxCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_IOutbox *IOutboxSession) L2ToL1Sender() (common.Address, error) {
	return _IOutbox.Contract.L2ToL1Sender(&_IOutbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_IOutbox *IOutboxCallerSession) L2ToL1Sender() (common.Address, error) {
	return _IOutbox.Contract.L2ToL1Sender(&_IOutbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_IOutbox *IOutboxCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_IOutbox *IOutboxSession) L2ToL1Timestamp() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1Timestamp(&_IOutbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_IOutbox *IOutboxCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1Timestamp(&_IOutbox.CallOpts)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_IOutbox *IOutboxTransactor) ProcessOutgoingMessages(opts *bind.TransactOpts, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _IOutbox.contract.Transact(opts, "processOutgoingMessages", sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_IOutbox *IOutboxSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _IOutbox.Contract.ProcessOutgoingMessages(&_IOutbox.TransactOpts, sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_IOutbox *IOutboxTransactorSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _IOutbox.Contract.ProcessOutgoingMessages(&_IOutbox.TransactOpts, sendsData, sendLengths)
}

// IOutboxOutBoxTransactionExecutedIterator is returned from FilterOutBoxTransactionExecuted and is used to iterate over the raw logs and unpacked data for OutBoxTransactionExecuted events raised by the IOutbox contract.
type IOutboxOutBoxTransactionExecutedIterator struct {
	Event *IOutboxOutBoxTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *IOutboxOutBoxTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOutboxOutBoxTransactionExecuted)
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
		it.Event = new(IOutboxOutBoxTransactionExecuted)
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
func (it *IOutboxOutBoxTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOutboxOutBoxTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOutboxOutBoxTransactionExecuted represents a OutBoxTransactionExecuted event raised by the IOutbox contract.
type IOutboxOutBoxTransactionExecuted struct {
	DestAddr         common.Address
	L2Sender         common.Address
	OutboxIndex      *big.Int
	TransactionIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutBoxTransactionExecuted is a free log retrieval operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxIndex, uint256 transactionIndex)
func (_IOutbox *IOutboxFilterer) FilterOutBoxTransactionExecuted(opts *bind.FilterOpts, destAddr []common.Address, l2Sender []common.Address, outboxIndex []*big.Int) (*IOutboxOutBoxTransactionExecutedIterator, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var outboxIndexRule []interface{}
	for _, outboxIndexItem := range outboxIndex {
		outboxIndexRule = append(outboxIndexRule, outboxIndexItem)
	}

	logs, sub, err := _IOutbox.contract.FilterLogs(opts, "OutBoxTransactionExecuted", destAddrRule, l2SenderRule, outboxIndexRule)
	if err != nil {
		return nil, err
	}
	return &IOutboxOutBoxTransactionExecutedIterator{contract: _IOutbox.contract, event: "OutBoxTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchOutBoxTransactionExecuted is a free log subscription operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxIndex, uint256 transactionIndex)
func (_IOutbox *IOutboxFilterer) WatchOutBoxTransactionExecuted(opts *bind.WatchOpts, sink chan<- *IOutboxOutBoxTransactionExecuted, destAddr []common.Address, l2Sender []common.Address, outboxIndex []*big.Int) (event.Subscription, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var outboxIndexRule []interface{}
	for _, outboxIndexItem := range outboxIndex {
		outboxIndexRule = append(outboxIndexRule, outboxIndexItem)
	}

	logs, sub, err := _IOutbox.contract.WatchLogs(opts, "OutBoxTransactionExecuted", destAddrRule, l2SenderRule, outboxIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOutboxOutBoxTransactionExecuted)
				if err := _IOutbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
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

// ParseOutBoxTransactionExecuted is a log parse operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxIndex, uint256 transactionIndex)
func (_IOutbox *IOutboxFilterer) ParseOutBoxTransactionExecuted(log types.Log) (*IOutboxOutBoxTransactionExecuted, error) {
	event := new(IOutboxOutBoxTransactionExecuted)
	if err := _IOutbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOutboxOutboxEntryCreatedIterator is returned from FilterOutboxEntryCreated and is used to iterate over the raw logs and unpacked data for OutboxEntryCreated events raised by the IOutbox contract.
type IOutboxOutboxEntryCreatedIterator struct {
	Event *IOutboxOutboxEntryCreated // Event containing the contract specifics and raw log

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
func (it *IOutboxOutboxEntryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOutboxOutboxEntryCreated)
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
		it.Event = new(IOutboxOutboxEntryCreated)
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
func (it *IOutboxOutboxEntryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOutboxOutboxEntryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOutboxOutboxEntryCreated represents a OutboxEntryCreated event raised by the IOutbox contract.
type IOutboxOutboxEntryCreated struct {
	BatchNum    *big.Int
	OutboxIndex *big.Int
	OutputRoot  [32]byte
	NumInBatch  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOutboxEntryCreated is a free log retrieval operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_IOutbox *IOutboxFilterer) FilterOutboxEntryCreated(opts *bind.FilterOpts, batchNum []*big.Int) (*IOutboxOutboxEntryCreatedIterator, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _IOutbox.contract.FilterLogs(opts, "OutboxEntryCreated", batchNumRule)
	if err != nil {
		return nil, err
	}
	return &IOutboxOutboxEntryCreatedIterator{contract: _IOutbox.contract, event: "OutboxEntryCreated", logs: logs, sub: sub}, nil
}

// WatchOutboxEntryCreated is a free log subscription operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_IOutbox *IOutboxFilterer) WatchOutboxEntryCreated(opts *bind.WatchOpts, sink chan<- *IOutboxOutboxEntryCreated, batchNum []*big.Int) (event.Subscription, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _IOutbox.contract.WatchLogs(opts, "OutboxEntryCreated", batchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOutboxOutboxEntryCreated)
				if err := _IOutbox.contract.UnpackLog(event, "OutboxEntryCreated", log); err != nil {
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

// ParseOutboxEntryCreated is a log parse operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_IOutbox *IOutboxFilterer) ParseOutboxEntryCreated(log types.Log) (*IOutboxOutboxEntryCreated, error) {
	event := new(IOutboxOutboxEntryCreated)
	if err := _IOutbox.contract.UnpackLog(event, "OutboxEntryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxABI is the input ABI used to generate the binding from.
const InboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxGasPrice\",\"type\":\"uint256\"}],\"name\":\"depositEthRetryable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InboxBin is the compiled bytecode used for deploying new contracts.
var InboxBin = "0x608060405234801561001057600080fd5b506000805460ff19166001179055610cb88061002d6000396000f3fe6080604052600436106100975760003560e01c80631fe927cf1461009c5780635075788b146101295780635e916758146101cc578063679b6ded1461025657806367ef3ab8146102ff5780636f791d291461038e5780638a631aa6146103b7578063ad9d4ba314610453578063b75436bb14610479578063b8c6a388146104f4578063c4d66de81461052c578063e78cea9214610561575b600080fd5b3480156100a857600080fd5b50610117600480360360208110156100bf57600080fd5b810190602081018135600160201b8111156100d957600080fd5b8201836020820111156100eb57600080fd5b803590602001918460018302840111600160201b8311171561010c57600080fd5b509092509050610592565b60408051918252519081900360200190f35b34801561013557600080fd5b50610117600480360360c081101561014c57600080fd5b8135916020810135916040820135916001600160a01b03606082013516916080820135919081019060c0810160a0820135600160201b81111561018e57600080fd5b8201836020820111156101a057600080fd5b803590602001918460018302840111600160201b831117156101c157600080fd5b509092509050610635565b610117600480360360808110156101e257600080fd5b8135916020810135916001600160a01b036040830135169190810190608081016060820135600160201b81111561021857600080fd5b82018360208201111561022a57600080fd5b803590602001918460018302840111600160201b8311171561024b57600080fd5b5090925090506106bd565b610117600480360361010081101561026d57600080fd5b6001600160a01b038235811692602081013592604082013592606083013581169260808101359091169160a08201359160c081013591810190610100810160e0820135600160201b8111156102c157600080fd5b8201836020820111156102d357600080fd5b803590602001918460018302840111600160201b831117156102f457600080fd5b50909250905061073b565b610117600480360360a081101561031557600080fd5b8135916020810135916040820135916001600160a01b036060820135169181019060a081016080820135600160201b81111561035057600080fd5b82018360208201111561036257600080fd5b803590602001918460018302840111600160201b8311171561038357600080fd5b5090925090506107f7565b34801561039a57600080fd5b506103a361087e565b604080519115158252519081900360200190f35b3480156103c357600080fd5b50610117600480360360a08110156103da57600080fd5b8135916020810135916001600160a01b036040830135169160608101359181019060a081016080820135600160201b81111561041557600080fd5b82018360208201111561042757600080fd5b803590602001918460018302840111600160201b8311171561044857600080fd5b509092509050610887565b6101176004803603602081101561046957600080fd5b50356001600160a01b03166108fb565b34801561048557600080fd5b506101176004803603602081101561049c57600080fd5b810190602081018135600160201b8111156104b657600080fd5b8201836020820111156104c857600080fd5b803590602001918460018302840111600160201b831117156104e957600080fd5b509092509050610958565b6101176004803603608081101561050a57600080fd5b506001600160a01b0381351690602081013590604081013590606001356109db565b34801561053857600080fd5b5061055f6004803603602081101561054f57600080fd5b50356001600160a01b0316610a92565b005b34801561056d57600080fd5b50610576610b0c565b604080516001600160a01b039092168252519081900360200190f35b60003332146105d6576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600061060060033386866040518083838082843760405192018290039091209350610b2092505050565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a29392505050565b60006106b160033360008b8b8b8b60601b60601c6001600160a01b03168b8b8b604051602001808960ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052610bbb565b98975050505050505050565b6000610731600733600189898960601b60601c6001600160a01b0316348a8a604051602001808860ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052610bbb565b9695505050505050565b60006107e96009338c60601b60601c6001600160a01b03168c348d8d60601b60601c6001600160a01b03168d60601b60601c6001600160a01b03168d8d8d8d90508e8e604051602001808c81526020018b81526020018a8152602001898152602001888152602001878152602001868152602001858152602001848152602001838380828437808301925050509b505050505050505050505050604051602081830303815290604052610bbb565b9a9950505050505050505050565b600061087360073360008a8a8a8a60601b60601c6001600160a01b0316348b8b604051602001808960ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052610bbb565b979650505050505050565b60005460ff1690565b600061087360033360018a8a8a60601b60601c6001600160a01b03168a8a8a604051602001808860ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052610bbb565b60408051600160f81b6020820152600060218201819052604182018190526001600160a01b0384166061830152346081808401919091528351808403909101815260a1909201909252610952906007908490610bbb565b92915050565b60008061098360033386866040518083838082843760405192018290039091209350610b2092505050565b905080600080516020610c63833981519152858560405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a29392505050565b6040805163679b6ded60e01b81526001600160a01b0386166004820152346024820152604481018590523360648201819052608482015260a4810184905260c4810183905261010060e4820152600061010482018190529151309163679b6ded9161014480830192602092919082900301818787803b158015610a5d57600080fd5b505af1158015610a71573d6000803e3d6000fd5b505050506040513d6020811015610a8757600080fd5b505195945050505050565b60005461010090046001600160a01b031615610ae4576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600080546001600160a01b0390921661010002610100600160a81b0319909216919091179055565b60005461010090046001600160a01b031681565b60008054604080516302bbfad160e01b815260ff871660048201526001600160a01b038681166024830152604482018690529151610100909304909116916302bbfad1913491606480830192602092919082900301818588803b158015610b8657600080fd5b505af1158015610b9a573d6000803e3d6000fd5b50505050506040513d6020811015610bb157600080fd5b5051949350505050565b600080610bd085858580519060200120610b20565b905080600080516020610c63833981519152846040518080602001828103825283818151815260200191508051906020019080838360005b83811015610c20578181015183820152602001610c08565b50505050905090810190601f168015610c4d5780820380516001836020036101000a031916815260200191505b509250505060405180910390a294935050505056feff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60ba2646970667358221220488558ba5fa286ea60ec52a8c54eed040a38bb3686c995260bf4769b82979aab64736f6c634300060c0033"

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCallerSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Inbox *InboxCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Inbox *InboxSession) IsMaster() (bool, error) {
	return _Inbox.Contract.IsMaster(&_Inbox.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Inbox *InboxCallerSession) IsMaster() (bool, error) {
	return _Inbox.Contract.IsMaster(&_Inbox.CallOpts)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) CreateRetryableTicket(opts *bind.TransactOpts, destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "createRetryableTicket", destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) CreateRetryableTicket(destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicket(&_Inbox.TransactOpts, destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) CreateRetryableTicket(destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicket(&_Inbox.TransactOpts, destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns(uint256)
func (_Inbox *InboxTransactor) DepositEth(opts *bind.TransactOpts, destAddr common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEth", destAddr)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns(uint256)
func (_Inbox *InboxSession) DepositEth(destAddr common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, destAddr)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns(uint256)
func (_Inbox *InboxTransactorSession) DepositEth(destAddr common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, destAddr)
}

// DepositEthRetryable is a paid mutator transaction binding the contract method 0xb8c6a388.
//
// Solidity: function depositEthRetryable(address destAddr, uint256 maxSubmissionCost, uint256 maxGas, uint256 maxGasPrice) payable returns(uint256)
func (_Inbox *InboxTransactor) DepositEthRetryable(opts *bind.TransactOpts, destAddr common.Address, maxSubmissionCost *big.Int, maxGas *big.Int, maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEthRetryable", destAddr, maxSubmissionCost, maxGas, maxGasPrice)
}

// DepositEthRetryable is a paid mutator transaction binding the contract method 0xb8c6a388.
//
// Solidity: function depositEthRetryable(address destAddr, uint256 maxSubmissionCost, uint256 maxGas, uint256 maxGasPrice) payable returns(uint256)
func (_Inbox *InboxSession) DepositEthRetryable(destAddr common.Address, maxSubmissionCost *big.Int, maxGas *big.Int, maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEthRetryable(&_Inbox.TransactOpts, destAddr, maxSubmissionCost, maxGas, maxGasPrice)
}

// DepositEthRetryable is a paid mutator transaction binding the contract method 0xb8c6a388.
//
// Solidity: function depositEthRetryable(address destAddr, uint256 maxSubmissionCost, uint256 maxGas, uint256 maxGasPrice) payable returns(uint256)
func (_Inbox *InboxTransactorSession) DepositEthRetryable(destAddr common.Address, maxSubmissionCost *big.Int, maxGas *big.Int, maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEthRetryable(&_Inbox.TransactOpts, destAddr, maxSubmissionCost, maxGas, maxGasPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_Inbox *InboxTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_Inbox *InboxSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_Inbox *InboxTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, _bridge)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendContractTransaction", maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedContractTransaction", maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// InboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the Inbox contract.
type InboxInboxMessageDeliveredIterator struct {
	Event *InboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDelivered)
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
		it.Event = new(InboxInboxMessageDelivered)
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
func (it *InboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the Inbox contract.
type InboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredIterator{contract: _Inbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDelivered)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxInboxMessageDelivered, error) {
	event := new(InboxInboxMessageDelivered)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredFromOriginIterator{contract: _Inbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDeliveredFromOrigin)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxInboxMessageDeliveredFromOrigin)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxABI is the input ABI used to generate the binding from.
const OutboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionIndex\",\"type\":\"uint256\"}],\"name\":\"OutBoxTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numInBatch\",\"type\":\"uint256\"}],\"name\":\"OutboxEntryCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"calculateItemHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"item\",\"type\":\"bytes32\"}],\"name\":\"calculateMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1EthBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outboxes\",\"outputs\":[{\"internalType\":\"contractOutboxEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outboxesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"processOutgoingMessages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OutboxBin is the compiled bytecode used for deploying new contracts.
var OutboxBin = "0x608060405234801561001057600080fd5b506000805460ff191660011790556118378061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80627436d3146100a857806305d3efe6146101605780630c726847146101685780634654779014610228578063485cc955146102305780636d5161ec1461025e5780636f791d291461029757806380648b02146102b35780638515bc6a146102bb5780639c5cfe0b146102c35780639f0c04bf146103bf578063b0f305371461045e575b600080fd5b61014e600480360360608110156100be57600080fd5b810190602081018135600160201b8111156100d857600080fd5b8201836020820111156100ea57600080fd5b803590602001918460208302840111600160201b8311171561010b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505082359350505060200135610466565b60408051918252519081900360200190f35b61014e6104a1565b6102266004803603604081101561017e57600080fd5b810190602081018135600160201b81111561019857600080fd5b8201836020820111156101aa57600080fd5b803590602001918460018302840111600160201b831117156101cb57600080fd5b919390929091602081019035600160201b8111156101e857600080fd5b8201836020820111156101fa57600080fd5b803590602001918460208302840111600160201b8311171561021b57600080fd5b5090925090506104a7565b005b61014e610593565b6102266004803603604081101561024657600080fd5b506001600160a01b03813581169160200135166105a2565b61027b6004803603602081101561027457600080fd5b5035610680565b604080516001600160a01b039092168252519081900360200190f35b61029f6106a7565b604080519115158252519081900360200190f35b61027b6106b0565b61014e6106bf565b61022660048036036101408110156102da57600080fd5b81359190810190604081016020820135600160201b8111156102fb57600080fd5b82018360208201111561030d57600080fd5b803590602001918460208302840111600160201b8311171561032e57600080fd5b919390928235926001600160a01b03602082013581169360408301359091169260608301359260808101359260a08201359260c08301359261010081019060e00135600160201b81111561038157600080fd5b82018360208201111561039357600080fd5b803590602001918460018302840111600160201b831117156103b457600080fd5b5090925090506106d5565b61014e600480360360e08110156103d557600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561042057600080fd5b82018360208201111561043257600080fd5b803590602001918460018302840111600160201b8311171561045357600080fd5b509092509050610898565b61014e610935565b60006104998484846040516020018082815260200191505060405160208183030381529060405280519060200120610944565b949350505050565b60035490565b60005461010090046001600160a01b031633146104f9576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b806000805b8281101561058a5761056887838888888681811061051857fe5b9050602002013586019261052e93929190611304565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a1292505050565b84848281811061057457fe5b60200291909101359290920191506001016104fe565b50505050505050565b6005546001600160801b031690565b60005461010090046001600160a01b0316156105f4576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b60008054610100600160a81b0319166101006001600160a01b038581169190910291909117909155600180546001600160a01b03191691831691909117905560405161063f906112f7565b604051809103906000f08015801561065b573d6000803e3d6000fd5b50600280546001600160a01b0319166001600160a01b03929092169190911790555050565b6003818154811061068d57fe5b6000918252602090912001546001600160a01b0316905081565b60005460ff1690565b6004546001600160a01b031690565b600554600160801b90046001600160801b031690565b60006106e78989898989898989610898565b905061072a8d8d8d808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508f9250869150610bc39050565b8c896001600160a01b0316896001600160a01b03167f20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab189648d6040518082815260200191505060405180910390a46004805460058054600680546001600160a01b038f81166001600160a01b03198716179096556001600160801b038c8116600160801b9081028f83166001600160801b0319808816919091178416919091179096558c821695831695909517909255604080516020601f8b018190048102820181019092528981529690951695828416959490930482169391169161082c918e918b918b908b9081908401838280828437600092019190915250610e1c92505050565b600480546001600160a01b03959095166001600160a01b031990951694909417909355600580546001600160801b03928316600160801b029383166001600160801b03199182161783169390931790556006805491909316911617905550505050505050505050505050565b600060038960601b60601c6001600160a01b03168960601b60601c6001600160a01b0316898989898989604051602001808a60ff1660f81b815260010189815260200188815260200187815260200186815260200185815260200184815260200183838082843780830192505050995050505050505050505060405160208183030381529060405280519060200120905098975050505050505050565b6006546001600160801b031690565b825160009061010081111561095857600080fd5b8260005b82811015610a0857600286066109b55786818151811061097857fe5b60200260200101518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091506109fa565b818782815181106109c257fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b60028604955060010161095c565b5095945050505050565b805160009082908290610a2157fe5b01602001516001600160f81b0319161415610bc0578051606114610a79576040805162461bcd60e51b815260206004820152600a6024820152690848288be988a9c8ea8960b31b604482015290519081900360640190fd5b6000610a86826001611039565b90506000610a95836021611039565b90506000610aa4846041611039565b600254909150600090610abf906001600160a01b0316611092565b9050806001600160a01b0316635b36c66b83856040518363ffffffff1660e01b81526004018083815260200182815260200192505050600060405180830381600087803b158015610b0f57600080fd5b505af1158015610b23573d6000803e3d6000fd5b5050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810180546001600160a01b0386166001600160a01b0319909116179055604080518281526020810187905280820188905290519193508792507fe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131919081900360600190a250505050505b50565b61010083511115610c0c576040805162461bcd60e51b815260206004820152600e60248201526d50524f4f465f544f4f5f4c4f4e4760901b604482015290519081900360640190fd5b825160020a8210610c57576040805162461bcd60e51b815260206004820152601060248201526f1410551217d393d517d352539253505360821b604482015290519081900360640190fd5b6000610c64848484610466565b9050600060038681548110610c7557fe5b6000918252602090912001546001600160a01b0316905080610cca576040805162461bcd60e51b815260206004820152600960248201526809c9ebe9eaaa8849eb60bb1b604482015290519081900360640190fd5b600084865160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012090506000826001600160a01b03166357d61c0b85846040518363ffffffff1660e01b81526004018083815260200182815260200192505050602060405180830381600087803b158015610d4e57600080fd5b505af1158015610d62573d6000803e3d6000fd5b505050506040513d6020811015610d7857600080fd5b5051905080610e1257826001600160a01b03166383197ef06040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610dbc57600080fd5b505af1158015610dd0573d6000803e3d6000fd5b50505050600060038981548110610de357fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b5050505050505050565b600154604051639e5d4c4960e01b81526001600160a01b03858116600483019081526024830186905260606044840181815286516064860152865160009692959490921693639e5d4c49938a938a938a93909160849091019060208501908083838e5b83811015610e97578181015183820152602001610e7f565b50505050905090810190601f168015610ec45780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610ee557600080fd5b505af1158015610ef9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015610f2257600080fd5b815160208301805160405192949293830192919084600160201b821115610f4857600080fd5b908301906020820185811115610f5d57600080fd5b8251600160201b811182820188101715610f7657600080fd5b82525081516020918201929091019080838360005b83811015610fa3578181015183820152602001610f8b565b50505050905090810190601f168015610fd05780820380516001836020036101000a031916815260200191505b50604052505050915091508161103257805115610ff05780518082602001fd5b6040805162461bcd60e51b81526020600482015260126024820152711094925111d157d0d0531317d1905253115160721b604482015290519081900360640190fd5b5050505050565b60008160200183511015611089576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000816110a7816001600160a01b0316611250565b604051806040016040528060188152602001772727afa1a7a72a2920a1aa2fa1a627a722afa6a0a9aa22a960411b815250906111615760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561112657818101518382015260200161110e565b50505050905090810190601f1680156111535780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b15801561119b57600080fd5b505afa1580156111af573d6000803e3d6000fd5b505050506040513d60208110156111c557600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b6020820152906112365760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561112657818101518382015260200161110e565b50611249836001600160a01b031661125a565b9392505050565b803b15155b919050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116611255576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b6104d58061132d83390190565b60008085851115611313578182fd5b8386111561131f578182fd5b505082019391909203915056fe608060405234801561001057600080fd5b506000805460ff191660011790556104a88061002d6000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80635780e4e71461007257806357d61c0b1461008c5780635b36c66b146100af5780636f791d29146100d457806383197ef0146100f05780639db9af81146100f8578063ebf0c71714610115575b600080fd5b61007a61011d565b60408051918252519081900360200190f35b61007a600480360360408110156100a257600080fd5b5080359060200135610123565b6100d2600480360360408110156100c557600080fd5b508035906020013561023b565b005b6100dc610330565b604080519115158252519081900360200190f35b6100d2610339565b6100dc6004803603602081101561010e57600080fd5b503561039a565b61007a6103af565b60025481565b6000805461010090046001600160a01b0316331461017a576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b60008281526003602052604090205460ff16156101ce576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b600154831461020f576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b506000818152600360205260409020805460ff1916600117905560028054600019019081905592915050565b60005461010090046001600160a01b03161561028d576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600154156102d1576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b8161030e576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b60008054610100600160a81b0319163361010002179055600191909155600255565b60005460ff1690565b60005461010090046001600160a01b0316331461038f576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b610398336103b5565b565b60036020526000908152604090205460ff1681565b60015481565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff16156104655760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561042a578181015183820152602001610412565b50505050905090810190601f1680156104575780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316fffea26469706673582212205ecec9da6ebed4eccad9344ae0a522e3085e41b86c9eccc6fb729d9bcb7618e464736f6c634300060c0033a26469706673582212204ccec08ea791b7560390c0551e8cd2460e272073f7af02bb1876e4e89da7ad9964736f6c634300060c0033"

// DeployOutbox deploys a new Ethereum contract, binding an instance of Outbox to it.
func DeployOutbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Outbox, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// Outbox is an auto generated Go binding around an Ethereum contract.
type Outbox struct {
	OutboxCaller     // Read-only binding to the contract
	OutboxTransactor // Write-only binding to the contract
	OutboxFilterer   // Log filterer for contract events
}

// OutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxSession struct {
	Contract     *Outbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxCallerSession struct {
	Contract *OutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxTransactorSession struct {
	Contract     *OutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxRaw struct {
	Contract *Outbox // Generic contract binding to access the raw methods on
}

// OutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxCallerRaw struct {
	Contract *OutboxCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxTransactorRaw struct {
	Contract *OutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutbox creates a new instance of Outbox, bound to a specific deployed contract.
func NewOutbox(address common.Address, backend bind.ContractBackend) (*Outbox, error) {
	contract, err := bindOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// NewOutboxCaller creates a new read-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxCaller(address common.Address, caller bind.ContractCaller) (*OutboxCaller, error) {
	contract, err := bindOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxCaller{contract: contract}, nil
}

// NewOutboxTransactor creates a new write-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxTransactor, error) {
	contract, err := bindOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxTransactor{contract: contract}, nil
}

// NewOutboxFilterer creates a new log filterer instance of Outbox, bound to a specific deployed contract.
func NewOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxFilterer, error) {
	contract, err := bindOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxFilterer{contract: contract}, nil
}

// bindOutbox binds a generic wrapper to an already deployed contract.
func bindOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.OutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transact(opts, method, params...)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) pure returns(bytes32)
func (_Outbox *OutboxCaller) CalculateItemHash(opts *bind.CallOpts, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "calculateItemHash", l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) pure returns(bytes32)
func (_Outbox *OutboxSession) CalculateItemHash(l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateItemHash(&_Outbox.CallOpts, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) pure returns(bytes32)
func (_Outbox *OutboxCallerSession) CalculateItemHash(l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateItemHash(&_Outbox.CallOpts, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxCaller) CalculateMerkleRoot(opts *bind.CallOpts, proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "calculateMerkleRoot", proof, path, item)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateMerkleRoot(&_Outbox.CallOpts, proof, path, item)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxCallerSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateMerkleRoot(&_Outbox.CallOpts, proof, path, item)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Outbox *OutboxCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Outbox *OutboxSession) IsMaster() (bool, error) {
	return _Outbox.Contract.IsMaster(&_Outbox.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Outbox *OutboxCallerSession) IsMaster() (bool, error) {
	return _Outbox.Contract.IsMaster(&_Outbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1EthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1EthBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1EthBlock() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1EthBlock(&_Outbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1EthBlock() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1EthBlock(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCallerSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// Outboxes is a free data retrieval call binding the contract method 0x6d5161ec.
//
// Solidity: function outboxes(uint256 ) view returns(address)
func (_Outbox *OutboxCaller) Outboxes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "outboxes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outboxes is a free data retrieval call binding the contract method 0x6d5161ec.
//
// Solidity: function outboxes(uint256 ) view returns(address)
func (_Outbox *OutboxSession) Outboxes(arg0 *big.Int) (common.Address, error) {
	return _Outbox.Contract.Outboxes(&_Outbox.CallOpts, arg0)
}

// Outboxes is a free data retrieval call binding the contract method 0x6d5161ec.
//
// Solidity: function outboxes(uint256 ) view returns(address)
func (_Outbox *OutboxCallerSession) Outboxes(arg0 *big.Int) (common.Address, error) {
	return _Outbox.Contract.Outboxes(&_Outbox.CallOpts, arg0)
}

// OutboxesLength is a free data retrieval call binding the contract method 0x05d3efe6.
//
// Solidity: function outboxesLength() view returns(uint256)
func (_Outbox *OutboxCaller) OutboxesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "outboxesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutboxesLength is a free data retrieval call binding the contract method 0x05d3efe6.
//
// Solidity: function outboxesLength() view returns(uint256)
func (_Outbox *OutboxSession) OutboxesLength() (*big.Int, error) {
	return _Outbox.Contract.OutboxesLength(&_Outbox.CallOpts)
}

// OutboxesLength is a free data retrieval call binding the contract method 0x05d3efe6.
//
// Solidity: function outboxesLength() view returns(uint256)
func (_Outbox *OutboxCallerSession) OutboxesLength() (*big.Int, error) {
	return _Outbox.Contract.OutboxesLength(&_Outbox.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x9c5cfe0b.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactor) ExecuteTransaction(opts *bind.TransactOpts, outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeTransaction", outboxIndex, proof, index, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x9c5cfe0b.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxSession) ExecuteTransaction(outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, proof, index, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x9c5cfe0b.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactorSession) ExecuteTransaction(outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, proof, index, l2Sender, destAddr, l2Block, l1Block, l2Timestamp, amount, calldataForL1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rollup, address _bridge) returns()
func (_Outbox *OutboxTransactor) Initialize(opts *bind.TransactOpts, _rollup common.Address, _bridge common.Address) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "initialize", _rollup, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rollup, address _bridge) returns()
func (_Outbox *OutboxSession) Initialize(_rollup common.Address, _bridge common.Address) (*types.Transaction, error) {
	return _Outbox.Contract.Initialize(&_Outbox.TransactOpts, _rollup, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rollup, address _bridge) returns()
func (_Outbox *OutboxTransactorSession) Initialize(_rollup common.Address, _bridge common.Address) (*types.Transaction, error) {
	return _Outbox.Contract.Initialize(&_Outbox.TransactOpts, _rollup, _bridge)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxTransactor) ProcessOutgoingMessages(opts *bind.TransactOpts, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "processOutgoingMessages", sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.Contract.ProcessOutgoingMessages(&_Outbox.TransactOpts, sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxTransactorSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.Contract.ProcessOutgoingMessages(&_Outbox.TransactOpts, sendsData, sendLengths)
}

// OutboxOutBoxTransactionExecutedIterator is returned from FilterOutBoxTransactionExecuted and is used to iterate over the raw logs and unpacked data for OutBoxTransactionExecuted events raised by the Outbox contract.
type OutboxOutBoxTransactionExecutedIterator struct {
	Event *OutboxOutBoxTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *OutboxOutBoxTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxOutBoxTransactionExecuted)
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
		it.Event = new(OutboxOutBoxTransactionExecuted)
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
func (it *OutboxOutBoxTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxOutBoxTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxOutBoxTransactionExecuted represents a OutBoxTransactionExecuted event raised by the Outbox contract.
type OutboxOutBoxTransactionExecuted struct {
	DestAddr         common.Address
	L2Sender         common.Address
	OutboxIndex      *big.Int
	TransactionIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutBoxTransactionExecuted is a free log retrieval operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxIndex, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) FilterOutBoxTransactionExecuted(opts *bind.FilterOpts, destAddr []common.Address, l2Sender []common.Address, outboxIndex []*big.Int) (*OutboxOutBoxTransactionExecutedIterator, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var outboxIndexRule []interface{}
	for _, outboxIndexItem := range outboxIndex {
		outboxIndexRule = append(outboxIndexRule, outboxIndexItem)
	}

	logs, sub, err := _Outbox.contract.FilterLogs(opts, "OutBoxTransactionExecuted", destAddrRule, l2SenderRule, outboxIndexRule)
	if err != nil {
		return nil, err
	}
	return &OutboxOutBoxTransactionExecutedIterator{contract: _Outbox.contract, event: "OutBoxTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchOutBoxTransactionExecuted is a free log subscription operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxIndex, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) WatchOutBoxTransactionExecuted(opts *bind.WatchOpts, sink chan<- *OutboxOutBoxTransactionExecuted, destAddr []common.Address, l2Sender []common.Address, outboxIndex []*big.Int) (event.Subscription, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var outboxIndexRule []interface{}
	for _, outboxIndexItem := range outboxIndex {
		outboxIndexRule = append(outboxIndexRule, outboxIndexItem)
	}

	logs, sub, err := _Outbox.contract.WatchLogs(opts, "OutBoxTransactionExecuted", destAddrRule, l2SenderRule, outboxIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxOutBoxTransactionExecuted)
				if err := _Outbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
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

// ParseOutBoxTransactionExecuted is a log parse operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed destAddr, address indexed l2Sender, uint256 indexed outboxIndex, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) ParseOutBoxTransactionExecuted(log types.Log) (*OutboxOutBoxTransactionExecuted, error) {
	event := new(OutboxOutBoxTransactionExecuted)
	if err := _Outbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxOutboxEntryCreatedIterator is returned from FilterOutboxEntryCreated and is used to iterate over the raw logs and unpacked data for OutboxEntryCreated events raised by the Outbox contract.
type OutboxOutboxEntryCreatedIterator struct {
	Event *OutboxOutboxEntryCreated // Event containing the contract specifics and raw log

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
func (it *OutboxOutboxEntryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxOutboxEntryCreated)
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
		it.Event = new(OutboxOutboxEntryCreated)
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
func (it *OutboxOutboxEntryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxOutboxEntryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxOutboxEntryCreated represents a OutboxEntryCreated event raised by the Outbox contract.
type OutboxOutboxEntryCreated struct {
	BatchNum    *big.Int
	OutboxIndex *big.Int
	OutputRoot  [32]byte
	NumInBatch  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOutboxEntryCreated is a free log retrieval operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) FilterOutboxEntryCreated(opts *bind.FilterOpts, batchNum []*big.Int) (*OutboxOutboxEntryCreatedIterator, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Outbox.contract.FilterLogs(opts, "OutboxEntryCreated", batchNumRule)
	if err != nil {
		return nil, err
	}
	return &OutboxOutboxEntryCreatedIterator{contract: _Outbox.contract, event: "OutboxEntryCreated", logs: logs, sub: sub}, nil
}

// WatchOutboxEntryCreated is a free log subscription operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) WatchOutboxEntryCreated(opts *bind.WatchOpts, sink chan<- *OutboxOutboxEntryCreated, batchNum []*big.Int) (event.Subscription, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Outbox.contract.WatchLogs(opts, "OutboxEntryCreated", batchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxOutboxEntryCreated)
				if err := _Outbox.contract.UnpackLog(event, "OutboxEntryCreated", log); err != nil {
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

// ParseOutboxEntryCreated is a log parse operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) ParseOutboxEntryCreated(log types.Log) (*OutboxOutboxEntryCreated, error) {
	event := new(OutboxOutboxEntryCreated)
	if err := _Outbox.contract.UnpackLog(event, "OutboxEntryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxEntryABI is the input ABI used to generate the binding from.
const OutboxEntryABI = "[{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_numInBatch\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"spendOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"spentOutput\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OutboxEntryBin is the compiled bytecode used for deploying new contracts.
var OutboxEntryBin = "0x608060405234801561001057600080fd5b506000805460ff191660011790556104a88061002d6000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80635780e4e71461007257806357d61c0b1461008c5780635b36c66b146100af5780636f791d29146100d457806383197ef0146100f05780639db9af81146100f8578063ebf0c71714610115575b600080fd5b61007a61011d565b60408051918252519081900360200190f35b61007a600480360360408110156100a257600080fd5b5080359060200135610123565b6100d2600480360360408110156100c557600080fd5b508035906020013561023b565b005b6100dc610330565b604080519115158252519081900360200190f35b6100d2610339565b6100dc6004803603602081101561010e57600080fd5b503561039a565b61007a6103af565b60025481565b6000805461010090046001600160a01b0316331461017a576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b60008281526003602052604090205460ff16156101ce576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b600154831461020f576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b506000818152600360205260409020805460ff1916600117905560028054600019019081905592915050565b60005461010090046001600160a01b03161561028d576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600154156102d1576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b8161030e576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b60008054610100600160a81b0319163361010002179055600191909155600255565b60005460ff1690565b60005461010090046001600160a01b0316331461038f576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b610398336103b5565b565b60036020526000908152604090205460ff1681565b60015481565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff16156104655760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561042a578181015183820152602001610412565b50505050905090810190601f1680156104575780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316fffea26469706673582212205ecec9da6ebed4eccad9344ae0a522e3085e41b86c9eccc6fb729d9bcb7618e464736f6c634300060c0033"

// DeployOutboxEntry deploys a new Ethereum contract, binding an instance of OutboxEntry to it.
func DeployOutboxEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OutboxEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// OutboxEntry is an auto generated Go binding around an Ethereum contract.
type OutboxEntry struct {
	OutboxEntryCaller     // Read-only binding to the contract
	OutboxEntryTransactor // Write-only binding to the contract
	OutboxEntryFilterer   // Log filterer for contract events
}

// OutboxEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxEntrySession struct {
	Contract     *OutboxEntry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxEntryCallerSession struct {
	Contract *OutboxEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OutboxEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxEntryTransactorSession struct {
	Contract     *OutboxEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OutboxEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxEntryRaw struct {
	Contract *OutboxEntry // Generic contract binding to access the raw methods on
}

// OutboxEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxEntryCallerRaw struct {
	Contract *OutboxEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxEntryTransactorRaw struct {
	Contract *OutboxEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutboxEntry creates a new instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntry(address common.Address, backend bind.ContractBackend) (*OutboxEntry, error) {
	contract, err := bindOutboxEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// NewOutboxEntryCaller creates a new read-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryCaller(address common.Address, caller bind.ContractCaller) (*OutboxEntryCaller, error) {
	contract, err := bindOutboxEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryCaller{contract: contract}, nil
}

// NewOutboxEntryTransactor creates a new write-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxEntryTransactor, error) {
	contract, err := bindOutboxEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryTransactor{contract: contract}, nil
}

// NewOutboxEntryFilterer creates a new log filterer instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxEntryFilterer, error) {
	contract, err := bindOutboxEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryFilterer{contract: contract}, nil
}

// bindOutboxEntry binds a generic wrapper to an already deployed contract.
func bindOutboxEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.OutboxEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transact(opts, method, params...)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntryCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntrySession) IsMaster() (bool, error) {
	return _OutboxEntry.Contract.IsMaster(&_OutboxEntry.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntryCallerSession) IsMaster() (bool, error) {
	return _OutboxEntry.Contract.IsMaster(&_OutboxEntry.CallOpts)
}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntryCaller) NumRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "numRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntrySession) NumRemaining() (*big.Int, error) {
	return _OutboxEntry.Contract.NumRemaining(&_OutboxEntry.CallOpts)
}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntryCallerSession) NumRemaining() (*big.Int, error) {
	return _OutboxEntry.Contract.NumRemaining(&_OutboxEntry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntryCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntrySession) Root() ([32]byte, error) {
	return _OutboxEntry.Contract.Root(&_OutboxEntry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntryCallerSession) Root() ([32]byte, error) {
	return _OutboxEntry.Contract.Root(&_OutboxEntry.CallOpts)
}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntryCaller) SpentOutput(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "spentOutput", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntrySession) SpentOutput(arg0 [32]byte) (bool, error) {
	return _OutboxEntry.Contract.SpentOutput(&_OutboxEntry.CallOpts, arg0)
}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntryCallerSession) SpentOutput(arg0 [32]byte) (bool, error) {
	return _OutboxEntry.Contract.SpentOutput(&_OutboxEntry.CallOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntryTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntrySession) Destroy() (*types.Transaction, error) {
	return _OutboxEntry.Contract.Destroy(&_OutboxEntry.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntryTransactorSession) Destroy() (*types.Transaction, error) {
	return _OutboxEntry.Contract.Destroy(&_OutboxEntry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b36c66b.
//
// Solidity: function initialize(bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntryTransactor) Initialize(opts *bind.TransactOpts, _root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "initialize", _root, _numInBatch)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b36c66b.
//
// Solidity: function initialize(bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntrySession) Initialize(_root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.Initialize(&_OutboxEntry.TransactOpts, _root, _numInBatch)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b36c66b.
//
// Solidity: function initialize(bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) Initialize(_root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.Initialize(&_OutboxEntry.TransactOpts, _root, _numInBatch)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns(uint256)
func (_OutboxEntry *OutboxEntryTransactor) SpendOutput(opts *bind.TransactOpts, _root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "spendOutput", _root, _id)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns(uint256)
func (_OutboxEntry *OutboxEntrySession) SpendOutput(_root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, _root, _id)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns(uint256)
func (_OutboxEntry *OutboxEntryTransactorSession) SpendOutput(_root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, _root, _id)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxySession) Receive() (*types.Transaction, error) {
	return _Proxy.Contract.Receive(&_Proxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _Proxy.Contract.Receive(&_Proxy.TransactOpts)
}

// ProxyAdminABI is the input ABI used to generate the binding from.
const ProxyAdminABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ProxyAdminBin is the compiled bytecode used for deploying new contracts.
var ProxyAdminBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6108658061007d6000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100bf5780637eff275e146100d65780638da5cb5b146101115780639623609d1461012657806399a88ec4146101e3578063f2fde38b1461021e578063f3b7dead14610251575b600080fd5b34801561007c57600080fd5b506100a36004803603602081101561009357600080fd5b50356001600160a01b0316610284565b604080516001600160a01b039092168252519081900360200190f35b3480156100cb57600080fd5b506100d4610316565b005b3480156100e257600080fd5b506100d4600480360360408110156100f957600080fd5b506001600160a01b03813581169160200135166103b0565b34801561011d57600080fd5b506100a361047d565b6100d46004803603606081101561013c57600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460018302840111600160201b831117156101a257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061048c945050505050565b3480156101ef57600080fd5b506100d46004803603604081101561020657600080fd5b506001600160a01b03813581169160200135166105c5565b34801561022a57600080fd5b506100d46004803603602081101561024157600080fd5b50356001600160a01b0316610676565b34801561025d57600080fd5b506100a36004803603602081101561027457600080fd5b50356001600160a01b0316610766565b6000806060836001600160a01b03166040518080635c60da1b60e01b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b606091505b5091509150816102f757600080fd5b80806020019051602081101561030c57600080fd5b5051949350505050565b61031e6107c5565b6001600160a01b031661032f61047d565b6001600160a01b031614610378576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610810833981519152908390a3600080546001600160a01b0319169055565b6103b86107c5565b6001600160a01b03166103c961047d565b6001600160a01b031614610412576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316638f283970826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b505af1158015610475573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b6104946107c5565b6001600160a01b03166104a561047d565b6001600160a01b0316146104ee576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b826001600160a01b0316634f1ef2863484846040518463ffffffff1660e01b815260040180836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561055b578181015183820152602001610543565b50505050905090810190601f1680156105885780820380516001836020036101000a031916815260200191505b5093505050506000604051808303818588803b1580156105a757600080fd5b505af11580156105bb573d6000803e3d6000fd5b5050505050505050565b6105cd6107c5565b6001600160a01b03166105de61047d565b6001600160a01b031614610627576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316633659cfe6826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b61067e6107c5565b6001600160a01b031661068f61047d565b6001600160a01b0316146106d8576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b6001600160a01b03811661071d5760405162461bcd60e51b81526004018080602001828103825260268152602001806107ca6026913960400191505060405180910390fd5b600080546040516001600160a01b038085169392169160008051602061081083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806060836001600160a01b031660405180806303e1469160e61b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122036b5070566a1002642d9947563bea5cf538675280563c6ec32ab459f386727a264736f6c634300060c0033"

// DeployProxyAdmin deploys a new Ethereum contract, binding an instance of ProxyAdmin to it.
func DeployProxyAdmin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProxyAdmin, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAdminABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyAdminBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// ProxyAdmin is an auto generated Go binding around an Ethereum contract.
type ProxyAdmin struct {
	ProxyAdminCaller     // Read-only binding to the contract
	ProxyAdminTransactor // Write-only binding to the contract
	ProxyAdminFilterer   // Log filterer for contract events
}

// ProxyAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyAdminSession struct {
	Contract     *ProxyAdmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyAdminCallerSession struct {
	Contract *ProxyAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProxyAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyAdminTransactorSession struct {
	Contract     *ProxyAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProxyAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyAdminRaw struct {
	Contract *ProxyAdmin // Generic contract binding to access the raw methods on
}

// ProxyAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyAdminCallerRaw struct {
	Contract *ProxyAdminCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyAdminTransactorRaw struct {
	Contract *ProxyAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyAdmin creates a new instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdmin(address common.Address, backend bind.ContractBackend) (*ProxyAdmin, error) {
	contract, err := bindProxyAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// NewProxyAdminCaller creates a new read-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminCaller(address common.Address, caller bind.ContractCaller) (*ProxyAdminCaller, error) {
	contract, err := bindProxyAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminCaller{contract: contract}, nil
}

// NewProxyAdminTransactor creates a new write-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyAdminTransactor, error) {
	contract, err := bindProxyAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminTransactor{contract: contract}, nil
}

// NewProxyAdminFilterer creates a new log filterer instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyAdminFilterer, error) {
	contract, err := bindProxyAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminFilterer{contract: contract}, nil
}

// bindProxyAdmin binds a generic wrapper to an already deployed contract.
func bindProxyAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.ProxyAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transact(opts, method, params...)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyAdmin(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyAdmin", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyImplementation(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyImplementation", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, proxy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "changeProxyAdmin", proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, proxy, newAdmin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactor) Upgrade(opts *bind.TransactOpts, proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgrade", proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, proxy, implementation)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactor) UpgradeAndCall(opts *bind.TransactOpts, proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgradeAndCall", proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, proxy, implementation, data)
}

// ProxyAdminOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferredIterator struct {
	Event *ProxyAdminOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProxyAdminOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAdminOwnershipTransferred)
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
		it.Event = new(ProxyAdminOwnershipTransferred)
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
func (it *ProxyAdminOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAdminOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAdminOwnershipTransferred represents a OwnershipTransferred event raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProxyAdminOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminOwnershipTransferredIterator{contract: _ProxyAdmin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyAdminOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAdminOwnershipTransferred)
				if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyAdminOwnershipTransferred, error) {
	event := new(ProxyAdminOwnershipTransferred)
	if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCoreABI is the input ABI used to generate the binding from.
const RollupCoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupCoreBin is the compiled bytecode used for deploying new contracts.
var RollupCoreBin = "0x608060405234801561001057600080fd5b50610570806100206000396000f3fe608060405234801561001057600080fd5b50600436106100e65760003560e01c80632f30cabd146100eb5780633e96576e146101235780634f0f4aa9146101495780636177fd181461018257806362a82d7d146101bc57806363721d6b146101d957806365f7f80d146101e157806369fd251c146101e95780637ba9534a1461020f5780638640ce5f1461021757806391c657e81461021f578063d01e660214610245578063d735e21d14610262578063dff697871461026a578063e8bd492214610272578063ef40a670146102ce578063f33e1fac146102f4578063f8d1f19414610311575b600080fd5b6101116004803603602081101561010157600080fd5b50356001600160a01b031661032e565b60408051918252519081900360200190f35b6101116004803603602081101561013957600080fd5b50356001600160a01b031661034d565b6101666004803603602081101561015f57600080fd5b503561036b565b604080516001600160a01b039092168252519081900360200190f35b6101a86004803603602081101561019857600080fd5b50356001600160a01b0316610386565b604080519115158252519081900360200190f35b610166600480360360208110156101d257600080fd5b50356103ae565b6101116103d8565b6101116103de565b610166600480360360208110156101ff57600080fd5b50356001600160a01b03166103e4565b610111610405565b61011161040b565b6101a86004803603602081101561023557600080fd5b50356001600160a01b0316610411565b6101666004803603602081101561025b57600080fd5b503561046b565b61011161049a565b6101116104a0565b6102986004803603602081101561028857600080fd5b50356001600160a01b03166104a6565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b610111600480360360208110156102e457600080fd5b50356001600160a01b03166104e2565b6101116004803603602081101561030a57600080fd5b5035610500565b6101116004803603602081101561032757600080fd5b5035610528565b6001600160a01b0381166000908152600960205260409020545b919050565b6001600160a01b031660009081526007602052604090206001015490565b6000908152600460205260409020546001600160a01b031690565b6001600160a01b0316600090815260076020526040902060030154600160a01b900460ff1690565b6000600682815481106103bd57fe5b6000918252602090912001546001600160a01b031692915050565b60085490565b60005490565b6001600160a01b039081166000908152600760205260409020600301541690565b60025490565b60035490565b6000805b600854811015610462576008818154811061042c57fe5b60009182526020909120600290910201546001600160a01b038481169116141561045a576001915050610348565b600101610415565b50600092915050565b60006008828154811061047a57fe5b60009182526020909120600290910201546001600160a01b031692915050565b60015490565b60065490565b6007602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6001600160a01b031660009081526007602052604090206002015490565b60006008828154811061050f57fe5b9060005260206000209060020201600101549050919050565b6000908152600560205260409020549056fea264697066735822122097938160b753385047689b820b8399c4620753e0cf7c5c5e4c9d3555f748a7b564736f6c634300060c0033"

// DeployRollupCore deploys a new Ethereum contract, binding an instance of RollupCore to it.
func DeployRollupCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupCore, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCore{RollupCoreCaller: RollupCoreCaller{contract: contract}, RollupCoreTransactor: RollupCoreTransactor{contract: contract}, RollupCoreFilterer: RollupCoreFilterer{contract: contract}}, nil
}

// RollupCore is an auto generated Go binding around an Ethereum contract.
type RollupCore struct {
	RollupCoreCaller     // Read-only binding to the contract
	RollupCoreTransactor // Write-only binding to the contract
	RollupCoreFilterer   // Log filterer for contract events
}

// RollupCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCoreSession struct {
	Contract     *RollupCore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCoreCallerSession struct {
	Contract *RollupCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCoreTransactorSession struct {
	Contract     *RollupCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCoreRaw struct {
	Contract *RollupCore // Generic contract binding to access the raw methods on
}

// RollupCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCoreCallerRaw struct {
	Contract *RollupCoreCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCoreTransactorRaw struct {
	Contract *RollupCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCore creates a new instance of RollupCore, bound to a specific deployed contract.
func NewRollupCore(address common.Address, backend bind.ContractBackend) (*RollupCore, error) {
	contract, err := bindRollupCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCore{RollupCoreCaller: RollupCoreCaller{contract: contract}, RollupCoreTransactor: RollupCoreTransactor{contract: contract}, RollupCoreFilterer: RollupCoreFilterer{contract: contract}}, nil
}

// NewRollupCoreCaller creates a new read-only instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreCaller(address common.Address, caller bind.ContractCaller) (*RollupCoreCaller, error) {
	contract, err := bindRollupCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCoreCaller{contract: contract}, nil
}

// NewRollupCoreTransactor creates a new write-only instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCoreTransactor, error) {
	contract, err := bindRollupCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCoreTransactor{contract: contract}, nil
}

// NewRollupCoreFilterer creates a new log filterer instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCoreFilterer, error) {
	contract, err := bindRollupCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCoreFilterer{contract: contract}, nil
}

// bindRollupCore binds a generic wrapper to an already deployed contract.
func bindRollupCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCore *RollupCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCore.Contract.RollupCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCore *RollupCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCore.Contract.RollupCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCore *RollupCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCore.Contract.RollupCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCore *RollupCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCore *RollupCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCore *RollupCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCore.Contract.contract.Transact(opts, method, params...)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "_stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _RollupCore.Contract.StakerMap(&_RollupCore.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _RollupCore.Contract.StakerMap(&_RollupCore.CallOpts, arg0)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.AmountStaked(&_RollupCore.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.AmountStaked(&_RollupCore.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _RollupCore.Contract.CurrentChallenge(&_RollupCore.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _RollupCore.Contract.CurrentChallenge(&_RollupCore.CallOpts, staker)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreSession) FirstUnresolvedNode() (*big.Int, error) {
	return _RollupCore.Contract.FirstUnresolvedNode(&_RollupCore.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _RollupCore.Contract.FirstUnresolvedNode(&_RollupCore.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetNode(&_RollupCore.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetNode(&_RollupCore.CallOpts, nodeNum)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _RollupCore.Contract.GetNodeHash(&_RollupCore.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _RollupCore.Contract.GetNodeHash(&_RollupCore.CallOpts, index)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetStakerAddress(&_RollupCore.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetStakerAddress(&_RollupCore.CallOpts, stakerNum)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreSession) IsStaked(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsStaked(&_RollupCore.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsStaked(&_RollupCore.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupCore *RollupCoreCaller) IsZombie(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "isZombie", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupCore *RollupCoreSession) IsZombie(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsZombie(&_RollupCore.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupCore *RollupCoreCallerSession) IsZombie(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsZombie(&_RollupCore.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreSession) LastStakeBlock() (*big.Int, error) {
	return _RollupCore.Contract.LastStakeBlock(&_RollupCore.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LastStakeBlock() (*big.Int, error) {
	return _RollupCore.Contract.LastStakeBlock(&_RollupCore.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestConfirmed() (*big.Int, error) {
	return _RollupCore.Contract.LatestConfirmed(&_RollupCore.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestConfirmed() (*big.Int, error) {
	return _RollupCore.Contract.LatestConfirmed(&_RollupCore.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestNodeCreated() (*big.Int, error) {
	return _RollupCore.Contract.LatestNodeCreated(&_RollupCore.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _RollupCore.Contract.LatestNodeCreated(&_RollupCore.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.LatestStakedNode(&_RollupCore.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.LatestStakedNode(&_RollupCore.CallOpts, staker)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreSession) StakerCount() (*big.Int, error) {
	return _RollupCore.Contract.StakerCount(&_RollupCore.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) StakerCount() (*big.Int, error) {
	return _RollupCore.Contract.StakerCount(&_RollupCore.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _RollupCore.Contract.WithdrawableFunds(&_RollupCore.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _RollupCore.Contract.WithdrawableFunds(&_RollupCore.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.ZombieAddress(&_RollupCore.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.ZombieAddress(&_RollupCore.CallOpts, zombieNum)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreSession) ZombieCount() (*big.Int, error) {
	return _RollupCore.Contract.ZombieCount(&_RollupCore.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) ZombieCount() (*big.Int, error) {
	return _RollupCore.Contract.ZombieCount(&_RollupCore.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _RollupCore.Contract.ZombieLatestStakedNode(&_RollupCore.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _RollupCore.Contract.ZombieLatestStakedNode(&_RollupCore.CallOpts, zombieNum)
}

// RollupCreatorNoProxyABI is the input ABI used to generate the binding from.
const RollupCreatorNoProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeCreator\",\"outputs\":[{\"internalType\":\"contractBridgeCreator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerDelaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"name\":\"createRollupNoProxy\",\"outputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTemplate\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBridgeCreator\",\"name\":\"_bridgeCreator\",\"type\":\"address\"},{\"internalType\":\"contractICloneable\",\"name\":\"_rollupTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"}],\"name\":\"setTemplates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupCreatorNoProxyBin is the compiled bytecode used for deploying new contracts.
var RollupCreatorNoProxyBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b610c808061007d6000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c80635dbaf68b14610088578063715018a6146100ac5780638689d996146100b65780638babc4d9146100be5780638da5cb5b1461017d578063c8a7cb2114610185578063d93fe9c4146101c3578063f2fde38b146101cb578063f860cefa146101f1575b600080fd5b6100906101f9565b604080516001600160a01b039092168252519081900360200190f35b6100b4610208565b005b6100906102a2565b61009060048036036101608110156100d557600080fd5b8135916020810135916040820135916060810135916080820135916001600160a01b0360a082013581169260c083013582169260e0810135909216916101008101359161012082013591908101906101608101610140820135600160201b81111561013f57600080fd5b82018360208201111561015157600080fd5b803590602001918460018302840111600160201b8311171561017257600080fd5b5090925090506102b1565b610090610367565b6100b46004803603608081101561019b57600080fd5b506001600160a01b038135811691602081013582169160408201358116916060013516610376565b610090610428565b6100b4600480360360208110156101e157600080fd5b50356001600160a01b0316610437565b610090610527565b6003546001600160a01b031681565b610210610536565b6001600160a01b0316610221610367565b6001600160a01b03161461026a576040805162461bcd60e51b81526020600482018190526024820152600080516020610c0b833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610c2b833981519152908390a3600080546001600160a01b0319169055565b6002546001600160a01b031681565b60006103566040518061016001604052808f81526020018e81526020018d81526020018c81526020018b81526020018a6001600160a01b03168152602001896001600160a01b03168152602001886001600160a01b0316815260200187815260200186815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505091525061053a565b9d9c50505050505050505050505050565b6000546001600160a01b031690565b61037e610536565b6001600160a01b031661038f610367565b6001600160a01b0316146103d8576040805162461bcd60e51b81526020600482018190526024820152600080516020610c0b833981519152604482015290519081900360640190fd5b600180546001600160a01b039586166001600160a01b0319918216179091556002805494861694821694909417909355600380549285169284169290921790915560048054919093169116179055565b6004546001600160a01b031681565b61043f610536565b6001600160a01b0316610450610367565b6001600160a01b031614610499576040805162461bcd60e51b81526020600482018190526024820152600080516020610c0b833981519152604482015290519081900360640190fd5b6001600160a01b0381166104de5760405162461bcd60e51b8152600401808060200182810382526026815260200180610be56026913960400191505060405180910390fd5b600080546040516001600160a01b0380851693921691600080516020610c2b83398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b3390565b6000610544610ba8565b600254610559906001600160a01b0316610947565b6001600160a01b0390811660c08084018290526001549086015160e08701516101008801516101208901516040805163ed33557960e01b815294881660048601526024850196909652918616604484015260648301526084820152915192169163ed3355799160a48082019260a0929091908290030181600087803b1580156105e157600080fd5b505af11580156105f5573d6000803e3d6000fd5b505050506040513d60a081101561060b57600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050856020018660400187606001886080018960a001856001600160a01b03166001600160a01b0316815250856001600160a01b03166001600160a01b0316815250856001600160a01b03166001600160a01b0316815250856001600160a01b03166001600160a01b0316815250856001600160a01b03166001600160a01b031681525050505050508060c001516001600160a01b031663fdaf5797846000015185602001518660400151876060015188608001518960a001518a60c001518b61014001516040518060c001604052808c602001516001600160a01b03166001600160a01b031681526020018c604001516001600160a01b03166001600160a01b031681526020018c60a001516001600160a01b03166001600160a01b031681526020018c608001516001600160a01b03166001600160a01b03168152602001600360009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600460009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152506040518a63ffffffff1660e01b8152600401808a8152602001898152602001888152602001878152602001868152602001856001600160a01b03168152602001846001600160a01b031681526020018060200183600660200280838360005b83811015610851578181015183820152602001610839565b50505050905001828103825284818151815260200191508051906020019080838360005b8381101561088d578181015183820152602001610875565b50505050905090810190601f1680156108ba5780820380516001836020036101000a031916815260200191505b509a5050505050505050505050600060405180830381600087803b1580156108e157600080fd5b505af11580156108f5573d6000803e3d6000fd5b50505060c0820151604080516001600160a01b039092168252517f84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c92509081900360200190a160c0015190505b919050565b60008161095c816001600160a01b0316610b05565b604051806040016040528060188152602001772727afa1a7a72a2920a1aa2fa1a627a722afa6a0a9aa22a960411b81525090610a165760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109db5781810151838201526020016109c3565b50505050905090810190601f168015610a085780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b158015610a5057600080fd5b505afa158015610a64573d6000803e3d6000fd5b505050506040513d6020811015610a7a57600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b602082015290610aeb5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109db5781810151838201526020016109c3565b50610afe836001600160a01b0316610b0b565b9392505050565b3b151590565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116610942576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091529056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122070834ca6a604b8da2f1a9467a28cc8eb9a931519e267f1f241558134056c581664736f6c634300060c0033"

// DeployRollupCreatorNoProxy deploys a new Ethereum contract, binding an instance of RollupCreatorNoProxy to it.
func DeployRollupCreatorNoProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupCreatorNoProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorNoProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCreatorNoProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCreatorNoProxy{RollupCreatorNoProxyCaller: RollupCreatorNoProxyCaller{contract: contract}, RollupCreatorNoProxyTransactor: RollupCreatorNoProxyTransactor{contract: contract}, RollupCreatorNoProxyFilterer: RollupCreatorNoProxyFilterer{contract: contract}}, nil
}

// RollupCreatorNoProxy is an auto generated Go binding around an Ethereum contract.
type RollupCreatorNoProxy struct {
	RollupCreatorNoProxyCaller     // Read-only binding to the contract
	RollupCreatorNoProxyTransactor // Write-only binding to the contract
	RollupCreatorNoProxyFilterer   // Log filterer for contract events
}

// RollupCreatorNoProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCreatorNoProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorNoProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCreatorNoProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorNoProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCreatorNoProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorNoProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCreatorNoProxySession struct {
	Contract     *RollupCreatorNoProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupCreatorNoProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCreatorNoProxyCallerSession struct {
	Contract *RollupCreatorNoProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RollupCreatorNoProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCreatorNoProxyTransactorSession struct {
	Contract     *RollupCreatorNoProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RollupCreatorNoProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCreatorNoProxyRaw struct {
	Contract *RollupCreatorNoProxy // Generic contract binding to access the raw methods on
}

// RollupCreatorNoProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCreatorNoProxyCallerRaw struct {
	Contract *RollupCreatorNoProxyCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCreatorNoProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCreatorNoProxyTransactorRaw struct {
	Contract *RollupCreatorNoProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCreatorNoProxy creates a new instance of RollupCreatorNoProxy, bound to a specific deployed contract.
func NewRollupCreatorNoProxy(address common.Address, backend bind.ContractBackend) (*RollupCreatorNoProxy, error) {
	contract, err := bindRollupCreatorNoProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxy{RollupCreatorNoProxyCaller: RollupCreatorNoProxyCaller{contract: contract}, RollupCreatorNoProxyTransactor: RollupCreatorNoProxyTransactor{contract: contract}, RollupCreatorNoProxyFilterer: RollupCreatorNoProxyFilterer{contract: contract}}, nil
}

// NewRollupCreatorNoProxyCaller creates a new read-only instance of RollupCreatorNoProxy, bound to a specific deployed contract.
func NewRollupCreatorNoProxyCaller(address common.Address, caller bind.ContractCaller) (*RollupCreatorNoProxyCaller, error) {
	contract, err := bindRollupCreatorNoProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyCaller{contract: contract}, nil
}

// NewRollupCreatorNoProxyTransactor creates a new write-only instance of RollupCreatorNoProxy, bound to a specific deployed contract.
func NewRollupCreatorNoProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCreatorNoProxyTransactor, error) {
	contract, err := bindRollupCreatorNoProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyTransactor{contract: contract}, nil
}

// NewRollupCreatorNoProxyFilterer creates a new log filterer instance of RollupCreatorNoProxy, bound to a specific deployed contract.
func NewRollupCreatorNoProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCreatorNoProxyFilterer, error) {
	contract, err := bindRollupCreatorNoProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyFilterer{contract: contract}, nil
}

// bindRollupCreatorNoProxy binds a generic wrapper to an already deployed contract.
func bindRollupCreatorNoProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorNoProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreatorNoProxy.Contract.RollupCreatorNoProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.RollupCreatorNoProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.RollupCreatorNoProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreatorNoProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.contract.Transact(opts, method, params...)
}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) BridgeCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "bridgeCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) BridgeCreator() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.BridgeCreator(&_RollupCreatorNoProxy.CallOpts)
}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) BridgeCreator() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.BridgeCreator(&_RollupCreatorNoProxy.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) ChallengeFactory() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.ChallengeFactory(&_RollupCreatorNoProxy.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) ChallengeFactory() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.ChallengeFactory(&_RollupCreatorNoProxy.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) NodeFactory() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.NodeFactory(&_RollupCreatorNoProxy.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) NodeFactory() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.NodeFactory(&_RollupCreatorNoProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) Owner() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.Owner(&_RollupCreatorNoProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) Owner() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.Owner(&_RollupCreatorNoProxy.CallOpts)
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) RollupTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "rollupTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) RollupTemplate() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.RollupTemplate(&_RollupCreatorNoProxy.CallOpts)
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) RollupTemplate() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.RollupTemplate(&_RollupCreatorNoProxy.CallOpts)
}

// CreateRollupNoProxy is a paid mutator transaction binding the contract method 0x8babc4d9.
//
// Solidity: function createRollupNoProxy(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactor) CreateRollupNoProxy(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.contract.Transact(opts, "createRollupNoProxy", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// CreateRollupNoProxy is a paid mutator transaction binding the contract method 0x8babc4d9.
//
// Solidity: function createRollupNoProxy(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) CreateRollupNoProxy(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.CreateRollupNoProxy(&_RollupCreatorNoProxy.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// CreateRollupNoProxy is a paid mutator transaction binding the contract method 0x8babc4d9.
//
// Solidity: function createRollupNoProxy(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorSession) CreateRollupNoProxy(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.CreateRollupNoProxy(&_RollupCreatorNoProxy.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.RenounceOwnership(&_RollupCreatorNoProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.RenounceOwnership(&_RollupCreatorNoProxy.TransactOpts)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactor) SetTemplates(opts *bind.TransactOpts, _bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.contract.Transact(opts, "setTemplates", _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) SetTemplates(_bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.SetTemplates(&_RollupCreatorNoProxy.TransactOpts, _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorSession) SetTemplates(_bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.SetTemplates(&_RollupCreatorNoProxy.TransactOpts, _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.TransferOwnership(&_RollupCreatorNoProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.TransferOwnership(&_RollupCreatorNoProxy.TransactOpts, newOwner)
}

// RollupCreatorNoProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RollupCreatorNoProxy contract.
type RollupCreatorNoProxyOwnershipTransferredIterator struct {
	Event *RollupCreatorNoProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RollupCreatorNoProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorNoProxyOwnershipTransferred)
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
		it.Event = new(RollupCreatorNoProxyOwnershipTransferred)
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
func (it *RollupCreatorNoProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorNoProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorNoProxyOwnershipTransferred represents a OwnershipTransferred event raised by the RollupCreatorNoProxy contract.
type RollupCreatorNoProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupCreatorNoProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreatorNoProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyOwnershipTransferredIterator{contract: _RollupCreatorNoProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupCreatorNoProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreatorNoProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorNoProxyOwnershipTransferred)
				if err := _RollupCreatorNoProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) ParseOwnershipTransferred(log types.Log) (*RollupCreatorNoProxyOwnershipTransferred, error) {
	event := new(RollupCreatorNoProxyOwnershipTransferred)
	if err := _RollupCreatorNoProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorNoProxyRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the RollupCreatorNoProxy contract.
type RollupCreatorNoProxyRollupCreatedIterator struct {
	Event *RollupCreatorNoProxyRollupCreated // Event containing the contract specifics and raw log

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
func (it *RollupCreatorNoProxyRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorNoProxyRollupCreated)
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
		it.Event = new(RollupCreatorNoProxyRollupCreated)
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
func (it *RollupCreatorNoProxyRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorNoProxyRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorNoProxyRollupCreated represents a RollupCreated event raised by the RollupCreatorNoProxy contract.
type RollupCreatorNoProxyRollupCreated struct {
	RollupAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*RollupCreatorNoProxyRollupCreatedIterator, error) {

	logs, sub, err := _RollupCreatorNoProxy.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyRollupCreatedIterator{contract: _RollupCreatorNoProxy.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupCreatorNoProxyRollupCreated) (event.Subscription, error) {

	logs, sub, err := _RollupCreatorNoProxy.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorNoProxyRollupCreated)
				if err := _RollupCreatorNoProxy.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) ParseRollupCreated(log types.Log) (*RollupCreatorNoProxyRollupCreated, error) {
	event := new(RollupCreatorNoProxyRollupCreated)
	if err := _RollupCreatorNoProxy.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupEventBridgeABI is the input ABI used to generate the binding from.
const RollupEventBridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"claimNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"nodeConfirmed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"}],\"name\":\"nodeCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"nodeRejected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraConfig\",\"type\":\"bytes\"}],\"name\":\"rollupInitialized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"stakeCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupEventBridgeBin is the compiled bytecode used for deploying new contracts.
var RollupEventBridgeBin = "0x608060405234801561001057600080fd5b506000805460ff19166001179055610a878061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c806316b9109b1461007d57806330a826b41461009c578063485cc955146100b957806364126c7c146100e75780636f791d29146101135780638b8ca1991461012f578063b0f2af2914610167578063f03c04a514610206575b600080fd5b61009a6004803603602081101561009357600080fd5b5035610232565b005b61009a600480360360208110156100b257600080fd5b50356102b3565b61009a600480360360408110156100cf57600080fd5b506001600160a01b0381358116916020013516610331565b61009a600480360360408110156100fd57600080fd5b50803590602001356001600160a01b03166103b6565b61011b6105dd565b604080519115158252519081900360200190f35b61009a6004803603608081101561014557600080fd5b50803590602081013590604081013590606001356001600160a01b03166105e6565b61009a600480360360e081101561017d57600080fd5b8135916020810135916040820135916060810135916001600160a01b03608083013581169260a08101359091169181019060e0810160c0820135600160201b8111156101c857600080fd5b8201836020820111156101da57600080fd5b803590602001918460018302840111600160201b831117156101fb57600080fd5b509092509050610687565b61009a6004803603604081101561021c57600080fd5b506001600160a01b038135169060200135610879565b6001546001600160a01b0316331461027f576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f81b602082015260218082018490528251808303909101815260419091019091526102b090610913565b50565b6001546001600160a01b03163314610300576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f91b602082015260218082018490528251808303909101815260419091019091526102b090610913565b6001546001600160a01b03161561037e576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b60008054610100600160a81b0319166101006001600160a01b0394851602179055600180546001600160a01b03191691909216179055565b6001546001600160a01b03163314610403576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60015460408051634f0f4aa960e01b81526004810185905290516001600160a01b03909216916000918391634f0f4aa991602480820192602092909190829003018186803b15801561045457600080fd5b505afa158015610468573d6000803e3d6000fd5b505050506040513d602081101561047e57600080fd5b5051604080516348b4573960e11b81526001600160a01b038681166004830152915192935090831691639168ae7291602480820192602092909190829003018186803b1580156104cd57600080fd5b505afa1580156104e1573d6000803e3d6000fd5b505050506040513d60208110156104f757600080fd5b5051610537576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b816001600160a01b0316632b2af0ab856040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561057b57600080fd5b505afa15801561058f573d6000803e3d6000fd5b505060408051600160fa1b6020820152602181018890526001600160a01b0387166041808301919091528251808303909101815260619091019091526105d792509050610913565b50505050565b60005460ff1690565b6001546001600160a01b03163314610633576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600060208201526021810186905260418101859052436061820152608181018490526001600160a01b03831660a1808301919091528251808303909101815260c19091019091526105d790610913565b6001546001600160a01b031633146106d4576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b6060888888888860601b60601c6001600160a01b03168860601b60601c6001600160a01b03168888604051602001808981526020018881526020018781526020018681526020018581526020018481526020018383808284376040805191909301818103601f190182528084526000805483516020808601919091206302bbfad160e01b855260048086015233602486015260448501529551939f50909d5061010090046001600160a01b03169b506302bbfad19a5060648082019a509398509096508690039091019350849250899150889050803b1580156107b657600080fd5b505af11580156107ca573d6000803e3d6000fd5b505050506040513d60208110156107e057600080fd5b505160408051602080825285518282015285519394508493600080516020610a32833981519152938793928392918301919085019080838360005b8381101561083357818101518382015260200161081b565b50505050905090810190601f1680156108605780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050505050505050565b6001546001600160a01b031633146108c6576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600360f81b60208201526001600160a01b0384166021820152604181018390524360618083019190915282518083039091018152608190910190915261090f90610913565b5050565b600080548251602080850191909120604080516302bbfad160e01b8152600860048201523360248201526044810192909252516101009093046001600160a01b0316936302bbfad193606480840194939192918390030190829087803b15801561097c57600080fd5b505af1158015610990573d6000803e3d6000fd5b505050506040513d60208110156109a657600080fd5b50516040805160208082528451828201528451600080516020610a32833981519152938693928392918301919085019080838360005b838110156109f45781810151838201526020016109dc565b50505050905090810190601f168015610a215780820380516001836020036101000a031916815260200191505b509250505060405180910390a25056feff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60ba26469706673582212208dc0296bd65615ead206d527d65af5d42dee111320c869c9002172ca197bed1164736f6c634300060c0033"

// DeployRollupEventBridge deploys a new Ethereum contract, binding an instance of RollupEventBridge to it.
func DeployRollupEventBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupEventBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupEventBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupEventBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupEventBridge{RollupEventBridgeCaller: RollupEventBridgeCaller{contract: contract}, RollupEventBridgeTransactor: RollupEventBridgeTransactor{contract: contract}, RollupEventBridgeFilterer: RollupEventBridgeFilterer{contract: contract}}, nil
}

// RollupEventBridge is an auto generated Go binding around an Ethereum contract.
type RollupEventBridge struct {
	RollupEventBridgeCaller     // Read-only binding to the contract
	RollupEventBridgeTransactor // Write-only binding to the contract
	RollupEventBridgeFilterer   // Log filterer for contract events
}

// RollupEventBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupEventBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupEventBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupEventBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupEventBridgeSession struct {
	Contract     *RollupEventBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RollupEventBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupEventBridgeCallerSession struct {
	Contract *RollupEventBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RollupEventBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupEventBridgeTransactorSession struct {
	Contract     *RollupEventBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RollupEventBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupEventBridgeRaw struct {
	Contract *RollupEventBridge // Generic contract binding to access the raw methods on
}

// RollupEventBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupEventBridgeCallerRaw struct {
	Contract *RollupEventBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// RollupEventBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupEventBridgeTransactorRaw struct {
	Contract *RollupEventBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupEventBridge creates a new instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridge(address common.Address, backend bind.ContractBackend) (*RollupEventBridge, error) {
	contract, err := bindRollupEventBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridge{RollupEventBridgeCaller: RollupEventBridgeCaller{contract: contract}, RollupEventBridgeTransactor: RollupEventBridgeTransactor{contract: contract}, RollupEventBridgeFilterer: RollupEventBridgeFilterer{contract: contract}}, nil
}

// NewRollupEventBridgeCaller creates a new read-only instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeCaller(address common.Address, caller bind.ContractCaller) (*RollupEventBridgeCaller, error) {
	contract, err := bindRollupEventBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeCaller{contract: contract}, nil
}

// NewRollupEventBridgeTransactor creates a new write-only instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupEventBridgeTransactor, error) {
	contract, err := bindRollupEventBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeTransactor{contract: contract}, nil
}

// NewRollupEventBridgeFilterer creates a new log filterer instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupEventBridgeFilterer, error) {
	contract, err := bindRollupEventBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeFilterer{contract: contract}, nil
}

// bindRollupEventBridge binds a generic wrapper to an already deployed contract.
func bindRollupEventBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupEventBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupEventBridge *RollupEventBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupEventBridge.Contract.RollupEventBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupEventBridge *RollupEventBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupEventBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupEventBridge *RollupEventBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupEventBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupEventBridge *RollupEventBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupEventBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupEventBridge *RollupEventBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupEventBridge *RollupEventBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.contract.Transact(opts, method, params...)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_RollupEventBridge *RollupEventBridgeCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RollupEventBridge.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_RollupEventBridge *RollupEventBridgeSession) IsMaster() (bool, error) {
	return _RollupEventBridge.Contract.IsMaster(&_RollupEventBridge.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_RollupEventBridge *RollupEventBridgeCallerSession) IsMaster() (bool, error) {
	return _RollupEventBridge.Contract.IsMaster(&_RollupEventBridge.CallOpts)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) ClaimNode(opts *bind.TransactOpts, nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "claimNode", nodeNum, staker)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeSession) ClaimNode(nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.ClaimNode(&_RollupEventBridge.TransactOpts, nodeNum, staker)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) ClaimNode(nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.ClaimNode(&_RollupEventBridge.TransactOpts, nodeNum, staker)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _rollup) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "initialize", _bridge, _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _rollup) returns()
func (_RollupEventBridge *RollupEventBridgeSession) Initialize(_bridge common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.Initialize(&_RollupEventBridge.TransactOpts, _bridge, _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _rollup) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) Initialize(_bridge common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.Initialize(&_RollupEventBridge.TransactOpts, _bridge, _rollup)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeConfirmed(opts *bind.TransactOpts, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeConfirmed", nodeNum)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeConfirmed(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeConfirmed(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeConfirmed(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeConfirmed(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeCreated(opts *bind.TransactOpts, nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeCreated", nodeNum, prev, deadline, asserter)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeCreated(nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeCreated(&_RollupEventBridge.TransactOpts, nodeNum, prev, deadline, asserter)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeCreated(nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeCreated(&_RollupEventBridge.TransactOpts, nodeNum, prev, deadline, asserter)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeRejected(opts *bind.TransactOpts, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeRejected", nodeNum)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeRejected(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeRejected(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeRejected(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeRejected(&_RollupEventBridge.TransactOpts, nodeNum)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) RollupInitialized(opts *bind.TransactOpts, confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "rollupInitialized", confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeSession) RollupInitialized(confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupInitialized(&_RollupEventBridge.TransactOpts, confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) RollupInitialized(confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupInitialized(&_RollupEventBridge.TransactOpts, confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) StakeCreated(opts *bind.TransactOpts, staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "stakeCreated", staker, nodeNum)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) StakeCreated(staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.StakeCreated(&_RollupEventBridge.TransactOpts, staker, nodeNum)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) StakeCreated(staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.StakeCreated(&_RollupEventBridge.TransactOpts, staker, nodeNum)
}

// RollupEventBridgeInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredIterator struct {
	Event *RollupEventBridgeInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupEventBridgeInboxMessageDelivered)
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
		it.Event = new(RollupEventBridgeInboxMessageDelivered)
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
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupEventBridgeInboxMessageDelivered represents a InboxMessageDelivered event raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*RollupEventBridgeInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeInboxMessageDeliveredIterator{contract: _RollupEventBridge.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *RollupEventBridgeInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupEventBridgeInboxMessageDelivered)
				if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) ParseInboxMessageDelivered(log types.Log) (*RollupEventBridgeInboxMessageDelivered, error) {
	event := new(RollupEventBridgeInboxMessageDelivered)
	if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupEventBridgeInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredFromOriginIterator struct {
	Event *RollupEventBridgeInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
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
		it.Event = new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
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
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupEventBridgeInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*RollupEventBridgeInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeInboxMessageDeliveredFromOriginIterator{contract: _RollupEventBridge.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *RollupEventBridgeInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
				if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*RollupEventBridgeInboxMessageDeliveredFromOrigin, error) {
	event := new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
	if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxABI is the input ABI used to generate the binding from.
const SequencerInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"afterAccAndDelayed\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"DelayedInboxForced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedInbox\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"l1BlockAndTimestamp\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_delayedInbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inboxCount\",\"type\":\"uint256\"}],\"name\":\"proveBatchContainsSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SequencerInboxBin is the compiled bytecode used for deploying new contracts.
var SequencerInboxBin = "0x608060405234801561001057600080fd5b506000805460ff1916600117905561154d8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c806306cc91b2146100a95780630a17a46414610130578063342025fa146102025780633dbcc8d11461021c5780635c1bba38146102245780636f791d29146102485780638af00545146102645780639afc500d146102b1578063b71939b114610381578063d9dd67ab14610389578063e367a2c1146103a6578063eb990c59146103ae575b600080fd5b610117600480360360408110156100bf57600080fd5b810190602081018135600160201b8111156100d957600080fd5b8201836020820111156100eb57600080fd5b803590602001918460018302840111600160201b8311171561010c57600080fd5b9193509150356103ea565b6040805192835260208301919091528051918290030190f35b610200600480360360c081101561014657600080fd5b810190602081018135600160201b81111561016057600080fd5b82018360208201111561017257600080fd5b803590602001918460018302840111600160201b8311171561019357600080fd5b919390929091602081019035600160201b8111156101b057600080fd5b8201836020820111156101c257600080fd5b803590602001918460208302840111600160201b831117156101e357600080fd5b91935091508035906020810135906040810135906060013561053c565b005b61020a610636565b60408051918252519081900360200190f35b61020a61063c565b61022c610642565b604080516001600160a01b039092168252519081900360200190f35b610250610651565b604080519115158252519081900360200190f35b610200600480360361010081101561027b57600080fd5b5080359060ff60208201351690604081019060808101359060a0810135906001600160a01b0360c0820135169060e0013561065a565b610200600480360360c08110156102c757600080fd5b810190602081018135600160201b8111156102e157600080fd5b8201836020820111156102f357600080fd5b803590602001918460018302840111600160201b8311171561031457600080fd5b919390929091602081019035600160201b81111561033157600080fd5b82018360208201111561034357600080fd5b803590602001918460208302840111600160201b8311171561036457600080fd5b9193509150803590602081013590604081013590606001356109c9565b61022c610a89565b61020a6004803603602081101561039f57600080fd5b5035610a98565b61020a610ab6565b610200600480360360808110156103c457600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135610abc565b600080826103fd57506000905080610534565b60008061043f87878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509250610b44915050565b90925090506000811561047a5761047188888560018087038154811061046157fe5b9060005260206000200154610bb2565b90935060010190505b60006001838154811061048957fe5b9060005260206000200154905060006104a48a8a8785610bb2565b90955090508288116104eb576040805162461bcd60e51b815260206004820152600b60248201526a10905510d217d4d510549560aa1b604482015290519081900360640190fd5b8088111561052c576040805162461bcd60e51b815260206004820152600960248201526810905510d217d1539160ba1b604482015290519081900360640190fd5b955093505050505b935093915050565b6002546000806105528b8b8b8b8b8b8b8b610d42565b9150915081837f43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063600254878f8f8f8f8f8f8f8c6001808054905003604051808c81526020018b8152602001806020018060200188815260200187815260200186815260200185815260200184815260200183810383528c8c82818152602001925080828437600083820152601f01601f191690910184810383528a8152602090810191508b908b0280828437600083820152604051601f909101601f19169092018290039f50909d5050505050505050505050505050a35050505050505050505050565b60075481565b60025481565b6005546001600160a01b031681565b60005460ff1690565b60035487116106a4576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b60006106ba8784883560208a01358989886111bf565b6006549091504387359091011061070b576040805162461bcd60e51b815260206004820152601060248201526f4d41585f44454c41595f424c4f434b5360801b604482015290519081900360640190fd5b6007544260208801359091011061075a576040805162461bcd60e51b815260206004820152600e60248201526d4d41585f44454c41595f54494d4560901b604482015290519081900360640190fd5b600060018911156107e457600480546040805163d9dd67ab60e01b81526001198d0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b1580156107b557600080fd5b505afa1580156107c9573d6000803e3d6000fd5b505050506040513d60208110156107df57600080fd5b505190505b6107ee8183611230565b600480546040805163d9dd67ab60e01b81526000198e0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561083e57600080fd5b505afa158015610852573d6000803e3d6000fd5b505050506040513d602081101561086857600080fd5b5051146108b2576040805162461bcd60e51b81526020600482015260136024820152722222a620aca2a22fa0a1a1aaa6aaa620aa27a960691b604482015290519081900360640190fd5b5050600254600154600090156108e2576001805460001981019081106108d457fe5b906000526020600020015490505b60008060006108f484868e434261125c565b92509250925060018290806001815401808255809150506001900390600052602060002001600090919091909150558060028190555083857f85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0838f60405180604001604052808881526020018981525060018080549050036040518085815260200184815260200183600260200280838360005b838110156109a0578181015183820152602001610988565b5050505090500182815260200194505050505060405180910390a3505050505050505050505050565b333214610a0b576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600254600080610a218b8b8b8b8b8b8b8b610d42565b6002546001546040805192835260208301899052828101849052600019909101606083015251929450909250839185917f90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812919081900360800190a35050505050505050505050565b6004546001600160a01b031681565b60018181548110610aa557fe5b600091825260209091200154905081565b60065481565b6004546001600160a01b031615610b09576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600480546001600160a01b039586166001600160a01b0319918216179091556005805494909516931692909217909255600691909155600755565b60008082845110158015610b5c575060208385510310155b610b99576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301610ba78585611496565b915091509250929050565b600080600080600080610bfc8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610b44915050565b8094508199505050610c458a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610b44915050565b8093508199505050610c8e8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610b44915050565b809250819950505082828260405160200180807129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b8152506012018481526020018381526020018281526020019350505050604051602081830303815290604052805190602001209250816001019350868314610d32576040805162461bcd60e51b815260206004820152600960248201526842415443485f41434360b81b604482015290519081900360640190fd5b5095989197509095505050505050565b60055460009081906001600160a01b03163314610d97576040805162461bcd60e51b815260206004820152600e60248201526d27a7262cafa9a2a8aaa2a721a2a960911b604482015290519081900360640190fd5b4360065487011015610de0576040805162461bcd60e51b815260206004820152600d60248201526c109313d0d2d7d513d3d7d3d311609a1b604482015290519081900360640190fd5b43861115610e25576040805162461bcd60e51b815260206004820152600d60248201526c424c4f434b5f544f4f5f4e455760981b604482015290519081900360640190fd5b4260075486011015610e6d576040805162461bcd60e51b815260206004820152600c60248201526b1512535157d513d3d7d3d31160a21b604482015290519081900360640190fd5b42851115610eb1576040805162461bcd60e51b815260206004820152600c60248201526b54494d455f544f4f5f4e455760a01b604482015290519081900360640190fd5b600354841015610efc576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b6001841015610f46576040805162461bcd60e51b8152602060048201526011602482015270135554d517d111531056515117d2539255607a1b604482015290519081900360640190fd5b6001600354101580610f56575086155b610fa1576040805162461bcd60e51b8152602060048201526017602482015276135554d517d111531056515117d253925517d4d5105495604a1b604482015290519081900360640190fd5b60015415610fc957600180546000198101908110610fbb57fe5b906000526020600020015491505b600254826000805b8a8110156110dd5760008e838f8f8f86818110610fea57fe5b90506020020135860192611000939291906114ef565b604051808383808284378083019250505092505050604051809103902090506000600390508d8d8481811061103157fe5b9050602002013560001415611044575060065b600061105682338f8f8b6000896111bf565b905085878260405160200180807129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b81525060120184815260200183815260200182815260200193505050506040516020818303038152906040528051906020012095508e8e858181106110ba57fe5b905060200201358501945086806001019750505050508080600101915050610fd1565b506110eb8284898c8c61125c565b600254929650945092508311611136576040805162461bcd60e51b815260206004820152600b60248201526a08a9aa0a8b2be8482a886960ab1b604482015290519081900360640190fd5b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60182905560028390558582146111af576040805162461bcd60e51b815260206004820152600960248201526841465445525f41434360b81b604482015290519081900360640190fd5b5050509850989650505050505050565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6001600160601b0319166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008060008060035487111561148c576004805460408051633dbcc8d160e01b815290516001600160a01b0390921692633dbcc8d1928282019260209290829003018186803b1580156112ae57600080fd5b505afa1580156112c2573d6000803e3d6000fd5b505050506040513d60208110156112d857600080fd5b5051871115611320576040805162461bcd60e51b815260206004820152600f60248201526e2222a620aca2a22faa27a7afa320a960891b604482015290519081900360640190fd5b600480546040805163d9dd67ab60e01b81526000198b0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561137057600080fd5b505afa158015611384573d6000803e3d6000fd5b505050506040513d602081101561139a57600080fd5b505160035460408051702232b630bcb2b21036b2b9b9b0b3b2b99d60791b602082810191909152603182019d909d52605181018c905260718101839052609181018b905260b18082018590528251808303909101815260d190910190915280519b019a909a2060608051608020919b8a03909a019991925060009061142b9060069083908b908b908f9084906111bf565b604080517129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b602080830191909152603282019e909e52605281018d905260728082019390935281518082039093018352609201905280519b019a909a2060038990559950506001909701965b9895505050505050565b600081602001835110156114e6576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b600080858511156114fe578182fd5b8386111561150a578182fd5b505082019391909203915056fea2646970667358221220aeba8c9d33854dd36c0e27f94d18d1ea3dbb28396ca1bc93ea42b2d2dacb9fc364736f6c634300060c0033"

// DeploySequencerInbox deploys a new Ethereum contract, binding an instance of SequencerInbox to it.
func DeploySequencerInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SequencerInbox, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerInboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SequencerInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SequencerInbox{SequencerInboxCaller: SequencerInboxCaller{contract: contract}, SequencerInboxTransactor: SequencerInboxTransactor{contract: contract}, SequencerInboxFilterer: SequencerInboxFilterer{contract: contract}}, nil
}

// SequencerInbox is an auto generated Go binding around an Ethereum contract.
type SequencerInbox struct {
	SequencerInboxCaller     // Read-only binding to the contract
	SequencerInboxTransactor // Write-only binding to the contract
	SequencerInboxFilterer   // Log filterer for contract events
}

// SequencerInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerInboxSession struct {
	Contract     *SequencerInbox   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SequencerInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerInboxCallerSession struct {
	Contract *SequencerInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SequencerInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerInboxTransactorSession struct {
	Contract     *SequencerInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SequencerInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerInboxRaw struct {
	Contract *SequencerInbox // Generic contract binding to access the raw methods on
}

// SequencerInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerInboxCallerRaw struct {
	Contract *SequencerInboxCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerInboxTransactorRaw struct {
	Contract *SequencerInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerInbox creates a new instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInbox(address common.Address, backend bind.ContractBackend) (*SequencerInbox, error) {
	contract, err := bindSequencerInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerInbox{SequencerInboxCaller: SequencerInboxCaller{contract: contract}, SequencerInboxTransactor: SequencerInboxTransactor{contract: contract}, SequencerInboxFilterer: SequencerInboxFilterer{contract: contract}}, nil
}

// NewSequencerInboxCaller creates a new read-only instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxCaller(address common.Address, caller bind.ContractCaller) (*SequencerInboxCaller, error) {
	contract, err := bindSequencerInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxCaller{contract: contract}, nil
}

// NewSequencerInboxTransactor creates a new write-only instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerInboxTransactor, error) {
	contract, err := bindSequencerInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxTransactor{contract: contract}, nil
}

// NewSequencerInboxFilterer creates a new log filterer instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerInboxFilterer, error) {
	contract, err := bindSequencerInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxFilterer{contract: contract}, nil
}

// bindSequencerInbox binds a generic wrapper to an already deployed contract.
func bindSequencerInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInbox *SequencerInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInbox.Contract.SequencerInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInbox *SequencerInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SequencerInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInbox *SequencerInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SequencerInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInbox *SequencerInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInbox *SequencerInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInbox *SequencerInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInbox.Contract.contract.Transact(opts, method, params...)
}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_SequencerInbox *SequencerInboxCaller) DelayedInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "delayedInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_SequencerInbox *SequencerInboxSession) DelayedInbox() (common.Address, error) {
	return _SequencerInbox.Contract.DelayedInbox(&_SequencerInbox.CallOpts)
}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_SequencerInbox *SequencerInboxCallerSession) DelayedInbox() (common.Address, error) {
	return _SequencerInbox.Contract.DelayedInbox(&_SequencerInbox.CallOpts)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_SequencerInbox *SequencerInboxCaller) InboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "inboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_SequencerInbox *SequencerInboxSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _SequencerInbox.Contract.InboxAccs(&_SequencerInbox.CallOpts, arg0)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _SequencerInbox.Contract.InboxAccs(&_SequencerInbox.CallOpts, arg0)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_SequencerInbox *SequencerInboxCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_SequencerInbox *SequencerInboxSession) IsMaster() (bool, error) {
	return _SequencerInbox.Contract.IsMaster(&_SequencerInbox.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_SequencerInbox *SequencerInboxCallerSession) IsMaster() (bool, error) {
	return _SequencerInbox.Contract.IsMaster(&_SequencerInbox.CallOpts)
}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) MaxDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "maxDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) MaxDelayBlocks() (*big.Int, error) {
	return _SequencerInbox.Contract.MaxDelayBlocks(&_SequencerInbox.CallOpts)
}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) MaxDelayBlocks() (*big.Int, error) {
	return _SequencerInbox.Contract.MaxDelayBlocks(&_SequencerInbox.CallOpts)
}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) MaxDelaySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "maxDelaySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) MaxDelaySeconds() (*big.Int, error) {
	return _SequencerInbox.Contract.MaxDelaySeconds(&_SequencerInbox.CallOpts)
}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) MaxDelaySeconds() (*big.Int, error) {
	return _SequencerInbox.Contract.MaxDelaySeconds(&_SequencerInbox.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) MessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "messageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) MessageCount() (*big.Int, error) {
	return _SequencerInbox.Contract.MessageCount(&_SequencerInbox.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) MessageCount() (*big.Int, error) {
	return _SequencerInbox.Contract.MessageCount(&_SequencerInbox.CallOpts)
}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxCaller) ProveBatchContainsSequenceNumber(opts *bind.CallOpts, proof []byte, inboxCount *big.Int) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "proveBatchContainsSequenceNumber", proof, inboxCount)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxSession) ProveBatchContainsSequenceNumber(proof []byte, inboxCount *big.Int) (*big.Int, [32]byte, error) {
	return _SequencerInbox.Contract.ProveBatchContainsSequenceNumber(&_SequencerInbox.CallOpts, proof, inboxCount)
}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) ProveBatchContainsSequenceNumber(proof []byte, inboxCount *big.Int) (*big.Int, [32]byte, error) {
	return _SequencerInbox.Contract.ProveBatchContainsSequenceNumber(&_SequencerInbox.CallOpts, proof, inboxCount)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_SequencerInbox *SequencerInboxCaller) Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_SequencerInbox *SequencerInboxSession) Sequencer() (common.Address, error) {
	return _SequencerInbox.Contract.Sequencer(&_SequencerInbox.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_SequencerInbox *SequencerInboxCallerSession) Sequencer() (common.Address, error) {
	return _SequencerInbox.Contract.Sequencer(&_SequencerInbox.CallOpts)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x0a17a464.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2Batch", transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x0a17a464.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2Batch(transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2Batch(&_SequencerInbox.TransactOpts, transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x0a17a464.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2Batch(transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2Batch(&_SequencerInbox.TransactOpts, transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x9afc500d.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2BatchFromOrigin(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2BatchFromOrigin", transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x9afc500d.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2BatchFromOrigin(transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInbox.TransactOpts, transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x9afc500d.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2BatchFromOrigin(transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInbox.TransactOpts, transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x8af00545.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x8af00545.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ForceInclusion(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x8af00545.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ForceInclusion(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, uint256 _maxDelayBlocks, uint256 _maxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxTransactor) Initialize(opts *bind.TransactOpts, _delayedInbox common.Address, _sequencer common.Address, _maxDelayBlocks *big.Int, _maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "initialize", _delayedInbox, _sequencer, _maxDelayBlocks, _maxDelaySeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, uint256 _maxDelayBlocks, uint256 _maxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxSession) Initialize(_delayedInbox common.Address, _sequencer common.Address, _maxDelayBlocks *big.Int, _maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, _delayedInbox, _sequencer, _maxDelayBlocks, _maxDelaySeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, uint256 _maxDelayBlocks, uint256 _maxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) Initialize(_delayedInbox common.Address, _sequencer common.Address, _maxDelayBlocks *big.Int, _maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, _delayedInbox, _sequencer, _maxDelayBlocks, _maxDelaySeconds)
}

// SequencerInboxDelayedInboxForcedIterator is returned from FilterDelayedInboxForced and is used to iterate over the raw logs and unpacked data for DelayedInboxForced events raised by the SequencerInbox contract.
type SequencerInboxDelayedInboxForcedIterator struct {
	Event *SequencerInboxDelayedInboxForced // Event containing the contract specifics and raw log

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
func (it *SequencerInboxDelayedInboxForcedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxDelayedInboxForced)
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
		it.Event = new(SequencerInboxDelayedInboxForced)
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
func (it *SequencerInboxDelayedInboxForcedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxDelayedInboxForcedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxDelayedInboxForced represents a DelayedInboxForced event raised by the SequencerInbox contract.
type SequencerInboxDelayedInboxForced struct {
	FirstMessageNum          *big.Int
	BeforeAcc                [32]byte
	NewMessageCount          *big.Int
	TotalDelayedMessagesRead *big.Int
	AfterAccAndDelayed       [2][32]byte
	SeqBatchIndex            *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDelayedInboxForced is a free log retrieval operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) FilterDelayedInboxForced(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*SequencerInboxDelayedInboxForcedIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "DelayedInboxForced", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxDelayedInboxForcedIterator{contract: _SequencerInbox.contract, event: "DelayedInboxForced", logs: logs, sub: sub}, nil
}

// WatchDelayedInboxForced is a free log subscription operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) WatchDelayedInboxForced(opts *bind.WatchOpts, sink chan<- *SequencerInboxDelayedInboxForced, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "DelayedInboxForced", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxDelayedInboxForced)
				if err := _SequencerInbox.contract.UnpackLog(event, "DelayedInboxForced", log); err != nil {
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

// ParseDelayedInboxForced is a log parse operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) ParseDelayedInboxForced(log types.Log) (*SequencerInboxDelayedInboxForced, error) {
	event := new(SequencerInboxDelayedInboxForced)
	if err := _SequencerInbox.contract.UnpackLog(event, "DelayedInboxForced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDeliveredIterator struct {
	Event *SequencerInboxSequencerBatchDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxSequencerBatchDelivered)
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
		it.Event = new(SequencerInboxSequencerBatchDelivered)
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
func (it *SequencerInboxSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDelivered struct {
	FirstMessageNum          *big.Int
	BeforeAcc                [32]byte
	NewMessageCount          *big.Int
	AfterAcc                 [32]byte
	Transactions             []byte
	Lengths                  []*big.Int
	L1BlockNumber            *big.Int
	Timestamp                *big.Int
	TotalDelayedMessagesRead *big.Int
	DelayedAcc               [32]byte
	SeqBatchIndex            *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDelivered is a free log retrieval operation binding the contract event 0x43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 totalDelayedMessagesRead, bytes32 delayedAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*SequencerInboxSequencerBatchDeliveredIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "SequencerBatchDelivered", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxSequencerBatchDeliveredIterator{contract: _SequencerInbox.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 totalDelayedMessagesRead, bytes32 delayedAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxSequencerBatchDelivered, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "SequencerBatchDelivered", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxSequencerBatchDelivered)
				if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
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

// ParseSequencerBatchDelivered is a log parse operation binding the contract event 0x43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 totalDelayedMessagesRead, bytes32 delayedAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) ParseSequencerBatchDelivered(log types.Log) (*SequencerInboxSequencerBatchDelivered, error) {
	event := new(SequencerInboxSequencerBatchDelivered)
	if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxSequencerBatchDeliveredFromOriginIterator is returned from FilterSequencerBatchDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for SequencerBatchDeliveredFromOrigin events raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDeliveredFromOriginIterator struct {
	Event *SequencerInboxSequencerBatchDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *SequencerInboxSequencerBatchDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxSequencerBatchDeliveredFromOrigin)
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
		it.Event = new(SequencerInboxSequencerBatchDeliveredFromOrigin)
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
func (it *SequencerInboxSequencerBatchDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxSequencerBatchDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxSequencerBatchDeliveredFromOrigin represents a SequencerBatchDeliveredFromOrigin event raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDeliveredFromOrigin struct {
	FirstMessageNum *big.Int
	BeforeAcc       [32]byte
	NewMessageCount *big.Int
	AfterAcc        [32]byte
	DelayedAcc      [32]byte
	SeqBatchIndex   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes32 delayedAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) FilterSequencerBatchDeliveredFromOrigin(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*SequencerInboxSequencerBatchDeliveredFromOriginIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "SequencerBatchDeliveredFromOrigin", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxSequencerBatchDeliveredFromOriginIterator{contract: _SequencerInbox.contract, event: "SequencerBatchDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDeliveredFromOrigin is a free log subscription operation binding the contract event 0x90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes32 delayedAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) WatchSequencerBatchDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *SequencerInboxSequencerBatchDeliveredFromOrigin, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "SequencerBatchDeliveredFromOrigin", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxSequencerBatchDeliveredFromOrigin)
				if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDeliveredFromOrigin", log); err != nil {
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

// ParseSequencerBatchDeliveredFromOrigin is a log parse operation binding the contract event 0x90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes32 delayedAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) ParseSequencerBatchDeliveredFromOrigin(log types.Log) (*SequencerInboxSequencerBatchDeliveredFromOrigin, error) {
	event := new(SequencerInboxSequencerBatchDeliveredFromOrigin)
	if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxyABI is the input ABI used to generate the binding from.
const TransparentUpgradeableProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// TransparentUpgradeableProxyBin is the compiled bytecode used for deploying new contracts.
var TransparentUpgradeableProxyBin = "0x608060405260405162000c7438038062000c74833981810160405260608110156200002957600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200005557600080fd5b9083019060208201858111156200006b57600080fd5b82516401000000008111828201881017156200008657600080fd5b82525081516020918201929091019080838360005b83811015620000b55781810151838201526020016200009b565b50505050905090810190601f168015620000e35780820380516001836020036101000a031916815260200191505b5060405250849150829050620000f98262000137565b8051156200011a57620001188282620001ae60201b620003821760201c565b505b50620001239050565b6200012e82620001dd565b505050620003bf565b6200014d816200020160201b620003ae1760201c565b6200018a5760405162461bcd60e51b815260040180806020018281038252603681526020018062000c186036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b6060620001d6838360405180606001604052806027815260200162000bf16027913962000207565b9392505050565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b3b151590565b6060620002148462000201565b620002515760405162461bcd60e51b815260040180806020018281038252602681526020018062000c4e6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b60208310620002915780518252601f19909201916020918201910162000270565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114620002f3576040519150601f19603f3d011682016040523d82523d6000602084013e620002f8565b606091505b5090925090506200030b82828662000315565b9695505050505050565b6060831562000326575081620001d6565b825115620003375782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200038357818101518382015260200162000369565b50505050905090810190601f168015620003b15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b61082280620003cf6000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100985780635c60da1b146101165780638f28397014610147578063f851a4401461017a5761005d565b3661005d5761005b61018f565b005b61005b61018f565b34801561007157600080fd5b5061005b6004803603602081101561008857600080fd5b50356001600160a01b03166101a9565b61005b600480360360408110156100ae57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156100d857600080fd5b8201836020820111156100ea57600080fd5b803590602001918460018302840111600160201b8311171561010b57600080fd5b5090925090506101e3565b34801561012257600080fd5b5061012b610260565b604080516001600160a01b039092168252519081900360200190f35b34801561015357600080fd5b5061005b6004803603602081101561016a57600080fd5b50356001600160a01b031661029d565b34801561018657600080fd5b5061012b610357565b6101976103b4565b6101a76101a2610414565b610427565b565b6101b161044b565b6001600160a01b0316336001600160a01b031614156101d8576101d38161045e565b6101e0565b6101e061018f565b50565b6101eb61044b565b6001600160a01b0316336001600160a01b031614156102535761020d8361045e565b61024d8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038292505050565b5061025b565b61025b61018f565b505050565b600061026a61044b565b6001600160a01b0316336001600160a01b031614156102925761028b610414565b905061029a565b61029a61018f565b90565b6102a561044b565b6001600160a01b0316336001600160a01b031614156101d8576001600160a01b0381166103035760405162461bcd60e51b815260040180806020018281038252603a8152602001806106ce603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61032c61044b565b604080516001600160a01b03928316815291841660208301528051918290030190a16101d38161049e565b600061036161044b565b6001600160a01b0316336001600160a01b031614156102925761028b61044b565b60606103a78383604051806060016040528060278152602001610728602791396104b0565b9392505050565b3b151590565b6103bc61044b565b6001600160a01b0316336001600160a01b0316141561040c5760405162461bcd60e51b81526004018080602001828103825260428152602001806107ab6042913960600191505060405180910390fd5b6101a76101a7565b6000805160206107088339815191525490565b3660008037600080366000845af43d6000803e808015610446573d6000f35b3d6000fd5b6000805160206106ae8339815191525490565b610467816105b3565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6000805160206106ae83398151915255565b60606104bb846103ae565b6104f65760405162461bcd60e51b81526004018080602001828103825260268152602001806107856026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106105345780518252601f199092019160209182019101610515565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610594576040519150601f19603f3d011682016040523d82523d6000602084013e610599565b606091505b50915091506105a9828286610609565b9695505050505050565b6105bc816103ae565b6105f75760405162461bcd60e51b815260040180806020018281038252603681526020018061074f6036913960400191505060405180910390fd5b60008051602061070883398151915255565b606083156106185750816103a7565b8251156106285782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561067257818101518382015260200161065a565b50505050905090810190601f16801561069f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfeb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f2061646472657373360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a264697066735822122081df10ca698e92e05c31196844f9e9c309ede61eb7e45463098c415a459a179564736f6c634300060c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e7472616374"

// DeployTransparentUpgradeableProxy deploys a new Ethereum contract, binding an instance of TransparentUpgradeableProxy to it.
func DeployTransparentUpgradeableProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, admin_ common.Address, _data []byte) (common.Address, *types.Transaction, *TransparentUpgradeableProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(TransparentUpgradeableProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransparentUpgradeableProxyBin), backend, _logic, admin_, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransparentUpgradeableProxy{TransparentUpgradeableProxyCaller: TransparentUpgradeableProxyCaller{contract: contract}, TransparentUpgradeableProxyTransactor: TransparentUpgradeableProxyTransactor{contract: contract}, TransparentUpgradeableProxyFilterer: TransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

// TransparentUpgradeableProxy is an auto generated Go binding around an Ethereum contract.
type TransparentUpgradeableProxy struct {
	TransparentUpgradeableProxyCaller     // Read-only binding to the contract
	TransparentUpgradeableProxyTransactor // Write-only binding to the contract
	TransparentUpgradeableProxyFilterer   // Log filterer for contract events
}

// TransparentUpgradeableProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransparentUpgradeableProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransparentUpgradeableProxySession struct {
	Contract     *TransparentUpgradeableProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TransparentUpgradeableProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransparentUpgradeableProxyCallerSession struct {
	Contract *TransparentUpgradeableProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// TransparentUpgradeableProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransparentUpgradeableProxyTransactorSession struct {
	Contract     *TransparentUpgradeableProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// TransparentUpgradeableProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransparentUpgradeableProxyRaw struct {
	Contract *TransparentUpgradeableProxy // Generic contract binding to access the raw methods on
}

// TransparentUpgradeableProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyCallerRaw struct {
	Contract *TransparentUpgradeableProxyCaller // Generic read-only contract binding to access the raw methods on
}

// TransparentUpgradeableProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyTransactorRaw struct {
	Contract *TransparentUpgradeableProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransparentUpgradeableProxy creates a new instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxy(address common.Address, backend bind.ContractBackend) (*TransparentUpgradeableProxy, error) {
	contract, err := bindTransparentUpgradeableProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxy{TransparentUpgradeableProxyCaller: TransparentUpgradeableProxyCaller{contract: contract}, TransparentUpgradeableProxyTransactor: TransparentUpgradeableProxyTransactor{contract: contract}, TransparentUpgradeableProxyFilterer: TransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

// NewTransparentUpgradeableProxyCaller creates a new read-only instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyCaller(address common.Address, caller bind.ContractCaller) (*TransparentUpgradeableProxyCaller, error) {
	contract, err := bindTransparentUpgradeableProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyCaller{contract: contract}, nil
}

// NewTransparentUpgradeableProxyTransactor creates a new write-only instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*TransparentUpgradeableProxyTransactor, error) {
	contract, err := bindTransparentUpgradeableProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyTransactor{contract: contract}, nil
}

// NewTransparentUpgradeableProxyFilterer creates a new log filterer instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*TransparentUpgradeableProxyFilterer, error) {
	contract, err := bindTransparentUpgradeableProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyFilterer{contract: contract}, nil
}

// bindTransparentUpgradeableProxy binds a generic wrapper to an already deployed contract.
func bindTransparentUpgradeableProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransparentUpgradeableProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgradeableProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Admin() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Admin(&_TransparentUpgradeableProxy.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Admin() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Admin(&_TransparentUpgradeableProxy.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.ChangeAdmin(&_TransparentUpgradeableProxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.ChangeAdmin(&_TransparentUpgradeableProxy.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Implementation() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Implementation(&_TransparentUpgradeableProxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Implementation() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Implementation(&_TransparentUpgradeableProxy.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeTo(&_TransparentUpgradeableProxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeTo(&_TransparentUpgradeableProxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeToAndCall(&_TransparentUpgradeableProxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeToAndCall(&_TransparentUpgradeableProxy.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Fallback(&_TransparentUpgradeableProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Fallback(&_TransparentUpgradeableProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Receive() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Receive(&_TransparentUpgradeableProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Receive(&_TransparentUpgradeableProxy.TransactOpts)
}

// TransparentUpgradeableProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyAdminChangedIterator struct {
	Event *TransparentUpgradeableProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyAdminChanged)
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
		it.Event = new(TransparentUpgradeableProxyAdminChanged)
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
func (it *TransparentUpgradeableProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyAdminChanged represents a AdminChanged event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TransparentUpgradeableProxyAdminChangedIterator, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyAdminChangedIterator{contract: _TransparentUpgradeableProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyAdminChanged)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseAdminChanged(log types.Log) (*TransparentUpgradeableProxyAdminChanged, error) {
	event := new(TransparentUpgradeableProxyAdminChanged)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyUpgradedIterator struct {
	Event *TransparentUpgradeableProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyUpgraded)
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
		it.Event = new(TransparentUpgradeableProxyUpgraded)
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
func (it *TransparentUpgradeableProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyUpgraded represents a Upgraded event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TransparentUpgradeableProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyUpgradedIterator{contract: _TransparentUpgradeableProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyUpgraded)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseUpgraded(log types.Log) (*TransparentUpgradeableProxyUpgraded, error) {
	event := new(TransparentUpgradeableProxyUpgraded)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableProxyABI is the input ABI used to generate the binding from.
const UpgradeableProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// UpgradeableProxyBin is the compiled bytecode used for deploying new contracts.
var UpgradeableProxyBin = "0x60806040526040516106c43803806106c48339818101604052604081101561002657600080fd5b81516020830180516040519294929383019291908464010000000082111561004d57600080fd5b90830190602082018581111561006257600080fd5b825164010000000081118282018810171561007c57600080fd5b82525081516020918201929091019080838360005b838110156100a9578181015183820152602001610091565b50505050905090810190601f1680156100d65780820380516001836020036101000a031916815260200191505b50604052506100e3915050565b6100ec82610110565b80511561010957610107828261018260201b6100311760201c565b505b505061035c565b610123816101ae60201b61005d1760201c565b61015e5760405162461bcd60e51b81526004018080602001828103825260368152602001806106686036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b60606101a78383604051806060016040528060278152602001610641602791396101b4565b9392505050565b3b151590565b60606101bf846101ae565b6101fa5760405162461bcd60e51b815260040180806020018281038252602681526020018061069e6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106102385780518252601f199092019160209182019101610219565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610298576040519150601f19603f3d011682016040523d82523d6000602084013e61029d565b606091505b5090925090506102ae8282866102b8565b9695505050505050565b606083156102c75750816101a7565b8251156102d75782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610321578181015183820152602001610309565b50505050905090810190601f16801561034e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b6102d68061036b6000396000f3fe60806040523661001357610011610017565b005b6100115b61001f61002f565b61002f61002a610063565b610088565b565b60606100568383604051806060016040528060278152602001610254602791396100ac565b9392505050565b3b151590565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156100a7573d6000f35b3d6000fd5b60606100b78461005d565b6100f25760405162461bcd60e51b815260040180806020018281038252602681526020018061027b6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106101305780518252601f199092019160209182019101610111565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610190576040519150601f19603f3d011682016040523d82523d6000602084013e610195565b606091505b50915091506101a58282866101af565b9695505050505050565b606083156101be575081610056565b8251156101ce5782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610218578181015183820152602001610200565b50505050905090810190601f1680156102455780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e7472616374a26469706673582212202cd40b6d0ce1676861ea6089545943e5feec9ba4c92cdcac0083607926df832b64736f6c634300060c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e7472616374"

// DeployUpgradeableProxy deploys a new Ethereum contract, binding an instance of UpgradeableProxy to it.
func DeployUpgradeableProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _data []byte) (common.Address, *types.Transaction, *UpgradeableProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeableProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradeableProxyBin), backend, _logic, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeableProxy{UpgradeableProxyCaller: UpgradeableProxyCaller{contract: contract}, UpgradeableProxyTransactor: UpgradeableProxyTransactor{contract: contract}, UpgradeableProxyFilterer: UpgradeableProxyFilterer{contract: contract}}, nil
}

// UpgradeableProxy is an auto generated Go binding around an Ethereum contract.
type UpgradeableProxy struct {
	UpgradeableProxyCaller     // Read-only binding to the contract
	UpgradeableProxyTransactor // Write-only binding to the contract
	UpgradeableProxyFilterer   // Log filterer for contract events
}

// UpgradeableProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeableProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeableProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeableProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeableProxySession struct {
	Contract     *UpgradeableProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeableProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeableProxyCallerSession struct {
	Contract *UpgradeableProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UpgradeableProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeableProxyTransactorSession struct {
	Contract     *UpgradeableProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UpgradeableProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeableProxyRaw struct {
	Contract *UpgradeableProxy // Generic contract binding to access the raw methods on
}

// UpgradeableProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeableProxyCallerRaw struct {
	Contract *UpgradeableProxyCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeableProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeableProxyTransactorRaw struct {
	Contract *UpgradeableProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeableProxy creates a new instance of UpgradeableProxy, bound to a specific deployed contract.
func NewUpgradeableProxy(address common.Address, backend bind.ContractBackend) (*UpgradeableProxy, error) {
	contract, err := bindUpgradeableProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxy{UpgradeableProxyCaller: UpgradeableProxyCaller{contract: contract}, UpgradeableProxyTransactor: UpgradeableProxyTransactor{contract: contract}, UpgradeableProxyFilterer: UpgradeableProxyFilterer{contract: contract}}, nil
}

// NewUpgradeableProxyCaller creates a new read-only instance of UpgradeableProxy, bound to a specific deployed contract.
func NewUpgradeableProxyCaller(address common.Address, caller bind.ContractCaller) (*UpgradeableProxyCaller, error) {
	contract, err := bindUpgradeableProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxyCaller{contract: contract}, nil
}

// NewUpgradeableProxyTransactor creates a new write-only instance of UpgradeableProxy, bound to a specific deployed contract.
func NewUpgradeableProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeableProxyTransactor, error) {
	contract, err := bindUpgradeableProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxyTransactor{contract: contract}, nil
}

// NewUpgradeableProxyFilterer creates a new log filterer instance of UpgradeableProxy, bound to a specific deployed contract.
func NewUpgradeableProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeableProxyFilterer, error) {
	contract, err := bindUpgradeableProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxyFilterer{contract: contract}, nil
}

// bindUpgradeableProxy binds a generic wrapper to an already deployed contract.
func bindUpgradeableProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeableProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableProxy *UpgradeableProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableProxy.Contract.UpgradeableProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableProxy *UpgradeableProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.UpgradeableProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableProxy *UpgradeableProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.UpgradeableProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableProxy *UpgradeableProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableProxy *UpgradeableProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableProxy *UpgradeableProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeableProxy *UpgradeableProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UpgradeableProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeableProxy *UpgradeableProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.Fallback(&_UpgradeableProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeableProxy *UpgradeableProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.Fallback(&_UpgradeableProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeableProxy *UpgradeableProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeableProxy *UpgradeableProxySession) Receive() (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.Receive(&_UpgradeableProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeableProxy *UpgradeableProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.Receive(&_UpgradeableProxy.TransactOpts)
}

// UpgradeableProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UpgradeableProxy contract.
type UpgradeableProxyUpgradedIterator struct {
	Event *UpgradeableProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *UpgradeableProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableProxyUpgraded)
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
		it.Event = new(UpgradeableProxyUpgraded)
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
func (it *UpgradeableProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableProxyUpgraded represents a Upgraded event raised by the UpgradeableProxy contract.
type UpgradeableProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableProxy *UpgradeableProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UpgradeableProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxyUpgradedIterator{contract: _UpgradeableProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableProxy *UpgradeableProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UpgradeableProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableProxyUpgraded)
				if err := _UpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableProxy *UpgradeableProxyFilterer) ParseUpgraded(log types.Log) (*UpgradeableProxyUpgraded, error) {
	event := new(UpgradeableProxyUpgraded)
	if err := _UpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
