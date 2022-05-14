package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	var traverse func(node *TreeNode) int

	traverse = func(node *TreeNode) int {

		if node == nil{
			return 0
		}

		leftDepth := traverse(node.Left)
		rightDepth := traverse(node.Right)

		if max < (leftDepth+rightDepth){
			max = leftDepth+rightDepth
		}
		if leftDepth > rightDepth{
			return leftDepth + 1

		}else{
			return rightDepth + 1
		}

	}

	traverse(root)

	return max
}