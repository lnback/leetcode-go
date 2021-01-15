package H200

func isBalanced(root *TreeNode) bool  {
	if root == nil{
		return true
	}

	return abs(MaxDepth(root.Left) , MaxDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func abs(a,b int) int  {
	if a > b{
		return a-b
	}
	return b-a
}
func Max(a,b int) int {
	if a > b{
		return a
	}
	return b
}
func MaxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return 1 + Max(maxDepth(root.Left),MaxDepth(root.Right))
}