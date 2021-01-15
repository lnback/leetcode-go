package H200

func climbStairs(n int) int  {
	var a,b,c int
	a = 0
	b = 0
	c = 1
	for  i := 1;i<=n;i++{
		a = b
		b = c
		c = a + b
	}
	return c
}
