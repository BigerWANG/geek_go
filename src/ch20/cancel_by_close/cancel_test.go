package cancel_by_close

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{})  {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{})  {
	close(cancelChan)
}

func TestCancelled(t *testing.T) {

	cancelChan := make(chan struct{}, 0)
	var wg sync.WaitGroup
	for i:=0; i<5; i++{
		wg.Add(1)
		go func(i int, cancel chan struct{}) {
			for {
				if isCancelled(cancel){
					break
				}
				wg.Done()
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)

	}
	//cancel_1(cancelChan)
	cancel_2(cancelChan)
	wg.Wait()
}

func TestWaiteGroup(t *testing.T) {
	/*
	WaiteGroup内部有一个计数器，最初从0开始
	它有三个方法：
	Add()将计数器设置为n
	Done()每次将计数器-1
	Waite()主协程阻塞等待直到计数器为0
	WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址
	*/
	var wg sync.WaitGroup
	for i:=0; i<10; i++{
		wg.Add(1) // 开启协程之前+1
		go func(i int) {
			fmt.Printf(strconv.Itoa(i))
			time.Sleep(time.Millisecond * 1000)
			wg.Done()  // 协程完成后 -1
		}(i)
	}
	wg.Wait() // 直到所有协程中的
}
