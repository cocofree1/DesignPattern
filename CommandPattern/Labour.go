package CommandPattern

import "fmt"

// 劳动力，真正执行命令

type LabourExecute struct{}

func (l* LabourExecute)ReadCommand(){
	fmt.Println("读命令成功")
}

func (l* LabourExecute)WriteCommand(){
	fmt.Println("写命令成功")
}
