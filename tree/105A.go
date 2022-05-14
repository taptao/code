package tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(preStart int, preEnd int, inStart int, inEnd int) *TreeNode
	inorderMap := make(map[int]int)

	for i,v := range inorder{
		inorderMap[v]= i
	}
	build = func(preStart int, preEnd int, inStart int, inEnd int) *TreeNode {
		if preStart > preEnd{
			return nil
		}
		node := &TreeNode{
			Val: preorder[preStart],
		}
		index := inorderMap[preorder[preStart]]

		leftNodeNum := index - inStart

		node.Left = build(preStart+1, preStart+leftNodeNum, inStart, index- 1)
		node.Right = build(preStart+leftNodeNum+1, preEnd, index+1, inEnd)

		return node
	}
	node := build(0, len(preorder)-1, 0, len(inorder)-1)
	return node

}
