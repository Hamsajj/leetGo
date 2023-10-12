// Package bt_level_order_traversal
// difficulty: medium
// link: https://leetcode.com/problems/binary-tree-level-order-traversal/
package bt_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeWithDepth struct {
	node  *TreeNode
	depth int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*NodeWithDepth{
		{
			node:  root,
			depth: 0,
		},
	}
	var res [][]int
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if len(res) > current.depth {
			res[current.depth] = append(res[current.depth], current.node.Val)
		} else {
			res = append(res, []int{current.node.Val})
		}
		if current.node.Left != nil {

			queue = append(queue, &NodeWithDepth{
				node:  current.node.Left,
				depth: current.depth + 1,
			})
		}
		if current.node.Right != nil {
			queue = append(queue, &NodeWithDepth{
				node:  current.node.Right,
				depth: current.depth + 1,
			})
		}
	}
	return res
}
