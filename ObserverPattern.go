package main

import "DesignPattern/ObserverPattern"

func main(){
	// 创建观察者
	o := ObserverPattern.NewObserver()
	// 创建被观察者
	r1 := ObserverPattern.NewReader("kk1")
	r2 := ObserverPattern.NewReader("kk2")
	r3 := ObserverPattern.NewReader("kk3")
	// 添加到观察队列
	o.AddPerson(r1)
	o.AddPerson(r2)
	o.AddPerson(r3)
	// 通知消息
	o.UpdateMessage("班主任来了")

}
