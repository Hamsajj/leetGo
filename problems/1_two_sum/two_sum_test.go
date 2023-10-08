package two_sum

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "testCase 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "testCase 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "testCase 2",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSum(tt.nums, tt.target)
			if !utils.AreSlicesEqualWithoutOrder(res, tt.expected) {
				t.Errorf("expected %v, but recieved %v", tt.expected, res)
			}
		})
	}
}
