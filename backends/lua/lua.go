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
	"os"
	"path/filepath"
	"strings"

	"github.com/alexript/installngo/fs"

	"github.com/bep/overlayfs"
	lua "github.com/yuin/gopher-lua"
)

var ofs *overlayfs.OverlayFs

func loFindFile(name string) (string, string) {
	messages := []string{}
	name = strings.Replace(name, ".", "/", -1)

	ex, err := os.Getwd()
	if err != nil {
		messages = append(messages, err.Error())
	} else {
		if ofs == nil {
			ofs = fs.New(ex, ex)
		}
		// fmt.Printf("package path: '%q'\n", name)
		if fi, err := ofs.Stat(name); err == nil {
			if fi.IsDir() {
				filename := filepath.Join(name, "init.lua")
				if fffi, err := ofs.Stat(filename); err == nil {
					if !fffi.IsDir() {
						return filename, ""
					} else {
						messages = append(messages, "Unable to find init.lua")
					}
				} else {
					messages = append(messages, err.Error())
				}
			} else {
				filename := name + ".lua"
				if ffi, err := ofs.Stat(filename); err == nil {
					if !ffi.IsDir() {
						return filename, ""
					} else {
						messages = append(messages, filename+" is a directory.")
					}
				} else {
					messages = append(messages, err.Error())
				}
			}
		} else {
			// fmt.Printf("err:: %q\n", err)
			messages = append(messages, err.Error())
		}
	}

	return "", strings.Join(messages, "\n\t")
}

func backendLoader(L *lua.LState) int {
	name := L.CheckString(1)
	// fmt.Printf("Require:: %q\n", name)
	path, msg := loFindFile(name)
	// fmt.Printf("Requred path:: %q\n", path)
	if ofs == nil {
		L.Push(lua.LString(`Unable to initialize filesystem abstractions`))
		return 1
	}
	if len(path) == 0 {
		L.Push(lua.LString(msg))
		return 1
	}
	file, err := ofs.Open(path)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	defer file.Close()
	fn, err1 := L.Load(file, path)
	if err1 != nil {
		L.RaiseError(err1.Error())
	}
	L.Push(fn)

	return 1
}

func newLState() *lua.LState {
	L := lua.NewState()

	loaders, ok := L.GetField(L.Get(lua.RegistryIndex), "_LOADERS").(*lua.LTable)
	if !ok {
		L.RaiseError("package.loaders must be a table")
	}

	//	fmt.Printf("%q\n", loaders)
	loaders.Append(L.NewFunction(backendLoader))

	L.SetField(L.Get(lua.RegistryIndex), "_LOADERS", loaders)

	return L
}

func closeOFS() {
	ofs = nil
}

func DoString(luastring string) {
	L := newLState()
	defer L.Close()
	defer closeOFS()
	if err := L.DoString(luastring); err != nil {
		panic(err)
	}
}

func DoFile(filename string) {
	L := newLState()
	defer L.Close()
	defer closeOFS()
	if err := L.DoFile(filename); err != nil {
		panic(err)
	}
}

func DoRequire(modulename string) {
	DoString(`require("` + modulename + `")`)
}
