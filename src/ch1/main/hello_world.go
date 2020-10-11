package main

import (
	"fmt"
	"os"
)

/*
应用程序入口
1, 必须是main包
2，必须是main方法
3，文件名不一定必须叫mian
 */
func main(){
	if len(os.Args) > 1{
		fmt.Println("fuck you bitch", os.Args[1])
	}
	os.Exit(255)
}

/*
退出返回值
main 方法不允许有返回值
需要使用os.Exit(number)指定返回值
 */