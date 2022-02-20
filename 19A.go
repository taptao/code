package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeNode := new(ListNode)
	fakeNode.Next = head
	start := fakeNode
	end := head
	for i:=1;i<n;i++{
		end = end.Next
	}
	for end.Next != nil{
		end = end.Next
		start = start.Next
	}

	start.Next = start.Next.Next
	return  fakeNode.Next
}
