package objectpoll

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// 创建一个对象池


type ReusableObj struct {

}

type ObjPool struct{
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int)  *ObjPool{
	ObjPool := ObjPool{}
	ObjPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i:=0; i<numOfObj; i++ {  // 创建一个对象池，将对象初始化好放入channel中
		ObjPool.bufChan <- &ReusableObj{}
	}
	return &ObjPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <- p.bufChan:  // 从对象池中获取一个已经初始化好的对象
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func(p *ObjPool) PutObj(obj *ReusableObj) error{
	select {
	case p.bufChan <- obj:  // 将已经用完的对象重新放入对象池子中
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i:=0; i<12;  i++{
		if v, err:=pool.GetObj(time.Second * 1); err!=nil{
			t.Error(err)
		}else {
			fmt.Printf("%T\n", v)
			if err := pool.PutObj(v); err!=nil{
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}




