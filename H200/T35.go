package H200

func searchInsert(nums []int,target int) int {
	left := 0
	right := len(nums)

	for left < right{
		mid := (left + right) / 2
		if nums[mid] >= target{
			right = mid
		}else {
			left = mid + 1
		}
	}
	return left
}
