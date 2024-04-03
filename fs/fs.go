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

func New(cwd string) *overlayfs.OverlayFs {

	newVirtualFS := overlayfs.New(overlayfs.Options{Fss: []afero.Fs{}, FirstWritable: false})

	osFS := afero.NewOsFs()
	cwdFS := afero.NewBasePathFs(osFS, cwd)

	baseLuaFS := cwdFS

	if cwdLuadirStat, err := cwdFS.Stat("/lua"); err == nil {
		if cwdLuadirStat.IsDir() {
			baseLuaFS = afero.NewBasePathFs(osFS, cwd+"/lua")
		}
	}

	newVirtualFS = newVirtualFS.Append(baseLuaFS)

	cwdir, err := cwdFS.Open("/")
	if err != nil {
		fmt.Printf("Failed bundles search: %q\n", err.Error())
		return newVirtualFS
	}
	defer cwdir.Close()
	filenamesInCwd, err := cwdir.Readdirnames(-1)
	if err != nil {
		fmt.Printf("Failed bundles search: %q\n", err.Error())
		return newVirtualFS
	}

	for _, filename := range filenamesInCwd {
		if strings.HasSuffix(strings.ToLower(filename), ".zip") {
			//fmt.Printf("name: %q\n", name)
			zrc, err := zip.OpenReader(filename)
			if err == nil {
				zipReaderFS := zipfs.New(&zrc.Reader)
				zipFS := &afero.Afero{Fs: zipReaderFS}
				if luadirInZipStat, err := zipFS.Stat("/lua"); err == nil {
					if luadirInZipStat.IsDir() {
						zipLuaFS := afero.NewBasePathFs(zipFS, "/lua")
						newVirtualFS = newVirtualFS.Append(zipLuaFS)
					}
				}
			}
		}
	}

	return newVirtualFS
}
