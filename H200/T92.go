package H200

func reverseBetween(head * ListNode,m int,n int) *ListNode {
	if m == 1{
		return reverseN(head,n)
	}
	head.Next = reverseBetween(head.Next,m-1,n-1)
	return head
}
var tmp *ListNode
func reverseN(head *ListNode,n int) *ListNode {
	if n == 1{
		tmp = head.Next
		return head
	}
	p := reverseN(head.Next,n-1)
	head.Next.Next = head
	head.Next = tmp
	return p
}
