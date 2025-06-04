package main

import (
	"testing"
)

// Mock data for testing
var mockUsers = map[string]string{
	"U123ABC": "alice.johnson",
	"U456DEF": "bob.smith", 
	"U789GHI": "charlie.brown",
	"U012JKL": "david.wilson",
	"U345MNO": "emma.davis",
}

var mockChannelMembers = map[string][]string{
	"C1234567890": {"U123ABC", "U456DEF", "U789GHI", "U012JKL"},    // alice, bob, charlie, david
	"C0987654321": {"U123ABC", "U456DEF", "U345MNO"},               // alice, bob, emma
}

// Test scenarios with mock data
func TestMockChannelComparison(t *testing.T) {
	channelA := "C1234567890"
	channelB := "C0987654321"
	
	membersA := mockChannelMembers[channelA]
	membersB := mockChannelMembers[channelB]
	
	// Test expected overlaps
	common := intersection(membersA, membersB)
	expected := []string{"U123ABC", "U456DEF"} // alice, bob
	
	if len(common) != len(expected) {
		t.Errorf("Expected %d common members, got %d", len(expected), len(common))
	}
	
	// Test users only in A
	onlyA := difference(membersA, membersB)
	expectedOnlyA := []string{"U789GHI", "U012JKL"} // charlie, david
	
	if len(onlyA) != len(expectedOnlyA) {
		t.Errorf("Expected %d members only in A, got %d", len(expectedOnlyA), len(onlyA))
	}
	
	// Test users only in B
	onlyB := difference(membersB, membersA)
	expectedOnlyB := []string{"U345MNO"} // emma
	
	if len(onlyB) != len(expectedOnlyB) {
		t.Errorf("Expected %d members only in B, got %d", len(expectedOnlyB), len(onlyB))
	}
}

// Test edge cases
func TestMockEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		channelA []string
		channelB []string
		desc     string
	}{
		{
			name:     "Empty channels",
			channelA: []string{},
			channelB: []string{},
			desc:     "Both channels empty",
		},
		{
			name:     "One empty channel",
			channelA: []string{"U123ABC"},
			channelB: []string{},
			desc:     "One channel empty",
		},
		{
			name:     "Identical channels",
			channelA: []string{"U123ABC", "U456DEF"},
			channelB: []string{"U123ABC", "U456DEF"},
			desc:     "Identical member lists",
		},
		{
			name:     "Large channel",
			channelA: generateLargeUserList(1000),
			channelB: generateLargeUserList(500),
			desc:     "Large channel test",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common := intersection(tt.channelA, tt.channelB)
			onlyA := difference(tt.channelA, tt.channelB)
			onlyB := difference(tt.channelB, tt.channelA)
			
			// Basic sanity checks
			totalMembers := len(common) + len(onlyA) + len(onlyB)
			uniqueMembers := len(removeDuplicates(append(tt.channelA, tt.channelB...)))
			
			if totalMembers != uniqueMembers {
				t.Errorf("Member count mismatch in %s: total=%d, unique=%d", 
					tt.desc, totalMembers, uniqueMembers)
			}
		})
	}
}

// Helper function to generate large user lists for testing
func generateLargeUserList(size int) []string {
	users := make([]string, size)
	for i := 0; i < size; i++ {
		users[i] = string(rune('U' + i%26)) + string(rune('A' + (i/26)%26)) + "123"
	}
	return users
}

// Helper function to remove duplicates
func removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	var result []string
	
	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}
	return result
}

// Test performance with large datasets
func TestMockPerformance(t *testing.T) {
	largeChannelA := generateLargeUserList(5000)
	largeChannelB := generateLargeUserList(3000)
	
	// This should complete quickly even with large datasets
	common := intersection(largeChannelA, largeChannelB)
	onlyA := difference(largeChannelA, largeChannelB)
	onlyB := difference(largeChannelB, largeChannelA)
	
	t.Logf("Large dataset test completed:")
	t.Logf("  Channel A: %d members", len(largeChannelA))
	t.Logf("  Channel B: %d members", len(largeChannelB))
	t.Logf("  Common: %d members", len(common))
	t.Logf("  Only A: %d members", len(onlyA))
	t.Logf("  Only B: %d members", len(onlyB))
} 