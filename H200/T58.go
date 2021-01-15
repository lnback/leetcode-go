package H200

func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1;i >= 0;i-- {
		ch := string(s[i])
		if ch != " "{
			ans ++
		}else if ans > 0{
			break
		}
	}
	return ans
}
