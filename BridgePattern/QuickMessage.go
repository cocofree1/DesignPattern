package BridgePattern

type QuickMessage struct{
	Message DifferentMessage
}

func NewQuickMessage(mode DifferentMessage)DiffSpeedMessage{
	return &QuickMessage{Message: mode}
}

func (q* QuickMessage)SendSpeedMessage(name,to string){
	q.Message.SendKindMessage(name,to)
}