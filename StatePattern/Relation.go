package StatePattern

import "fmt"

// 饱的状态
type FullState struct{}

func (f* FullState)NowStateDeal(){
	fmt.Println("当前是饱的状态")
}

// 指向当前状态的指针保存下一个状态
func (f* FullState)NextState(next *NowState){
	next.state = &HalfFullState{}
}

// 半饱状态
type HalfFullState struct {}

func (h* HalfFullState)NowStateDeal(){
	fmt.Println("当前是半饱的状态")
}

func (h* HalfFullState)NextState(next *NowState){
	next.state = &HungerState{}
}

// 饥饿状态
type HungerState struct{}

func (h* HungerState)NowStateDeal(){
	fmt.Println("当前是饥饿的状态")
}

func (h* HungerState)NextState(next *NowState){
	next.state = &FullState{}
}