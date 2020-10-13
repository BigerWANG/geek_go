package _func

import (
	"fmt"
	"testing"
	"time"
)

// golang 的装饰器
func TimeSpent(inner func(op int) int)  func(op int) int {  // 接收一个函数作为参数，
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("Time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func dosth(op int) int{
	fmt.Println("I'm start working...")
	time.Sleep(time.Second * 1)
	fmt.Println("I'm done")
	return op
}

func TestFunc(t *testing.T) {

	ret := TimeSpent(dosth)  // 返回被装饰后的函数
	t.Log(ret(100))

}
