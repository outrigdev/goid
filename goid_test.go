// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

import (
	"fmt"
	"sync"
	"testing"
)

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


func TestGoroutineIDUniquenessAndConsistency(t *testing.T) {
	const numGoroutines = 100
	
	// Channel to collect results from goroutines
	type result struct {
		goroutineIndex int
		getID          uint64
		getFromStackID uint64
		err            error
	}
	
	results := make(chan result, numGoroutines)
	var wg sync.WaitGroup
	
	// Launch 100 goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			
			// Get IDs using both methods
			getID := Get()
			getFromStackID := GetFromStack()
			
			// Verify both methods return the same value
			var err error
			if getID != getFromStackID {
				err = fmt.Errorf("goroutine %d: Get()=%d != GetFromStack()=%d", 
					index, getID, getFromStackID)
			}
			
			results <- result{
				goroutineIndex: index,
				getID:          getID,
				getFromStackID: getFromStackID,
				err:            err,
			}
		}(i)
	}
	
	// Wait for all goroutines to complete
	wg.Wait()
	close(results)
	
	// Collect and analyze results
	seenIDs := make(map[uint64][]int) // map[goid][]goroutineIndexes
	var errors []error
	
	for res := range results {
		if res.err != nil {
			errors = append(errors, res.err)
			continue
		}
		
		// Track which goroutines had this ID
		seenIDs[res.getID] = append(seenIDs[res.getID], res.goroutineIndex)
	}
	
	// Report any consistency errors between Get() and GetFromStack()
	if len(errors) > 0 {
		for _, err := range errors {
			t.Error(err)
		}
	}
	
	// Verify uniqueness - each goroutine should have a unique ID
	var duplicates []string
	for goid, goroutineIndexes := range seenIDs {
		if len(goroutineIndexes) > 1 {
			duplicates = append(duplicates, fmt.Sprintf("goid %d used by goroutines %v", goid, goroutineIndexes))
		}
	}
	
	if len(duplicates) > 0 {
		t.Errorf("Found duplicate goroutine IDs:\n%s", fmt.Sprintf("  %s\n", duplicates))
	}
	
	// Verify we got results from all goroutines
	if len(seenIDs) != numGoroutines {
		t.Errorf("Expected %d unique goroutine IDs, got %d", numGoroutines, len(seenIDs))
	}
	
	// Log success summary
	t.Logf("Successfully tested %d goroutines:", numGoroutines)
	t.Logf("  - All goroutine IDs are unique: %v", len(duplicates) == 0)
	t.Logf("  - Get() and GetFromStack() are consistent: %v", len(errors) == 0)
	t.Logf("  - ID range: %d to %d", minID(seenIDs), maxID(seenIDs))
}

// Helper function to find minimum ID
func minID(seenIDs map[uint64][]int) uint64 {
	var min uint64 = ^uint64(0) // max uint64
	for id := range seenIDs {
		if id < min {
			min = id
		}
	}
	return min
}

// Helper function to find maximum ID
func maxID(seenIDs map[uint64][]int) uint64 {
	var max uint64
	for id := range seenIDs {
		if id > max {
			max = id
		}
	}
	return max
}
