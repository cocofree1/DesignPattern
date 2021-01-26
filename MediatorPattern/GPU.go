package MediatorPattern

import "fmt"

type GPU struct{
	data string
}

func NewGpu(data string)*GPU{
	return &GPU{data:data}
}

func (g* GPU)Calc(data string){
	fmt.Println("GPU is calc ",data)
}
