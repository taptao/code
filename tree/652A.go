package tree

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	treeMap := make(map[string]int)
	treeRes := []*TreeNode{}

	var traverse func(node *TreeNode) string

	traverse = func(node *TreeNode) string {
		if node == nil{
			return "#"
		}
		leftTree := traverse(node.Left)
		rightTree := traverse(node.Right)

		nodeTree := leftTree + "," + rightTree + "," + strconv.Itoa(node.Val)

		if s,exist := treeMap[nodeTree]; exist {
			if s == 0 {
				treeRes = append(treeRes, node)
			}
			s++
			treeMap[nodeTree] = s
		}else{
			treeMap[nodeTree] = 0
		}

		return nodeTree
	}

	traverse(root)

	return treeRes
}
