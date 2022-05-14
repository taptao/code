package tree

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	
	var str []byte
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil{
			str = append(str, ",#"...)
			return
		}

		str = append(str, ","+strconv.Itoa(node.Val)...)

		traverse(node.Left)
		traverse(node.Right)
	}
	if root == nil{
		return ""
	}

	str = append(str, strconv.Itoa(root.Val)...)

	traverse(root.Left)
	traverse(root.Right)
	return string(str)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 0{
		return nil
	}
	str := strings.Split(data, ",")

	var traverse func(index int) (int, *TreeNode)

	traverse = func(index int) (int, *TreeNode) {
		if index >= len(str) || str[index] == "#"{
			return 1, nil
		}
		node := &TreeNode{}
		node.Val, _ = strconv.Atoi(str[index])
		var i,j int
		i, node.Left = traverse(index+1)
		j, node.Right = traverse(index+1+i)
		return i+j+1,node
	}


	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(str[0])
	var i int
	i, root.Left = traverse(1)
	_, root.Right = traverse(1+i)
	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */