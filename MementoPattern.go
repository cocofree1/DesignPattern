package main

import (
	"DesignPattern/MementoPattern"
	"fmt"
)

func main(){
	memento := MementoPattern.NewMemento("test code")
	memento.Write()
	fmt.Println("当前状态:",memento.GetState(),"当前代码:",memento.GetCode())
	memento.Commit()
	memento.Submit()
	fmt.Println("当前状态:",memento.GetState(),"当前代码:",memento.GetCode())
	// 打印所有历史状态
	for _,value := range MementoPattern.StateArray{
		fmt.Println("状态:",value.GetState(),"代码:",value.GetCode())
	}
}