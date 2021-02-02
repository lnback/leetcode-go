package H200

func deleteDuplicates2(head *ListNode)  *ListNode{
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val{
			for head.Val == head.Next.Val && head != nil && head.Next != nil {
				head = head.Next
			}
			pre.Next = head.Next
		}else{
			head = head.Next
			pre = pre.Next
		}

	}
	
	return dummy.Next

}
