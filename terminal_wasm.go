// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build wasm

package conf

// isTerminal returns true if the given file descriptor is a terminal.
func isTerminal(fd int) bool {
    return false
}
