// Package merge_sorted_linked_list
// difficulty: easy
// link: https://leetcode.com/problems/merge-two-sorted-lists/
package merge_sorted_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	root := &ListNode{}
	currNode := root
	for head1 != nil && head2 != nil {
		currNode.Next = &ListNode{}
		currNode = currNode.Next
		if head1.Val < head2.Val {
			currNode.Val = head1.Val
			head1 = head1.Next
		} else if head2.Val < head1.Val {
			currNode.Val = head2.Val
			head2 = head2.Next
		} else {
			currNode.Val = head1.Val
			currNode.Next = &ListNode{Val: head2.Val}
			currNode = currNode.Next
			head2 = head2.Next
			head1 = head1.Next
		}
	}

	if head1 != nil {
		currNode.Next = head1
	} else if head2 != nil {
		currNode.Next = head2
	}
	return root.Next
}
