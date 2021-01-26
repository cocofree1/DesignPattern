package InterpreterPattern

type AddNode struct {
	left,right Node
}

func (a* AddNode)Interpreter()int{
	return a.left.Interpreter() + a.right.Interpreter()
}
