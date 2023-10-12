// Package max_depth
// difficulty: easy
// link: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
package max_depth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}

type NodeDepth struct {
	node  *TreeNode
	depth int
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := 0
	queue := []NodeDepth{NodeDepth{
		node:  root,
		depth: 1,
	}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.depth > maxDepth {
			maxDepth = node.depth
		}
		if node.node.Left != nil {
			queue = append(queue, NodeDepth{
				node:  node.node.Left,
				depth: node.depth + 1,
			})
		}

		if node.node.Right != nil {
			queue = append(queue, NodeDepth{
				node:  node.node.Right,
				depth: node.depth + 1,
			})
		}
	}
	return maxDepth
}
