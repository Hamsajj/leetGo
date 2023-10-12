// Package linked_list_cycle
// difficulty: easy
// link: https://leetcode.com/problems/linked-list-cycle/description/
package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycleMap(head *ListNode) bool {
	seen := make(map[*ListNode]struct{})

	for head != nil {
		if _, exists := seen[head]; exists {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}
