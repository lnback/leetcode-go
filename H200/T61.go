package H200


func rotateRight(head *ListNode,k int) *ListNode {
	if head == nil || k == 0{
		return head
	}
	tail := head
	n := 1
	//先把链表构造成环 tail.next = head
	for tail.Next != nil {
		n++
		tail = tail.Next
	}
	tail.Next = head
	//通过计算 发现新链表的头结点是n-(k%n)
	k %= n
	//遍历到相应节点
	for i := 0; i < n-k;i++ {
		tail = tail.Next
	}
	//此时head变为tail.next即可
	head,tail.Next = tail.Next,nil
	return head

}
