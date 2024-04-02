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
	"runtime"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func isWindows(L *lua.LState) int {
	if runtime.GOOS == "windows" {
		L.Push(lua.LTrue)
	} else {
		L.Push(lua.LFalse)
	}
	return 1
}

func isLinux(L *lua.LState) int {
	if runtime.GOOS == "linux" {
		L.Push(lua.LTrue)
	} else {
		L.Push(lua.LFalse)
	}
	return 1
}

func isMacos(L *lua.LState) int {
	if runtime.GOOS == "macos" {
		L.Push(lua.LTrue)
	} else {
		L.Push(lua.LFalse)
	}
	return 1
}

type platform struct {
	OsName    string
	Arch      string
	IsWindows *lua.LFunction
	IsLinux   *lua.LFunction
	IsMacOS   *lua.LFunction
}

func platformLoader(L *lua.LState) int {
	tbl := luar.New(L, &platform{
		OsName:    runtime.GOOS,
		Arch:      runtime.GOARCH,
		IsWindows: L.NewFunction(isWindows),
		IsLinux:   L.NewFunction(isLinux),
		IsMacOS:   L.NewFunction(isMacos),
	})
	L.Push(tbl)
	return 1
}
func AttachPlatform(L *lua.LState) {
	L.PreloadModule("platform", platformLoader)
}
