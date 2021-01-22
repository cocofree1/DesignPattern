package ChainPattern

import "fmt"

type DepManager struct{}

func NewDepManagerChain()*RequestChain{
	return &RequestChain{&DepManager{},nil}
}

func (p* DepManager)HavePermission(money int)bool{
	if money < 5000 {
		return true
	}
	return false
}
func (p* DepManager)DealRequest(name string, money int){
	msg := fmt.Sprintf("DepManager agree %s %d request",name,money)
	fmt.Println(msg)
}
