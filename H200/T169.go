package H200

func majorityElement(nums []int) int{
	cnt := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0{
			ans = nums[i]
		}
		if ans == nums[i]{
			cnt++
		}else{
			cnt--
		}
	}
	return ans
}