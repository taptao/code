package tree

/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Left *Node
*     Right *Node
*     Next *Node
* }
*/

func connect(root *Node) *Node {


	var c func(lNode *Node, rNode *Node)

	c = func(lNode *Node, rNode *Node) {

		if lNode==nil || rNode ==nil{
			return
		}

		lNode.Next = rNode

		c(lNode.Left, lNode.Right)
		c(lNode.Right, rNode.Left)
		c(rNode.Left, rNode.Right)
	}
	if root != nil{
		c(root.Left, root.Right)
	}

	return root
}