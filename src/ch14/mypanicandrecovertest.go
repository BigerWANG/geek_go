package main

import "time"

func main() {
	defer println("in main")  // 未被执行
	go func() { // main函数中的defer并没有被执行，panic 只会触发当前 Goroutine 的延迟函数调用
		defer println("in goroutine")  // 被执行
		panic("got panic")
	}()

	time.Sleep(1 * time.Second)
}