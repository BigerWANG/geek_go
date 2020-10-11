package array_test

import "testing"

func TestArrayInit(t *testing.T) {

	var array [3]int // 声明一个定长的数组

	array1 := [...]int{1,2,3,6} // 声明一个变长的数组, TODO： 这里跟python要区分，虽然是变长数组，定义好了之后也不能再变长了
	//array1[4] = 100  TODO 这个写法是错的

	array2 := [5]int{1,2,3}  // 这种方式，既初始化变量，又带了初始化值，数组长度已经定义好
	array3 := [...]int{1, 2, 3, 4, 5}  // 这种方式，既初始化变量，又带了初始化值，数组长度已经定义好



	t.Log(array[1], array[0])




	t.Log(array1)



	t.Log(array2[1], array2[0], array2[2])
	t.Log(array3)




	for idx, val := range array3{
		t.Log(idx, val)
	}

	// {1,2,3,4,5}
	// 数组截取 a[开始索引位置(包含): 结束索引位置(不包含)]
	// 不支持负数索引
	t.Log(array3[1:4])
	t.Log(array3[:4])


}


