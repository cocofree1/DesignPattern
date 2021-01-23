package StatePattern

// 状态接口
type State interface {
	NowStateDeal()        // 当前状态处理
	NextState(next *NowState)          // 下一个状态
}

// 控制当前状态
type NowState struct{
	state State
}

func NewNowState()*NowState{
	return &NowState{&HungerState{}}
}

func (n* NowState)StateDeal(){
	n.state.NowStateDeal()
}

// 当前指针保存下一个状态
func (n* NowState)NextState(){
	n.state.NextState(n)
}