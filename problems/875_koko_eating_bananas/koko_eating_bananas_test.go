package koko_eating_bananas

import (
	"fmt"
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "testCase 1",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "testCase 2",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "testCase 2",
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := minEatingSpeed(tt.piles, tt.h)
			if tt.expected != got {
				fmt.Errorf("expected %d, but recieved %d", tt.expected, got)
			}
		})
	}
}
