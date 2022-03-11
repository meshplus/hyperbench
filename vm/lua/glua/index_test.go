package glua

import (
	idex "github.com/meshplus/hyperbench/plugins/index"
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
	"testing"
)

func Test_index(t *testing.T) {
	L := lua.NewState()
	defer L.Close()
	mt := L.NewTypeMetatable("case")
	L.SetGlobal("case", mt)
	passIdx := &idex.Index{1, 1, 1, 1}

	cLua := newIdexIndex(L, passIdx)
	passIdx.Tx = 2
	L.SetField(mt, "index", cLua)
	scripts := []string{`
		function run()
			print("tx:",case.index.Tx)
            case.index.Worker=2
            case.index.VM=2
            case.index.Engine=2
            case.index.Tx=2
			return case.index
		end
	`}
	for _, script := range scripts {
		lvalue, err := runLuaRunFunc(L, script)
		assert.Nil(t, err)
		idx := &idex.Index{}
		l1, err := Lua2Go(lvalue)
		idx, ok := l1.(*idex.Index)
		assert.True(t, ok)
		assert.Nil(t, err)
		assert.Equal(t, idx, &idex.Index{Worker: 2, VM: 2, Engine: 2, Tx: 2})
	}

}
