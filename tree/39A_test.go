package tree

import (
	"testing"
)


var  A39 = []struct {
	nums []int
	target int
}{
	{[]int{2,3,5}, 8},
}

func TestWombinatinSum(t *testing.T){
	for _, testItem := range A39 {
		combinationSum(testItem.nums, testItem.target)
	}
}


var A105 = []struct{
	preOrder []int
	inOrder []int
	target Node
}{
	{[]int{3,9,20,15,7}, []int{9,3,15,20,7}, Node{}},
}

func Test39A(t *testing.T){
	for _, testItem := range A105{
		buildTree(testItem.preOrder, testItem.inOrder)
	}
}