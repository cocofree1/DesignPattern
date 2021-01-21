package StrategyPattern

import "fmt"

// 控制选择
type SelectContent struct {
	Mod  string
	Str Strategy
}

func NewSelectContent(mod string,str Strategy)* SelectContent{
	return &SelectContent{Mod: mod,Str: str}
}

// 不同的选择不同的策略
func (s *SelectContent) Travel() string{
	return s.Str.Travel(s)
}

// 策略选择
type Strategy interface {
	Travel(selectWay *SelectContent)string
}

type Bicycle struct{}

func (b *Bicycle) Travel(selectWay *SelectContent)string{
	return fmt.Sprintf("自行车出行,心情：%s",selectWay.Mod)
}

type Car struct{}

func (c *Car) Travel(selectWay *SelectContent)string{
	return fmt.Sprintf("小车出行,心情：%s",selectWay.Mod)
}