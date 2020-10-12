package map_test

import (
	"testing"
)

func TestMap(t *testing.T) {
	// map 声明 方式
	m := map[string]int{"one": 1, "two": 2}
	t.Log(m)

	m1 := map[string]int{}
	m1["one"] = 1
	//m1["three"] = 0

	m2 := make(map[string]int, 10) // 使用make 初始化，指定Capcaity
	/*
	TODO:为什么使用make创建map 不需要指定len？
	因为指定len初始化的时候 golang会指定零值(默认值)
	但是创建map时无法指定 k,v

	TODO:有什么好处？
	指定固定容量，优化性能
	*/

	m2["three"] = 3
	m2["four"] = 4
	m2["five"] = 5
	t.Log(m2["three"])
	t.Log(m2["four"])
	t.Log(m2["five"])
	t.Log(len(m2))
	t.Log("six", m2["six"])  // 访问不存在的key时 会返回一个默认零值

	t.Log("m1['three'] =", m1["three"])
	v, ok := m1["three"]

	t.Log(ok)
	t.Log("m1['three'] =", v)

	if v, ok := m1["three"]; ok{  // 使用此方法判断是否存在某个key, 如果存在 ok 为 true
		t.Log("exist key three")
		t.Log("value is ", v)
	}else{
		t.Log("key three is not exist")
	}

}

func TestTravelMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	for k, v := range m{
		t.Log(k, v)
	}
}


