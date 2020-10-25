package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	count := 0
	for i:=0; i<=500; i++{
		go func() {
			count ++ // 此处不是线程安全的操作，在并发执行中会出错
		}()
	}
	t.Log(count)
}

func TestCounterThreadSafe(t *testing.T) {
	count := 0
	var mut sync.Mutex
	for i:=0; i<5000; i++{
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++

		}()
	}
	time.Sleep(time.Second * 1) // sleep是为了主线程等待子协程执行完毕，因为goroutine是默认主协程执行完毕直接退出
	t.Log(count)
}

func TestWiteGroup(t *testing.T) {
	//使用witeGroup等待子协程执行完毕，相当于Python中的join
	count := 0
	var mut sync.Mutex  // 声明互斥锁
	var wg sync.WaitGroup  // 声明一个wait group

	for i:=0; i<5000; i++{
		wg.Add(1)  // 在开启协程之前 将计数器加一
		go func() {
			defer mut.Unlock()
			mut.Lock()
			count ++
			wg.Done()  // 执行完成，释放wg计数器

		}()
	}

	wg.Wait()   //如果计数器变为零，则释放等待时阻塞的所有goroutine
	t.Log(count)

	
}


