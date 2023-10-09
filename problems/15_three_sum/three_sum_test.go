package three_sum

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "testCase 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "testCase 2 - no answer",
			nums:     []int{1, 0, 1},
			expected: [][]int{},
		},
		{
			name:     "testCase 3",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "testCase 4",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if !utils.AreArrayOfArraysEqual(got, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, got)
			}
		})
	}
}
