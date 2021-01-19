package main

import (
	"DesignPattern/FactoryPattern"
	"fmt"
)

func main(){
	var sub FactoryPattern.OperateFactory
	// sub = &FactoryPattern.AddOperateFactory{}
	sub = &FactoryPattern.SubOperateFactory{}
	operate := sub.Create()
	operate.SetLeft(10)
	operate.SetRight(20)
	fmt.Println(operate.Result())
}
