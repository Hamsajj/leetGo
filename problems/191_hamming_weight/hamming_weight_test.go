package hamming_weight

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    uint32
		expected int
	}{
		{
			name:     "testCase 1",
			input:    11,
			expected: 3,
		},
		{
			name:     "testCase 2",
			input:    4294967293,
			expected: 31,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.input)
			if got != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, got)
			}
		})
	}
}
