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

package fs

import (
	"archive/zip"
	"fmt"
	"strings"

	"github.com/bep/overlayfs"
	"github.com/spf13/afero"
	"github.com/spf13/afero/zipfs"
)

func New(cwd string, bundlesPath string) *overlayfs.OverlayFs {

	newFS := overlayfs.New(overlayfs.Options{Fss: []afero.Fs{}, FirstWritable: false})

	osFS := afero.NewOsFs()
	basepathFS := afero.NewBasePathFs(osFS, cwd)
	newFS = newFS.Append(basepathFS)

	dir, err := basepathFS.Open("/")
	if err != nil {
		fmt.Printf("Failed bundles search: %q\n", err.Error())
		return newFS
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		fmt.Printf("Failed bundles search: %q\n", err.Error())
		return newFS
	}

	for _, name := range names {
		if strings.HasSuffix(strings.ToLower(name), ".zip") {
			//			fmt.Printf("name: %q\n", name)
			zrc, err := zip.OpenReader(name)
			if err == nil {
				zfs := zipfs.New(&zrc.Reader)
				fs := &afero.Afero{Fs: zfs}
				if fi, err := fs.Stat("/lua"); err == nil {
					if fi.IsDir() {
						luaFS := afero.NewBasePathFs(fs, "/lua")
						newFS = newFS.Append(luaFS)
						//
						// if ffi, err := luaFS.Stat("/init.lua"); err == nil {
						// 	if !ffi.IsDir() {
						//
						// 		file, err := luaFS.Open("/init.lua")
						// 		if err == nil {
						// 			defer file.Close()
						// 			fn, err1 := L.Load(file, "/init.lua")
						// 			if err1 != nil {
						// 				L.RaiseError(err1.Error())
						// 			}
						// 			L.Push(fn)
						// 		}
						//
						// 	}
						// }
					}
				}
			}
		}
	}

	return newFS
}
