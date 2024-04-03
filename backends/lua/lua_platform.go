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

	lual "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func isOsName(L *lual.LState, osname string) int {
	if runtime.GOOS == osname {
		L.Push(lual.LTrue)
	} else {
		L.Push(lual.LFalse)
	}
	return 1
}
func isWindows(L *lual.LState) int {
	return isOsName(L, "windows")
}

func isLinux(L *lual.LState) int {
	return isOsName(L, "linux")
}

func isMacos(L *lual.LState) int {
	return isOsName(L, "macos")
}

type platform struct {
	OsName    string
	Arch      string
	IsWindows *lual.LFunction
	IsLinux   *lual.LFunction
	IsMacOS   *lual.LFunction
}

func platformLoader(L *lual.LState) int {
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
func AttachPlatform(L *lual.LState) {
	L.PreloadModule("platform", platformLoader)
}
