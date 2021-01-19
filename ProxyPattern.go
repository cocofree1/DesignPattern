package main

import "fmt"

// 代理模式

type Subject interface {
	Do()string
}

// 真实操作
type RealSub struct{}

func (r *RealSub)Do()string{
	return "业务处理"
}

// 代理操作
type ProxySub struct {
	Read RealSub
	Pwd string
}

func (p *ProxySub)Do()string{
	if p.Pwd == "12345wsx" {
		return p.Read.Do()
	} else {
		return "操作失败"
	}
}

func main(){
	var sub Subject
	sub = &ProxySub{Pwd: "12345wsx"}
	fmt.Println(sub.Do())
}