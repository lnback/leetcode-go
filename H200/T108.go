package H200

func sortArrayToBST(nums []int) *TreeNode {
	return solve1(nums,0,len(nums)-1)

}

func solve1(nums []int,left,right int) *TreeNode {
	if left > right{
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = solve1(nums,left,mid-1)
	root.Right = solve1(nums,mid + 1,right)

	return root
}
