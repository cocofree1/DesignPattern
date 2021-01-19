package main

import (
	"fmt"
	"sync"
)
var single *Single
var once sync.Once  // 内核信号，只会执行一次

type Single struct {
	data  string
}

func NewSingle()*Single{
	once.Do(func(){single = &Single{data :"kkkkk"}})
	return single
}

func main() {
	k1 := NewSingle()
	k2 := NewSingle()
	fmt.Println(k1 == k2)
}
