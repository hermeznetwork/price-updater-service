// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// HezrollupABI is the input ABI used to generate the binding from.
const HezrollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenID\",\"type\":\"uint32\"}],\"name\":\"AddToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"batchNum\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"l1UserTxsLen\",\"type\":\"uint16\"}],\"name\":\"ForgeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"forgeL1L2BatchTimeout\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAddToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"withdrawalDelay\",\"type\":\"uint64\"}],\"name\":\"InitializeHermezEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"queueIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"position\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"l1UserTx\",\"type\":\"bytes\"}],\"name\":\"L1UserTxEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SafeMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"numBucket\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockStamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawals\",\"type\":\"uint256\"}],\"name\":\"UpdateBucketWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"arrayBuckets\",\"type\":\"uint256[]\"}],\"name\":\"UpdateBucketsParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeAddToken\",\"type\":\"uint256\"}],\"name\":\"UpdateFeeAddToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newForgeL1L2BatchTimeout\",\"type\":\"uint8\"}],\"name\":\"UpdateForgeL1L2BatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addressArray\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"valueArray\",\"type\":\"uint64[]\"}],\"name\":\"UpdateTokenExchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newWithdrawalDelay\",\"type\":\"uint64\"}],\"name\":\"UpdateWithdrawalDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint48\",\"name\":\"idx\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"numExitRoot\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"instantWithdraw\",\"type\":\"bool\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"hermezV2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ABSOLUTE_MAX_L1L2BATCHTIMEOUT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ACCOUNT_CREATION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUTHORISE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HERMEZ_NETWORK_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"babyPubKey\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"fromIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint40\",\"name\":\"loadAmountF\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"amountF\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"tokenID\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"toIdx\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"addL1Transaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"buckets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"name\":\"exitNullifierMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"exitRootsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newLastIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"newStRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newExitRoot\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedL1CoordinatorTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"l1L2TxsData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeIdxCoordinator\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"verifierIdx\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"l1Batch\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"proofB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofC\",\"type\":\"uint256[2]\"}],\"name\":\"forgeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeL1L2BatchTimeout\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hermezAuctionContract\",\"outputs\":[{\"internalType\":\"contractIHermezAuctionProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hermezGovernanceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifiers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_verifiersParams\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_withdrawVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermezAuctionContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenHEZ\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_forgeL1L2BatchTimeout\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_feeAddToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_poseidon2Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon3Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon4Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermezGovernanceAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_withdrawalDelay\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_withdrawDelayerContract\",\"type\":\"address\"}],\"name\":\"initializeHermez\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"amount\",\"type\":\"uint192\"}],\"name\":\"instantWithdrawalViewer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"l1L2TxsDataHashMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForgedBatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastIdx\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastL1L2Batch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"mapL1TxQueue\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nBuckets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextL1FillingQueue\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextL1ToForgeQueue\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ceilUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateWithdrawals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxWithdrawals\",\"type\":\"uint256\"}],\"name\":\"packBucket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ret\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupVerifiers\",\"outputs\":[{\"internalType\":\"contractVerifierRollupInterface\",\"name\":\"verifierInterface\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nLevels\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupVerifiersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"stateRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenExchange\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenHEZ\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bucket\",\"type\":\"uint256\"}],\"name\":\"unpackBucket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ceilUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateWithdrawals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxWithdrawals\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"newBuckets\",\"type\":\"uint256[]\"}],\"name\":\"updateBucketsParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeAddToken\",\"type\":\"uint256\"}],\"name\":\"updateFeeAddToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newForgeL1L2BatchTimeout\",\"type\":\"uint8\"}],\"name\":\"updateForgeL1L2BatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addressArray\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"valueArray\",\"type\":\"uint64[]\"}],\"name\":\"updateTokenExchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newWithdrawalDelay\",\"type\":\"uint64\"}],\"name\":\"updateWithdrawalDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"proofA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"proofB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint32\",\"name\":\"tokenID\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"amount\",\"type\":\"uint192\"},{\"internalType\":\"uint32\",\"name\":\"numExitRoot\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"idx\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"instantWithdraw\",\"type\":\"bool\"}],\"name\":\"withdrawCircuit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDelayerContract\",\"outputs\":[{\"internalType\":\"contractIWithdrawalDelayer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tokenID\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"amount\",\"type\":\"uint192\"},{\"internalType\":\"uint256\",\"name\":\"babyPubKey\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numExitRoot\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint48\",\"name\":\"idx\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"instantWithdraw\",\"type\":\"bool\"}],\"name\":\"withdrawMerkleProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawVerifier\",\"outputs\":[{\"internalType\":\"contractVerifierWithdrawInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Hezrollup is an auto generated Go binding around an Ethereum contract.
type Hezrollup struct {
	HezrollupCaller     // Read-only binding to the contract
	HezrollupTransactor // Write-only binding to the contract
	HezrollupFilterer   // Log filterer for contract events
}

// HezrollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type HezrollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HezrollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HezrollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HezrollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HezrollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HezrollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HezrollupSession struct {
	Contract     *Hezrollup        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HezrollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HezrollupCallerSession struct {
	Contract *HezrollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HezrollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HezrollupTransactorSession struct {
	Contract     *HezrollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HezrollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type HezrollupRaw struct {
	Contract *Hezrollup // Generic contract binding to access the raw methods on
}

// HezrollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HezrollupCallerRaw struct {
	Contract *HezrollupCaller // Generic read-only contract binding to access the raw methods on
}

// HezrollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HezrollupTransactorRaw struct {
	Contract *HezrollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHezrollup creates a new instance of Hezrollup, bound to a specific deployed contract.
func NewHezrollup(address common.Address, backend bind.ContractBackend) (*Hezrollup, error) {
	contract, err := bindHezrollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hezrollup{HezrollupCaller: HezrollupCaller{contract: contract}, HezrollupTransactor: HezrollupTransactor{contract: contract}, HezrollupFilterer: HezrollupFilterer{contract: contract}}, nil
}

// NewHezrollupCaller creates a new read-only instance of Hezrollup, bound to a specific deployed contract.
func NewHezrollupCaller(address common.Address, caller bind.ContractCaller) (*HezrollupCaller, error) {
	contract, err := bindHezrollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HezrollupCaller{contract: contract}, nil
}

// NewHezrollupTransactor creates a new write-only instance of Hezrollup, bound to a specific deployed contract.
func NewHezrollupTransactor(address common.Address, transactor bind.ContractTransactor) (*HezrollupTransactor, error) {
	contract, err := bindHezrollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HezrollupTransactor{contract: contract}, nil
}

// NewHezrollupFilterer creates a new log filterer instance of Hezrollup, bound to a specific deployed contract.
func NewHezrollupFilterer(address common.Address, filterer bind.ContractFilterer) (*HezrollupFilterer, error) {
	contract, err := bindHezrollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HezrollupFilterer{contract: contract}, nil
}

// bindHezrollup binds a generic wrapper to an already deployed contract.
func bindHezrollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HezrollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hezrollup *HezrollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hezrollup.Contract.HezrollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hezrollup *HezrollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hezrollup.Contract.HezrollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hezrollup *HezrollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hezrollup.Contract.HezrollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hezrollup *HezrollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hezrollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hezrollup *HezrollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hezrollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hezrollup *HezrollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hezrollup.Contract.contract.Transact(opts, method, params...)
}

// ABSOLUTEMAXL1L2BATCHTIMEOUT is a free data retrieval call binding the contract method 0x95a09f2a.
//
// Solidity: function ABSOLUTE_MAX_L1L2BATCHTIMEOUT() view returns(uint8)
func (_Hezrollup *HezrollupCaller) ABSOLUTEMAXL1L2BATCHTIMEOUT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "ABSOLUTE_MAX_L1L2BATCHTIMEOUT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ABSOLUTEMAXL1L2BATCHTIMEOUT is a free data retrieval call binding the contract method 0x95a09f2a.
//
// Solidity: function ABSOLUTE_MAX_L1L2BATCHTIMEOUT() view returns(uint8)
func (_Hezrollup *HezrollupSession) ABSOLUTEMAXL1L2BATCHTIMEOUT() (uint8, error) {
	return _Hezrollup.Contract.ABSOLUTEMAXL1L2BATCHTIMEOUT(&_Hezrollup.CallOpts)
}

