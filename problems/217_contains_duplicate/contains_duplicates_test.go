package contains_duplicates

import (
	"testing"
)

func TestContainsDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "contains duplicate",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "does not contain duplicate",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "zero length",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "length of 1",
			nums:     []int{1},
			expected: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if containsDuplicate(tt.nums) != tt.expected {
				t.Errorf("expected result to be %t", tt.expected)
			}
		})
	}
}
