//go:build go1.25

package goid

// Go 1.25+ is not supported yet - the runtime.g struct may have changed
type g struct {
	// Placeholder - this version is not supported yet
	unsupportedGoVersion [0]int
	goid                 uint64
}

// Get returns the current goroutine ID using stack trace method for Go 1.25+
func Get() uint64 {
	return GetFromStack()
}
