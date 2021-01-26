package VisitorPattern

// 不变的结构，增加新的功能，增加新的访问者和接受者
type Receiver interface {
	Accept(visitor Visitor)
}

type Visitor interface {
	Visit(receiver Receiver)
}