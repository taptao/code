package tree

func maxDepth(root *TreeNode) int {
    max := 0
    depth := 0

    var traverse func(node *TreeNode)
    traverse = func(node *TreeNode) {
    	if node == nil{
    		return
		}
    	if max < depth{
    		max = depth
		}

		depth ++
		traverse(node.Left)
    	traverse(node.Right)
    	depth --
	}
	depth++
	traverse(root)

	return max
}