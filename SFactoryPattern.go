package main

import "fmt"

// 简单工厂

type PersonAPI interface{
	Say(name string)string
}

func NewPersonAPI(str string)PersonAPI{
	if str == "cn" {
		return &Chinese{}
	} else if str == "en" {
		return &English{}
	} else {
		return nil
	}
}

type Chinese struct {}

func (*Chinese)Say(name string)string{
	return "你好 " + name
}

type English struct {}

func (*English)Say(name string)string{
	return "hello " + name
}

func main(){
	api := NewPersonAPI("en")
	fmt.Println(api.Say("yaya"))
}