package H200

func lengthOfLongestSubstrings(s string) int  {
	ans := 0
	right := -1
	set := make(map[string]bool)
	for i := 0;i < len(s);i++{
		if i != 0 {
			set[string(s[i-1])] = false
		}
		for right + 1 < len(s) && !set[string(s[right+1])] {
			set[string(s[right+1])] = true
			right ++
		}
		ans = max(ans ,right-i + 1)
	}
	return ans
}

