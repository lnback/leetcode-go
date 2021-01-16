package H200

import "fmt"

func singleNumber(nums []int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	fmt.Println(ans)

	return ans
}