//go:build !go1.23 || go1.25 || (go1.23 && !go1.25 && !(amd64 || arm64))

// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

// Get returns the current goroutine ID using the fastest available method.
//
// On Go 1.23-1.24 with amd64/arm64 architectures, this function uses an optimized
// assembly implementation that directly accesses the goroutine structure for
// maximum performance.
//
// On all other Go versions or architectures, it automatically falls back to a
// stack trace parsing method that provides universal compatibility across all
// platforms while maintaining good performance through optimized parsing.
//
// The implementation is automatically selected at build time using Go build
// constraints, ensuring optimal performance when available and reliable
// compatibility everywhere else.
//
// Returns 0 if the goroutine ID cannot be determined.
func Get() uint64 {
	return getFromStack()
}
