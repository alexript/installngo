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
	lual "github.com/yuin/gopher-lua"
)

var ofs *overlayfs.OverlayFs

func createCurrentFS() {
	if ofs == nil {
		cwd, _ := os.Getwd()
		ofs = fs.New(cwd)
	}
}

func GetCurrentFS() *overlayfs.OverlayFs {
	createCurrentFS()
	return ofs
}

func loFindFile(name string) (string, string) {
	messages := []string{}
	requiredfilename := strings.Replace(name, ".", "/", -1)

	createCurrentFS()
	if requiredNameStat, err := ofs.Stat(requiredfilename); err == nil {
		if requiredNameStat.IsDir() {
			initLua := filepath.Join(requiredfilename, "init.lua")
			if initLuaStat, err := ofs.Stat(initLua); err == nil {
				if !initLuaStat.IsDir() {
					return initLua, ""
				} else {
					messages = append(messages, "Unable to find init.lua")
				}
			} else {
				messages = append(messages, err.Error())
			}
		} else {
			luaFilename := requiredfilename + ".lua"
			if luaFileStat, err := ofs.Stat(luaFilename); err == nil {
				if !luaFileStat.IsDir() {
					return luaFilename, ""
				} else {
					messages = append(messages, luaFilename+" is a directory.")
				}
			} else {
				messages = append(messages, err.Error())
			}
		}
	} else {
		// fmt.Printf("err:: %q\n", err)
		messages = append(messages, err.Error())
	}

	return "", strings.Join(messages, "\n\t")
}

func CloseOFS() {
	ofs = nil
}

func VFSLoader(L *lual.LState) int {
	requiredName := L.CheckString(1)
	fileName, msg := loFindFile(requiredName)
	if ofs == nil {
		L.Push(lual.LString(`Unable to initialize filesystem abstractions`))
		return 1
	}
	if len(fileName) == 0 {
		L.Push(lual.LString(msg))
		return 1
	}
	file, err := ofs.Open(fileName)
	if err != nil {
		L.Push(lual.LString(err.Error()))
		return 1
	}
	defer file.Close()
	fn, err1 := L.Load(file, fileName)
	if err1 != nil {
		L.RaiseError(err1.Error())
	}
	L.Push(fn)

	return 1
}
