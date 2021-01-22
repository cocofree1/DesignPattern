package BridgePattern

type CommMessage struct{
	Message DifferentMessage
}

func NewCommMessage(mode DifferentMessage)DiffSpeedMessage{
	return &CommMessage{Message: mode}
}

func (c* CommMessage)SendSpeedMessage(name,to string){
	c.Message.SendKindMessage(name,to)
}
