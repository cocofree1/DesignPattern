package CommandPattern

// 发送命令
type SendCommand struct{
	read  Command
	write Command
}

func NewSendCommand(read, write Command)*SendCommand{
	return &SendCommand{read: read, write: write}
}

func (s* SendCommand)Read(){
	s.read.Execute()
}

func (s* SendCommand)Write(){
	s.write.Execute()
}
