package MediatorPattern

import "fmt"

type Mediator struct{
	CPU *CPU
	GPU *GPU
	MEM *MEM
}

var mediator *Mediator

func NewMediator()*Mediator{
	if mediator == nil{
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator)Change(data interface{}){
	switch inst := data.(type) {
	case *CPU:
		m.CPU.Process(inst.data)
	case *GPU:
		m.GPU.Calc(inst.data)
	case *MEM:
		m.MEM.Save(inst.data)
	default:
		fmt.Println("暂时不能处理")
	}
}