package H200
//分治法：先获取链表的长度O(n)，再用数组的做法将其分开建树
var globalHead *ListNode
func sortedListToBST(head *ListNode) *TreeNode  {
	globalHead = head
	length := getLength(head)
	return buildTree(0,length-1)

}

func getLength(head *ListNode) int {
	res := 0

	for ;head != nil;head = head.Next {
		res ++
	}
	return res
}
func buildTree(left,right int) *TreeNode {
	if left > right{
		return nil
	}
	mid := (left + right + 1)/2
	root := &TreeNode{}
	root.Left = buildTree(left,mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildTree(mid +1,right)
	return root
}
