package BuilderPattern

type Builder interface {
	Part1()
	Part2()
}

type MealCombo struct{
	builder Builder
}

// 创建套餐组合
func NewMealCombo(builder Builder)*MealCombo{
	return &MealCombo{builder: builder}
}

// 显示组合
func (m* MealCombo)ComboShow(){
	m.builder.Part1()
	m.builder.Part2()
}