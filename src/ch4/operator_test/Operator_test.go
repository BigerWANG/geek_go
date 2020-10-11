package operator_test

import "testing"


const (
	Readable = 1 << iota
	Writeable
	Executeable
)


func TestOperator(t *testing.T)  {

	/*
	用== 比较数组

	相同维度数且含有相同个数元素的数组才可以比较
	每个元素都相同时，两个数组才相等
	 */

	a := [...]int{1,2,3,4}  //数组的初始化
	b := [...]int{1,2,3,5}
	c := [...]int{1,2,3,4}
	if a ==b || a==c{
		t.Log("hahah")
	}else{
		t.Error("hehehe")
	}

}
