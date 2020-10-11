package constant_test

import "testing"


const (
	Monday = 1+iota  // 常量的连续赋值
	Tuesday
	Wednesy
)

const (
	Readable = 1 << iota
	Writeable
	Executeable
)


func TestConstant(t *testing.T){
	t.Log(Monday, Tuesday, Wednesy)
	a := 1
	t.Log(a&Readable, a&Writeable, a&Executeable)
	
}
