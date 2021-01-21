package main

import "fmt"

// 目标接口
type Target interface {
	Request(a int,b int)string
}

type TargetInstance struct {
	Special
}

func (t *TargetInstance) Request(a int,b int)string{
	return t.SpecialRequest(a,b)
}

// 接口适配
func NewTargetInstance(special Special) Target{
	return &TargetInstance{special}
}

// 特殊接口
type Special interface {
	SpecialRequest(a int,b int)string
}

type SpecialInstance struct {}

func (s *SpecialInstance)SpecialRequest(a int,b int)string{
	fmt.Println(a,b)
	return "this is special request"
}

func NewSpacialInstance()Special{
	return &SpecialInstance{}
}

func main(){
	special := NewSpacialInstance()
	target := NewTargetInstance(special)
	fmt.Println(target.Request(1,2))
}