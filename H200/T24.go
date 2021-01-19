package H200

func swapPairs(head * ListNode) *ListNode  {

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for head != nil && head.Next != nil {
		next := head.Next
		//更换
		head.Next = next.Next
		next.Next = head
		pre.Next = next

		//前进
		pre = head
		head = head.Next
	}
	return dummy.Next
}
