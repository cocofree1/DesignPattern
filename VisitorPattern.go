package main

import "DesignPattern/VisitorPattern"

// 访问者模式，不同的访问者不同的接待者去处理
func main(){
	// 接待者
	var receiver1 VisitorPattern.Receiver = &VisitorPattern.CommReceiver{}
	var receiver2 VisitorPattern.Receiver = &VisitorPattern.BossReceiver{}

	// 拜访者
	contrl := VisitorPattern.Visitors{}
	contrl.Add(VisitorPattern.NewBossVisitor("boss1"))
	contrl.Add(VisitorPattern.NewCommVisitor("comm1"))
	contrl.Add(VisitorPattern.NewBossVisitor("boss2"))
	contrl.Add(VisitorPattern.NewCommVisitor("comm2"))
	contrl.Visit(receiver1)
	contrl.Visit(receiver2)
}


