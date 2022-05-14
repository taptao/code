package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0{
        return nil
    }
    i := findMax(nums)

    node := new(TreeNode)
    node.Val = nums[i]

    node.Left = constructMaximumBinaryTree(nums[:i])
    node.Right = constructMaximumBinaryTree(nums[i+1:])
	return node
}


func findMax(nums []int) int{
    max := nums[0]
    index := 0
	for i := range nums{
	    if max < nums[i]{
	        max = nums[i]
	        index = i
        }
    }
    return index
}