// Package remove_nth_from_last
// difficulty: medium
// link: https://leetcode.com/problems/remove-nth-node-from-end-of-list
package remove_nth_from_last

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		if n != 1 {
			return head
		} else if n == 1 {
			return nil
		}
	}
	curr := head
	nthFromLast := head
	for i := 0; i < n; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return head.Next
	}
	for curr.Next != nil {
		curr = curr.Next
		nthFromLast = nthFromLast.Next
	}
	nthFromLast.Next = nthFromLast.Next.Next
	return head
}
