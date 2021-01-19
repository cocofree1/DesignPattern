package FactoryPattern

type AddOperateFactory struct{
	*OperateBase
}

func (a *AddOperateFactory) Result()int{
	return a.a + a.b
}

func (a *AddOperateFactory) Create()Operate{
	return &AddOperateFactory{&OperateBase{}}
}