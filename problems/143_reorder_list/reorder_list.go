// Package reorder_list
// difficulty: medium
// link: https://leetcode.com/problems/reorder-list/
package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

type stack []*ListNode

func (s stack) Push(v *ListNode) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, *ListNode) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Len() int {
	return len(s)
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var s stack
	//seen := make(map[*ListNode]struct{})
	curr := head
	for curr != nil {
		s = s.Push(curr)
		curr = curr.Next
	}
	curr = head
	for {
		var last *ListNode
		s, last = s.Pop()
		if last == curr.Next || last == curr {
			last.Next = nil
			break
		}
		last.Next = curr.Next
		curr.Next = last
		curr = last.Next
	}
}
