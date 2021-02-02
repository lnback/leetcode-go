package H200

func maxArea(height []int) int{
	ans := 0
	left , right := 0,len(height)-1

	for left < right {
		tmp := min(height[left],height[right]) * (right-left)

		ans = max(ans,tmp)

		if height[left] <= height[right]{
			left ++
		}else {
			right --
		}
	}
	return ans
}
