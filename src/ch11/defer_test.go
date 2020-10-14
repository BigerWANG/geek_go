package ch11

import (
	"testing"
	"fmt"
)


func Clear(){
	fmt.Println("clear resource...")
}

func TestDefer(t *testing.T) {

	// 延迟执行, 函数执行完成后再执行defer语句
	defer Clear()
	fmt.Printf("start...\n")
	panic("error") // defer 会在panic后边执行

}

