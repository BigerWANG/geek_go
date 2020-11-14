package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func firstResponse()  string{
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i:=0; i<numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <- ch // 只要channel中有数据，就可以立马返回
}


func Test(t *testing.T) {

	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(firstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After: ", runtime.NumGoroutine())

}
