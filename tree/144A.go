package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    ret := []int{}

    var traverse func(node *TreeNode)

    traverse = func(node *TreeNode) {
        if node == nil{
            return
        }

        ret = append(ret, node.Val)
        traverse(node.Left)
        traverse(node.Right)
    }

    traverse(root)
    return ret
}