// ABSOLUTEMAXL1L2BATCHTIMEOUT is a free data retrieval call binding the contract method 0x95a09f2a.
//
// Solidity: function ABSOLUTE_MAX_L1L2BATCHTIMEOUT() view returns(uint8)
func (_Hezrollup *HezrollupCallerSession) ABSOLUTEMAXL1L2BATCHTIMEOUT() (uint8, error) {
	return _Hezrollup.Contract.ABSOLUTEMAXL1L2BATCHTIMEOUT(&_Hezrollup.CallOpts)
}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCaller) ACCOUNTCREATIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "ACCOUNT_CREATION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupSession) ACCOUNTCREATIONHASH() ([32]byte, error) {
	return _Hezrollup.Contract.ACCOUNTCREATIONHASH(&_Hezrollup.CallOpts)
}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCallerSession) ACCOUNTCREATIONHASH() ([32]byte, error) {
	return _Hezrollup.Contract.ACCOUNTCREATIONHASH(&_Hezrollup.CallOpts)
}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Hezrollup *HezrollupCaller) AUTHORISETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "AUTHORISE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Hezrollup *HezrollupSession) AUTHORISETYPEHASH() ([32]byte, error) {
	return _Hezrollup.Contract.AUTHORISETYPEHASH(&_Hezrollup.CallOpts)
}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Hezrollup *HezrollupCallerSession) AUTHORISETYPEHASH() ([32]byte, error) {
	return _Hezrollup.Contract.AUTHORISETYPEHASH(&_Hezrollup.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Hezrollup *HezrollupCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Hezrollup *HezrollupSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Hezrollup.Contract.DOMAINSEPARATOR(&_Hezrollup.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Hezrollup *HezrollupCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Hezrollup.Contract.DOMAINSEPARATOR(&_Hezrollup.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "EIP712DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Hezrollup.Contract.EIP712DOMAINHASH(&_Hezrollup.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Hezrollup.Contract.EIP712DOMAINHASH(&_Hezrollup.CallOpts)
}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCaller) HERMEZNETWORKHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "HERMEZ_NETWORK_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupSession) HERMEZNETWORKHASH() ([32]byte, error) {
	return _Hezrollup.Contract.HERMEZNETWORKHASH(&_Hezrollup.CallOpts)
}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCallerSession) HERMEZNETWORKHASH() ([32]byte, error) {
	return _Hezrollup.Contract.HERMEZNETWORKHASH(&_Hezrollup.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCaller) NAMEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "NAME_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupSession) NAMEHASH() ([32]byte, error) {
	return _Hezrollup.Contract.NAMEHASH(&_Hezrollup.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCallerSession) NAMEHASH() ([32]byte, error) {
	return _Hezrollup.Contract.NAMEHASH(&_Hezrollup.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCaller) VERSIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "VERSION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupSession) VERSIONHASH() ([32]byte, error) {
	return _Hezrollup.Contract.VERSIONHASH(&_Hezrollup.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Hezrollup *HezrollupCallerSession) VERSIONHASH() ([32]byte, error) {
	return _Hezrollup.Contract.VERSIONHASH(&_Hezrollup.CallOpts)
}

// Buckets is a free data retrieval call binding the contract method 0x061d0964.
//
// Solidity: function buckets(int256 ) view returns(uint256)
func (_Hezrollup *HezrollupCaller) Buckets(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "buckets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buckets is a free data retrieval call binding the contract method 0x061d0964.
//
// Solidity: function buckets(int256 ) view returns(uint256)
func (_Hezrollup *HezrollupSession) Buckets(arg0 *big.Int) (*big.Int, error) {
	return _Hezrollup.Contract.Buckets(&_Hezrollup.CallOpts, arg0)
}

// Buckets is a free data retrieval call binding the contract method 0x061d0964.
//
// Solidity: function buckets(int256 ) view returns(uint256)
func (_Hezrollup *HezrollupCallerSession) Buckets(arg0 *big.Int) (*big.Int, error) {
	return _Hezrollup.Contract.Buckets(&_Hezrollup.CallOpts, arg0)
}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Hezrollup *HezrollupCaller) ExitNullifierMap(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "exitNullifierMap", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Hezrollup *HezrollupSession) ExitNullifierMap(arg0 uint32, arg1 *big.Int) (bool, error) {
	return _Hezrollup.Contract.ExitNullifierMap(&_Hezrollup.CallOpts, arg0, arg1)
}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Hezrollup *HezrollupCallerSession) ExitNullifierMap(arg0 uint32, arg1 *big.Int) (bool, error) {
	return _Hezrollup.Contract.ExitNullifierMap(&_Hezrollup.CallOpts, arg0, arg1)
}

// ExitRootsMap is a free data retrieval call binding the contract method 0x3ee641ea.
//
// Solidity: function exitRootsMap(uint32 ) view returns(uint256)
func (_Hezrollup *HezrollupCaller) ExitRootsMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "exitRootsMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitRootsMap is a free data retrieval call binding the contract method 0x3ee641ea.
//
// Solidity: function exitRootsMap(uint32 ) view returns(uint256)
func (_Hezrollup *HezrollupSession) ExitRootsMap(arg0 uint32) (*big.Int, error) {
	return _Hezrollup.Contract.ExitRootsMap(&_Hezrollup.CallOpts, arg0)
}

// ExitRootsMap is a free data retrieval call binding the contract method 0x3ee641ea.
//
// Solidity: function exitRootsMap(uint32 ) view returns(uint256)
func (_Hezrollup *HezrollupCallerSession) ExitRootsMap(arg0 uint32) (*big.Int, error) {
	return _Hezrollup.Contract.ExitRootsMap(&_Hezrollup.CallOpts, arg0)
}

// FeeAddToken is a free data retrieval call binding the contract method 0xbded9bb8.
//
// Solidity: function feeAddToken() view returns(uint256)
func (_Hezrollup *HezrollupCaller) FeeAddToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "feeAddToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAddToken is a free data retrieval call binding the contract method 0xbded9bb8.
//
// Solidity: function feeAddToken() view returns(uint256)
func (_Hezrollup *HezrollupSession) FeeAddToken() (*big.Int, error) {
	return _Hezrollup.Contract.FeeAddToken(&_Hezrollup.CallOpts)
}

// FeeAddToken is a free data retrieval call binding the contract method 0xbded9bb8.
//
// Solidity: function feeAddToken() view returns(uint256)
func (_Hezrollup *HezrollupCallerSession) FeeAddToken() (*big.Int, error) {
	return _Hezrollup.Contract.FeeAddToken(&_Hezrollup.CallOpts)
}

// ForgeL1L2BatchTimeout is a free data retrieval call binding the contract method 0xa3275838.
//
// Solidity: function forgeL1L2BatchTimeout() view returns(uint8)
func (_Hezrollup *HezrollupCaller) ForgeL1L2BatchTimeout(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "forgeL1L2BatchTimeout")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ForgeL1L2BatchTimeout is a free data retrieval call binding the contract method 0xa3275838.
//
// Solidity: function forgeL1L2BatchTimeout() view returns(uint8)
func (_Hezrollup *HezrollupSession) ForgeL1L2BatchTimeout() (uint8, error) {
	return _Hezrollup.Contract.ForgeL1L2BatchTimeout(&_Hezrollup.CallOpts)
}

// ForgeL1L2BatchTimeout is a free data retrieval call binding the contract method 0xa3275838.
//
// Solidity: function forgeL1L2BatchTimeout() view returns(uint8)
func (_Hezrollup *HezrollupCallerSession) ForgeL1L2BatchTimeout() (uint8, error) {
	return _Hezrollup.Contract.ForgeL1L2BatchTimeout(&_Hezrollup.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256 chainId)
func (_Hezrollup *HezrollupCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256 chainId)
func (_Hezrollup *HezrollupSession) GetChainId() (*big.Int, error) {
	return _Hezrollup.Contract.GetChainId(&_Hezrollup.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256 chainId)
func (_Hezrollup *HezrollupCallerSession) GetChainId() (*big.Int, error) {
	return _Hezrollup.Contract.GetChainId(&_Hezrollup.CallOpts)
}

// HermezAuctionContract is a free data retrieval call binding the contract method 0x2bd83626.
//
// Solidity: function hermezAuctionContract() view returns(address)
func (_Hezrollup *HezrollupCaller) HermezAuctionContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "hermezAuctionContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HermezAuctionContract is a free data retrieval call binding the contract method 0x2bd83626.
//
// Solidity: function hermezAuctionContract() view returns(address)
func (_Hezrollup *HezrollupSession) HermezAuctionContract() (common.Address, error) {
	return _Hezrollup.Contract.HermezAuctionContract(&_Hezrollup.CallOpts)
}

// HermezAuctionContract is a free data retrieval call binding the contract method 0x2bd83626.
//
// Solidity: function hermezAuctionContract() view returns(address)
func (_Hezrollup *HezrollupCallerSession) HermezAuctionContract() (common.Address, error) {
	return _Hezrollup.Contract.HermezAuctionContract(&_Hezrollup.CallOpts)
}

// HermezGovernanceAddress is a free data retrieval call binding the contract method 0x013f7852.
//
// Solidity: function hermezGovernanceAddress() view returns(address)
func (_Hezrollup *HezrollupCaller) HermezGovernanceAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "hermezGovernanceAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HermezGovernanceAddress is a free data retrieval call binding the contract method 0x013f7852.
//
// Solidity: function hermezGovernanceAddress() view returns(address)
func (_Hezrollup *HezrollupSession) HermezGovernanceAddress() (common.Address, error) {
	return _Hezrollup.Contract.HermezGovernanceAddress(&_Hezrollup.CallOpts)
}

// HermezGovernanceAddress is a free data retrieval call binding the contract method 0x013f7852.
//
// Solidity: function hermezGovernanceAddress() view returns(address)
func (_Hezrollup *HezrollupCallerSession) HermezGovernanceAddress() (common.Address, error) {
	return _Hezrollup.Contract.HermezGovernanceAddress(&_Hezrollup.CallOpts)
}

// InstantWithdrawalViewer is a free data retrieval call binding the contract method 0x375110aa.
//
// Solidity: function instantWithdrawalViewer(address tokenAddress, uint192 amount) view returns(bool)
func (_Hezrollup *HezrollupCaller) InstantWithdrawalViewer(opts *bind.CallOpts, tokenAddress common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "instantWithdrawalViewer", tokenAddress, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InstantWithdrawalViewer is a free data retrieval call binding the contract method 0x375110aa.
//
// Solidity: function instantWithdrawalViewer(address tokenAddress, uint192 amount) view returns(bool)
func (_Hezrollup *HezrollupSession) InstantWithdrawalViewer(tokenAddress common.Address, amount *big.Int) (bool, error) {
	return _Hezrollup.Contract.InstantWithdrawalViewer(&_Hezrollup.CallOpts, tokenAddress, amount)
}

// InstantWithdrawalViewer is a free data retrieval call binding the contract method 0x375110aa.
//
// Solidity: function instantWithdrawalViewer(address tokenAddress, uint192 amount) view returns(bool)
func (_Hezrollup *HezrollupCallerSession) InstantWithdrawalViewer(tokenAddress common.Address, amount *big.Int) (bool, error) {
	return _Hezrollup.Contract.InstantWithdrawalViewer(&_Hezrollup.CallOpts, tokenAddress, amount)
}

// L1L2TxsDataHashMap is a free data retrieval call binding the contract method 0xce5ec65a.
//
// Solidity: function l1L2TxsDataHashMap(uint32 ) view returns(bytes32)
func (_Hezrollup *HezrollupCaller) L1L2TxsDataHashMap(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "l1L2TxsDataHashMap", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1L2TxsDataHashMap is a free data retrieval call binding the contract method 0xce5ec65a.
//
// Solidity: function l1L2TxsDataHashMap(uint32 ) view returns(bytes32)
func (_Hezrollup *HezrollupSession) L1L2TxsDataHashMap(arg0 uint32) ([32]byte, error) {
	return _Hezrollup.Contract.L1L2TxsDataHashMap(&_Hezrollup.CallOpts, arg0)
}

// L1L2TxsDataHashMap is a free data retrieval call binding the contract method 0xce5ec65a.
//
// Solidity: function l1L2TxsDataHashMap(uint32 ) view returns(bytes32)
func (_Hezrollup *HezrollupCallerSession) L1L2TxsDataHashMap(arg0 uint32) ([32]byte, error) {
	return _Hezrollup.Contract.L1L2TxsDataHashMap(&_Hezrollup.CallOpts, arg0)
}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Hezrollup *HezrollupCaller) LastForgedBatch(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "lastForgedBatch")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Hezrollup *HezrollupSession) LastForgedBatch() (uint32, error) {
	return _Hezrollup.Contract.LastForgedBatch(&_Hezrollup.CallOpts)
}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Hezrollup *HezrollupCallerSession) LastForgedBatch() (uint32, error) {
	return _Hezrollup.Contract.LastForgedBatch(&_Hezrollup.CallOpts)
}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Hezrollup *HezrollupCaller) LastIdx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "lastIdx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Hezrollup *HezrollupSession) LastIdx() (*big.Int, error) {
	return _Hezrollup.Contract.LastIdx(&_Hezrollup.CallOpts)
}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Hezrollup *HezrollupCallerSession) LastIdx() (*big.Int, error) {
	return _Hezrollup.Contract.LastIdx(&_Hezrollup.CallOpts)
}

// LastL1L2Batch is a free data retrieval call binding the contract method 0x84ef9ed4.
//
// Solidity: function lastL1L2Batch() view returns(uint64)
func (_Hezrollup *HezrollupCaller) LastL1L2Batch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "lastL1L2Batch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastL1L2Batch is a free data retrieval call binding the contract method 0x84ef9ed4.
//
// Solidity: function lastL1L2Batch() view returns(uint64)
func (_Hezrollup *HezrollupSession) LastL1L2Batch() (uint64, error) {
	return _Hezrollup.Contract.LastL1L2Batch(&_Hezrollup.CallOpts)
}

// LastL1L2Batch is a free data retrieval call binding the contract method 0x84ef9ed4.
//
// Solidity: function lastL1L2Batch() view returns(uint64)
func (_Hezrollup *HezrollupCallerSession) LastL1L2Batch() (uint64, error) {
	return _Hezrollup.Contract.LastL1L2Batch(&_Hezrollup.CallOpts)
}

