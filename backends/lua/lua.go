package lua

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func loFindFile(L *lua.LState, name, pname string) (string, string) {
	messages := []string{}
	name = strings.Replace(name, ".", string(os.PathSeparator), -1)

	ex, err := os.Getwd()
	if err != nil {
		messages = append(messages, err.Error())
	} else {
		dirPath := filepath.Join(ex, name)
		fmt.Printf("package path: '%q'\n", dirPath)
		if fi, err := os.Stat(dirPath); err == nil {
			if fi.IsDir() {
				filename := filepath.Join(dirPath, "init.lua")
				return filename, ""
			} else {
				filename := dirPath + ".lua"
				return filename, ""
			}
		} else {
			messages = append(messages, err.Error())
		}
	}

	return "", strings.Join(messages, "\n\t")
}

func backendLoader(L *lua.LState) int {
	name := L.CheckString(1)
	fmt.Printf("Require:: %q\n", name)
	path, msg := loFindFile(L, name, "path")
	fmt.Printf("Requred path:: %q\n", path)
	if len(path) == 0 {
		L.Push(lua.LString(msg))
		return 1
	}
	fn, err1 := L.LoadFile(path) //L.Load(reader, filename)
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

func DoString(luastring string) {
	L := newLState()
	defer L.Close()
	if err := L.DoString(luastring); err != nil {
		panic(err)
	}
}

func DoFile(filename string) {
	L := newLState()
	defer L.Close()
	if err := L.DoFile(filename); err != nil {
		panic(err)
	}
}

func DoRequire(modulename string) {
	DoString(`require("` + modulename + `")`)

}
