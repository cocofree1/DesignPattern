package VisitorPattern

// 批量处理，方便显示结果使用

type Visitors struct{
	visitors []Visitor
}

func (v* Visitors)Add(visitor Visitor){
	v.visitors = append(v.visitors, visitor)
}

func (v* Visitors)Visit(receiver Receiver){
	for _,visitor := range v.visitors{
		visitor.Visit(receiver)
	}
}