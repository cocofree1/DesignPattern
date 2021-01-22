package CompositePattern

import "fmt"

type CombinateNode struct{
	Common
	Childs []Composite  //叶子节点
}

func NewCombinateNode()*CombinateNode{
	return &CombinateNode{Childs: make([]Composite,0)}
}

// 重写 AddChild
func (c* CombinateNode)AddChild(child Composite){
	child.SetNode(child)
	c.Childs = append(c.Childs, child)
}

// 重写 Print
func (c* CombinateNode)Print(pre string){
	fmt.Println(pre,c.Name)
	pre += "   "
	for _,item := range c.Childs{
		item.Print(pre)
	}
}