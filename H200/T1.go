package H200


func twoSum(nums []int,target int) []int  {
	maps := make(map[int]int)
	ans := []int{}
	for i,v := range nums{
		if value,exist := maps[target-v];exist{
			ans = append(ans,value)
			ans = append(ans,i)
		}
		maps[v] = i
	}
	return ans
}
