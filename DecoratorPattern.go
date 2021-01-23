package main

import (
	"DesignPattern/DecoratorPattern"
	"fmt"
)

//装饰器模式

func main(){
	var decorator DecoratorPattern.Decorator
	decorator = &DecoratorPattern.OriginFunction{}
	// decorator = DecoratorPattern.WrapAddFunction(decorator,10)
	decorator = DecoratorPattern.WrapMulFunction(decorator, 10)
	fmt.Println(decorator.Calc())
}
