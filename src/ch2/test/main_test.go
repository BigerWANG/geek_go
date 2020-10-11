package test

import "testing"

/*
编写测试程序
1， 源码文件以 xxx_test.go结尾

2，测试方法名称以 Test 开头： func Testxxx(t *testing.T) {.....}
 */


func TestHaha(t *testing.T){  // 默认参数
	t.Log("My first try!!")

}