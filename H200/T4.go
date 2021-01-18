package H200

import "fmt"

func findMedianSortedArrays(nums1 []int,nums2 []int) float64  {
	l1 := len(nums1)
	l2 := len(nums2)

	mid := (l1 + l2) / 2
	cnt := 0
	nums3 := make([]int,l1 + l2 - 1)
	i,j := 0,0
	for i < l1 && j < l2{
		if nums1[i] < nums2[j]{
			nums3[cnt] = nums1[i]
			i++
			cnt++
		}else {
			nums3[cnt] = nums2[j]
			j++
			cnt++
		}
	}
	if i != l1{
		nums3 = append(nums3,nums1[l1-i])
	}
	if j != l2{
		nums3 = append(nums3,nums2[l2-j])
	}
	fmt.Println(nums3)
	if (l1 + l2) % 2 != 0{
		return float64((nums3[mid] + nums3[mid+1])/2)
	}else{
		return float64(nums3[mid])
	}
}
