package tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {

	var f func(node *TreeNode)

	f = func(node *TreeNode) {
		if node == nil{
			return
		}

		f(node.Left)
		f(node.Right)
		nodeTmp := node.Left
		right := node.Right
		if nodeTmp != nil{
			for ;nodeTmp.Right!=nil;{
				nodeTmp = nodeTmp.Right
			}
			nodeTmp.Right = right
			node.Right = node.Left
			node.Left = nil
		}
	}

	f(root)
}
