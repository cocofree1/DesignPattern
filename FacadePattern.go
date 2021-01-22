package main

import "fmt"

// 外观模式
//--------------------------------------
// a 测试接口
type TestAAPI interface {
	TestA()
}

type TestA struct {}

func (t* TestA)TestA(){
	fmt.Println("测试A")
}

func NewTestAPI()TestAAPI{
	return &TestA{}
}

//-----------------------

// b 测试接口
type TestBAPI interface {
	TestB()
}

type TestB struct {}
func (t* TestB)TestB(){
	fmt.Println("测试B")
}

func NewTestBAPI()TestBAPI{
	return &TestB{}
}

//------------------------------
// 统一测试接口
type TestAllAPI interface {
	Test()
}

type TestAll struct {
	A TestAAPI
	B TestBAPI
}

func (t* TestAll)Test(){
	t.A.TestA()
	t.B.TestB()
}

func NewTestAllAPI()TestAllAPI{
	return &TestAll{NewTestAPI(),NewTestBAPI()}
}

func main(){
	api := NewTestAllAPI()
	api.Test()
}
