package remove_nth_from_last

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestName(t *testing.T) {
	var tests = []struct {
		name     string
		head     []int
		n        int
		expected []int
	}{
		{
			name:     "testCase 1",
			head:     []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "testCase 1",
			head:     []int{1, 2, 3, 4, 5},
			n:        5,
			expected: []int{2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(intListToNodeList(tt.head), tt.n)
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
