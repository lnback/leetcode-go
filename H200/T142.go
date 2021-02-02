package H200

func detectCycle(head *ListNode) *ListNode  {
	slow := head
	fast := head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil{
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow{
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
