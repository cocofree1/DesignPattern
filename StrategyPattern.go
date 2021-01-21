package main

import (
	"DesignPattern/StrategyPattern"
	"fmt"
)

func main(){
	strategy1 := StrategyPattern.NewSelectContent("大好",&StrategyPattern.Bicycle{})
	strategy2 := StrategyPattern.NewSelectContent("一般",&StrategyPattern.Car{})
	fmt.Println(strategy1.Travel())
	fmt.Println(strategy2.Travel())
}
