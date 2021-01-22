package main

import (
	"DesignPattern/FlyweightPattern"
	"fmt"
)

// 享元模式
func main(){
	source := FlyweightPattern.NewShowSource("yaya")
	source.Display()
	source1 := FlyweightPattern.NewShowSource("yaya")
	if source.SingleSource == source1.SingleSource{
		fmt.Println("节约资源")
	} else {
		fmt.Println("浪费资源")
	}
}
