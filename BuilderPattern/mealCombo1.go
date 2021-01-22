package BuilderPattern

import "fmt"

// 套餐1
type MealCombo1 struct{}

func (m* MealCombo1)Part1(){
	fmt.Println("牛肉汉堡")
}

func (m* MealCombo1)Part2(){
	fmt.Println("可口可乐")
}