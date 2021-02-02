package H200

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0;i < len(nums) - 2;i++{
		n1 := nums[i]
		//因为是排好序的数组，所以此时如果第一个数大于0的话说明后面也没有负数了 也就不能三数之和为0了
		if n1 > 0{
			break
		}
		//避免重复相同的数
		if i > 0 && n1 == nums[i-1]{
			continue
		}
		left , right := i + 1,len(nums) - 1
		for left < right{
			n2,n3 := nums[left],nums[right]
			if n1 + n2 + n3 == 0{
				res = append(res,[]int{n1,n2,n3})
				//避免重复
				for left < right && nums[left] == n2 {
					left ++
				}
				for left < right && nums[right] == n3{
					right --
				}
			}else if n1 + n2 + n3 < 0{//说明数小了 左指针移动
				left ++
			}else{ //相加大了 右指针移动
				right --
			}
		}
	}

	return res

}
