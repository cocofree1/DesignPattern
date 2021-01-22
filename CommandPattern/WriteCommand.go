package CommandPattern

//实现写命令

type WriteCommand struct {
	labour LabourExecute
}

func NewWriteCommand(labour LabourExecute)*WriteCommand{
	return &WriteCommand{labour: labour}
}

func (w* WriteCommand)Execute(){
	w.labour.WriteCommand()
}