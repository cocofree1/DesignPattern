package main

import "DesignPattern/MediatorPattern"

func main(){
	mediator := MediatorPattern.NewMediator()
	cpu := MediatorPattern.NewCpu("test data")
	gpu := MediatorPattern.NewGpu("test data")
	mem := MediatorPattern.NewMem("test data")
	mediator.CPU = cpu
	mediator.GPU = gpu
	mediator.MEM = mem
	mediator.Change(cpu)
	mediator.Change(gpu)
	mediator.Change(mem)
}