// MapL1TxQueue is a free data retrieval call binding the contract method 0xdc3e718e.
//
// Solidity: function mapL1TxQueue(uint32 ) view returns(bytes)
func (_Hezrollup *HezrollupCaller) MapL1TxQueue(opts *bind.CallOpts, arg0 uint32) ([]byte, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "mapL1TxQueue", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MapL1TxQueue is a free data retrieval call binding the contract method 0xdc3e718e.
//
// Solidity: function mapL1TxQueue(uint32 ) view returns(bytes)
func (_Hezrollup *HezrollupSession) MapL1TxQueue(arg0 uint32) ([]byte, error) {
	return _Hezrollup.Contract.MapL1TxQueue(&_Hezrollup.CallOpts, arg0)
}

// MapL1TxQueue is a free data retrieval call binding the contract method 0xdc3e718e.
//
// Solidity: function mapL1TxQueue(uint32 ) view returns(bytes)
func (_Hezrollup *HezrollupCallerSession) MapL1TxQueue(arg0 uint32) ([]byte, error) {
	return _Hezrollup.Contract.MapL1TxQueue(&_Hezrollup.CallOpts, arg0)
}

// NBuckets is a free data retrieval call binding the contract method 0x07feef6e.
//
// Solidity: function nBuckets() view returns(uint256)
func (_Hezrollup *HezrollupCaller) NBuckets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "nBuckets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NBuckets is a free data retrieval call binding the contract method 0x07feef6e.
//
// Solidity: function nBuckets() view returns(uint256)
func (_Hezrollup *HezrollupSession) NBuckets() (*big.Int, error) {
	return _Hezrollup.Contract.NBuckets(&_Hezrollup.CallOpts)
}

// NBuckets is a free data retrieval call binding the contract method 0x07feef6e.
//
// Solidity: function nBuckets() view returns(uint256)
func (_Hezrollup *HezrollupCallerSession) NBuckets() (*big.Int, error) {
	return _Hezrollup.Contract.NBuckets(&_Hezrollup.CallOpts)
}

// NextL1FillingQueue is a free data retrieval call binding the contract method 0x0ee8e52b.
//
// Solidity: function nextL1FillingQueue() view returns(uint32)
func (_Hezrollup *HezrollupCaller) NextL1FillingQueue(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "nextL1FillingQueue")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextL1FillingQueue is a free data retrieval call binding the contract method 0x0ee8e52b.
//
// Solidity: function nextL1FillingQueue() view returns(uint32)
func (_Hezrollup *HezrollupSession) NextL1FillingQueue() (uint32, error) {
	return _Hezrollup.Contract.NextL1FillingQueue(&_Hezrollup.CallOpts)
}

// NextL1FillingQueue is a free data retrieval call binding the contract method 0x0ee8e52b.
//
// Solidity: function nextL1FillingQueue() view returns(uint32)
func (_Hezrollup *HezrollupCallerSession) NextL1FillingQueue() (uint32, error) {
	return _Hezrollup.Contract.NextL1FillingQueue(&_Hezrollup.CallOpts)
}

// NextL1ToForgeQueue is a free data retrieval call binding the contract method 0xd0f32e67.
//
// Solidity: function nextL1ToForgeQueue() view returns(uint32)
func (_Hezrollup *HezrollupCaller) NextL1ToForgeQueue(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "nextL1ToForgeQueue")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextL1ToForgeQueue is a free data retrieval call binding the contract method 0xd0f32e67.
//
// Solidity: function nextL1ToForgeQueue() view returns(uint32)
func (_Hezrollup *HezrollupSession) NextL1ToForgeQueue() (uint32, error) {
	return _Hezrollup.Contract.NextL1ToForgeQueue(&_Hezrollup.CallOpts)
}

// NextL1ToForgeQueue is a free data retrieval call binding the contract method 0xd0f32e67.
//
// Solidity: function nextL1ToForgeQueue() view returns(uint32)
func (_Hezrollup *HezrollupCallerSession) NextL1ToForgeQueue() (uint32, error) {
	return _Hezrollup.Contract.NextL1ToForgeQueue(&_Hezrollup.CallOpts)
}

// PackBucket is a free data retrieval call binding the contract method 0xccd226a7.
//
// Solidity: function packBucket(uint256 ceilUSD, uint256 blockStamp, uint256 withdrawals, uint256 rateBlocks, uint256 rateWithdrawals, uint256 maxWithdrawals) pure returns(uint256 ret)
func (_Hezrollup *HezrollupCaller) PackBucket(opts *bind.CallOpts, ceilUSD *big.Int, blockStamp *big.Int, withdrawals *big.Int, rateBlocks *big.Int, rateWithdrawals *big.Int, maxWithdrawals *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "packBucket", ceilUSD, blockStamp, withdrawals, rateBlocks, rateWithdrawals, maxWithdrawals)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackBucket is a free data retrieval call binding the contract method 0xccd226a7.
//
// Solidity: function packBucket(uint256 ceilUSD, uint256 blockStamp, uint256 withdrawals, uint256 rateBlocks, uint256 rateWithdrawals, uint256 maxWithdrawals) pure returns(uint256 ret)
func (_Hezrollup *HezrollupSession) PackBucket(ceilUSD *big.Int, blockStamp *big.Int, withdrawals *big.Int, rateBlocks *big.Int, rateWithdrawals *big.Int, maxWithdrawals *big.Int) (*big.Int, error) {
	return _Hezrollup.Contract.PackBucket(&_Hezrollup.CallOpts, ceilUSD, blockStamp, withdrawals, rateBlocks, rateWithdrawals, maxWithdrawals)
}

// PackBucket is a free data retrieval call binding the contract method 0xccd226a7.
//
// Solidity: function packBucket(uint256 ceilUSD, uint256 blockStamp, uint256 withdrawals, uint256 rateBlocks, uint256 rateWithdrawals, uint256 maxWithdrawals) pure returns(uint256 ret)
func (_Hezrollup *HezrollupCallerSession) PackBucket(ceilUSD *big.Int, blockStamp *big.Int, withdrawals *big.Int, rateBlocks *big.Int, rateWithdrawals *big.Int, maxWithdrawals *big.Int) (*big.Int, error) {
	return _Hezrollup.Contract.PackBucket(&_Hezrollup.CallOpts, ceilUSD, blockStamp, withdrawals, rateBlocks, rateWithdrawals, maxWithdrawals)
}

// RegisterTokensCount is a free data retrieval call binding the contract method 0x9f34e9a3.
//
// Solidity: function registerTokensCount() view returns(uint256)
func (_Hezrollup *HezrollupCaller) RegisterTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "registerTokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegisterTokensCount is a free data retrieval call binding the contract method 0x9f34e9a3.
//
// Solidity: function registerTokensCount() view returns(uint256)
func (_Hezrollup *HezrollupSession) RegisterTokensCount() (*big.Int, error) {
	return _Hezrollup.Contract.RegisterTokensCount(&_Hezrollup.CallOpts)
}

// RegisterTokensCount is a free data retrieval call binding the contract method 0x9f34e9a3.
//
// Solidity: function registerTokensCount() view returns(uint256)
func (_Hezrollup *HezrollupCallerSession) RegisterTokensCount() (*big.Int, error) {
	return _Hezrollup.Contract.RegisterTokensCount(&_Hezrollup.CallOpts)
}

// RollupVerifiers is a free data retrieval call binding the contract method 0x38330200.
//
// Solidity: function rollupVerifiers(uint256 ) view returns(address verifierInterface, uint256 maxTx, uint256 nLevels)
func (_Hezrollup *HezrollupCaller) RollupVerifiers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevels           *big.Int
}, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "rollupVerifiers", arg0)

	outstruct := new(struct {
		VerifierInterface common.Address
		MaxTx             *big.Int
		NLevels           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VerifierInterface = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MaxTx = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NLevels = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RollupVerifiers is a free data retrieval call binding the contract method 0x38330200.
//
// Solidity: function rollupVerifiers(uint256 ) view returns(address verifierInterface, uint256 maxTx, uint256 nLevels)
func (_Hezrollup *HezrollupSession) RollupVerifiers(arg0 *big.Int) (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevels           *big.Int
}, error) {
	return _Hezrollup.Contract.RollupVerifiers(&_Hezrollup.CallOpts, arg0)
}

// RollupVerifiers is a free data retrieval call binding the contract method 0x38330200.
//
// Solidity: function rollupVerifiers(uint256 ) view returns(address verifierInterface, uint256 maxTx, uint256 nLevels)
func (_Hezrollup *HezrollupCallerSession) RollupVerifiers(arg0 *big.Int) (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevels           *big.Int
}, error) {
	return _Hezrollup.Contract.RollupVerifiers(&_Hezrollup.CallOpts, arg0)
}

// RollupVerifiersLength is a free data retrieval call binding the contract method 0x7ba3a5e0.
//
// Solidity: function rollupVerifiersLength() view returns(uint256)
func (_Hezrollup *HezrollupCaller) RollupVerifiersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "rollupVerifiersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupVerifiersLength is a free data retrieval call binding the contract method 0x7ba3a5e0.
//
// Solidity: function rollupVerifiersLength() view returns(uint256)
func (_Hezrollup *HezrollupSession) RollupVerifiersLength() (*big.Int, error) {
	return _Hezrollup.Contract.RollupVerifiersLength(&_Hezrollup.CallOpts)
}

// RollupVerifiersLength is a free data retrieval call binding the contract method 0x7ba3a5e0.
//
// Solidity: function rollupVerifiersLength() view returns(uint256)
func (_Hezrollup *HezrollupCallerSession) RollupVerifiersLength() (*big.Int, error) {
	return _Hezrollup.Contract.RollupVerifiersLength(&_Hezrollup.CallOpts)
}

// StateRootMap is a free data retrieval call binding the contract method 0x9e00d7ea.
//
// Solidity: function stateRootMap(uint32 ) view returns(uint256)
func (_Hezrollup *HezrollupCaller) StateRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "stateRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateRootMap is a free data retrieval call binding the contract method 0x9e00d7ea.
//
// Solidity: function stateRootMap(uint32 ) view returns(uint256)
func (_Hezrollup *HezrollupSession) StateRootMap(arg0 uint32) (*big.Int, error) {
	return _Hezrollup.Contract.StateRootMap(&_Hezrollup.CallOpts, arg0)
}

// StateRootMap is a free data retrieval call binding the contract method 0x9e00d7ea.
//
// Solidity: function stateRootMap(uint32 ) view returns(uint256)
func (_Hezrollup *HezrollupCallerSession) StateRootMap(arg0 uint32) (*big.Int, error) {
	return _Hezrollup.Contract.StateRootMap(&_Hezrollup.CallOpts, arg0)
}

