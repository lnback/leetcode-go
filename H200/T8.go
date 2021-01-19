package H200

import "math"

func myAtoi(s string) int {
	flag := false
	var ans int = 0

	for i:=0 ; i<len(s) ; i++ {
		ch := string(s[i])
		if ch == "-" || ch == "+" || (ch >= "0" && ch <= "9"){
			flag = true
		}
	}

	if ans > math.MaxInt32{
		return math.MaxInt32
	}
	if ans < math.MinInt32{
		return math.MinInt32
	}

	return ans
}

