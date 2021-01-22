package BridgePattern

import "fmt"

type PhoneMessage struct{}

func NewPhoneMessage()DifferentMessage{
	return &PhoneMessage{}
}

func (p* PhoneMessage)SendKindMessage(name,to string){
	msg := fmt.Sprintf("%s send phone message to %s",name,to)
	fmt.Println(msg)
}
