package tapping_rain_water

import "testing"

func TestTrapRain(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "testCase 1",
			input:    []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "testCase 2",
			input:    []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.input)
			if got != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, got)
			}
		})
	}
}
