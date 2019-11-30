// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v1

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/accounts/abi"
	"github.com/FISCO-BCOS/go-sdk/accounts/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/common"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = common.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// V1ABI is the input ABI used to generate the binding from.
const V1ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Task2Previous\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TaskActivated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Terminated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Task2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Task1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Terminate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Task1Next\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Task1Previous\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Task2Next\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Task1Finished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Task2Finished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ProcessFinished\",\"type\":\"event\"}]"

// V1Bin is the compiled bytecode used for deploying new contracts.
var V1Bin = "0x60806040526000600160006101000a81548160ff02191690831515021790555034801561002b57600080fd5b5061065e8061003b6000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063185cde79146100a9578063502ad655146100ea57806356a9f39a1461012f57806357a86f7d1461015e57806371e54f81146101755780638cc1eb001461018c5780639445eb3a146101a3578063b5ab9956146101ba578063ddf06b8c146101fb578063f749731a1461023c575b600080fd5b3480156100b557600080fd5b506100d46004803603810190808035906020019092919050505061027d565b6040518082815260200191505060405180910390f35b3480156100f657600080fd5b50610115600480360381019080803590602001909291905050506102a0565b604051808215151515815260200191505060405180910390f35b34801561013b57600080fd5b506101446102c9565b604051808215151515815260200191505060405180910390f35b34801561016a57600080fd5b506101736102dc565b005b34801561018157600080fd5b5061018a61033e565b005b34801561019857600080fd5b506101a1610436565b005b3480156101af57600080fd5b506101b861056e565b005b3480156101c657600080fd5b506101e5600480360381019080803590602001909291905050506105c9565b6040518082815260200191505060405180910390f35b34801561020757600080fd5b50610226600480360381019080803590602001909291905050506105ec565b6040518082815260200191505060405180910390f35b34801561024857600080fd5b506102676004803603810190808035906020019092919050505061060f565b6040518082815260200191505060405180910390f35b60038181548110151561028c57fe5b906000526020600020016000915090505481565b6000816002811015156102af57fe5b60209182820401919006915054906101000a900460ff1681565b600160009054906101000a900460ff1681565b60016000806002811015156102ed57fe5b602091828204019190066101000a81548160ff021916908315150217905550600080600160028110151561031d57fe5b602091828204019190066101000a81548160ff021916908315150217905550565b6000600115156000600160028110151561035457fe5b602091828204019190069054906101000a900460ff16151514151561037857600080fd5b600090505b6003805490508110156103e257600080600160038481548110151561039e57fe5b9060005260206000200154036002811015156103b657fe5b602091828204019190066101000a81548160ff021916908315150217905550808060010191505061037d565b7f82cc44c0ecef1b0b842308a49ac8534cea1aeff84119a06d81750f83907f0a67426040518082815260200191505060405180910390a160018060006101000a81548160ff02191690831515021790555050565b60006001151560008060028110151561044b57fe5b602091828204019190069054906101000a900460ff16151514151561046f57600080fd5b600460029080600181540180825580915050906001820390600052602060002001600090919290919091505550600360019080600181540180825580915050906001820390600052602060002001600090919290919091505550600090505b600480549050811015610534576001600060016004848154811015156104f057fe5b90600052602060002001540360028110151561050857fe5b602091828204019190066101000a81548160ff02191690831515021790555080806001019150506104ce565b7f235cc65f7a69f382cb90f6741cb943ebed2a7eb35095eda55bb6c9ff27a98092426040518082815260200191505060405180910390a150565b60011515600160009054906101000a900460ff16151514151561059057600080fd5b7f9784791c8cbd375f82799dcf406f4f2d5292fb9df9d1bc20d38fe3d9d80a1ff6426040518082815260200191505060405180910390a1565b6004818154811015156105d857fe5b906000526020600020016000915090505481565b6002818154811015156105fb57fe5b906000526020600020016000915090505481565b60058181548110151561061e57fe5b9060005260206000200160009150905054815600a165627a7a7230582047c8af2911746868848054793eac96acf2f6588945a18cbf3b78b4e006bd56ff0029"

