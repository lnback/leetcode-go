package H200

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode,l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	pre := &ListNode{}
	ans := pre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val{
			pre.Next = l1
			l1 = l1.Next
		}else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil{
		pre.Next = l1
	}else {
		pre.Next = l2
	}
	return ans.Next
}
