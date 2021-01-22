package CompositePattern

import "fmt"

type LeafNode struct{
	Common
}

func NewLeafNode()*LeafNode{
	return &LeafNode{}
}

// 重写print
func (l* LeafNode)Print(pre string){
	fmt.Println(pre,l.Name)
}
