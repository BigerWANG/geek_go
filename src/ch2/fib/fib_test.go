package fib

import (
	"fmt"
	"testing"
)


func TestFib(t *testing.T)  {

	res := fib(5)  // 同一个包中不同文件可以直接使用 不用导入

	fmt.Println(res)
}