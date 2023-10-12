// Package LCA_binary_tree
// difficulty: medium
// link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
package LCA_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	currentLCA := root
	for {
		if p.Val > currentLCA.Val && q.Val > currentLCA.Val {
			currentLCA = currentLCA.Right
		} else if p.Val < currentLCA.Val && q.Val < currentLCA.Val {
			currentLCA = currentLCA.Left
		} else {
			break
		}
	}
	return currentLCA
}
