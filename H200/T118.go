package H200

func generate(numRows int) [][]int {
	if numRows == 0{
		return [][]int{}
	}
	ans := [][]int{{1}}

	for i := 1;i<numRows ;i++{
		m := []int{0}
		m = append(m,ans[i-1]...)
		for j:=0;j<len(m)-1;j++ {
			m[j] = m[j] + m[j+1]
		}
		ans = append(ans,m)
	}

	return ans

}
