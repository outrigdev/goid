# goid

A high-performance Go package for retrieving the current goroutine ID.

## Version Compatibility & Optimization

- **Optimized for Go 1.23 and 1.24** on **amd64** and **arm64** architectures
- **Future-proofed**: Automatic fallback to stack trace method for newer and older Go versions or unsupported architectures
- **Universal compatibility**: Works on all Go versions and architectures (via slower fallback for unsupported versions)

## Overview

The `goid` package provides ultra-fast access to the current goroutine ID using optimized assembly implementations for Go 1.23-1.24 on amd64/arm64, with automatic fallback to a stack trace method for maximum compatibility.

## Features

- **Ultra-fast performance**: Direct assembly access to goroutine ID on Go 1.23-1.24 (amd64/arm64)
- **Automatic fallback**: Stack trace method for all other Go versions or architectures
- **Future-proofed**: Graceful degradation for newer Go versions until optimizations are added
- **Zero dependencies**: Pure Go implementation with no external dependencies

## Installation

```bash
go get github.com/outrigdev/goid
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/outrigdev/goid"
)

func main() {
    // Get the current goroutine ID
    id := goid.Get()
    fmt.Printf("Current goroutine ID: %d\n", id)
    
    // Force use of stack trace method (for comparison)
    id2 := goid.GetFromStack()
    fmt.Printf("Goroutine ID from stack: %d\n", id2)
}
```

## API

### `Get() uint64`

Returns the current goroutine ID using the fastest available method:
- On Go 1.23+ with amd64/arm64: Uses optimized assembly implementation
- On other versions/architectures: Falls back to stack trace method

### `GetFromStack() uint64`

Returns the current goroutine ID by parsing the stack trace. This is the fallback method used when the optimized assembly implementation is not available.

## Performance

Benchmark results on Apple M3 Pro (Go 1.24, darwin/arm64):

```
goos: darwin
goarch: arm64
pkg: github.com/outrigdev/goid
cpu: Apple M3 Pro
BenchmarkGet-12             	769931396	         1.334 ns/op
BenchmarkGetFromStack-12    	  733010	      1605 ns/op
PASS
ok  	github.com/outrigdev/goid	2.697s
```

The optimized `Get()` function is approximately **1000x+ faster** than the stack trace method, demonstrating the significant performance benefit of the assembly optimization.

This substantial speed-up is particularly beneficial for debugging, tracing, logging, or other scenarios where retrieving goroutine IDs frequently occurs.

## Supported Platforms

### Optimized Assembly Implementation
- Go 1.23, 1.24 on amd64 architecture
- Go 1.23, 1.24 on arm64 architecture

### Stack Trace Fallback
- All Go versions on all architectures
- Automatically used when assembly implementation is not available

## How It Works

The package uses build constraints to select the appropriate implementation:

1. **Assembly Method**: Directly accesses the goroutine structure in memory using assembly code to read the `goid` field
2. **Stack Trace Method**: Parses the first line of `runtime.Stack()` output to extract the goroutine ID using regex

The assembly method is significantly faster but requires knowledge of the internal Go runtime structures, which can change between Go versions. The stack trace method is slower but works across all Go versions.

## Technical Implementation

For the optimized assembly implementation, this package includes exact copies of the internal `g` (goroutine) struct definitions from the Go runtime source code for each supported version:

- **Go 1.23**: Complete `g` struct with 161 lines including all fields like `goid`, `stack`, `atomicstatus`, etc.
- **Go 1.24**: Updated `g` struct with additional fields like `fipsIndicator` and `syncGroup`

**Critical: Version-Locked Compatibility**

This package uses Go build constraints to version-lock each struct definition to its corresponding Go release. This is essential because the internal `g` struct layout can change between Go versions, and accessing the wrong memory offset would cause crashes or incorrect results.

Many other goroutine ID libraries attempt to use these internal structures but fail to properly version-lock them, leading to compatibility issues when Go releases update the runtime internals. This package solves that problem by:

1. **Exact struct copies**: Each struct is an exact copy from the corresponding Go runtime source
2. **Build constraints**: `//go:build go1.23 && !go1.24` ensures each struct is only used with its intended Go version
3. **Automatic fallback**: When the optimized version isn't available, the package automatically uses the universal stack trace method

The assembly functions (`getg()`) retrieve the current goroutine pointer, and the Go code then accesses the `goid` field directly from the version-matched struct, ensuring both safety and performance.

## Implementation Matrix

| Go Version | Architecture | getg() Assembly | Optimized | Implementation |
|------------|--------------|-----------------|-----------|----------------|
| pre-1.23   | all          | ❌              | ❌        | Stack trace fallback |
| 1.23       | amd64        | ✅              | ✅        | Assembly + struct access |
| 1.23       | arm64        | ✅              | ✅        | Assembly + struct access |
| 1.23       | other        | ❌              | ❌        | Stack trace fallback |
| 1.24       | amd64        | ✅              | ✅        | Assembly + struct access |
| 1.24       | arm64        | ✅              | ✅        | Assembly + struct access |
| 1.24       | other        | ❌              | ❌        | Stack trace fallback |
| 1.25+      | all          | ❌              | ❌        | Stack trace fallback |

**Note**: Go 1.25 optimization will be added as soon as Go 1.25 stable is released.

## License

Apache-2.0

## Copyright

Copyright 2025, Command Line Inc.