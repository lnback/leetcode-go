package H200

func maxProfit(prices []int) int {

	n := len(prices)
	if n < 2{
		return 0
	}
	ans := 0
	min := prices[0]

	for _,v := range prices{
		if v < min{
			min = v
		}else if v - min > ans {
			ans = v - min
		}
	}
	return ans
}
