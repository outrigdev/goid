// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

import (
	"regexp"
	"runtime"
	"strconv"
)

var goroutineIDRegexp = regexp.MustCompile(`goroutine (\d+) \[`)

// GetFromStack extracts goroutine ID from stack trace (fallback method)
func GetFromStack() uint64 {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	// Format of the first line of stack trace is "goroutine N [status]:"
	matches := goroutineIDRegexp.FindSubmatch(buf[:n])
	if len(matches) < 2 {
		return 0
	}
	id, err := strconv.ParseInt(string(matches[1]), 10, 64)
	if err != nil {
		return 0
	}
	return uint64(id)
}