package H200

// 按K为一组旋转链表 分为三步
// 0.先设定一个前置节点 用来保存head 后面返回也是返回这个前置节点的next
// 1.先找到[n,n+k]范围内的链表节点 如果节点不够 直接返回hair.next 这时候要保存第[n+k]节点的next 后面需要接一下next
// 2.将[n,n+k]的头尾节点反转 参考反转链表的迭代写法
// 3.得到反转后的链表后，将其与后面的链表连接起来
func reverseKGroup(head *ListNode,k int) *ListNode {
	hair := &ListNode{Next: head}
	//前置节点
	pre := hair
	for head != nil{
		tail := pre
		for i := 0;i < k;i++{
			tail = tail.Next
			if tail == nil{
				return hair.Next
			}
		}
		//记录n+k的下一个节点 后驱节点
		next := tail.Next
		//反转链表
		head , tail = myReverse(head,tail)
		//连接前后节点
		pre.Next = head
		tail.Next = next
		//前置节点转换到尾结点
		pre = tail
		//头节点变为尾结点的下一个（也可以理解为下一组）
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head,tail *ListNode) (*ListNode,*ListNode)  {
	pre := tail.Next
	p := head
	for pre != tail {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return tail,head
}
