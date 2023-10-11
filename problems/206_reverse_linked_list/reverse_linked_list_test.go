package reverse_linked_list

import "testing"

func TestName(t *testing.T) {
	var tests = []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name: "testCase 1",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "testCase 2",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			expected: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		{
			name: "testCase 2",
			head: &ListNode{
				Val:  1,
				Next: nil,
			},
			expected: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.head)
			if !areLinkedListsEqual(tt.expected, got) {
				t.Errorf("Not matching")
			}
		})
	}
}

func areLinkedListsEqual(head1 *ListNode, head2 *ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	if head1 != nil || head2 != nil {
		return false
	}
	return true
}
