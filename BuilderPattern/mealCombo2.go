package BuilderPattern

import "fmt"

// 套餐2
type MealCombo2 struct{}

func (m* MealCombo2)Part1(){
	fmt.Println("鸡肉汉堡")
}

func (m* MealCombo2)Part2(){
	fmt.Println("百事可乐")
}