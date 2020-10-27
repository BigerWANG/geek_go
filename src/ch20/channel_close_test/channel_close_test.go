package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

/*
channel的关闭

向关闭的channel发送数据，会导致panic
v, ok <- ch; ok 为bool值，true表示正常接受，false表示通道关闭

所有的channel接受者都会在channel关闭时，立刻从阻塞等待中返回且channel的ok值为false
这个广播机制常被利用，进行向多个订阅者同时发送信号，比如：退出信号
*/

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 4; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReveicer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReveicer(ch, &wg)

	wg.Wait()
}
