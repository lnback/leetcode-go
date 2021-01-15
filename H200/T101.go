package H200

func isSymmetric(root *TreeNode) bool {
	return dfs(root,root)
}

func dfs(left ,right *TreeNode) bool {
	if left == nil && right == nil{
		return true
	}
	if left == nil || right == nil{
		return false
	}
	return left.Val == right.Val && dfs(left.Left,right.Right) && dfs(left.Right,right.Left)
}