// TokenExchange is a free data retrieval call binding the contract method 0x0dd94b96.
//
// Solidity: function tokenExchange(address ) view returns(uint64)
func (_Hezrollup *HezrollupCaller) TokenExchange(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "tokenExchange", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TokenExchange is a free data retrieval call binding the contract method 0x0dd94b96.
//
// Solidity: function tokenExchange(address ) view returns(uint64)
func (_Hezrollup *HezrollupSession) TokenExchange(arg0 common.Address) (uint64, error) {
	return _Hezrollup.Contract.TokenExchange(&_Hezrollup.CallOpts, arg0)
}

// TokenExchange is a free data retrieval call binding the contract method 0x0dd94b96.
//
// Solidity: function tokenExchange(address ) view returns(uint64)
func (_Hezrollup *HezrollupCallerSession) TokenExchange(arg0 common.Address) (uint64, error) {
	return _Hezrollup.Contract.TokenExchange(&_Hezrollup.CallOpts, arg0)
}

// TokenHEZ is a free data retrieval call binding the contract method 0x79a135e3.
//
// Solidity: function tokenHEZ() view returns(address)
func (_Hezrollup *HezrollupCaller) TokenHEZ(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "tokenHEZ")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenHEZ is a free data retrieval call binding the contract method 0x79a135e3.
//
// Solidity: function tokenHEZ() view returns(address)
func (_Hezrollup *HezrollupSession) TokenHEZ() (common.Address, error) {
	return _Hezrollup.Contract.TokenHEZ(&_Hezrollup.CallOpts)
}

// TokenHEZ is a free data retrieval call binding the contract method 0x79a135e3.
//
// Solidity: function tokenHEZ() view returns(address)
func (_Hezrollup *HezrollupCallerSession) TokenHEZ() (common.Address, error) {
	return _Hezrollup.Contract.TokenHEZ(&_Hezrollup.CallOpts)
}

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(address)
func (_Hezrollup *HezrollupCaller) TokenList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "tokenList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(address)
func (_Hezrollup *HezrollupSession) TokenList(arg0 *big.Int) (common.Address, error) {
	return _Hezrollup.Contract.TokenList(&_Hezrollup.CallOpts, arg0)
}

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(address)
func (_Hezrollup *HezrollupCallerSession) TokenList(arg0 *big.Int) (common.Address, error) {
	return _Hezrollup.Contract.TokenList(&_Hezrollup.CallOpts, arg0)
}

// TokenMap is a free data retrieval call binding the contract method 0x004aca6e.
//
// Solidity: function tokenMap(address ) view returns(uint256)
func (_Hezrollup *HezrollupCaller) TokenMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "tokenMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMap is a free data retrieval call binding the contract method 0x004aca6e.
//
// Solidity: function tokenMap(address ) view returns(uint256)
func (_Hezrollup *HezrollupSession) TokenMap(arg0 common.Address) (*big.Int, error) {
	return _Hezrollup.Contract.TokenMap(&_Hezrollup.CallOpts, arg0)
}

// TokenMap is a free data retrieval call binding the contract method 0x004aca6e.
//
// Solidity: function tokenMap(address ) view returns(uint256)
func (_Hezrollup *HezrollupCallerSession) TokenMap(arg0 common.Address) (*big.Int, error) {
	return _Hezrollup.Contract.TokenMap(&_Hezrollup.CallOpts, arg0)
}

// UnpackBucket is a free data retrieval call binding the contract method 0x3f267155.
//
// Solidity: function unpackBucket(uint256 bucket) pure returns(uint256 ceilUSD, uint256 blockStamp, uint256 withdrawals, uint256 rateBlocks, uint256 rateWithdrawals, uint256 maxWithdrawals)
func (_Hezrollup *HezrollupCaller) UnpackBucket(opts *bind.CallOpts, bucket *big.Int) (struct {
	CeilUSD         *big.Int
	BlockStamp      *big.Int
	Withdrawals     *big.Int
	RateBlocks      *big.Int
	RateWithdrawals *big.Int
	MaxWithdrawals  *big.Int
}, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "unpackBucket", bucket)

	outstruct := new(struct {
		CeilUSD         *big.Int
		BlockStamp      *big.Int
		Withdrawals     *big.Int
		RateBlocks      *big.Int
		RateWithdrawals *big.Int
		MaxWithdrawals  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CeilUSD = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockStamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Withdrawals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RateBlocks = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RateWithdrawals = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxWithdrawals = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnpackBucket is a free data retrieval call binding the contract method 0x3f267155.
//
// Solidity: function unpackBucket(uint256 bucket) pure returns(uint256 ceilUSD, uint256 blockStamp, uint256 withdrawals, uint256 rateBlocks, uint256 rateWithdrawals, uint256 maxWithdrawals)
func (_Hezrollup *HezrollupSession) UnpackBucket(bucket *big.Int) (struct {
	CeilUSD         *big.Int
	BlockStamp      *big.Int
	Withdrawals     *big.Int
	RateBlocks      *big.Int
	RateWithdrawals *big.Int
	MaxWithdrawals  *big.Int
}, error) {
	return _Hezrollup.Contract.UnpackBucket(&_Hezrollup.CallOpts, bucket)
}

// UnpackBucket is a free data retrieval call binding the contract method 0x3f267155.
//
// Solidity: function unpackBucket(uint256 bucket) pure returns(uint256 ceilUSD, uint256 blockStamp, uint256 withdrawals, uint256 rateBlocks, uint256 rateWithdrawals, uint256 maxWithdrawals)
func (_Hezrollup *HezrollupCallerSession) UnpackBucket(bucket *big.Int) (struct {
	CeilUSD         *big.Int
	BlockStamp      *big.Int
	Withdrawals     *big.Int
	RateBlocks      *big.Int
	RateWithdrawals *big.Int
	MaxWithdrawals  *big.Int
}, error) {
	return _Hezrollup.Contract.UnpackBucket(&_Hezrollup.CallOpts, bucket)
}

// WithdrawDelayerContract is a free data retrieval call binding the contract method 0x1b0a8223.
//
// Solidity: function withdrawDelayerContract() view returns(address)
func (_Hezrollup *HezrollupCaller) WithdrawDelayerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "withdrawDelayerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawDelayerContract is a free data retrieval call binding the contract method 0x1b0a8223.
//
// Solidity: function withdrawDelayerContract() view returns(address)
func (_Hezrollup *HezrollupSession) WithdrawDelayerContract() (common.Address, error) {
	return _Hezrollup.Contract.WithdrawDelayerContract(&_Hezrollup.CallOpts)
}

// WithdrawDelayerContract is a free data retrieval call binding the contract method 0x1b0a8223.
//
// Solidity: function withdrawDelayerContract() view returns(address)
func (_Hezrollup *HezrollupCallerSession) WithdrawDelayerContract() (common.Address, error) {
	return _Hezrollup.Contract.WithdrawDelayerContract(&_Hezrollup.CallOpts)
}

// WithdrawVerifier is a free data retrieval call binding the contract method 0x864eb164.
//
// Solidity: function withdrawVerifier() view returns(address)
func (_Hezrollup *HezrollupCaller) WithdrawVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "withdrawVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawVerifier is a free data retrieval call binding the contract method 0x864eb164.
//
// Solidity: function withdrawVerifier() view returns(address)
func (_Hezrollup *HezrollupSession) WithdrawVerifier() (common.Address, error) {
	return _Hezrollup.Contract.WithdrawVerifier(&_Hezrollup.CallOpts)
}

// WithdrawVerifier is a free data retrieval call binding the contract method 0x864eb164.
//
// Solidity: function withdrawVerifier() view returns(address)
func (_Hezrollup *HezrollupCallerSession) WithdrawVerifier() (common.Address, error) {
	return _Hezrollup.Contract.WithdrawVerifier(&_Hezrollup.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint64)
func (_Hezrollup *HezrollupCaller) WithdrawalDelay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Hezrollup.contract.Call(opts, &out, "withdrawalDelay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint64)
func (_Hezrollup *HezrollupSession) WithdrawalDelay() (uint64, error) {
	return _Hezrollup.Contract.WithdrawalDelay(&_Hezrollup.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint64)
func (_Hezrollup *HezrollupCallerSession) WithdrawalDelay() (uint64, error) {
	return _Hezrollup.Contract.WithdrawalDelay(&_Hezrollup.CallOpts)
}

// AddL1Transaction is a paid mutator transaction binding the contract method 0xc7273053.
//
// Solidity: function addL1Transaction(uint256 babyPubKey, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint32 tokenID, uint48 toIdx, bytes permit) payable returns()
func (_Hezrollup *HezrollupTransactor) AddL1Transaction(opts *bind.TransactOpts, babyPubKey *big.Int, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, tokenID uint32, toIdx *big.Int, permit []byte) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "addL1Transaction", babyPubKey, fromIdx, loadAmountF, amountF, tokenID, toIdx, permit)
}

// AddL1Transaction is a paid mutator transaction binding the contract method 0xc7273053.
//
// Solidity: function addL1Transaction(uint256 babyPubKey, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint32 tokenID, uint48 toIdx, bytes permit) payable returns()
func (_Hezrollup *HezrollupSession) AddL1Transaction(babyPubKey *big.Int, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, tokenID uint32, toIdx *big.Int, permit []byte) (*types.Transaction, error) {
	return _Hezrollup.Contract.AddL1Transaction(&_Hezrollup.TransactOpts, babyPubKey, fromIdx, loadAmountF, amountF, tokenID, toIdx, permit)
}

// AddL1Transaction is a paid mutator transaction binding the contract method 0xc7273053.
//
// Solidity: function addL1Transaction(uint256 babyPubKey, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint32 tokenID, uint48 toIdx, bytes permit) payable returns()
func (_Hezrollup *HezrollupTransactorSession) AddL1Transaction(babyPubKey *big.Int, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, tokenID uint32, toIdx *big.Int, permit []byte) (*types.Transaction, error) {
	return _Hezrollup.Contract.AddL1Transaction(&_Hezrollup.TransactOpts, babyPubKey, fromIdx, loadAmountF, amountF, tokenID, toIdx, permit)
}

// AddToken is a paid mutator transaction binding the contract method 0x70c2f1c0.
//
// Solidity: function addToken(address tokenAddress, bytes permit) returns()
func (_Hezrollup *HezrollupTransactor) AddToken(opts *bind.TransactOpts, tokenAddress common.Address, permit []byte) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "addToken", tokenAddress, permit)
}

// AddToken is a paid mutator transaction binding the contract method 0x70c2f1c0.
//
// Solidity: function addToken(address tokenAddress, bytes permit) returns()
func (_Hezrollup *HezrollupSession) AddToken(tokenAddress common.Address, permit []byte) (*types.Transaction, error) {
	return _Hezrollup.Contract.AddToken(&_Hezrollup.TransactOpts, tokenAddress, permit)
}

// AddToken is a paid mutator transaction binding the contract method 0x70c2f1c0.
//
// Solidity: function addToken(address tokenAddress, bytes permit) returns()
func (_Hezrollup *HezrollupTransactorSession) AddToken(tokenAddress common.Address, permit []byte) (*types.Transaction, error) {
	return _Hezrollup.Contract.AddToken(&_Hezrollup.TransactOpts, tokenAddress, permit)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x6e7e1365.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newStRoot, uint256 newExitRoot, bytes encodedL1CoordinatorTx, bytes l1L2TxsData, bytes feeIdxCoordinator, uint8 verifierIdx, bool l1Batch, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Hezrollup *HezrollupTransactor) ForgeBatch(opts *bind.TransactOpts, newLastIdx *big.Int, newStRoot *big.Int, newExitRoot *big.Int, encodedL1CoordinatorTx []byte, l1L2TxsData []byte, feeIdxCoordinator []byte, verifierIdx uint8, l1Batch bool, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "forgeBatch", newLastIdx, newStRoot, newExitRoot, encodedL1CoordinatorTx, l1L2TxsData, feeIdxCoordinator, verifierIdx, l1Batch, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x6e7e1365.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newStRoot, uint256 newExitRoot, bytes encodedL1CoordinatorTx, bytes l1L2TxsData, bytes feeIdxCoordinator, uint8 verifierIdx, bool l1Batch, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Hezrollup *HezrollupSession) ForgeBatch(newLastIdx *big.Int, newStRoot *big.Int, newExitRoot *big.Int, encodedL1CoordinatorTx []byte, l1L2TxsData []byte, feeIdxCoordinator []byte, verifierIdx uint8, l1Batch bool, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Hezrollup.Contract.ForgeBatch(&_Hezrollup.TransactOpts, newLastIdx, newStRoot, newExitRoot, encodedL1CoordinatorTx, l1L2TxsData, feeIdxCoordinator, verifierIdx, l1Batch, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x6e7e1365.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newStRoot, uint256 newExitRoot, bytes encodedL1CoordinatorTx, bytes l1L2TxsData, bytes feeIdxCoordinator, uint8 verifierIdx, bool l1Batch, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Hezrollup *HezrollupTransactorSession) ForgeBatch(newLastIdx *big.Int, newStRoot *big.Int, newExitRoot *big.Int, encodedL1CoordinatorTx []byte, l1L2TxsData []byte, feeIdxCoordinator []byte, verifierIdx uint8, l1Batch bool, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Hezrollup.Contract.ForgeBatch(&_Hezrollup.TransactOpts, newLastIdx, newStRoot, newExitRoot, encodedL1CoordinatorTx, l1L2TxsData, feeIdxCoordinator, verifierIdx, l1Batch, proofA, proofB, proofC)
}

// InitializeHermez is a paid mutator transaction binding the contract method 0x599897e3.
//
// Solidity: function initializeHermez(address[] _verifiers, uint256[] _verifiersParams, address _withdrawVerifier, address _hermezAuctionContract, address _tokenHEZ, uint8 _forgeL1L2BatchTimeout, uint256 _feeAddToken, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements, address _hermezGovernanceAddress, uint64 _withdrawalDelay, address _withdrawDelayerContract) returns()
func (_Hezrollup *HezrollupTransactor) InitializeHermez(opts *bind.TransactOpts, _verifiers []common.Address, _verifiersParams []*big.Int, _withdrawVerifier common.Address, _hermezAuctionContract common.Address, _tokenHEZ common.Address, _forgeL1L2BatchTimeout uint8, _feeAddToken *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address, _hermezGovernanceAddress common.Address, _withdrawalDelay uint64, _withdrawDelayerContract common.Address) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "initializeHermez", _verifiers, _verifiersParams, _withdrawVerifier, _hermezAuctionContract, _tokenHEZ, _forgeL1L2BatchTimeout, _feeAddToken, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements, _hermezGovernanceAddress, _withdrawalDelay, _withdrawDelayerContract)
}

// InitializeHermez is a paid mutator transaction binding the contract method 0x599897e3.
//
// Solidity: function initializeHermez(address[] _verifiers, uint256[] _verifiersParams, address _withdrawVerifier, address _hermezAuctionContract, address _tokenHEZ, uint8 _forgeL1L2BatchTimeout, uint256 _feeAddToken, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements, address _hermezGovernanceAddress, uint64 _withdrawalDelay, address _withdrawDelayerContract) returns()
func (_Hezrollup *HezrollupSession) InitializeHermez(_verifiers []common.Address, _verifiersParams []*big.Int, _withdrawVerifier common.Address, _hermezAuctionContract common.Address, _tokenHEZ common.Address, _forgeL1L2BatchTimeout uint8, _feeAddToken *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address, _hermezGovernanceAddress common.Address, _withdrawalDelay uint64, _withdrawDelayerContract common.Address) (*types.Transaction, error) {
	return _Hezrollup.Contract.InitializeHermez(&_Hezrollup.TransactOpts, _verifiers, _verifiersParams, _withdrawVerifier, _hermezAuctionContract, _tokenHEZ, _forgeL1L2BatchTimeout, _feeAddToken, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements, _hermezGovernanceAddress, _withdrawalDelay, _withdrawDelayerContract)
}

// InitializeHermez is a paid mutator transaction binding the contract method 0x599897e3.
//
// Solidity: function initializeHermez(address[] _verifiers, uint256[] _verifiersParams, address _withdrawVerifier, address _hermezAuctionContract, address _tokenHEZ, uint8 _forgeL1L2BatchTimeout, uint256 _feeAddToken, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements, address _hermezGovernanceAddress, uint64 _withdrawalDelay, address _withdrawDelayerContract) returns()
func (_Hezrollup *HezrollupTransactorSession) InitializeHermez(_verifiers []common.Address, _verifiersParams []*big.Int, _withdrawVerifier common.Address, _hermezAuctionContract common.Address, _tokenHEZ common.Address, _forgeL1L2BatchTimeout uint8, _feeAddToken *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address, _hermezGovernanceAddress common.Address, _withdrawalDelay uint64, _withdrawDelayerContract common.Address) (*types.Transaction, error) {
	return _Hezrollup.Contract.InitializeHermez(&_Hezrollup.TransactOpts, _verifiers, _verifiersParams, _withdrawVerifier, _hermezAuctionContract, _tokenHEZ, _forgeL1L2BatchTimeout, _feeAddToken, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements, _hermezGovernanceAddress, _withdrawalDelay, _withdrawDelayerContract)
}

// SafeMode is a paid mutator transaction binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() returns()
func (_Hezrollup *HezrollupTransactor) SafeMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "safeMode")
}

