package FactoryPattern

// 具体的实现
type Operate interface {
	SetLeft(left int)
	SetRight(right int)
	Result()int
}

// 基本的数据结构
type OperateBase struct{
	a,b int
}

func (o *OperateBase)SetLeft(left int){
	o.a = left
}

func (o *OperateBase)SetRight(right int){
	o.b = right
}