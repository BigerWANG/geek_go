package string_fun_test

import (
	"strconv" // 字符串转换函数
	"strings"  // 字符串内置方法库
	"testing"
)

func TestStrFun(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")  // 字符串分隔
	t.Log(parts)
	str := strings.Join(parts, "|")  //字符串连接
	t.Log(str)
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(100) // int 转 str
	t.Log("str: "+s)

	s1 := "100"
	if ints, ok := strconv.Atoi(s1); ok==nil{
		t.Logf("this ok to convert: %d", ints)
	}
	
}


