package DecoratorPattern

type Decorator interface {
	Calc()int
}

// 原有的功能
type OriginFunction struct {}

func (o* OriginFunction)Calc()int{
	return 7
}
