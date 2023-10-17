package glua

import (
	"testing"

	"github.com/meshplus/hyperbench-common/common"
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
)

func Test_Go2Lua(t *testing.T) {

	t.Run("demo", func(t *testing.T) {
		str := "demo"
		L := lua.NewState()
		strLua := Go2Lua(L, str)
		assert.Equal(t, strLua, lua.LString("demo"))
	})
}

func Test_TableLua2GoStruct(t *testing.T) {
	l := lua.NewState()
	err := l.DoString(
		`return {
			caller = fromAddr,
			contract = "ERC20",
			func = "transfer",
			args = {"toAddr", 111},
	}`)
	assert.Nil(t, err)
	luaValue := l.CheckTable(1)
	result := &common.Invoke{}
	err = TableLua2GoStruct(luaValue, result)
	assert.Nil(t, err)
}
