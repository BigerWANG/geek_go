package type_test

import (
	"testing"
)

func TestPoint(t *testing.T){
	// 0xc4200162c0
	a := 1  // a 是值类型
	aPrt := &a  // aPrt 是指针类型 & 符号是取值符
	a += 1
	*aPrt += 1 // 对于指针类型的操作 需要在前边加*号
	t.Log(a, aPrt)
}