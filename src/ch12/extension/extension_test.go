package extension

import (
	"fmt"
	"testing"
)

/*
继承
 一个结构体嵌套到另一个结构体，称作组合

匿名和组合的区别
如果一个struct A嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的方法，实现"继承"
如果一个 struct 嵌套了另一个 「有名字」的结构体，那么这个模式叫做 组合
如果一个struct嵌套了多个匿名结构体，那么这个结构体可以访问多个匿名结构体，从而实现多重继承
*/

type Car struct {
	weight    int
	name      string
	isRunning bool
}

func (c *Car)Run(){
	c.isRunning = true
	fmt.Println("car is running")
}

func (c *Car)Stop(){
	c.isRunning = false
	fmt.Println("car is stop")
}

type Bike struct {
	car Car  // 组合模式
	wheel int
}

type SuperCar struct {
	Car // 匿名模式
}
func(b *Bike)Run(){
	fmt.Println("bike is running")
}

func(t *SuperCar)Run(){
	t.Car.Run() // 直接调用Car中的方法
}

func TestS(t *testing.T) {
	var train SuperCar
	train.Run()
	train.Stop()
	t.Log(train.Car.isRunning)
}





