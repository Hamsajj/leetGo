package container_with_most_water

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "testCase 1",
			heights:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "testCase 2",
			heights:  []int{1, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.heights)
			if got != tt.expected {
				t.Errorf("Expected %d, recieved %d", tt.expected, got)
			}
		})
	}
}
