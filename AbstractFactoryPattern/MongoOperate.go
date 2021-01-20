package AbstractFactoryPattern

import "fmt"

// Mongo order
type MongoOrderInfo struct {}

func (m *MongoOrderInfo) GetOrderInfo(){
	fmt.Println("Mongo get order info")
}

func (m *MongoOrderInfo) SaveOrderInfo(){
	fmt.Println("Mongo save order info")
}

// Mongo order详情
type MongoOrderDetailInfo struct{}

func (m *MongoOrderDetailInfo) GetOrderDetailInfo(){
	fmt.Println("Mongo get order detail info")
}

func (m *MongoOrderDetailInfo) SaveOrderDetailInfo(){
	fmt.Println("Mongo save order detail info")
}

