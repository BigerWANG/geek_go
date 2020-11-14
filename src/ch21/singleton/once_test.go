package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"

)

type Single struct {

}


var singleInstance *Single

var once  sync.Once

func GetSingletonObj()  *Single{
	once.Do(func() {
		fmt.Println("create object")
		singleInstance = &Single{}

	})
	return singleInstance

}

func TestSingleTon(t *testing.T) {
	var wg sync.WaitGroup  // 等待所有的协程执行完毕
	for i:=0; i<=10; i++{
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
