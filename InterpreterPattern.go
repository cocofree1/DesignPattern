package main

import (
	"DesignPattern/InterpreterPattern"
	"fmt"
)

// 解释器模式，将一种语法按一定的规则解释成另一种语法
func main(){
	str := "1 + 4 - 3"
	fmt.Println(str)
	interpreter := &InterpreterPattern.Interpreter{}
	interpreter.Parse(str)
	fmt.Println(interpreter.Result().Interpreter())
	// interface套interface，执行实现的函数，会递归执行
}

