package min_cost_climbing

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "testCase 1",
			input:    []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "testCase 2",
			input:    []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := minCostClimbingStairs(tt.input)
			if got != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, got)
			}
		})
	}
}
