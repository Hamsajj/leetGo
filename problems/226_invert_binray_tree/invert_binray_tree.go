// Package invert_binray_tree
// difficulty: easy
// link: https://leetcode.com/problems/invert-binary-tree/description/
package invert_binray_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

func invertTreeStack(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var s stack
	s = s.Push(root)
	for s.Len() > 0 {
		var node *TreeNode
		s, node = s.Pop()
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			s = s.Push(node.Left)
		}
		if node.Right != nil {
			s = s.Push(node.Right)
		}
	}
	return root
}

type stack []*TreeNode

func (s stack) Push(v *TreeNode) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, *TreeNode) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Len() int {
	return len(s)
}
