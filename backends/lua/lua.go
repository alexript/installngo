// The MIT License (MIT)
//
// Copyright © 2014 Alex 'Ript' Malyshev
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the “Software”), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package lua

import (
	lua "github.com/yuin/gopher-lua"
)

func attachLibsAndObjects(L *lua.LState) {
	AttachPlatform(L)
	AttachLuaLibs(L)
}

func newLState() *lua.LState {
	L := lua.NewState()

	loaders, ok := L.GetField(L.Get(lua.RegistryIndex), "_LOADERS").(*lua.LTable)
	if !ok {
		L.RaiseError("package.loaders must be a table")
	}

	loaders.Append(L.NewFunction(VFSLoader))

	L.SetField(L.Get(lua.RegistryIndex), "_LOADERS", loaders)
	attachLibsAndObjects(L)

	return L
}

func DoString(luastring string) {
	L := newLState()
	defer L.Close()
	defer CloseOFS()
	if err := L.DoString(luastring); err != nil {
		panic(err)
	}
}

func DoFile(filename string) {
	L := newLState()
	defer L.Close()
	defer CloseOFS()
	if err := L.DoFile(filename); err != nil {
		panic(err)
	}
}

func DoRequire(modulename string) {
	DoString(`require("` + modulename + `")`)
}