// SafeMode is a paid mutator transaction binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() returns()
func (_Hezrollup *HezrollupSession) SafeMode() (*types.Transaction, error) {
	return _Hezrollup.Contract.SafeMode(&_Hezrollup.TransactOpts)
}

// SafeMode is a paid mutator transaction binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() returns()
func (_Hezrollup *HezrollupTransactorSession) SafeMode() (*types.Transaction, error) {
	return _Hezrollup.Contract.SafeMode(&_Hezrollup.TransactOpts)
}

// UpdateBucketsParameters is a paid mutator transaction binding the contract method 0xac300ec9.
//
// Solidity: function updateBucketsParameters(uint256[] newBuckets) returns()
func (_Hezrollup *HezrollupTransactor) UpdateBucketsParameters(opts *bind.TransactOpts, newBuckets []*big.Int) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "updateBucketsParameters", newBuckets)
}

// UpdateBucketsParameters is a paid mutator transaction binding the contract method 0xac300ec9.
//
// Solidity: function updateBucketsParameters(uint256[] newBuckets) returns()
func (_Hezrollup *HezrollupSession) UpdateBucketsParameters(newBuckets []*big.Int) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateBucketsParameters(&_Hezrollup.TransactOpts, newBuckets)
}

// UpdateBucketsParameters is a paid mutator transaction binding the contract method 0xac300ec9.
//
// Solidity: function updateBucketsParameters(uint256[] newBuckets) returns()
func (_Hezrollup *HezrollupTransactorSession) UpdateBucketsParameters(newBuckets []*big.Int) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateBucketsParameters(&_Hezrollup.TransactOpts, newBuckets)
}

// UpdateFeeAddToken is a paid mutator transaction binding the contract method 0x314e5eda.
//
// Solidity: function updateFeeAddToken(uint256 newFeeAddToken) returns()
func (_Hezrollup *HezrollupTransactor) UpdateFeeAddToken(opts *bind.TransactOpts, newFeeAddToken *big.Int) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "updateFeeAddToken", newFeeAddToken)
}

// UpdateFeeAddToken is a paid mutator transaction binding the contract method 0x314e5eda.
//
// Solidity: function updateFeeAddToken(uint256 newFeeAddToken) returns()
func (_Hezrollup *HezrollupSession) UpdateFeeAddToken(newFeeAddToken *big.Int) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateFeeAddToken(&_Hezrollup.TransactOpts, newFeeAddToken)
}

// UpdateFeeAddToken is a paid mutator transaction binding the contract method 0x314e5eda.
//
// Solidity: function updateFeeAddToken(uint256 newFeeAddToken) returns()
func (_Hezrollup *HezrollupTransactorSession) UpdateFeeAddToken(newFeeAddToken *big.Int) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateFeeAddToken(&_Hezrollup.TransactOpts, newFeeAddToken)
}

// UpdateForgeL1L2BatchTimeout is a paid mutator transaction binding the contract method 0xcbd7b5fb.
//
// Solidity: function updateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout) returns()
func (_Hezrollup *HezrollupTransactor) UpdateForgeL1L2BatchTimeout(opts *bind.TransactOpts, newForgeL1L2BatchTimeout uint8) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "updateForgeL1L2BatchTimeout", newForgeL1L2BatchTimeout)
}

// UpdateForgeL1L2BatchTimeout is a paid mutator transaction binding the contract method 0xcbd7b5fb.
//
// Solidity: function updateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout) returns()
func (_Hezrollup *HezrollupSession) UpdateForgeL1L2BatchTimeout(newForgeL1L2BatchTimeout uint8) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateForgeL1L2BatchTimeout(&_Hezrollup.TransactOpts, newForgeL1L2BatchTimeout)
}

// UpdateForgeL1L2BatchTimeout is a paid mutator transaction binding the contract method 0xcbd7b5fb.
//
// Solidity: function updateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout) returns()
func (_Hezrollup *HezrollupTransactorSession) UpdateForgeL1L2BatchTimeout(newForgeL1L2BatchTimeout uint8) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateForgeL1L2BatchTimeout(&_Hezrollup.TransactOpts, newForgeL1L2BatchTimeout)
}

// UpdateTokenExchange is a paid mutator transaction binding the contract method 0x1a748c2d.
//
// Solidity: function updateTokenExchange(address[] addressArray, uint64[] valueArray) returns()
func (_Hezrollup *HezrollupTransactor) UpdateTokenExchange(opts *bind.TransactOpts, addressArray []common.Address, valueArray []uint64) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "updateTokenExchange", addressArray, valueArray)
}

// UpdateTokenExchange is a paid mutator transaction binding the contract method 0x1a748c2d.
//
// Solidity: function updateTokenExchange(address[] addressArray, uint64[] valueArray) returns()
func (_Hezrollup *HezrollupSession) UpdateTokenExchange(addressArray []common.Address, valueArray []uint64) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateTokenExchange(&_Hezrollup.TransactOpts, addressArray, valueArray)
}

// UpdateTokenExchange is a paid mutator transaction binding the contract method 0x1a748c2d.
//
// Solidity: function updateTokenExchange(address[] addressArray, uint64[] valueArray) returns()
func (_Hezrollup *HezrollupTransactorSession) UpdateTokenExchange(addressArray []common.Address, valueArray []uint64) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateTokenExchange(&_Hezrollup.TransactOpts, addressArray, valueArray)
}

// UpdateVerifiers is a paid mutator transaction binding the contract method 0x960207c0.
//
// Solidity: function updateVerifiers() returns()
func (_Hezrollup *HezrollupTransactor) UpdateVerifiers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "updateVerifiers")
}

// UpdateVerifiers is a paid mutator transaction binding the contract method 0x960207c0.
//
// Solidity: function updateVerifiers() returns()
func (_Hezrollup *HezrollupSession) UpdateVerifiers() (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateVerifiers(&_Hezrollup.TransactOpts)
}

// UpdateVerifiers is a paid mutator transaction binding the contract method 0x960207c0.
//
// Solidity: function updateVerifiers() returns()
func (_Hezrollup *HezrollupTransactorSession) UpdateVerifiers() (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateVerifiers(&_Hezrollup.TransactOpts)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0xef4a5c4a.
//
// Solidity: function updateWithdrawalDelay(uint64 newWithdrawalDelay) returns()
func (_Hezrollup *HezrollupTransactor) UpdateWithdrawalDelay(opts *bind.TransactOpts, newWithdrawalDelay uint64) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "updateWithdrawalDelay", newWithdrawalDelay)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0xef4a5c4a.
//
// Solidity: function updateWithdrawalDelay(uint64 newWithdrawalDelay) returns()
func (_Hezrollup *HezrollupSession) UpdateWithdrawalDelay(newWithdrawalDelay uint64) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateWithdrawalDelay(&_Hezrollup.TransactOpts, newWithdrawalDelay)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0xef4a5c4a.
//
// Solidity: function updateWithdrawalDelay(uint64 newWithdrawalDelay) returns()
func (_Hezrollup *HezrollupTransactorSession) UpdateWithdrawalDelay(newWithdrawalDelay uint64) (*types.Transaction, error) {
	return _Hezrollup.Contract.UpdateWithdrawalDelay(&_Hezrollup.TransactOpts, newWithdrawalDelay)
}

// WithdrawCircuit is a paid mutator transaction binding the contract method 0x9ce2ad42.
//
// Solidity: function withdrawCircuit(uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint32 tokenID, uint192 amount, uint32 numExitRoot, uint48 idx, bool instantWithdraw) returns()
func (_Hezrollup *HezrollupTransactor) WithdrawCircuit(opts *bind.TransactOpts, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, tokenID uint32, amount *big.Int, numExitRoot uint32, idx *big.Int, instantWithdraw bool) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "withdrawCircuit", proofA, proofB, proofC, tokenID, amount, numExitRoot, idx, instantWithdraw)
}

// WithdrawCircuit is a paid mutator transaction binding the contract method 0x9ce2ad42.
//
// Solidity: function withdrawCircuit(uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint32 tokenID, uint192 amount, uint32 numExitRoot, uint48 idx, bool instantWithdraw) returns()
func (_Hezrollup *HezrollupSession) WithdrawCircuit(proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, tokenID uint32, amount *big.Int, numExitRoot uint32, idx *big.Int, instantWithdraw bool) (*types.Transaction, error) {
	return _Hezrollup.Contract.WithdrawCircuit(&_Hezrollup.TransactOpts, proofA, proofB, proofC, tokenID, amount, numExitRoot, idx, instantWithdraw)
}

