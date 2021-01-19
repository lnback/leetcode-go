package H200

func removeNthFromEnd(head *ListNode,n int) *ListNode  {
	ans := &ListNode{}
	ans.Next = head
	pre := &ListNode{}
	cur := ans
	i := 1
	//head先走n-1步，cur再走
	for head != nil {
		if i >= n{
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return ans.Next
}
