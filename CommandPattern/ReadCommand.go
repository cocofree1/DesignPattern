package CommandPattern

// 实现读命令
type ReadCommand struct{
	labour  LabourExecute
}

func NewReadCommand(labour LabourExecute)*ReadCommand{
	return &ReadCommand{labour: labour}
}

func (r* ReadCommand)Execute(){
	r.labour.ReadCommand()
}
