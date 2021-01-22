package main

import "DesignPattern/ChainPattern"

// 责任链模式:链表操作
// projectManager: <1000
// depManager: <5000
// boss: all
func main(){
	c1 := ChainPattern.NewProjectManagerChain()
	c2 := ChainPattern.NewDepManagerChain()
	c3 := ChainPattern.NewBossChain()

	c1.SetNextChain(c2)
	c2.SetNextChain(c3)

	c1.DealChain("lqh",100)
	c1.DealChain("lqh",1500)
	c1.DealChain("lqh",6000)
}
