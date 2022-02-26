package TwoPoint

func rotateRight(head *ListNode, k int) *ListNode {
	if k <= 0{
		return head
	}

	if head == nil || head.Next == nil{
		return head
	}

	length := 1
	first := head
	var node *ListNode
	for node=head;node.Next!=nil;node=node.Next{
		length++
	}

	node.Next = head
	breakNum := length - k % length

	for i:=0;i<breakNum-1;i++{
		first = first.Next
	}
	last := first.Next
	first.Next = nil
	return last
}
