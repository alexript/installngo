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
	"github.com/vadv/gopher-lua-libs/argparse"
	"github.com/vadv/gopher-lua-libs/base64"
	"github.com/vadv/gopher-lua-libs/crypto"
	"github.com/vadv/gopher-lua-libs/db"
	"github.com/vadv/gopher-lua-libs/filepath"
	"github.com/vadv/gopher-lua-libs/humanize"
	"github.com/vadv/gopher-lua-libs/inspect"
	"github.com/vadv/gopher-lua-libs/json"
	"github.com/vadv/gopher-lua-libs/log"
	plugin "github.com/vadv/gopher-lua-libs/plugin"
	"github.com/vadv/gopher-lua-libs/regexp"
	"github.com/vadv/gopher-lua-libs/storage"
	"github.com/vadv/gopher-lua-libs/strings"
	"github.com/vadv/gopher-lua-libs/template"
	"github.com/vadv/gopher-lua-libs/time"
	"github.com/vadv/gopher-lua-libs/xmlpath"
	"github.com/vadv/gopher-lua-libs/yaml"
	lua "github.com/yuin/gopher-lua"
)

func AttachLuaLibs(L *lua.LState) {
	//	libs.Preload(L)
	plugin.Preload(L)

	argparse.Preload(L)
	base64.Preload(L)
	//	cert_util.Preload(L)
	//	chef.Preload(L)
	//	cloudwatch.Preload(L)
	//	cmd.Preload(L)
	crypto.Preload(L)
	db.Preload(L)
	filepath.Preload(L)
	//	goos.Preload(L)
	//	http.Preload(L)
	humanize.Preload(L)
	inspect.Preload(L)
	//	ioutil.Preload(L)
	json.Preload(L)
	log.Preload(L)
	//	pb.Preload(L)
	//	pprof.Preload(L)
	//	prometheus.Preload(L)
	regexp.Preload(L)
	///	runtime.Preload(L)
	//	shellescape.Preload(L)
	//	stats.Preload(L)
	storage.Preload(L)
	strings.Preload(L)
	//	tac.Preload(L)
	//	tcp.Preload(L)
	//	telegram.Preload(L)
	template.Preload(L)
	time.Preload(L)
	xmlpath.Preload(L)
	yaml.Preload(L)
	// zabbix.Preload(L)
}
