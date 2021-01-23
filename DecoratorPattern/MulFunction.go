package DecoratorPattern

// 新增乘法功能
type MulFunction struct{
	Decorator
	num int
}

func WrapMulFunction(origin Decorator, num int)*MulFunction{
	return &MulFunction{origin,num}
}

func (m* MulFunction)Calc()int{
	return m.Decorator.Calc() * m.num
}
