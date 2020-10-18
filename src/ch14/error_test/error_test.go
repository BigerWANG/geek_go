package error_test

import (
	"errors"
	"testing"
)

var lessThanZeroError = errors.New("lessThanZeroError ")
var largerThanHundredError = errors.New("largerThanHundredError ")

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

func TestFeb(t *testing.T) {
	v, err := GetFeb(50)
	if err == largerThanHundredError{
		t.Error(err)
		return

	}
	if err == lessThanZeroError{
		t.Error(err)
		return
	}
	t.Log(v)
}
