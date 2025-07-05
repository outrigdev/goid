//go:build !go1.23 || go1.25 || (go1.23 && !go1.25 && !(amd64 || arm64))

// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

// Get returns the current goroutine ID using stack trace method
func Get() uint64 {
	return getFromStack()
}
