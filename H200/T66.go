package H200

func plusOne(digits []int) []int{
	tmp := 1
	for i := len(digits) - 1; i>= 0;i--{
		digits[i] += tmp
		if digits[i] == 10 {
			digits[i] = 0
			tmp = 1
		}else {
			break
		}
		if i == 0 && tmp == 1{
			digits = append([]int{1},digits...)
		}
	}
	return digits
}
