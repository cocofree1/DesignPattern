package InterpreterPattern

type SubNode struct{
	left,right Node
}

func (s* SubNode)Interpreter()int{
	return s.left.Interpreter() - s.right.Interpreter()
}
