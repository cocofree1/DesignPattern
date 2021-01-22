package ChainPattern

import "fmt"

type ProjectManager struct{}

func NewProjectManagerChain()*RequestChain{
	return &RequestChain{&ProjectManager{},nil}
}

func (p* ProjectManager)HavePermission(money int)bool{
	if money < 1000 {
		return true
	}
	return false
}
func (p* ProjectManager)DealRequest(name string, money int){
	msg := fmt.Sprintf("ProjectManager agree %s %d request",name,money)
	fmt.Println(msg)
}
