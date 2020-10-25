package csp

import (
	"fmt"
	"testing"
	"time"
)

func service()  string {
	time.Sleep(time.Millisecond * 500)
	return "service Done"
}


func OtherTask(){
	fmt.Println("work on something else...")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Other work is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	OtherTask()
	
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
 		retCh <- ret  // 往channel中放数据
		fmt.Println("service is exist...")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	OtherTask()
	fmt.Println(<-retCh)  // 从channel中取数据
}
