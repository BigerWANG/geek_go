package fib

func fib(num int) []int{
	// 实现一个斐波那契数列


	var res []int

	if num < 1 {
		res = append(res, 0)
		return res
	}
	if num <= 2 {
		return append(res, num)
	}


	return res

}