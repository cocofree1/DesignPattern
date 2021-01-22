package FlyweightPattern

import "fmt"

// 资源展示

type ShowSource struct {
	*SingleSource
}

func NewShowSource(name string)*ShowSource{
	sourceFactory := GetFactoryMap().GetSourceByName(name)
	return &ShowSource{sourceFactory}
}

func (s* ShowSource)Display(){
	fmt.Println("display:",s.GetData())
}
