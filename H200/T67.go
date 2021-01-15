package H200

import "strconv"

func addBinary(a string,b string) string  {
	var ans = ""

	if len(a) < len(b){
		a,b = b,a
	}
	flag := 0
	lenAB := len(a) - len(b)
	if len(a) != len(b){
		for i := 0; i < lenAB;i++ {
			b = "0" + b
		}
	}
	for i := len(a) - 1;i>=0;i-- {
		aNum := a[i] - '0'
		bNum := b[i] - '0'
		sum := int(aNum + bNum) + flag

		if sum == 2{
			ans = "0" + ans
			flag = 1
		}else if sum == 1  || sum == 0{
			ans = strconv.Itoa(sum) + ans
			flag = 0
		}else if sum == 3{
			flag =1
			ans = "1" + ans
		}
	}
	if flag == 1{
		ans = "1" + ans
	}
	return ans
}
