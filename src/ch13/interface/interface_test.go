package _interface

import (
	"fmt"
	"testing"
)

// 接口定义
/*
与其他主要编程语言的差异
1, 接口为非入侵性，实现不依赖于接口定义
2，所以接口的定义可以包含在接口使用者包内
*/

type Programmer interface {
	WriteHelloWorld(string string)string
}


// 接口实现不依赖与接口的定义
type GoProgrammer struct {  // 并不需要显式的指定实现某个接口

	Title func(string string)string
}

// 只需要实现与接口指定的同名的方法就实现了这个接口
// 接口实现不依赖与接口的定义
func (p *GoProgrammer) WriteHelloWorld(op string)string{
	return fmt.Sprintf("hello world :%s", p.Title(op))
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = &GoProgrammer{Title: func(input string) string {
		return fmt.Sprintf("input string is %s", input)
	}} // 相当于new方法，接口的具体实现

	//p = new(GoProgrammer)

	ret := p.WriteHelloWorld("haha")
	t.Log(ret)
}
