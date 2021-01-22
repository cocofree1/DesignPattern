package CompositePattern

type Composite interface {
	GetNode()Composite
	SetNode(node Composite)
	GetName()string
	SetName(name string)
	AddChild(child Composite)
	Print(pre string)
}
type Common struct{
	Node Composite
	Name   string
}

const (
	Node = iota
	Leaf
)

func NewCommonNode(kind int, name string)Composite{
	var common Composite
	switch kind {
	case Node:
		common = NewCombinateNode()
	case Leaf:
		common = NewLeafNode()
	}
	common.SetName(name)
	return common
}

func (c* Common)GetNode()Composite{
	return c.Node
}

func (c* Common)SetNode(node Composite){
	c.Node = node
}

func (c* Common)GetName()string{
	return c.Name
}

func (c* Common)SetName(name string){
	c.Name = name
}

func (c* Common)AddChild(child Composite){

}

func (c* Common)Print(pre string){

}

