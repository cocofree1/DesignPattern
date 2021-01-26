package InterpreterPattern

//节点
type Node interface {
	Interpreter()int
}

type ValNode struct {
	val int
}

func (v* ValNode)Interpreter()int{
	return v.val
}