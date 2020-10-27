package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "service Done"
}

func OtherTask() {
	fmt.Println("work on something else...")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Other work is done")
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		retCh <- ret // 往channel中放数据
		fmt.Println("service is exist...")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {

	OtherTask()
	select {
	case retCh := <-AsyncService():
		t.Log(retCh)
	case <-time.After(time.Millisecond * 500):
		t.Log("time out")
	}
}
