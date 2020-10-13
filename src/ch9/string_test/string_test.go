package string_test

import "testing"



func TestString(t *testing.T) {

	/*
	字符串与其他主流编程语言的差异
	1， string是数据类型，不是引用或者数据类型
	2，string 是只读的byte slice, len 函数可以返回它所包含的byte数
	3，string 的byte可以存放任何数据
	*/
	var s string
	t.Log(s) // 默认值是一个空字符串

	s = "\xE4\xB8\xA1" // 可以存储任何二进制数据
	t.Log(s)

	s = "中"
	t.Log(len(s))  // 返回包含的byte个数

}

/*
Unicode 和 utf8的关系

1，unicode是一种字符集编码(code point)
2，utf8是unicode的存储实现(转换为字节序列的规则)，也就是定义Unicode怎么存储在内存空间上的存储规则

 */

func TestStrToRune(t *testing.T) {
	s := "中国"
	c := []rune(s)

	t.Logf("中 Unicode %x", c)
	t.Logf("中 UTF-8 %x", s)
}


