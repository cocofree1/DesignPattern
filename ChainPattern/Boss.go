package ChainPattern

import "fmt"

type Boss struct{}

func NewBossChain()*RequestChain{
	return &RequestChain{&Boss{},nil}
}

func (p* Boss)HavePermission(money int)bool{
	return true
}
func (p* Boss)DealRequest(name string, money int){
	msg := fmt.Sprintf("Boss agree %s %d request",name,money)
	fmt.Println(msg)
}
