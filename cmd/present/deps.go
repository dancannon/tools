// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"golang.org/x/tools/godoc/static"
)

var depsScripts = []string{"jquery.js", "jquery-ui.js"}

// depsScript registers an HTTP handler at /deps.js that serves all the
// scripts specified by the variable above.
func depsScript(root, transport string) {
	modTime := time.Now()
	var buf bytes.Buffer
	for _, p := range depsScripts {
		if s, ok := static.Files[p]; ok {
			buf.WriteString(s)
			continue
		}
		b, err := ioutil.ReadFile(filepath.Join(root, "static", p))
		if err != nil {
			panic(err)
		}
		buf.Write(b)
	}
	b := buf.Bytes()
	http.HandleFunc("/deps.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/javascript")
		http.ServeContent(w, r, "", modTime, bytes.NewReader(b))
	})
}
