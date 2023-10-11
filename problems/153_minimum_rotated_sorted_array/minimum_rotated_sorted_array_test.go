package minimum_rotated_sorted_array

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "testCase 1",
			input:    []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "testCase 2",
			input:    []int{0, 1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "testCase 3",
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "testCase 4",
			input:    []int{4, 5, 6, 7, 8, 2},
			expected: 2,
		},
		{
			name:     "testCase 5",
			input:    []int{5, 1, 2, 3, 4},
			expected: 1,
		},
		{
			name:     "testCase 6",
			input:    []int{4, 5, 1, 2, 3},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.input)
			if got != tt.expected {
				t.Errorf("Expected %d, but received %d", tt.expected, got)
			}
		})
	}
}
