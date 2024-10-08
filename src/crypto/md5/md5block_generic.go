// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (!386 && !amd64 && !arm && !arm64 && !loong64 && !ppc64 && !ppc64le && !s390x) || purego

package md5

const haveAsm = false

func block(dig *digest, p []byte) {
	blockGeneric(dig, p)
}
