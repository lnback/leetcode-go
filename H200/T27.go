package H200

func RemoveElement(nums []int,val int) int {
	len1 := len(nums)
	if len1 == 0{
		return 0
	}
	i := 0
	j := 0
	for j < len1 {
		if nums[j] == val{
			j++
		}else {
			nums[i] = nums[j]
			i++
			j++
		}
	}

	return len1 - (j - i)
}
