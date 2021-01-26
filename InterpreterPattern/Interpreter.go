package InterpreterPattern

import (
	"strconv"
	"strings"
)

type Interpreter struct{
	exp []string
	index int
	Pre Node
}

func (i* Interpreter)Parse(str string){
	i.exp = strings.Split(str," ")
	i.index = 0
	for i.index < len(i.exp){
		switch i.exp[i.index] {
		case "+":
			i.Pre = i.NewAddNode()
		case "-":
			i.Pre = i.NewSubNode()
		default:
			i.Pre = i.NewValNode()
		}
	}
}

func (i* Interpreter)NewValNode()Node{
	num,_ := strconv.Atoi(i.exp[i.index])
	i.index++
	return &ValNode{val: num}
}

func (i* Interpreter)NewAddNode()Node{
	i.index++
	return &AddNode{left: i.Pre,right:i.NewValNode()}
}

func (i* Interpreter)NewSubNode()Node{
	i.index++
	return &SubNode{left: i.Pre,right: i.NewValNode()}
}

func (i* Interpreter)Result()Node{
	return i.Pre
}