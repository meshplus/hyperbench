package glua

import (
	"github.com/meshplus/hyperbench/plugins/toolkit"
	lua "github.com/yuin/gopher-lua"
)

func hexLuaFunction(L *lua.LState, kit *toolkit.ToolKit) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		firstArgIdx := 1
		if checkToolKitByIdx(state, 1) {
			firstArgIdx++
		}

		input := state.CheckString(firstArgIdx)
		ret := kit.Hex(input)
		state.Push(lua.LString(ret))
		return 1
	})
}

func randStrLuaFunction(L *lua.LState, kit *toolkit.ToolKit) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		firstArgIdx := 1
		if checkToolKitByIdx(state, 1) {
			firstArgIdx++
		}
		size := state.CheckInt(firstArgIdx)
		ret := kit.RandStr(uint(size))
		L.Push(lua.LString(ret))
		return 1
	})
}

func randIntLuaFunction(L *lua.LState, kit *toolkit.ToolKit) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		firstArgIdx := 1
		if checkToolKitByIdx(state, 1) {
			firstArgIdx++
		}
		min := state.CheckInt(firstArgIdx)
		max := state.CheckInt(firstArgIdx + 1)
		ret := kit.RandInt(min, max)
		L.Push(lua.LNumber(ret))
		return 1
	})
}

func stringLuaFunction(L *lua.LState, kit *toolkit.ToolKit) *lua.LFunction {
	return L.NewFunction(func(state *lua.LState) int {
		firstArgIdx := 1
		if checkToolKitByIdx(state, 1) {
			firstArgIdx++
		}

		argLength := state.GetTop()
		if argLength < 1 {
			panic("args are less than 1")
		}
		input := state.CheckAny(firstArgIdx)
		inputGo, err := Lua2Go(input)
		if err != nil {
			panic(err.Error())
		}
		if argLength == firstArgIdx {
			ret := kit.String(inputGo)
			L.Push(lua.LString(ret))
			return 1
		}
		var offsets []int
		for i := firstArgIdx + 1; i < argLength; i++ {
			offset := state.CheckInt(i)
			offsets = append(offsets, offset)
		}
		ret := kit.String(inputGo, offsets...)
		L.Push(lua.LString(ret))
		return 1
	})
}

func newToolKit(L *lua.LState, kit *toolkit.ToolKit) lua.LValue {
	toolkitTable := L.NewTable()
	toolkitTable.RawSetString("Hex", hexLuaFunction(L, kit))
	toolkitTable.RawSetString("RandStr", randStrLuaFunction(L, kit))
	toolkitTable.RawSetString("RandInt", randIntLuaFunction(L, kit))
	toolkitTable.RawSetString("String", stringLuaFunction(L, kit))
	return toolkitTable
}

func checkToolKitByIdx(state *lua.LState, idx int) bool {
	if state.GetTop() < idx {
		return false
	}
	idxValue := state.CheckAny(idx)
	_, ok := idxValue.(*lua.LTable)
	if !ok {
		return false
	}
	return true
}
