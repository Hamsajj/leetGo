package longest_consecutive_sequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "testCase 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "testCase 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "testCase 3",
			nums:     []int{1, 122, 121, 2, 123},
			expected: 3,
		},
		{
			name:     "testCase 4 - len 0",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "testCase 5 - len 1",
			nums:     []int{2},
			expected: 1,
		},
		{
			name:     "testCase 5 - len 1",
			nums:     []int{1, 2, 0, 1},
			expected: 3,
		},
		{
			name:     "testCase 6",
			nums:     []int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7},
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run("longestConsecutiveSet_"+tt.name, func(t *testing.T) {
			got := longestConsecutiveSet(tt.nums)
			if got != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, got)
			}
		})

		t.Run("longestConsecutive_"+tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if got != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, got)
			}
		})
	}
}
