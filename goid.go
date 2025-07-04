// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

import (
	"bytes"
	"runtime"
)

// GetFromStack extracts goroutine ID from stack trace (fallback method)
// optimized to avoid allocations and regexps.
func GetFromStack() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	b := buf[:n]

	const prefix = "goroutine "
	// fast HasPrefix (no alloc)
	if !bytes.HasPrefix(b, []byte(prefix)) {
		return 0
	}
	b = b[len(prefix):]

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
