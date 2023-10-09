package best_time_to_buy_sell

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "testCase 1",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "testCase 2 - no profit",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, got)
			}
		})
	}
}
