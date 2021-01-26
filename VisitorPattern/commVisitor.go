package VisitorPattern

type CommVisitor struct{
	name string
}

func NewCommVisitor(name string)*CommVisitor{
	return &CommVisitor{name: name}
}

func (c* CommVisitor)Visit(receiver Receiver){
	receiver.Accept(c)
}