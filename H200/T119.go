package H200

func getRow(rowIndex int) []int  {
	pre := []int{1}
	cur := []int{}

	for i := 1;i<= rowIndex;i++{
		m := []int{0}
		m = append(m,pre...)
		for j := 0;j < len(m)-1;j++ {
			m[j] = m[j] + m[j+1]
		}
		cur = m
		pre = cur
	}
	return cur
}
