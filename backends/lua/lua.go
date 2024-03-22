package lua

import (
	lua "github.com/yuin/gopher-lua"
)

func DoString(luastring string) {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(luastring); err != nil {
		panic(err)
	}
}
