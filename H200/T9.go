package H200


func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0  && x != 0){
		return false
	}
	tmp := 0
	for x > tmp {
		tmp = tmp * 10 + x % 10
		x /= 10
	}

	return tmp == x || tmp / 10 == x
}


