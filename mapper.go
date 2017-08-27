// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"secondary\",\"type\":\"address\"}],\"name\":\"mapAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"secondaryInUse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"primaryToSecondary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"primary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"secondary\",\"type\":\"address\"}],\"name\":\"AddressMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"code\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Error\",\"type\":\"event\"}]"

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// PrimaryToSecondary is a free data retrieval call binding the contract method 0xd352221e.
//
// Solidity: function primaryToSecondary( address) constant returns(address)
func (_Main *MainCaller) PrimaryToSecondary(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "primaryToSecondary", arg0)
	return *ret0, err
}

// PrimaryToSecondary is a free data retrieval call binding the contract method 0xd352221e.
//
// Solidity: function primaryToSecondary( address) constant returns(address)
func (_Main *MainSession) PrimaryToSecondary(arg0 common.Address) (common.Address, error) {
	return _Main.Contract.PrimaryToSecondary(&_Main.CallOpts, arg0)
}

// PrimaryToSecondary is a free data retrieval call binding the contract method 0xd352221e.
//
// Solidity: function primaryToSecondary( address) constant returns(address)
func (_Main *MainCallerSession) PrimaryToSecondary(arg0 common.Address) (common.Address, error) {
	return _Main.Contract.PrimaryToSecondary(&_Main.CallOpts, arg0)
}

// SecondaryInUse is a free data retrieval call binding the contract method 0x6a50654a.
//
// Solidity: function secondaryInUse( address) constant returns(bool)
func (_Main *MainCaller) SecondaryInUse(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "secondaryInUse", arg0)
	return *ret0, err
}

// SecondaryInUse is a free data retrieval call binding the contract method 0x6a50654a.
//
// Solidity: function secondaryInUse( address) constant returns(bool)
func (_Main *MainSession) SecondaryInUse(arg0 common.Address) (bool, error) {
	return _Main.Contract.SecondaryInUse(&_Main.CallOpts, arg0)
}

// SecondaryInUse is a free data retrieval call binding the contract method 0x6a50654a.
//
// Solidity: function secondaryInUse( address) constant returns(bool)
func (_Main *MainCallerSession) SecondaryInUse(arg0 common.Address) (bool, error) {
	return _Main.Contract.SecondaryInUse(&_Main.CallOpts, arg0)
}

// MapAddress is a paid mutator transaction binding the contract method 0x4a270f47.
//
// Solidity: function mapAddress(secondary address) returns()
func (_Main *MainTransactor) MapAddress(opts *bind.TransactOpts, secondary common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "mapAddress", secondary)
}

// MapAddress is a paid mutator transaction binding the contract method 0x4a270f47.
//
// Solidity: function mapAddress(secondary address) returns()
func (_Main *MainSession) MapAddress(secondary common.Address) (*types.Transaction, error) {
	return _Main.Contract.MapAddress(&_Main.TransactOpts, secondary)
}

// MapAddress is a paid mutator transaction binding the contract method 0x4a270f47.
//
// Solidity: function mapAddress(secondary address) returns()
func (_Main *MainTransactorSession) MapAddress(secondary common.Address) (*types.Transaction, error) {
	return _Main.Contract.MapAddress(&_Main.TransactOpts, secondary)
}
