// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

import (
	"bytes"
	"runtime"
)

const goroutinePrefix = "goroutine "

var goroutinePrefixBytes = []byte(goroutinePrefix)

// getFromStack extracts goroutine ID from stack trace (fallback method)
// optimized to avoid allocations and regexps.
func getFromStack() uint64 {
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

// GetFromStack returns the current goroutine ID by parsing the stack trace.
//
// This function provides a universal fallback method that works on all Go versions
// and architectures. It extracts the goroutine ID from the first line of the stack
// trace output using optimized parsing that avoids allocations and regular expressions.
//
// While significantly slower than the optimized assembly implementations available
// for Go 1.23-1.24 on amd64/arm64, this method provides reliable compatibility
// across all platforms and Go versions.
//
// This function is automatically used as a fallback by Get() when the optimized
// assembly implementation is not available, but can also be called directly for
// testing or comparison purposes.
//
// Returns 0 if the goroutine ID cannot be parsed from the stack trace.
func GetFromStack() uint64 {
	// Use the optimized stack trace method to get goroutine ID
	return getFromStack()
}
