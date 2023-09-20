package glua

import (
	"fmt"

	fcom "github.com/meshplus/hyperbench-common/common"
	lua "github.com/yuin/gopher-lua"
)

func newBlockchain(L *lua.LState, client fcom.Blockchain) lua.LValue {
	clientTable := L.NewTable()
	clientTable.RawSetString("DeployContract", deployContractLuaFunction(L, client))
	clientTable.RawSetString("Invoke", invokeLuaFunction(L, client))
	clientTable.RawSetString("Transfer", transferLuaFunction(L, client))
	clientTable.RawSetString("Confirm", confirmLuaFunction(L, client))
	clientTable.RawSetString("Query", queryLuaFunction(L, client))
	clientTable.RawSetString("Option", optionLuaFunction(L, client))
	clientTable.RawSetString("GetContext", getContextLuaFunction(L, client))
	clientTable.RawSetString("SetContext", setContextLuaFunction(L, client))
	clientTable.RawSetString("ResetContext", resetContextLuaFunction(L, client))
	clientTable.RawSetString("GetRandomAccountByGroup", getRandomAccountByGroupLuaFunction(L, client))
	clientTable.RawSetString("GetRandomAccount", getRandomAccountLuaFunction(L, client))
	clientTable.RawSetString("GetAccount", getAccountLuaFunction(L, client))
	clientTable.RawSetString("GetContractAddrByName", getContractAddrByNameLuaFunction(L, client))
	//clientTable.RawSetString("Statistic",nil)
	return clientTable
}

func setContextLuaFunction(L *lua.LState, client fcom.Blockchain) lua.LValue {
	return L.NewFunction(func(state *lua.LState) int {
		firstArgIndex := 1
		// check first arg is fcom.Blockchain
		if checkBlockChainByIdx(state, 1) {
			firstArgIndex++
		}
		text := state.CheckString(firstArgIndex)
		err := client.SetContext(text)
		if err != nil {
			state.Push(lua.LString(err.Error()))
			return 1
		}
		state.Push(lua.LString(""))
		return 1
	})
}

func getContextLuaFunction(L *lua.LState, client fcom.Blockchain) lua.LValue {
	return L.NewFunction(func(state *lua.LState) int {
		text, err := client.GetContext()
		if err != nil {
			state.Push(lua.LString(err.Error()))
			return 1
		}
		state.Push(lua.LString(text))
		return 1
	})
}

func resetContextLuaFunction(L *lua.LState, client fcom.Blockchain) lua.LValue {
	return L.NewFunction(func(state *lua.LState) int {
		err := client.ResetContext()
		if err != nil {
			state.Push(lua.LString(err.Error()))
			return 1
		}
		state.Push(lua.LString(""))
		return 1
	})
}

func getRandomAccountByGroupLuaFunction(L *lua.LState, client fcom.Blockchain) lua.LValue {
	return L.NewFunction(func(state *lua.LState) int {
		account := client.GetRandomAccountByGroup()
		state.Push(lua.LString(account))
		return 1
	})
}

func getRandomAccountLuaFunction(L *lua.LState, client fcom.Blockchain) lua.LValue {
	return L.NewFunction(func(state *lua.LState) int {
		firstArgIndex := 1
		// check first arg is fcom.Blockchain
		if checkBlockChainByIdx(state, 1) {
			firstArgIndex++
		}
		text := state.CheckString(firstArgIndex)
		account := client.GetRandomAccount(text)
		state.Push(lua.LString(account))
		return 1
	})
}

func getAccountLuaFunction(L *lua.LState, client fcom.Blockchain) lua.LValue {
	return L.NewFunction(func(state *lua.LState) int {
		firstArgIndex := 1
		// check first arg is fcom.Blockchain
		if checkBlockChainByIdx(state, 1) {
			firstArgIndex++
		}
		index := state.CheckInt64(firstArgIndex)
		account := client.GetAccount(uint64(index))
		state.Push(lua.LString(account))
		return 1
	})
}

