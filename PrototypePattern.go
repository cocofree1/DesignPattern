package main

import "fmt"

// 原型模式
type CloneBar interface {
	Clone()CloneBar
}

type Type1 struct {
	Name  string
}

func (t *Type1)Clone()CloneBar{
	// 浅拷贝
	// return t

	// 深拷贝
	clone := *t
	return &clone
}

// ---------------原型模式管理------------------------
type PrototypeManage struct {
	prototype map[string]CloneBar
}

func NewPrototypeManage()* PrototypeManage{
	return &PrototypeManage{make(map[string]CloneBar)}
}

func (p *PrototypeManage)Get(name string)CloneBar{
	return p.prototype[name]
}

func (p *PrototypeManage)Set(name string,clone CloneBar){
	p.prototype[name] = clone
}

func main(){
	prototype := NewPrototypeManage()
	prototype.Set("tt",&Type1{})
	t1 := prototype.Get("tt")
	t2 := t1.Clone()
	if t1 != t2 {
		fmt.Println("深拷贝")
	}else{
		fmt.Println("浅拷贝")
	}
}