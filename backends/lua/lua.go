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
	lual "github.com/yuin/gopher-lua"
)

func initLuaFile(L *lual.LState) {
	luaFS := GetCurrentFS()
	if luaFS != nil {
		if ffi, err := luaFS.Stat("/init.lua"); err == nil {
			if !ffi.IsDir() {
				file, err := luaFS.Open("/init.lua")
				if err == nil {
					defer file.Close()
					fn, err1 := L.Load(file, "/init.lua")
					if err1 != nil {
						L.RaiseError(err1.Error())
					}
					L.Push(fn)
					L.PCall(0, lual.MultRet, nil)
				}

			}
		}
	}

}

func attachLibsAndObjects(L *lual.LState) {
	AttachPlatform(L)
	AttachLuaLibs(L)
	initLuaFile(L)
}

func newLState() *lual.LState {
	L := lual.NewState()

	loaders, ok := L.GetField(L.Get(lual.RegistryIndex), "_LOADERS").(*lual.LTable)
	if !ok {
		L.RaiseError("package.loaders must be a table")
	}

	loaders.Append(L.NewFunction(VFSLoader))

	L.SetField(L.Get(lual.RegistryIndex), "_LOADERS", loaders)
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