func optionLuaFunction(L *lua.LState, client fcom.Blockchain) lua.LValue {
	return L.NewFunction(func(state *lua.LState) int {
		var map1 fcom.Option
		// case.blockchain:Invoke() --> first arg is fcom.Blockchain
		// case.blockchain.Invoke  ----> first arg is normal
		firstArgIndex := 1
		// check first arg is fcom.Blockchain
		if checkBlockChainByIdx(state, 1) {
			firstArgIndex++
		}
		invokeTable := state.CheckTable(firstArgIndex)
		err := TableLua2GoStruct(invokeTable, &map1)
		if err != nil {
			state.ArgError(1, "common.Option expected")
		}
		err = client.Option(map1)
		if err != nil {
			state.Push(lua.LString(err.Error()))
		}
		state.Push(lua.LString(""))
		return 1
	})
}

func invokeLuaFunction(L *lua.LState, client fcom.Blockchain) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		var invoke fcom.Invoke
		res := blockchainLuaFunction(L, client, invoke, func(b fcom.Blockchain, b2 interface{}, option ...fcom.Option) interface{} {
			return b.Invoke(b2.(fcom.Invoke), option...)
		})
		return res
	})
}

func transferLuaFunction(L *lua.LState, client fcom.Blockchain) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		var transfer fcom.Transfer
		res := blockchainLuaFunction(L, client, transfer, func(b fcom.Blockchain, b2 interface{}, option ...fcom.Option) interface{} {
			return b.Transfer(b2.(fcom.Transfer), option...)
		})
		return res
	})
}

func queryLuaFunction(L *lua.LState, client fcom.Blockchain) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		var query fcom.Query
		res := blockchainLuaFunction(L, client, query, func(b fcom.Blockchain, b2 interface{}, option ...fcom.Option) interface{} {
			return b.Query(b2.(fcom.Query), option...)
		})
		return res
	})
}

func confirmLuaFunction(L *lua.LState, client fcom.Blockchain) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		var confirm *fcom.Result
		res := blockchainLuaFunction(L, client, confirm, func(b fcom.Blockchain, b2 interface{}, option ...fcom.Option) interface{} {
			return b.Confirm(b2.(*fcom.Result), option...)
		})
		return res
	})
}

func blockchainLuaFunction(state *lua.LState, cli fcom.Blockchain, arg1Type interface{}, fn func(fcom.Blockchain, interface{}, ...fcom.Option) interface{}) int {
	firstArgIndex := 1
	// check first arg is fcom.Blockchain
	if checkBlockChainByIdx(state, 1) {
		firstArgIndex++
	}
	invokeTable := state.CheckTable(firstArgIndex)
	err := TableLua2GoStruct(invokeTable, &arg1Type)
	if err != nil {
		state.ArgError(1, fmt.Sprintf("interface. expected, %s", err))
	}
	if state.GetTop() == 1+firstArgIndex {
		ret := fn(cli, arg1Type)
		state.Push(go2Lua(state, ret))
		return 1
	}
	var opts []fcom.Option
	for i := 1 + firstArgIndex; i <= state.GetTop(); i++ {
		table := state.CheckTable(i)
		var map1 fcom.Option
		err := TableLua2GoStruct(table, &map1)
		if err != nil {
			state.ArgError(1, "common.Option expected")
		}
		opts = append(opts, map1)
	}
	ret := fn(cli, arg1Type, opts...)
	state.Push(go2Lua(state, ret))
	return 1
}

func checkBlockChainByIdx(state *lua.LState, idx int) bool {
	if state.GetTop() < idx {
		return false
	}
	idxValue := state.CheckAny(idx)
	lvalue, ok := idxValue.(*lua.LTable)
	if !ok {
		return false
	}
	k, _ := lvalue.Next(lua.LString("Invok"))
	if k.String() != "Invoke" {
		return false
	}
	return true
}

func deployContractLuaFunction(L *lua.LState, client fcom.Blockchain) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		err := client.DeployContract()
		if err != nil {
			state.Push(lua.LString(err.Error()))
			return 1
		}
		state.Push(lua.LString(""))
		return 1
	})
}

func getContractAddrByNameLuaFunction(L *lua.LState, client fcom.Blockchain) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		firstArgIndex := 1
		if checkBlockChainByIdx(state, 1) {
			firstArgIndex++
		}
		text := state.CheckString(firstArgIndex)
		addr := client.GetRandomAccount(text)
		state.Push(lua.LString(addr))
		return 1
	})
}
