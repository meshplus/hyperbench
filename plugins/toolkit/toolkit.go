package toolkit

import (
	"encoding/hex"
	"encoding/json"
	"reflect"
)

// ToolKit is the set of tool plugins.
// exported methods and fields can be inject into vm.
type ToolKit struct {
}

// NewToolKit create a new ToolKit instant as a plugin to vm
func NewToolKit() *ToolKit {
	return &ToolKit{}
}

// RandStr generate a random string with specific length
func (t *ToolKit) RandStr(l uint) string {
	return randomString(l)
}

// RandInt generate a random int in specific range
func (t *ToolKit) RandInt(min, max int) int {
	return randomInt(min, max)
}

// String convert to string
func (t *ToolKit) String(input interface{}, offsets ...int) string {
	v := reflect.ValueOf(input)
	switch v.Kind() {
	case reflect.Ptr:
		return t.String(v.Elem().Interface(), offsets...)
	case reflect.Slice:
		bs := transferToBts(input)
		l := len(bs)
		start, end := 0, l
		if len(offsets) > 0 && offsets[0] <= l {
			start = offsets[0]
		}
		if len(offsets) > 1 && offsets[1] <= l+1 {
			start = offsets[1]
		}
		return string(bs[start:end])
	case reflect.Array:
		bs := transferToBts(input)
		l := v.Type().Len()
		start, end := 0, l
		if len(offsets) > 0 && offsets[0] <= l {
			start = offsets[0]
		}
		if len(offsets) > 1 && offsets[1] <= l+1 {
			start = offsets[1]
		}
		ss := make([]byte, 0, end-start)
		for i := start; i < end; i++ {
			ss = append(ss, bs[i])
		}
		return string(ss)
	}
	return ""
}

func transferToBts(input interface{}) []byte {
	bs, ok := input.([]byte)
	if !ok {
		bs = []byte{}
		jsData, _ := json.Marshal(input)
		err := json.Unmarshal(jsData, &bs)
		if err != nil {
			return []byte{}
		}
	}
	return bs
}

// Hex encode string as hex string
func (t *ToolKit) Hex(input string) string {
	return hex.EncodeToString([]byte(input))
}
