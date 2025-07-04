// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

import "testing"

func BenchmarkGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Get()
	}
}

func BenchmarkGetFromStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = GetFromStack()
	}
}