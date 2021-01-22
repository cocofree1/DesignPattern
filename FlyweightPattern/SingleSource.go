package FlyweightPattern

import "fmt"

// 资源处理

type SingleSource struct{
	data string
}

func NewSingleSource(name string)*SingleSource{
	data := fmt.Sprintf("source name: %s",name)
	return &SingleSource{data}
}

func (s* SingleSource)GetData()string{
	return s.data
}