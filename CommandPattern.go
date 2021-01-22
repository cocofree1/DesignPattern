package main

import "DesignPattern/CommandPattern"

func main(){
	labour := CommandPattern.LabourExecute{}
	cmd1 := CommandPattern.NewReadCommand(labour)
	cmd2 := CommandPattern.NewWriteCommand(labour)

	sender := CommandPattern.NewSendCommand(cmd1,cmd2)
	sender.Read()
	sender.Write()
}
