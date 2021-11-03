//// Code generated - DO NOT EDIT.
//// This file is a generated binding and any manual changes will be lost.
//
package store

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x608060405234801561001057600080fd5b5060405161072138038061072183398181016040528101906100329190610162565b806000908051906020019061004892919061004f565b505061031a565b82805461005b90610234565b90600052602060002090601f01602090048101928261007d57600085556100c4565b82601f1061009657805160ff19168380011785556100c4565b828001600101855582156100c4579182015b828111156100c35782518255916020019190600101906100a8565b5b5090506100d191906100d5565b5090565b5b808211156100ee5760008160009055506001016100d6565b5090565b6000610105610100846101d0565b6101ab565b905082815260208101848484011115610121576101206102fa565b5b61012c848285610201565b509392505050565b600082601f830112610149576101486102f5565b5b81516101598482602086016100f2565b91505092915050565b60006020828403121561017857610177610304565b5b600082015167ffffffffffffffff811115610196576101956102ff565b5b6101a284828501610134565b91505092915050565b60006101b56101c6565b90506101c18282610266565b919050565b6000604051905090565b600067ffffffffffffffff8211156101eb576101ea6102c6565b5b6101f482610309565b9050602081019050919050565b60005b8381101561021f578082015181840152602081019050610204565b8381111561022e576000848401525b50505050565b6000600282049050600182168061024c57607f821691505b602082108114156102605761025f610297565b5b50919050565b61026f82610309565b810181811067ffffffffffffffff8211171561028e5761028d6102c6565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6103f8806103296000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806348f343f31461004657806354fd4d5014610076578063f56256c714610094575b600080fd5b610060600480360381019061005b91906101c0565b6100b0565b60405161006d9190610275565b60405180910390f35b61007e6100c8565b60405161008b91906102b9565b60405180910390f35b6100ae60048036038101906100a991906101ed565b610156565b005b60016020528060005260406000206000915090505481565b600080546100d590610334565b80601f016020809104026020016040519081016040528092919081815260200182805461010190610334565b801561014e5780601f106101235761010080835404028352916020019161014e565b820191906000526020600020905b81548152906001019060200180831161013157829003601f168201915b505050505081565b8060016000848152602001908152602001600020819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4828260405161019f929190610290565b60405180910390a15050565b6000813590506101ba816103ab565b92915050565b6000602082840312156101d6576101d5610395565b5b60006101e4848285016101ab565b91505092915050565b6000806040838503121561020457610203610395565b5b6000610212858286016101ab565b9250506020610223858286016101ab565b9150509250929050565b610236816102f7565b82525050565b6000610247826102db565b61025181856102e6565b9350610261818560208601610301565b61026a8161039a565b840191505092915050565b600060208201905061028a600083018461022d565b92915050565b60006040820190506102a5600083018561022d565b6102b2602083018461022d565b9392505050565b600060208201905081810360008301526102d3818461023c565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050919050565b60005b8381101561031f578082015181840152602081019050610304565b8381111561032e576000848401525b50505050565b6000600282049050600182168061034c57607f821691505b602082108114156103605761035f610366565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b6103b4816102f7565b81146103bf57600080fd5b5056fea264697066735822122005a5951e5067dae6b917ec5a4c3ccb95f0179d7371d220df636bc264a789ee6e64736f6c63430008060033"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreCaller) Items(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreCallerSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreCallerSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreTransactor) SetItem(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreTransactorSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// StoreItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Store contract.
type StoreItemSetIterator struct {
	Event *StoreItemSet // Event containing the contract specifics and raw log

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
func (it *StoreItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreItemSet)
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
		it.Event = new(StoreItemSet)
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
func (it *StoreItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreItemSet represents a ItemSet event raised by the Store contract.
type StoreItemSet struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) FilterItemSet(opts *bind.FilterOpts) (*StoreItemSetIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &StoreItemSetIterator{contract: _Store.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *StoreItemSet) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreItemSet)
				if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) ParseItemSet(log types.Log) (*StoreItemSet, error) {
	event := new(StoreItemSet)
	if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
