package ObserverPattern

// 观察者
type ObserverControl struct {
	Person       []Observer
	UpdateMsg    string
}

// 新建对象
func NewObserver()*ObserverControl{
	return &ObserverControl{Person: make([]Observer,0)}
}

// 添加观察人员
func (o *ObserverControl)AddPerson(obs Observer){
	o.Person = append(o.Person, obs)
}

// 通知信息
func (o *ObserverControl) Notify(){
	for _,item := range o.Person{
		item.UpdateMsg(o)
	}
}

// 更新消息
func (o *ObserverControl) UpdateMessage(context string){
	o.UpdateMsg = context
	o.Notify()
}
