package MementoPattern

type Memento struct{
	code  string
	state string
}

func NewMemento(code string)*Memento{
	return &Memento{code:code,state: ""}
}

func (m* Memento)GetCode()string{
	return m.code
}

func (m* Memento)GetState()string{
	return m.state
}

var StateArray []*Memento

func (m* Memento)Write(){
	m.state = "编写状态"
	memento := &Memento{code:m.code,state:"编写状态"}
	StateArray = append(StateArray,memento)
}

func (m* Memento)Commit(){
	m.state = "commit状态"
	memento := &Memento{code:m.code,state:"commit状态"}
	StateArray = append(StateArray,memento)
}

func (m* Memento)Submit(){
	m.state = "提交状态"
	memento := &Memento{code:m.code,state:"提交状态"}
	StateArray = append(StateArray,memento)
}
