package until_all_done


import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func allResponse()  string{
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i:=0; i<numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	var finalRet string
	for j:=0; j<numOfRunner; j++ {
		finalRet += <- ch + "\n"
	}
	return finalRet
}

func allResponseWithWaiteGroup() string{

	// 使用waitGroup等待所有任务完成

	numOfRunner := 10
	var wg sync.WaitGroup

	ch := make(chan string, numOfRunner)
	for i:=0; i<numOfRunner; i++ {
		wg.Add(1)
		go func(i int) {
			ret := runTask(i)
			fmt.Println(ret)
			ch <- ret
			wg.Done()
		}(i)
	}
	wg.Wait()
	for len(ch)!=0{
		fmt.Println(<-ch)
	}
	fmt.Println("channel len>>>", len(ch))
	close(ch)
	return "all task done..."
}



func Test(t *testing.T) {

	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(allResponseWithWaiteGroup())
	time.Sleep(time.Second * 1)
	t.Log("After: ", runtime.NumGoroutine())

}
