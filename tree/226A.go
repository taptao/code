package tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	var invert func(node *TreeNode)

	invert = func(node *TreeNode) {
		if node == nil{
			return
		}

		invert(node.Left)
		invert(node.Right)

		left := node.Left
		node.Left = node.Right
		node.Right = left
	}

	invert(root)

	return root
}
