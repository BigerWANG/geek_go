package panic_test

import (
	"fmt"
	"testing"
)

/*
panic用于不可恢复的错误
panic退出前会执行defer指定的内容

panic 和 os.Exit的区别
os.Exit 退出时不会调用defer指定的函数
os.Exit 退出时不会输出当前调用栈的信息
 */

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("resource clean~~~~")
	}()
	//panic(errors.New("somthing error"))
	//os.Exit(-1)
}

func TestPanic(t *testing.T) {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("error")
	}()
}




