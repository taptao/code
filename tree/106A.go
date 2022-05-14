package tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromInAndPost(inorder []int, postorder []int) *TreeNode {
 	var build func(postStart int, postEnd int, inStart int, inEnd int) *TreeNode
	inorderMap := make(map[int]int)

	for i,v := range inorder{
		inorderMap[v]= i
	}
	build = func(postStart int, postEnd int, inStart int, inEnd int) *TreeNode {
		if postStart > postEnd{
			return nil
		}
		node := &TreeNode{
			Val: postorder[postEnd],
		}
		index := inorderMap[postorder[postEnd]]

		leftNodeNum := index - inStart

		node.Left = build(postStart, postStart+leftNodeNum - 1, inStart, index- 1)
		node.Right = build(postStart+leftNodeNum, postEnd - 1, index+1, inEnd)

		return node
	}
	node := build(0, len(postorder)-1, 0, len(inorder)-1)
	return node

}