// DeployV1 deploys a new Ethereum contract, binding an instance of V1 to it.
func DeployV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.RawTransaction, *V1, error) {
	parsed, err := abi.JSON(strings.NewReader(V1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(V1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &V1{V1Caller: V1Caller{contract: contract}, V1Transactor: V1Transactor{contract: contract}, V1Filterer: V1Filterer{contract: contract}}, nil
}

// V1 is an auto generated Go binding around an Ethereum contract.
type V1 struct {
	V1Caller     // Read-only binding to the contract
	V1Transactor // Write-only binding to the contract
	V1Filterer   // Log filterer for contract events
}

// V1Caller is an auto generated read-only Go binding around an Ethereum contract.
type V1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type V1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V1Session struct {
	Contract     *V1               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V1CallerSession struct {
	Contract *V1Caller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// V1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V1TransactorSession struct {
	Contract     *V1Transactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1Raw is an auto generated low-level Go binding around an Ethereum contract.
type V1Raw struct {
	Contract *V1 // Generic contract binding to access the raw methods on
}

// V1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V1CallerRaw struct {
	Contract *V1Caller // Generic read-only contract binding to access the raw methods on
}

// V1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V1TransactorRaw struct {
	Contract *V1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewV1 creates a new instance of V1, bound to a specific deployed contract.
func NewV1(address common.Address, backend bind.ContractBackend) (*V1, error) {
	contract, err := bindV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V1{V1Caller: V1Caller{contract: contract}, V1Transactor: V1Transactor{contract: contract}, V1Filterer: V1Filterer{contract: contract}}, nil
}

// NewV1Caller creates a new read-only instance of V1, bound to a specific deployed contract.
func NewV1Caller(address common.Address, caller bind.ContractCaller) (*V1Caller, error) {
	contract, err := bindV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V1Caller{contract: contract}, nil
}

// NewV1Transactor creates a new write-only instance of V1, bound to a specific deployed contract.
func NewV1Transactor(address common.Address, transactor bind.ContractTransactor) (*V1Transactor, error) {
	contract, err := bindV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V1Transactor{contract: contract}, nil
}

// NewV1Filterer creates a new log filterer instance of V1, bound to a specific deployed contract.
func NewV1Filterer(address common.Address, filterer bind.ContractFilterer) (*V1Filterer, error) {
	contract, err := bindV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V1Filterer{contract: contract}, nil
}

// bindV1 binds a generic wrapper to an already deployed contract.
func bindV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1 *V1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _V1.Contract.V1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1 *V1Raw) Transfer(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _V1.Contract.V1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1 *V1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.RawTransaction, error) {
	return _V1.Contract.V1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1 *V1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _V1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1 *V1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _V1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1 *V1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.RawTransaction, error) {
	return _V1.Contract.contract.Transact(opts, method, params...)
}

// Task1Next is a free data retrieval call binding the contract method 0xb5ab9956.
//
// Solidity: function Task1Next(uint256 ) constant returns(uint256)
func (_V1 *V1Caller) Task1Next(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _V1.contract.Call(opts, out, "Task1Next", arg0)
	return *ret0, err
}

// Task1Next is a free data retrieval call binding the contract method 0xb5ab9956.
//
// Solidity: function Task1Next(uint256 ) constant returns(uint256)
func (_V1 *V1Session) Task1Next(arg0 *big.Int) (*big.Int, error) {
	return _V1.Contract.Task1Next(&_V1.CallOpts, arg0)
}

// Task1Next is a free data retrieval call binding the contract method 0xb5ab9956.
//
// Solidity: function Task1Next(uint256 ) constant returns(uint256)
func (_V1 *V1CallerSession) Task1Next(arg0 *big.Int) (*big.Int, error) {
	return _V1.Contract.Task1Next(&_V1.CallOpts, arg0)
}

// Task1Previous is a free data retrieval call binding the contract method 0xddf06b8c.
//
// Solidity: function Task1Previous(uint256 ) constant returns(uint256)
func (_V1 *V1Caller) Task1Previous(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _V1.contract.Call(opts, out, "Task1Previous", arg0)
	return *ret0, err
}

// Task1Previous is a free data retrieval call binding the contract method 0xddf06b8c.
//
// Solidity: function Task1Previous(uint256 ) constant returns(uint256)
func (_V1 *V1Session) Task1Previous(arg0 *big.Int) (*big.Int, error) {
	return _V1.Contract.Task1Previous(&_V1.CallOpts, arg0)
}

// Task1Previous is a free data retrieval call binding the contract method 0xddf06b8c.
//
// Solidity: function Task1Previous(uint256 ) constant returns(uint256)
func (_V1 *V1CallerSession) Task1Previous(arg0 *big.Int) (*big.Int, error) {
	return _V1.Contract.Task1Previous(&_V1.CallOpts, arg0)
}

// Task2Next is a free data retrieval call binding the contract method 0xf749731a.
//
// Solidity: function Task2Next(uint256 ) constant returns(uint256)
func (_V1 *V1Caller) Task2Next(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _V1.contract.Call(opts, out, "Task2Next", arg0)
	return *ret0, err
}

// Task2Next is a free data retrieval call binding the contract method 0xf749731a.
//
// Solidity: function Task2Next(uint256 ) constant returns(uint256)
func (_V1 *V1Session) Task2Next(arg0 *big.Int) (*big.Int, error) {
	return _V1.Contract.Task2Next(&_V1.CallOpts, arg0)
}

// Task2Next is a free data retrieval call binding the contract method 0xf749731a.
//
// Solidity: function Task2Next(uint256 ) constant returns(uint256)
func (_V1 *V1CallerSession) Task2Next(arg0 *big.Int) (*big.Int, error) {
	return _V1.Contract.Task2Next(&_V1.CallOpts, arg0)
}

// Task2Previous is a free data retrieval call binding the contract method 0x185cde79.
//
// Solidity: function Task2Previous(uint256 ) constant returns(uint256)
func (_V1 *V1Caller) Task2Previous(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _V1.contract.Call(opts, out, "Task2Previous", arg0)
	return *ret0, err
}

// Task2Previous is a free data retrieval call binding the contract method 0x185cde79.
//
// Solidity: function Task2Previous(uint256 ) constant returns(uint256)
func (_V1 *V1Session) Task2Previous(arg0 *big.Int) (*big.Int, error) {
	return _V1.Contract.Task2Previous(&_V1.CallOpts, arg0)
}

// Task2Previous is a free data retrieval call binding the contract method 0x185cde79.
//
// Solidity: function Task2Previous(uint256 ) constant returns(uint256)
func (_V1 *V1CallerSession) Task2Previous(arg0 *big.Int) (*big.Int, error) {
	return _V1.Contract.Task2Previous(&_V1.CallOpts, arg0)
}

// TaskActivated is a free data retrieval call binding the contract method 0x502ad655.
//
// Solidity: function TaskActivated(uint256 ) constant returns(bool)
func (_V1 *V1Caller) TaskActivated(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _V1.contract.Call(opts, out, "TaskActivated", arg0)
	return *ret0, err
}

// TaskActivated is a free data retrieval call binding the contract method 0x502ad655.
//
// Solidity: function TaskActivated(uint256 ) constant returns(bool)
func (_V1 *V1Session) TaskActivated(arg0 *big.Int) (bool, error) {
	return _V1.Contract.TaskActivated(&_V1.CallOpts, arg0)
}

// TaskActivated is a free data retrieval call binding the contract method 0x502ad655.
//
// Solidity: function TaskActivated(uint256 ) constant returns(bool)
func (_V1 *V1CallerSession) TaskActivated(arg0 *big.Int) (bool, error) {
	return _V1.Contract.TaskActivated(&_V1.CallOpts, arg0)
}

// Terminated is a free data retrieval call binding the contract method 0x56a9f39a.
//
// Solidity: function Terminated() constant returns(bool)
func (_V1 *V1Caller) Terminated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _V1.contract.Call(opts, out, "Terminated")
	return *ret0, err
}

// Terminated is a free data retrieval call binding the contract method 0x56a9f39a.
//
// Solidity: function Terminated() constant returns(bool)
func (_V1 *V1Session) Terminated() (bool, error) {
	return _V1.Contract.Terminated(&_V1.CallOpts)
}

// Terminated is a free data retrieval call binding the contract method 0x56a9f39a.
//
// Solidity: function Terminated() constant returns(bool)
func (_V1 *V1CallerSession) Terminated() (bool, error) {
	return _V1.Contract.Terminated(&_V1.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0x57a86f7d.
//
// Solidity: function Init() returns()
func (_V1 *V1Transactor) Init(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _V1.contract.Transact(opts, "Init")
}

// Init is a paid mutator transaction binding the contract method 0x57a86f7d.
//
// Solidity: function Init() returns()
func (_V1 *V1Session) Init() (*types.RawTransaction, error) {
	return _V1.Contract.Init(&_V1.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x57a86f7d.
//
// Solidity: function Init() returns()
func (_V1 *V1TransactorSession) Init() (*types.RawTransaction, error) {
	return _V1.Contract.Init(&_V1.TransactOpts)
}

// Task1 is a paid mutator transaction binding the contract method 0x8cc1eb00.
//
// Solidity: function Task1() returns()
func (_V1 *V1Transactor) Task1(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _V1.contract.Transact(opts, "Task1")
}

// Task1 is a paid mutator transaction binding the contract method 0x8cc1eb00.
//
// Solidity: function Task1() returns()
func (_V1 *V1Session) Task1() (*types.RawTransaction, error) {
	return _V1.Contract.Task1(&_V1.TransactOpts)
}

// Task1 is a paid mutator transaction binding the contract method 0x8cc1eb00.
//
// Solidity: function Task1() returns()
func (_V1 *V1TransactorSession) Task1() (*types.RawTransaction, error) {
	return _V1.Contract.Task1(&_V1.TransactOpts)
}

// Task2 is a paid mutator transaction binding the contract method 0x71e54f81.
//
// Solidity: function Task2() returns()
func (_V1 *V1Transactor) Task2(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _V1.contract.Transact(opts, "Task2")
}

// Task2 is a paid mutator transaction binding the contract method 0x71e54f81.
//
// Solidity: function Task2() returns()
func (_V1 *V1Session) Task2() (*types.RawTransaction, error) {
	return _V1.Contract.Task2(&_V1.TransactOpts)
}

// Task2 is a paid mutator transaction binding the contract method 0x71e54f81.
//
// Solidity: function Task2() returns()
func (_V1 *V1TransactorSession) Task2() (*types.RawTransaction, error) {
	return _V1.Contract.Task2(&_V1.TransactOpts)
}

// Terminate is a paid mutator transaction binding the contract method 0x9445eb3a.
//
// Solidity: function Terminate() returns()
func (_V1 *V1Transactor) Terminate(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _V1.contract.Transact(opts, "Terminate")
}

// Terminate is a paid mutator transaction binding the contract method 0x9445eb3a.
//
// Solidity: function Terminate() returns()
func (_V1 *V1Session) Terminate() (*types.RawTransaction, error) {
	return _V1.Contract.Terminate(&_V1.TransactOpts)
}

// Terminate is a paid mutator transaction binding the contract method 0x9445eb3a.
//
// Solidity: function Terminate() returns()
func (_V1 *V1TransactorSession) Terminate() (*types.RawTransaction, error) {
	return _V1.Contract.Terminate(&_V1.TransactOpts)
}

// V1ProcessFinishedIterator is returned from FilterProcessFinished and is used to iterate over the raw logs and unpacked data for ProcessFinished events raised by the V1 contract.
type V1ProcessFinishedIterator struct {
	Event *V1ProcessFinished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  common.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *V1ProcessFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1ProcessFinished)
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
		it.Event = new(V1ProcessFinished)
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
func (it *V1ProcessFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1ProcessFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1ProcessFinished represents a ProcessFinished event raised by the V1 contract.
type V1ProcessFinished struct {
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProcessFinished is a free log retrieval operation binding the contract event 0x9784791c8cbd375f82799dcf406f4f2d5292fb9df9d1bc20d38fe3d9d80a1ff6.
//
// Solidity: event ProcessFinished(uint256 time)
func (_V1 *V1Filterer) FilterProcessFinished(opts *bind.FilterOpts) (*V1ProcessFinishedIterator, error) {

	logs, sub, err := _V1.contract.FilterLogs(opts, "ProcessFinished")
	if err != nil {
		return nil, err
	}
	return &V1ProcessFinishedIterator{contract: _V1.contract, event: "ProcessFinished", logs: logs, sub: sub}, nil
}

// WatchProcessFinished is a free log subscription operation binding the contract event 0x9784791c8cbd375f82799dcf406f4f2d5292fb9df9d1bc20d38fe3d9d80a1ff6.
//
// Solidity: event ProcessFinished(uint256 time)
func (_V1 *V1Filterer) WatchProcessFinished(opts *bind.WatchOpts, sink chan<- *V1ProcessFinished) (event.Subscription, error) {

	logs, sub, err := _V1.contract.WatchLogs(opts, "ProcessFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1ProcessFinished)
				if err := _V1.contract.UnpackLog(event, "ProcessFinished", log); err != nil {
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

// ParseProcessFinished is a log parse operation binding the contract event 0x9784791c8cbd375f82799dcf406f4f2d5292fb9df9d1bc20d38fe3d9d80a1ff6.
//
// Solidity: event ProcessFinished(uint256 time)
func (_V1 *V1Filterer) ParseProcessFinished(log types.Log) (*V1ProcessFinished, error) {
	event := new(V1ProcessFinished)
	if err := _V1.contract.UnpackLog(event, "ProcessFinished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// V1Task1FinishedIterator is returned from FilterTask1Finished and is used to iterate over the raw logs and unpacked data for Task1Finished events raised by the V1 contract.
type V1Task1FinishedIterator struct {
	Event *V1Task1Finished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  common.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *V1Task1FinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1Task1Finished)
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
		it.Event = new(V1Task1Finished)
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
func (it *V1Task1FinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1Task1FinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1Task1Finished represents a Task1Finished event raised by the V1 contract.
type V1Task1Finished struct {
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTask1Finished is a free log retrieval operation binding the contract event 0x235cc65f7a69f382cb90f6741cb943ebed2a7eb35095eda55bb6c9ff27a98092.
//
// Solidity: event Task1Finished(uint256 time)
func (_V1 *V1Filterer) FilterTask1Finished(opts *bind.FilterOpts) (*V1Task1FinishedIterator, error) {

	logs, sub, err := _V1.contract.FilterLogs(opts, "Task1Finished")
	if err != nil {
		return nil, err
	}
	return &V1Task1FinishedIterator{contract: _V1.contract, event: "Task1Finished", logs: logs, sub: sub}, nil
}

// WatchTask1Finished is a free log subscription operation binding the contract event 0x235cc65f7a69f382cb90f6741cb943ebed2a7eb35095eda55bb6c9ff27a98092.
//
// Solidity: event Task1Finished(uint256 time)
func (_V1 *V1Filterer) WatchTask1Finished(opts *bind.WatchOpts, sink chan<- *V1Task1Finished) (event.Subscription, error) {

	logs, sub, err := _V1.contract.WatchLogs(opts, "Task1Finished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1Task1Finished)
				if err := _V1.contract.UnpackLog(event, "Task1Finished", log); err != nil {
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

// ParseTask1Finished is a log parse operation binding the contract event 0x235cc65f7a69f382cb90f6741cb943ebed2a7eb35095eda55bb6c9ff27a98092.
//
// Solidity: event Task1Finished(uint256 time)
func (_V1 *V1Filterer) ParseTask1Finished(log types.Log) (*V1Task1Finished, error) {
	event := new(V1Task1Finished)
	if err := _V1.contract.UnpackLog(event, "Task1Finished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// V1Task2FinishedIterator is returned from FilterTask2Finished and is used to iterate over the raw logs and unpacked data for Task2Finished events raised by the V1 contract.
type V1Task2FinishedIterator struct {
	Event *V1Task2Finished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  common.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *V1Task2FinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1Task2Finished)
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
		it.Event = new(V1Task2Finished)
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
func (it *V1Task2FinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1Task2FinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1Task2Finished represents a Task2Finished event raised by the V1 contract.
type V1Task2Finished struct {
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTask2Finished is a free log retrieval operation binding the contract event 0x82cc44c0ecef1b0b842308a49ac8534cea1aeff84119a06d81750f83907f0a67.
//
// Solidity: event Task2Finished(uint256 time)
func (_V1 *V1Filterer) FilterTask2Finished(opts *bind.FilterOpts) (*V1Task2FinishedIterator, error) {

	logs, sub, err := _V1.contract.FilterLogs(opts, "Task2Finished")
	if err != nil {
		return nil, err
	}
	return &V1Task2FinishedIterator{contract: _V1.contract, event: "Task2Finished", logs: logs, sub: sub}, nil
}

// WatchTask2Finished is a free log subscription operation binding the contract event 0x82cc44c0ecef1b0b842308a49ac8534cea1aeff84119a06d81750f83907f0a67.
//
// Solidity: event Task2Finished(uint256 time)
func (_V1 *V1Filterer) WatchTask2Finished(opts *bind.WatchOpts, sink chan<- *V1Task2Finished) (event.Subscription, error) {

	logs, sub, err := _V1.contract.WatchLogs(opts, "Task2Finished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1Task2Finished)
				if err := _V1.contract.UnpackLog(event, "Task2Finished", log); err != nil {
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

// ParseTask2Finished is a log parse operation binding the contract event 0x82cc44c0ecef1b0b842308a49ac8534cea1aeff84119a06d81750f83907f0a67.
//
// Solidity: event Task2Finished(uint256 time)
func (_V1 *V1Filterer) ParseTask2Finished(log types.Log) (*V1Task2Finished, error) {
	event := new(V1Task2Finished)
	if err := _V1.contract.UnpackLog(event, "Task2Finished", log); err != nil {
		return nil, err
	}
	return event, nil
}
