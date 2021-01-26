package VisitorPattern

import "fmt"

// 只接受boss访问
type BossReceiver struct{}

func (b* BossReceiver)Accept(visitor Visitor){
	switch inst:= visitor.(type) {
	case *BossVisitor:
		fmt.Println("bossReceiver accept",inst.name)
	}
}
