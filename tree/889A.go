package tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {

	var build func(preStart int, preEnd int, postStart int, postEnd int) *TreeNode
	postorderMap := make(map[int]int)

	for i,v := range postorder{
		postorderMap[v]= i
	}
	build = func(preStart int, preEnd int, postStart int, postEnd int) *TreeNode {
		if preStart > preEnd{
			return nil
		}
		node := &TreeNode{
			Val: preorder[preStart],
		}
		if preStart + 1 > preEnd{
			return node
		}
		index := postorderMap[preorder[preStart+1]]

		leftNodeNum := index - postStart + 1

		node.Left = build(preStart+1, preStart+leftNodeNum, postStart, index)
		node.Right = build(preStart+leftNodeNum+1, preEnd, index+1, postEnd-1)

		return node
	}
	node := build(0, len(preorder)-1, 0, len(postorder)-1)
	return node
}