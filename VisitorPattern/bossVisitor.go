package VisitorPattern
// boss拜访者
type BossVisitor struct{
	name string
}

func NewBossVisitor(name string)*BossVisitor{
	return &BossVisitor{name:name}
}

func (b* BossVisitor)Visit(receiver Receiver){
	receiver.Accept(b)
}
