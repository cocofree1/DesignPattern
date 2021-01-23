package DecoratorPattern

// 新增加法功能
type AddFunction struct {
	Decorator
	num int
}

func WrapAddFunction(origin Decorator, num int)*AddFunction{
	return &AddFunction{origin,num}
}

func (a* AddFunction)Calc()int{
	return a.Decorator.Calc() + a.num
}

