package ch12

import (
	"fmt"
	"testing"
)

// 定义一个封装结构体类型
// 定义一个结构体对象
type Employee struct {
	Id string
	Name string
	Age int
}


// 这种实例方法定义方式在实例对应方法被调用时，实例的成员变量会进行值复制
// 此时会新建一个Employee实例
func (e Employee) String() string{
	return fmt.Sprintf("ID: %s-Name: %s-Age: %d", e.Id, e.Name, e.Age)
}


// 通常情况下为了避免内存拷贝，使用这种方式定义，定义一个指针变量，直接指向结构体的内存地址，避免值复制
// 在这个方法中修改也会影响
func (e *Employee) String1() string{
	return fmt.Sprintf("ID: %s-Name: %s-Age: %d", e.Id, e.Name, e.Age)
}

func TestStruct(t *testing.T) {
	e := Employee{"100", "wangmengyu04", 100}  // 实例化对象
	t.Log(e.String1())
	t.Log(e.String())
}
