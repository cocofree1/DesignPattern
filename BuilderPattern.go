package main

import "DesignPattern/BuilderPattern"

//建造者模式

func main(){
	//builder := &BuilderPattern.MealCombo1{}
	builder := &BuilderPattern.MealCombo2{}
	combo := BuilderPattern.NewMealCombo(builder)
	combo.ComboShow()
}