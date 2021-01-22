package BridgePattern

import "fmt"

type EmailMessage struct {}

func NewEmailMessage()DifferentMessage{
	return &EmailMessage{}
}

func (e* EmailMessage)SendKindMessage(name,to string){
	msg := fmt.Sprintf("%s send email message to %s",name,to)
	fmt.Println(msg)
}
