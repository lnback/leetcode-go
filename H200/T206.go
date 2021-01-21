package H200

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	pre.Next = head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre

}
