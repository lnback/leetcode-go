package H200

func maxSubArray(nums []int) int {
	ans := nums[0]
	tmp := 0
	max := func(a,b int) int{
		if a > b{
			return a
		}
		return b
	}
	for i := 0; i < len(nums) ;i++{
		if tmp > 0{
			tmp += nums[i]
		}else {
			tmp = nums[i]
		}
		ans = max(tmp,ans)
	}

	return ans
}
