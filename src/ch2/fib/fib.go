package fib

func fib(num int) int{
	// 实现一个斐波那契数列


	var res int

	if num < 1 {
		return 1
	}
	if num <= 2 {
		return num
	}

	res = fib(num-1) + fib(num-2)

	return res

}