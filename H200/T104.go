package H200

func max(a,b int) int {
	if a > b{
		return a
	}
	return b
}
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return 1 + max(maxDepth(root.Left),maxDepth(root.Right))
}
