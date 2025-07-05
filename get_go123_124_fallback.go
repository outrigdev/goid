//go:build go1.23 && !go1.25 && !(amd64 || arm64)

// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

import (
	"bytes"
	"runtime"
)

// Get returns the current goroutine ID using stack trace method for Go 1.23-1.24 on non-amd64/arm64 architectures
//
//go:inline
func Get() uint64 {
	// copied from goid.go GetFromStack function (pre go 1.23 will not inline this if we call GetFromStack() directly, so we inline it manually here, ~10% faster)
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	b := buf[:n]

	// fast HasPrefix (no alloc)
	if !bytes.HasPrefix(b, goroutinePrefixBytes) {
		return 0
	}
	b = b[len(goroutinePrefix):]

	// find end of ID (space before “[status]”)
	idx := bytes.IndexByte(b, ' ')
	if idx < 0 {
		return 0
	}

	// parse digits directly
	var id uint64
	for _, c := range b[:idx] {
		id = id*10 + uint64(c-'0')
	}
	return id
}
