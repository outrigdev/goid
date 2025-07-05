//go:build go1.23 && !go1.25 && !(amd64 || arm64)

// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

// Get returns the current goroutine ID using stack trace method for Go 1.23-1.24 on non-amd64/arm64 architectures
//
//go:inline
func Get() uint64 {
	return GetFromStack()
}
