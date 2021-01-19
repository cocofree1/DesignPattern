package FactoryPattern

type SubOperateFactory struct{
	*OperateBase
}

func (s *SubOperateFactory) Result()int{
	return s.a - s.b
}

func (s *SubOperateFactory) Create()Operate{
	return &SubOperateFactory{&OperateBase{}}
}
