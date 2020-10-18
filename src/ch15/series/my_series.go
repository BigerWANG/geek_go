package series

import (
	"errors"
	"fmt"
)
var lessThanZeroError = errors.New("lessThanZeroError ")
var largerThanHundredError = errors.New("largerThanHundredError ")

func init()  {
	fmt.Println("init1")
}

func init()  {
	fmt.Println("init2")
}

func GetFeb(n int) ([]int, error){
	if n < 0{  // 直接错误
		return nil, lessThanZeroError
	}
	if n>=100{
		return nil, largerThanHundredError
	}
	fibList := []int{1, 1}
	for i:=2; i < n; i++{
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList, nil
}

func Suqre(n int) int{
	return n * n
}