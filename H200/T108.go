package H200

func sortArrayToBST(nums []int) *TreeNode {
	return solve(nums,0,len(nums)-1)

}

func solve(nums []int,left,right int) *TreeNode {
	if left > right{
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = solve(nums,left,mid-1)
	root.Right = solve(nums,mid + 1,right)

	return root
}
