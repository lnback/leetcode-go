package H200

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2

	ans := RemoveElement(nums,val)

	fmt.Println(ans)
}

func TestStrStr(t *testing.T)  {
	s1 := "aaaaa"
	s2 := "bbb"
	
	ans := strStr(s1,s2)
	fmt.Println(ans)
}

func TestFun(t *testing.T)  {

	nums1 := []int{1,2,3}
	nums2 := []int{2,3,4}
	findMedianSortedArrays(nums1,nums2)


}



