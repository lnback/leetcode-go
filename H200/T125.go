package H200

import (
	"strings"
)

func isAorNum(c byte) bool  {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c<='Z') || (c >= '0' && c<='9')
}
func isPalindrome2(s string) bool {
	s = strings.ToLower(s)

	for i,j := 0,len(s) -1 ; i < j;{
		if !isAorNum(s[i]){
			i++
			continue
		}
		if !isAorNum(s[j]){
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}