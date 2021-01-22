package main

import "DesignPattern/CompositePattern"

func main(){
	root := CompositePattern.NewCommonNode(CompositePattern.Node,"root")
	r1 := CompositePattern.NewCommonNode(CompositePattern.Node,"r1")
	r2 := CompositePattern.NewCommonNode(CompositePattern.Node,"r2")
	r3 := CompositePattern.NewCommonNode(CompositePattern.Node,"r3")

	l1 := CompositePattern.NewCommonNode(CompositePattern.Leaf,"l1")
	l2 := CompositePattern.NewCommonNode(CompositePattern.Leaf,"l2")
	l3 := CompositePattern.NewCommonNode(CompositePattern.Leaf,"l3")

	root.AddChild(r1)
	root.AddChild(r2)
	root.AddChild(r3)
	r1.AddChild(l1)
	r2.AddChild(l2)
	r3.AddChild(l3)
	root.Print("")
}
