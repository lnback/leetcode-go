package H200

func addTwoNumbers(l1 *ListNode,l2 *ListNode) *ListNode  {
	ans := &ListNode{}
	pre := ans
	curA,curB := l1,l2
	tmp := 0
	for curA != nil && curB != nil{
		val := curA.Val + curB.Val + tmp
		if val >= 10{
			tmp = 1
			val = val - 10
		}
		cur := &ListNode{
			Val: val,
		}
		pre.Next = cur
		pre = ans.Next
	}
	if curA != nil{
		
	}
	
	return ans.Next

}