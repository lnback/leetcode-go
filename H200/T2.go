package H200

func addTwoNumbers(l1 *ListNode,l2 *ListNode)*ListNode  {
	head := new(ListNode)
	cur := head
	tmp := 0
	for l1 != nil || l2 != nil || tmp > 0{
		cur.Next = new(ListNode)
		cur = cur.Next
		if l1 != nil{
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			tmp += l2.Val
			l2 = l2.Next
		}
		cur.Val = tmp % 10
		tmp /= 10
	}

	return head.Next

}

