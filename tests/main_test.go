package main

import (
	"reflect"
	"testing"
)

// Test the difference function
func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		a        []string
		b        []string
		expected []string
	}{
		{
			name:     "Simple difference",
			a:        []string{"alice", "bob", "charlie"},
			b:        []string{"bob", "david"},
			expected: []string{"alice", "charlie"},
		},
		{
			name:     "No difference",
			a:        []string{"alice", "bob"},
			b:        []string{"alice", "bob", "charlie"},
			expected: []string{},
		},
		{
			name:     "Empty first array",
			a:        []string{},
			b:        []string{"alice"},
			expected: []string{},
		},
		{
			name:     "All different",
			a:        []string{"alice", "bob"},
			b:        []string{"charlie", "david"},
			expected: []string{"alice", "bob"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := difference(tt.a, tt.b)
			if !slicesEqual(result, tt.expected) {
				t.Errorf("difference(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Test the intersection function
func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		a        []string
		b        []string
		expected []string
	}{
		{
			name:     "Simple intersection",
			a:        []string{"alice", "bob", "charlie"},
			b:        []string{"bob", "charlie", "david"},
			expected: []string{"bob", "charlie"},
		},
		{
			name:     "No intersection",
			a:        []string{"alice", "bob"},
			b:        []string{"charlie", "david"},
			expected: []string{},
		},
		{
			name:     "Empty arrays",
			a:        []string{},
			b:        []string{"alice"},
			expected: []string{},
		},
		{
			name:     "Complete intersection",
			a:        []string{"alice", "bob"},
			b:        []string{"alice", "bob"},
			expected: []string{"alice", "bob"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intersection(tt.a, tt.b)
			if !slicesEqual(result, tt.expected) {
				t.Errorf("intersection(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Helper function to compare slices regardless of order
func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	
	// Create maps for comparison
	mapA := make(map[string]int)
	mapB := make(map[string]int)
	
	for _, item := range a {
		mapA[item]++
	}
	for _, item := range b {
		mapB[item]++
	}
	
	return reflect.DeepEqual(mapA, mapB)
}

// Benchmark tests
func BenchmarkDifference(b *testing.B) {
	a := make([]string, 1000)
	bArray := make([]string, 1000)
	
	for i := 0; i < 1000; i++ {
		a[i] = string(rune('a' + i%26))
		bArray[i] = string(rune('b' + i%26))
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		difference(a, bArray)
	}
}

func BenchmarkIntersection(b *testing.B) {
	a := make([]string, 1000)
	bArray := make([]string, 1000)
	
	for i := 0; i < 1000; i++ {
		a[i] = string(rune('a' + i%26))
		bArray[i] = string(rune('a' + i%26))
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersection(a, bArray)
	}
} 