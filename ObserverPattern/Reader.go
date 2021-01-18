package ObserverPattern

import "fmt"

// 观察
type Observer interface {
	UpdateMsg(observer *ObserverControl)
}

// 被观察者
type Reader struct {
	Name string
}

// 创建Reader
func NewReader(name string)*Reader {
	return &Reader{Name: name}
}

// 被通知消息
func (r *Reader) UpdateMsg(observer *ObserverControl){
	str := fmt.Sprintf("%s ,%s", r.Name,observer.UpdateMsg)
	fmt.Println(str)
}


