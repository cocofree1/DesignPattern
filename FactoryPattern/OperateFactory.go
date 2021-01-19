package FactoryPattern

// 工厂
type OperateFactory interface {
	Create()Operate
}