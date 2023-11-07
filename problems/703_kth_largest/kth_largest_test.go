package kth_largest

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestKthLargest(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		k        int
		init     []int
		adds     []int
		expected []int
	}{
		{
			name:     "testCase 1",
			k:        3,
			init:     []int{4, 5, 8, 2},
			adds:     []int{3, 5, 10, 9, 4},
			expected: []int{4, 5, 5, 8, 8},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			k := Constructor(tt.k, tt.init)
			got := make([]int, 0, len(tt.adds))
			for _, add := range tt.adds {
				got = append(got, k.Add(add))
			}
			if !utils.AreSlicesEqual(got, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, got)
			}
		})
	}
}
