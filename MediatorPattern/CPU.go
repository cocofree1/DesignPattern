package MediatorPattern

import "fmt"

type CPU struct{
	data string
}

func NewCpu(data string)*CPU{
	return &CPU{data:data}
}

func (c* CPU)Process(data string){
	fmt.Println("CPU is process ",data)
}
