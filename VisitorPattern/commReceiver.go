package VisitorPattern

import "fmt"

// 只接收普通的访问
type CommReceiver struct{}

func (c* CommReceiver)Accept(visitor Visitor){
	switch inst := visitor.(type) {
	case *CommVisitor:
		fmt.Println("commReceiver accept",inst.name)
	}
}
