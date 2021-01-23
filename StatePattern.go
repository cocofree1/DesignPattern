package main

import "DesignPattern/StatePattern"

func main(){
	state := StatePattern.NewNowState()
	NowAndNext := func(){
		state.StateDeal()
		state.NextState()   // 当前指针状态指向下一个
	}
	for i := 0; i < 10; i++{
		NowAndNext()
	}
}
