package map_ext_test

import (
	"testing"
	"fmt"
)

func TestMapWithFunValue(t *testing.T) {  // 指定函数为value

	m := map[int]func(op int)int{}  // 定义一个匿名函数

	m[1] = func(op int) int{return op}
	m[2] = func(op int) int{return op*op}
	m[3] = func(op int) int{return op*op*op}

	t.Log(m[1](2), m[2](2), m[3](3))  // 通过key 访问map中对应的不同的函数
}


var set map[int]bool  // 声明一个空map

func mySet(value int)(bool){ // 使用map实现一个set
	set = make(map[int]bool, 10)  // 初始化这个map

	if set[value] {
		fmt.Printf("the value %d is exist", value)
		return false
	}else{
		set[value] = true
		return true
	}
}

func TestMapForSet(t *testing.T) {
	// 使用map实现set map[type]bool
	n:=1
	if ok:=mySet(n); !ok{
		t.Logf("%d is existing", n)
	}else{
		t.Logf("%d is not existing", n)
		t.Log("the set is ", set)
	}
}
