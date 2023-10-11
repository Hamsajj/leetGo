package reorder_list

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "testCase 1",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 4, 2, 3},
		},
		{
			name:     "testCase 2",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "testCase 3",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 5, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intListToNodeList(tt.input)
			reorderList(got)
			gotNums := nodeListToIntList(got)
			if !utils.AreSlicesEqual(tt.expected, gotNums) {
				t.Errorf("expected %v, but recieved %v", tt.expected, gotNums)
			}
		})
	}
}

func intListToNodeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	curr := root
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{
			Val: nums[i],
		}
		curr = curr.Next
	}
	return root
}

func nodeListToIntList(node *ListNode) []int {
	nums := make([]int, 0, 1)
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	return nums
}
