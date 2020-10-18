package package_test

import "testing"
import "geek_go/src/ch15/series" // 在同一个gopath下可以写相对路径直接引入

func TestPackage(t *testing.T) {
	t.Log(series.GetFeb(20))
	t.Log(series.Suqre(100))
}


