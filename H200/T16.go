package H200

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int,target int) int {
	sort.Ints(nums)
	
	ans := math.MaxInt32

	for i := 0; i < len(nums) - 2;i++ {
		n1 := nums[i]

		left,right := i + 1,len(nums)-1

		for left < right {
			 n2 , n3 := nums[left] ,nums[right]
			 tmp := n1 + n2 + n3
			 if tmp > target{
			 	right --
			 }else{
			 	left ++
			 }
			 if Abs(tmp-target) < Abs(ans-target){
			 	ans = tmp
			 }
		}
	}
	
	return ans
}

func Abs(a int) int {
	if a < 0{
		return -a
	}
	return a
}
