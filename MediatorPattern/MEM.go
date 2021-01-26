package MediatorPattern

import "fmt"

type MEM struct{
	data string
}

func NewMem(data string)*MEM{
	return &MEM{data:data}
}

func (m* MEM)Save(data string){
	fmt.Println("MEM is save ",data)
}
