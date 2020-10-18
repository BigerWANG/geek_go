package _interface

import "fmt"

/*
Go接口最佳实践
 */


// 1 倾向于使用小的接口定义

type Reader interface {
	Read(p []byte)(n int, err error)
}

type Writer interface {
	Write(p []byte)(n int, err error)
}

// 2 较大的接口可以有多个小接口定义组合而成
type ReadWrite interface {
	Reader
	Writer
}


// 3 接口实现只依赖于必要功能的最小接口(接口隔离原则)
type MyStuct struct {
}

//
func (s *MyStuct) Read(myp []byte)(n int, err error){
	fmt.Println(myp)
	return 0, fmt.Errorf("nothing to do")
}

func SortData(r Reader)  error{

	
}



