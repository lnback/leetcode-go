package H200

import "strings"

func convert(s string,numRows int) string  {
	if numRows == 1{
		return s
	}

	rows := make([]string,numRows)

	n := 2 * numRows - 2

	for i,c := range s{
		x := i % n
		rows[min(x,n-x)] += string(c)
	}
	return strings.Join(rows,"")

}