// WithdrawCircuit is a paid mutator transaction binding the contract method 0x9ce2ad42.
//
// Solidity: function withdrawCircuit(uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint32 tokenID, uint192 amount, uint32 numExitRoot, uint48 idx, bool instantWithdraw) returns()
func (_Hezrollup *HezrollupTransactorSession) WithdrawCircuit(proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, tokenID uint32, amount *big.Int, numExitRoot uint32, idx *big.Int, instantWithdraw bool) (*types.Transaction, error) {
	return _Hezrollup.Contract.WithdrawCircuit(&_Hezrollup.TransactOpts, proofA, proofB, proofC, tokenID, amount, numExitRoot, idx, instantWithdraw)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0xd9d4ca44.
//
// Solidity: function withdrawMerkleProof(uint32 tokenID, uint192 amount, uint256 babyPubKey, uint32 numExitRoot, uint256[] siblings, uint48 idx, bool instantWithdraw) returns()
func (_Hezrollup *HezrollupTransactor) WithdrawMerkleProof(opts *bind.TransactOpts, tokenID uint32, amount *big.Int, babyPubKey *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int, instantWithdraw bool) (*types.Transaction, error) {
	return _Hezrollup.contract.Transact(opts, "withdrawMerkleProof", tokenID, amount, babyPubKey, numExitRoot, siblings, idx, instantWithdraw)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0xd9d4ca44.
//
// Solidity: function withdrawMerkleProof(uint32 tokenID, uint192 amount, uint256 babyPubKey, uint32 numExitRoot, uint256[] siblings, uint48 idx, bool instantWithdraw) returns()
func (_Hezrollup *HezrollupSession) WithdrawMerkleProof(tokenID uint32, amount *big.Int, babyPubKey *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int, instantWithdraw bool) (*types.Transaction, error) {
	return _Hezrollup.Contract.WithdrawMerkleProof(&_Hezrollup.TransactOpts, tokenID, amount, babyPubKey, numExitRoot, siblings, idx, instantWithdraw)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0xd9d4ca44.
//
// Solidity: function withdrawMerkleProof(uint32 tokenID, uint192 amount, uint256 babyPubKey, uint32 numExitRoot, uint256[] siblings, uint48 idx, bool instantWithdraw) returns()
func (_Hezrollup *HezrollupTransactorSession) WithdrawMerkleProof(tokenID uint32, amount *big.Int, babyPubKey *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int, instantWithdraw bool) (*types.Transaction, error) {
	return _Hezrollup.Contract.WithdrawMerkleProof(&_Hezrollup.TransactOpts, tokenID, amount, babyPubKey, numExitRoot, siblings, idx, instantWithdraw)
}

// HezrollupAddTokenIterator is returned from FilterAddToken and is used to iterate over the raw logs and unpacked data for AddToken events raised by the Hezrollup contract.
type HezrollupAddTokenIterator struct {
	Event *HezrollupAddToken // Event containing the contract specifics and raw log

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
func (it *HezrollupAddTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupAddToken)
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
		it.Event = new(HezrollupAddToken)
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
func (it *HezrollupAddTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupAddTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupAddToken represents a AddToken event raised by the Hezrollup contract.
type HezrollupAddToken struct {
	TokenAddress common.Address
	TokenID      uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddToken is a free log retrieval operation binding the contract event 0xcb73d161edb7cd4fb1d92fedfd2555384fd997fd44ab507656f8c81e15747dde.
//
// Solidity: event AddToken(address indexed tokenAddress, uint32 tokenID)
func (_Hezrollup *HezrollupFilterer) FilterAddToken(opts *bind.FilterOpts, tokenAddress []common.Address) (*HezrollupAddTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "AddToken", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &HezrollupAddTokenIterator{contract: _Hezrollup.contract, event: "AddToken", logs: logs, sub: sub}, nil
}

// WatchAddToken is a free log subscription operation binding the contract event 0xcb73d161edb7cd4fb1d92fedfd2555384fd997fd44ab507656f8c81e15747dde.
//
// Solidity: event AddToken(address indexed tokenAddress, uint32 tokenID)
func (_Hezrollup *HezrollupFilterer) WatchAddToken(opts *bind.WatchOpts, sink chan<- *HezrollupAddToken, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "AddToken", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupAddToken)
				if err := _Hezrollup.contract.UnpackLog(event, "AddToken", log); err != nil {
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

// ParseAddToken is a log parse operation binding the contract event 0xcb73d161edb7cd4fb1d92fedfd2555384fd997fd44ab507656f8c81e15747dde.
//
// Solidity: event AddToken(address indexed tokenAddress, uint32 tokenID)
func (_Hezrollup *HezrollupFilterer) ParseAddToken(log types.Log) (*HezrollupAddToken, error) {
	event := new(HezrollupAddToken)
	if err := _Hezrollup.contract.UnpackLog(event, "AddToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupForgeBatchIterator is returned from FilterForgeBatch and is used to iterate over the raw logs and unpacked data for ForgeBatch events raised by the Hezrollup contract.
type HezrollupForgeBatchIterator struct {
	Event *HezrollupForgeBatch // Event containing the contract specifics and raw log

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
func (it *HezrollupForgeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupForgeBatch)
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
		it.Event = new(HezrollupForgeBatch)
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
func (it *HezrollupForgeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupForgeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupForgeBatch represents a ForgeBatch event raised by the Hezrollup contract.
type HezrollupForgeBatch struct {
	BatchNum     uint32
	L1UserTxsLen uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterForgeBatch is a free log retrieval operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Hezrollup *HezrollupFilterer) FilterForgeBatch(opts *bind.FilterOpts, batchNum []uint32) (*HezrollupForgeBatchIterator, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "ForgeBatch", batchNumRule)
	if err != nil {
		return nil, err
	}
	return &HezrollupForgeBatchIterator{contract: _Hezrollup.contract, event: "ForgeBatch", logs: logs, sub: sub}, nil
}

// WatchForgeBatch is a free log subscription operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Hezrollup *HezrollupFilterer) WatchForgeBatch(opts *bind.WatchOpts, sink chan<- *HezrollupForgeBatch, batchNum []uint32) (event.Subscription, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "ForgeBatch", batchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupForgeBatch)
				if err := _Hezrollup.contract.UnpackLog(event, "ForgeBatch", log); err != nil {
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

// ParseForgeBatch is a log parse operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Hezrollup *HezrollupFilterer) ParseForgeBatch(log types.Log) (*HezrollupForgeBatch, error) {
	event := new(HezrollupForgeBatch)
	if err := _Hezrollup.contract.UnpackLog(event, "ForgeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupInitializeHermezEventIterator is returned from FilterInitializeHermezEvent and is used to iterate over the raw logs and unpacked data for InitializeHermezEvent events raised by the Hezrollup contract.
type HezrollupInitializeHermezEventIterator struct {
	Event *HezrollupInitializeHermezEvent // Event containing the contract specifics and raw log

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
func (it *HezrollupInitializeHermezEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupInitializeHermezEvent)
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
		it.Event = new(HezrollupInitializeHermezEvent)
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
func (it *HezrollupInitializeHermezEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupInitializeHermezEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupInitializeHermezEvent represents a InitializeHermezEvent event raised by the Hezrollup contract.
type HezrollupInitializeHermezEvent struct {
	ForgeL1L2BatchTimeout uint8
	FeeAddToken           *big.Int
	WithdrawalDelay       uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHermezEvent is a free log retrieval operation binding the contract event 0xc5272ad4c8d9f2e9af2f9555c11ead049be22b6e45c16975adc82371b7cd1040.
//
// Solidity: event InitializeHermezEvent(uint8 forgeL1L2BatchTimeout, uint256 feeAddToken, uint64 withdrawalDelay)
func (_Hezrollup *HezrollupFilterer) FilterInitializeHermezEvent(opts *bind.FilterOpts) (*HezrollupInitializeHermezEventIterator, error) {

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "InitializeHermezEvent")
	if err != nil {
		return nil, err
	}
	return &HezrollupInitializeHermezEventIterator{contract: _Hezrollup.contract, event: "InitializeHermezEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHermezEvent is a free log subscription operation binding the contract event 0xc5272ad4c8d9f2e9af2f9555c11ead049be22b6e45c16975adc82371b7cd1040.
//
// Solidity: event InitializeHermezEvent(uint8 forgeL1L2BatchTimeout, uint256 feeAddToken, uint64 withdrawalDelay)
func (_Hezrollup *HezrollupFilterer) WatchInitializeHermezEvent(opts *bind.WatchOpts, sink chan<- *HezrollupInitializeHermezEvent) (event.Subscription, error) {

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "InitializeHermezEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupInitializeHermezEvent)
				if err := _Hezrollup.contract.UnpackLog(event, "InitializeHermezEvent", log); err != nil {
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

// ParseInitializeHermezEvent is a log parse operation binding the contract event 0xc5272ad4c8d9f2e9af2f9555c11ead049be22b6e45c16975adc82371b7cd1040.
//
// Solidity: event InitializeHermezEvent(uint8 forgeL1L2BatchTimeout, uint256 feeAddToken, uint64 withdrawalDelay)
func (_Hezrollup *HezrollupFilterer) ParseInitializeHermezEvent(log types.Log) (*HezrollupInitializeHermezEvent, error) {
	event := new(HezrollupInitializeHermezEvent)
	if err := _Hezrollup.contract.UnpackLog(event, "InitializeHermezEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupL1UserTxEventIterator is returned from FilterL1UserTxEvent and is used to iterate over the raw logs and unpacked data for L1UserTxEvent events raised by the Hezrollup contract.
type HezrollupL1UserTxEventIterator struct {
	Event *HezrollupL1UserTxEvent // Event containing the contract specifics and raw log

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
func (it *HezrollupL1UserTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupL1UserTxEvent)
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
		it.Event = new(HezrollupL1UserTxEvent)
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
func (it *HezrollupL1UserTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupL1UserTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupL1UserTxEvent represents a L1UserTxEvent event raised by the Hezrollup contract.
type HezrollupL1UserTxEvent struct {
	QueueIndex uint32
	Position   uint8
	L1UserTx   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL1UserTxEvent is a free log retrieval operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Hezrollup *HezrollupFilterer) FilterL1UserTxEvent(opts *bind.FilterOpts, queueIndex []uint32, position []uint8) (*HezrollupL1UserTxEventIterator, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "L1UserTxEvent", queueIndexRule, positionRule)
	if err != nil {
		return nil, err
	}
	return &HezrollupL1UserTxEventIterator{contract: _Hezrollup.contract, event: "L1UserTxEvent", logs: logs, sub: sub}, nil
}

// WatchL1UserTxEvent is a free log subscription operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Hezrollup *HezrollupFilterer) WatchL1UserTxEvent(opts *bind.WatchOpts, sink chan<- *HezrollupL1UserTxEvent, queueIndex []uint32, position []uint8) (event.Subscription, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "L1UserTxEvent", queueIndexRule, positionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupL1UserTxEvent)
				if err := _Hezrollup.contract.UnpackLog(event, "L1UserTxEvent", log); err != nil {
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

// ParseL1UserTxEvent is a log parse operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Hezrollup *HezrollupFilterer) ParseL1UserTxEvent(log types.Log) (*HezrollupL1UserTxEvent, error) {
	event := new(HezrollupL1UserTxEvent)
	if err := _Hezrollup.contract.UnpackLog(event, "L1UserTxEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupSafeModeIterator is returned from FilterSafeMode and is used to iterate over the raw logs and unpacked data for SafeMode events raised by the Hezrollup contract.
type HezrollupSafeModeIterator struct {
	Event *HezrollupSafeMode // Event containing the contract specifics and raw log

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
func (it *HezrollupSafeModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupSafeMode)
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
		it.Event = new(HezrollupSafeMode)
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
func (it *HezrollupSafeModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupSafeModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupSafeMode represents a SafeMode event raised by the Hezrollup contract.
type HezrollupSafeMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSafeMode is a free log retrieval operation binding the contract event 0x0410e6ef2bd89ecf5b2dc2f62157f9863e09e89cb7c7f1abb7d4ec43a6019d1e.
//
// Solidity: event SafeMode()
func (_Hezrollup *HezrollupFilterer) FilterSafeMode(opts *bind.FilterOpts) (*HezrollupSafeModeIterator, error) {

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "SafeMode")
	if err != nil {
		return nil, err
	}
	return &HezrollupSafeModeIterator{contract: _Hezrollup.contract, event: "SafeMode", logs: logs, sub: sub}, nil
}

// WatchSafeMode is a free log subscription operation binding the contract event 0x0410e6ef2bd89ecf5b2dc2f62157f9863e09e89cb7c7f1abb7d4ec43a6019d1e.
//
// Solidity: event SafeMode()
func (_Hezrollup *HezrollupFilterer) WatchSafeMode(opts *bind.WatchOpts, sink chan<- *HezrollupSafeMode) (event.Subscription, error) {

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "SafeMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupSafeMode)
				if err := _Hezrollup.contract.UnpackLog(event, "SafeMode", log); err != nil {
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

// ParseSafeMode is a log parse operation binding the contract event 0x0410e6ef2bd89ecf5b2dc2f62157f9863e09e89cb7c7f1abb7d4ec43a6019d1e.
//
// Solidity: event SafeMode()
func (_Hezrollup *HezrollupFilterer) ParseSafeMode(log types.Log) (*HezrollupSafeMode, error) {
	event := new(HezrollupSafeMode)
	if err := _Hezrollup.contract.UnpackLog(event, "SafeMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupUpdateBucketWithdrawIterator is returned from FilterUpdateBucketWithdraw and is used to iterate over the raw logs and unpacked data for UpdateBucketWithdraw events raised by the Hezrollup contract.
type HezrollupUpdateBucketWithdrawIterator struct {
	Event *HezrollupUpdateBucketWithdraw // Event containing the contract specifics and raw log

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
func (it *HezrollupUpdateBucketWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupUpdateBucketWithdraw)
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
		it.Event = new(HezrollupUpdateBucketWithdraw)
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
func (it *HezrollupUpdateBucketWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupUpdateBucketWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupUpdateBucketWithdraw represents a UpdateBucketWithdraw event raised by the Hezrollup contract.
type HezrollupUpdateBucketWithdraw struct {
	NumBucket   uint8
	BlockStamp  *big.Int
	Withdrawals *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateBucketWithdraw is a free log retrieval operation binding the contract event 0xa35fe9a9e21cdbbc4774aa8a56e7b97ea9c06afc09ffb06af593d26951e350aa.
//
// Solidity: event UpdateBucketWithdraw(uint8 indexed numBucket, uint256 indexed blockStamp, uint256 withdrawals)
func (_Hezrollup *HezrollupFilterer) FilterUpdateBucketWithdraw(opts *bind.FilterOpts, numBucket []uint8, blockStamp []*big.Int) (*HezrollupUpdateBucketWithdrawIterator, error) {

	var numBucketRule []interface{}
	for _, numBucketItem := range numBucket {
		numBucketRule = append(numBucketRule, numBucketItem)
	}
	var blockStampRule []interface{}
	for _, blockStampItem := range blockStamp {
		blockStampRule = append(blockStampRule, blockStampItem)
	}

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "UpdateBucketWithdraw", numBucketRule, blockStampRule)
	if err != nil {
		return nil, err
	}
	return &HezrollupUpdateBucketWithdrawIterator{contract: _Hezrollup.contract, event: "UpdateBucketWithdraw", logs: logs, sub: sub}, nil
}

// WatchUpdateBucketWithdraw is a free log subscription operation binding the contract event 0xa35fe9a9e21cdbbc4774aa8a56e7b97ea9c06afc09ffb06af593d26951e350aa.
//
// Solidity: event UpdateBucketWithdraw(uint8 indexed numBucket, uint256 indexed blockStamp, uint256 withdrawals)
func (_Hezrollup *HezrollupFilterer) WatchUpdateBucketWithdraw(opts *bind.WatchOpts, sink chan<- *HezrollupUpdateBucketWithdraw, numBucket []uint8, blockStamp []*big.Int) (event.Subscription, error) {

	var numBucketRule []interface{}
	for _, numBucketItem := range numBucket {
		numBucketRule = append(numBucketRule, numBucketItem)
	}
	var blockStampRule []interface{}
	for _, blockStampItem := range blockStamp {
		blockStampRule = append(blockStampRule, blockStampItem)
	}

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "UpdateBucketWithdraw", numBucketRule, blockStampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupUpdateBucketWithdraw)
				if err := _Hezrollup.contract.UnpackLog(event, "UpdateBucketWithdraw", log); err != nil {
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

// ParseUpdateBucketWithdraw is a log parse operation binding the contract event 0xa35fe9a9e21cdbbc4774aa8a56e7b97ea9c06afc09ffb06af593d26951e350aa.
//
// Solidity: event UpdateBucketWithdraw(uint8 indexed numBucket, uint256 indexed blockStamp, uint256 withdrawals)
func (_Hezrollup *HezrollupFilterer) ParseUpdateBucketWithdraw(log types.Log) (*HezrollupUpdateBucketWithdraw, error) {
	event := new(HezrollupUpdateBucketWithdraw)
	if err := _Hezrollup.contract.UnpackLog(event, "UpdateBucketWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupUpdateBucketsParametersIterator is returned from FilterUpdateBucketsParameters and is used to iterate over the raw logs and unpacked data for UpdateBucketsParameters events raised by the Hezrollup contract.
type HezrollupUpdateBucketsParametersIterator struct {
	Event *HezrollupUpdateBucketsParameters // Event containing the contract specifics and raw log

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
func (it *HezrollupUpdateBucketsParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupUpdateBucketsParameters)
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
		it.Event = new(HezrollupUpdateBucketsParameters)
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
func (it *HezrollupUpdateBucketsParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupUpdateBucketsParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupUpdateBucketsParameters represents a UpdateBucketsParameters event raised by the Hezrollup contract.
type HezrollupUpdateBucketsParameters struct {
	ArrayBuckets []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateBucketsParameters is a free log retrieval operation binding the contract event 0xd4904145d7eae889c5493798579680417459783db0fa67398bea50e56859075f.
//
// Solidity: event UpdateBucketsParameters(uint256[] arrayBuckets)
func (_Hezrollup *HezrollupFilterer) FilterUpdateBucketsParameters(opts *bind.FilterOpts) (*HezrollupUpdateBucketsParametersIterator, error) {

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "UpdateBucketsParameters")
	if err != nil {
		return nil, err
	}
	return &HezrollupUpdateBucketsParametersIterator{contract: _Hezrollup.contract, event: "UpdateBucketsParameters", logs: logs, sub: sub}, nil
}

// WatchUpdateBucketsParameters is a free log subscription operation binding the contract event 0xd4904145d7eae889c5493798579680417459783db0fa67398bea50e56859075f.
//
// Solidity: event UpdateBucketsParameters(uint256[] arrayBuckets)
func (_Hezrollup *HezrollupFilterer) WatchUpdateBucketsParameters(opts *bind.WatchOpts, sink chan<- *HezrollupUpdateBucketsParameters) (event.Subscription, error) {

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "UpdateBucketsParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupUpdateBucketsParameters)
				if err := _Hezrollup.contract.UnpackLog(event, "UpdateBucketsParameters", log); err != nil {
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

// ParseUpdateBucketsParameters is a log parse operation binding the contract event 0xd4904145d7eae889c5493798579680417459783db0fa67398bea50e56859075f.
//
// Solidity: event UpdateBucketsParameters(uint256[] arrayBuckets)
func (_Hezrollup *HezrollupFilterer) ParseUpdateBucketsParameters(log types.Log) (*HezrollupUpdateBucketsParameters, error) {
	event := new(HezrollupUpdateBucketsParameters)
	if err := _Hezrollup.contract.UnpackLog(event, "UpdateBucketsParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupUpdateFeeAddTokenIterator is returned from FilterUpdateFeeAddToken and is used to iterate over the raw logs and unpacked data for UpdateFeeAddToken events raised by the Hezrollup contract.
type HezrollupUpdateFeeAddTokenIterator struct {
	Event *HezrollupUpdateFeeAddToken // Event containing the contract specifics and raw log

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
func (it *HezrollupUpdateFeeAddTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupUpdateFeeAddToken)
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
		it.Event = new(HezrollupUpdateFeeAddToken)
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
func (it *HezrollupUpdateFeeAddTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupUpdateFeeAddTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupUpdateFeeAddToken represents a UpdateFeeAddToken event raised by the Hezrollup contract.
type HezrollupUpdateFeeAddToken struct {
	NewFeeAddToken *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeAddToken is a free log retrieval operation binding the contract event 0xd1c873cd16013f0dc5f37992c0d12794389698512895ec036a568e393b46e3c1.
//
// Solidity: event UpdateFeeAddToken(uint256 newFeeAddToken)
func (_Hezrollup *HezrollupFilterer) FilterUpdateFeeAddToken(opts *bind.FilterOpts) (*HezrollupUpdateFeeAddTokenIterator, error) {

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "UpdateFeeAddToken")
	if err != nil {
		return nil, err
	}
	return &HezrollupUpdateFeeAddTokenIterator{contract: _Hezrollup.contract, event: "UpdateFeeAddToken", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeAddToken is a free log subscription operation binding the contract event 0xd1c873cd16013f0dc5f37992c0d12794389698512895ec036a568e393b46e3c1.
//
// Solidity: event UpdateFeeAddToken(uint256 newFeeAddToken)
func (_Hezrollup *HezrollupFilterer) WatchUpdateFeeAddToken(opts *bind.WatchOpts, sink chan<- *HezrollupUpdateFeeAddToken) (event.Subscription, error) {

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "UpdateFeeAddToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupUpdateFeeAddToken)
				if err := _Hezrollup.contract.UnpackLog(event, "UpdateFeeAddToken", log); err != nil {
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

// ParseUpdateFeeAddToken is a log parse operation binding the contract event 0xd1c873cd16013f0dc5f37992c0d12794389698512895ec036a568e393b46e3c1.
//
// Solidity: event UpdateFeeAddToken(uint256 newFeeAddToken)
func (_Hezrollup *HezrollupFilterer) ParseUpdateFeeAddToken(log types.Log) (*HezrollupUpdateFeeAddToken, error) {
	event := new(HezrollupUpdateFeeAddToken)
	if err := _Hezrollup.contract.UnpackLog(event, "UpdateFeeAddToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupUpdateForgeL1L2BatchTimeoutIterator is returned from FilterUpdateForgeL1L2BatchTimeout and is used to iterate over the raw logs and unpacked data for UpdateForgeL1L2BatchTimeout events raised by the Hezrollup contract.
type HezrollupUpdateForgeL1L2BatchTimeoutIterator struct {
	Event *HezrollupUpdateForgeL1L2BatchTimeout // Event containing the contract specifics and raw log

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
func (it *HezrollupUpdateForgeL1L2BatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupUpdateForgeL1L2BatchTimeout)
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
		it.Event = new(HezrollupUpdateForgeL1L2BatchTimeout)
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
func (it *HezrollupUpdateForgeL1L2BatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupUpdateForgeL1L2BatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupUpdateForgeL1L2BatchTimeout represents a UpdateForgeL1L2BatchTimeout event raised by the Hezrollup contract.
type HezrollupUpdateForgeL1L2BatchTimeout struct {
	NewForgeL1L2BatchTimeout uint8
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUpdateForgeL1L2BatchTimeout is a free log retrieval operation binding the contract event 0xff6221781ac525b04585dbb55cd2ebd2a92c828ca3e42b23813a1137ac974431.
//
// Solidity: event UpdateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout)
func (_Hezrollup *HezrollupFilterer) FilterUpdateForgeL1L2BatchTimeout(opts *bind.FilterOpts) (*HezrollupUpdateForgeL1L2BatchTimeoutIterator, error) {

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "UpdateForgeL1L2BatchTimeout")
	if err != nil {
		return nil, err
	}
	return &HezrollupUpdateForgeL1L2BatchTimeoutIterator{contract: _Hezrollup.contract, event: "UpdateForgeL1L2BatchTimeout", logs: logs, sub: sub}, nil
}

// WatchUpdateForgeL1L2BatchTimeout is a free log subscription operation binding the contract event 0xff6221781ac525b04585dbb55cd2ebd2a92c828ca3e42b23813a1137ac974431.
//
// Solidity: event UpdateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout)
func (_Hezrollup *HezrollupFilterer) WatchUpdateForgeL1L2BatchTimeout(opts *bind.WatchOpts, sink chan<- *HezrollupUpdateForgeL1L2BatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "UpdateForgeL1L2BatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupUpdateForgeL1L2BatchTimeout)
				if err := _Hezrollup.contract.UnpackLog(event, "UpdateForgeL1L2BatchTimeout", log); err != nil {
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

// ParseUpdateForgeL1L2BatchTimeout is a log parse operation binding the contract event 0xff6221781ac525b04585dbb55cd2ebd2a92c828ca3e42b23813a1137ac974431.
//
// Solidity: event UpdateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout)
func (_Hezrollup *HezrollupFilterer) ParseUpdateForgeL1L2BatchTimeout(log types.Log) (*HezrollupUpdateForgeL1L2BatchTimeout, error) {
	event := new(HezrollupUpdateForgeL1L2BatchTimeout)
	if err := _Hezrollup.contract.UnpackLog(event, "UpdateForgeL1L2BatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupUpdateTokenExchangeIterator is returned from FilterUpdateTokenExchange and is used to iterate over the raw logs and unpacked data for UpdateTokenExchange events raised by the Hezrollup contract.
type HezrollupUpdateTokenExchangeIterator struct {
	Event *HezrollupUpdateTokenExchange // Event containing the contract specifics and raw log

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
func (it *HezrollupUpdateTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupUpdateTokenExchange)
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
		it.Event = new(HezrollupUpdateTokenExchange)
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
func (it *HezrollupUpdateTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupUpdateTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupUpdateTokenExchange represents a UpdateTokenExchange event raised by the Hezrollup contract.
type HezrollupUpdateTokenExchange struct {
	AddressArray []common.Address
	ValueArray   []uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenExchange is a free log retrieval operation binding the contract event 0x10ff643ebeca3e33002e61b76fa85e7e10091e30afa39295f91af9838b3033b3.
//
// Solidity: event UpdateTokenExchange(address[] addressArray, uint64[] valueArray)
func (_Hezrollup *HezrollupFilterer) FilterUpdateTokenExchange(opts *bind.FilterOpts) (*HezrollupUpdateTokenExchangeIterator, error) {

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "UpdateTokenExchange")
	if err != nil {
		return nil, err
	}
	return &HezrollupUpdateTokenExchangeIterator{contract: _Hezrollup.contract, event: "UpdateTokenExchange", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenExchange is a free log subscription operation binding the contract event 0x10ff643ebeca3e33002e61b76fa85e7e10091e30afa39295f91af9838b3033b3.
//
// Solidity: event UpdateTokenExchange(address[] addressArray, uint64[] valueArray)
func (_Hezrollup *HezrollupFilterer) WatchUpdateTokenExchange(opts *bind.WatchOpts, sink chan<- *HezrollupUpdateTokenExchange) (event.Subscription, error) {

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "UpdateTokenExchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupUpdateTokenExchange)
				if err := _Hezrollup.contract.UnpackLog(event, "UpdateTokenExchange", log); err != nil {
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

// ParseUpdateTokenExchange is a log parse operation binding the contract event 0x10ff643ebeca3e33002e61b76fa85e7e10091e30afa39295f91af9838b3033b3.
//
// Solidity: event UpdateTokenExchange(address[] addressArray, uint64[] valueArray)
func (_Hezrollup *HezrollupFilterer) ParseUpdateTokenExchange(log types.Log) (*HezrollupUpdateTokenExchange, error) {
	event := new(HezrollupUpdateTokenExchange)
	if err := _Hezrollup.contract.UnpackLog(event, "UpdateTokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupUpdateWithdrawalDelayIterator is returned from FilterUpdateWithdrawalDelay and is used to iterate over the raw logs and unpacked data for UpdateWithdrawalDelay events raised by the Hezrollup contract.
type HezrollupUpdateWithdrawalDelayIterator struct {
	Event *HezrollupUpdateWithdrawalDelay // Event containing the contract specifics and raw log

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
func (it *HezrollupUpdateWithdrawalDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupUpdateWithdrawalDelay)
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
		it.Event = new(HezrollupUpdateWithdrawalDelay)
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
func (it *HezrollupUpdateWithdrawalDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupUpdateWithdrawalDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupUpdateWithdrawalDelay represents a UpdateWithdrawalDelay event raised by the Hezrollup contract.
type HezrollupUpdateWithdrawalDelay struct {
	NewWithdrawalDelay uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawalDelay is a free log retrieval operation binding the contract event 0x9db302c4547a21fb20a3a794e5f63ee87eb6e4afc3325ebdadba2d1fb4a90737.
//
// Solidity: event UpdateWithdrawalDelay(uint64 newWithdrawalDelay)
func (_Hezrollup *HezrollupFilterer) FilterUpdateWithdrawalDelay(opts *bind.FilterOpts) (*HezrollupUpdateWithdrawalDelayIterator, error) {

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "UpdateWithdrawalDelay")
	if err != nil {
		return nil, err
	}
	return &HezrollupUpdateWithdrawalDelayIterator{contract: _Hezrollup.contract, event: "UpdateWithdrawalDelay", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawalDelay is a free log subscription operation binding the contract event 0x9db302c4547a21fb20a3a794e5f63ee87eb6e4afc3325ebdadba2d1fb4a90737.
//
// Solidity: event UpdateWithdrawalDelay(uint64 newWithdrawalDelay)
func (_Hezrollup *HezrollupFilterer) WatchUpdateWithdrawalDelay(opts *bind.WatchOpts, sink chan<- *HezrollupUpdateWithdrawalDelay) (event.Subscription, error) {

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "UpdateWithdrawalDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupUpdateWithdrawalDelay)
				if err := _Hezrollup.contract.UnpackLog(event, "UpdateWithdrawalDelay", log); err != nil {
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

// ParseUpdateWithdrawalDelay is a log parse operation binding the contract event 0x9db302c4547a21fb20a3a794e5f63ee87eb6e4afc3325ebdadba2d1fb4a90737.
//
// Solidity: event UpdateWithdrawalDelay(uint64 newWithdrawalDelay)
func (_Hezrollup *HezrollupFilterer) ParseUpdateWithdrawalDelay(log types.Log) (*HezrollupUpdateWithdrawalDelay, error) {
	event := new(HezrollupUpdateWithdrawalDelay)
	if err := _Hezrollup.contract.UnpackLog(event, "UpdateWithdrawalDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the Hezrollup contract.
type HezrollupWithdrawEventIterator struct {
	Event *HezrollupWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *HezrollupWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupWithdrawEvent)
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
		it.Event = new(HezrollupWithdrawEvent)
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
func (it *HezrollupWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupWithdrawEvent represents a WithdrawEvent event raised by the Hezrollup contract.
type HezrollupWithdrawEvent struct {
	Idx             *big.Int
	NumExitRoot     uint32
	InstantWithdraw bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x69177d798b38e27bcc4e0338307e4f1490e12d1006729d0e6e9cc82a8732f415.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot, bool indexed instantWithdraw)
func (_Hezrollup *HezrollupFilterer) FilterWithdrawEvent(opts *bind.FilterOpts, idx []*big.Int, numExitRoot []uint32, instantWithdraw []bool) (*HezrollupWithdrawEventIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}
	var numExitRootRule []interface{}
	for _, numExitRootItem := range numExitRoot {
		numExitRootRule = append(numExitRootRule, numExitRootItem)
	}
	var instantWithdrawRule []interface{}
	for _, instantWithdrawItem := range instantWithdraw {
		instantWithdrawRule = append(instantWithdrawRule, instantWithdrawItem)
	}

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "WithdrawEvent", idxRule, numExitRootRule, instantWithdrawRule)
	if err != nil {
		return nil, err
	}
	return &HezrollupWithdrawEventIterator{contract: _Hezrollup.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x69177d798b38e27bcc4e0338307e4f1490e12d1006729d0e6e9cc82a8732f415.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot, bool indexed instantWithdraw)
func (_Hezrollup *HezrollupFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *HezrollupWithdrawEvent, idx []*big.Int, numExitRoot []uint32, instantWithdraw []bool) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}
	var numExitRootRule []interface{}
	for _, numExitRootItem := range numExitRoot {
		numExitRootRule = append(numExitRootRule, numExitRootItem)
	}
	var instantWithdrawRule []interface{}
	for _, instantWithdrawItem := range instantWithdraw {
		instantWithdrawRule = append(instantWithdrawRule, instantWithdrawItem)
	}

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "WithdrawEvent", idxRule, numExitRootRule, instantWithdrawRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupWithdrawEvent)
				if err := _Hezrollup.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
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

// ParseWithdrawEvent is a log parse operation binding the contract event 0x69177d798b38e27bcc4e0338307e4f1490e12d1006729d0e6e9cc82a8732f415.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot, bool indexed instantWithdraw)
func (_Hezrollup *HezrollupFilterer) ParseWithdrawEvent(log types.Log) (*HezrollupWithdrawEvent, error) {
	event := new(HezrollupWithdrawEvent)
	if err := _Hezrollup.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HezrollupHermezV2Iterator is returned from FilterHermezV2 and is used to iterate over the raw logs and unpacked data for HermezV2 events raised by the Hezrollup contract.
type HezrollupHermezV2Iterator struct {
	Event *HezrollupHermezV2 // Event containing the contract specifics and raw log

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
func (it *HezrollupHermezV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HezrollupHermezV2)
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
		it.Event = new(HezrollupHermezV2)
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
func (it *HezrollupHermezV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HezrollupHermezV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HezrollupHermezV2 represents a HermezV2 event raised by the Hezrollup contract.
type HezrollupHermezV2 struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHermezV2 is a free log retrieval operation binding the contract event 0xd5303fa2e7ece2a0fe77fbba1df5bb224b461198dd7bfd7fe0071f964c86c673.
//
// Solidity: event hermezV2()
func (_Hezrollup *HezrollupFilterer) FilterHermezV2(opts *bind.FilterOpts) (*HezrollupHermezV2Iterator, error) {

	logs, sub, err := _Hezrollup.contract.FilterLogs(opts, "hermezV2")
	if err != nil {
		return nil, err
	}
	return &HezrollupHermezV2Iterator{contract: _Hezrollup.contract, event: "hermezV2", logs: logs, sub: sub}, nil
}

// WatchHermezV2 is a free log subscription operation binding the contract event 0xd5303fa2e7ece2a0fe77fbba1df5bb224b461198dd7bfd7fe0071f964c86c673.
//
// Solidity: event hermezV2()
func (_Hezrollup *HezrollupFilterer) WatchHermezV2(opts *bind.WatchOpts, sink chan<- *HezrollupHermezV2) (event.Subscription, error) {

	logs, sub, err := _Hezrollup.contract.WatchLogs(opts, "hermezV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HezrollupHermezV2)
				if err := _Hezrollup.contract.UnpackLog(event, "hermezV2", log); err != nil {
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

// ParseHermezV2 is a log parse operation binding the contract event 0xd5303fa2e7ece2a0fe77fbba1df5bb224b461198dd7bfd7fe0071f964c86c673.
//
// Solidity: event hermezV2()
func (_Hezrollup *HezrollupFilterer) ParseHermezV2(log types.Log) (*HezrollupHermezV2, error) {
	event := new(HezrollupHermezV2)
	if err := _Hezrollup.contract.UnpackLog(event, "hermezV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
