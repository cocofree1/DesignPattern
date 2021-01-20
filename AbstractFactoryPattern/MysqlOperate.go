package AbstractFactoryPattern

import "fmt"

// mysql order
type MysqlOrderInfo struct {}

func (m *MysqlOrderInfo) GetOrderInfo(){
	fmt.Println("mysql get order info")
}

func (m *MysqlOrderInfo) SaveOrderInfo(){
	fmt.Println("mysql save order info")
}

// mysql order详情
type MysqlOrderDetailInfo struct{}

func (m *MysqlOrderDetailInfo) GetOrderDetailInfo(){
	fmt.Println("mysql get order detail info")
}

func (m *MysqlOrderDetailInfo) SaveOrderDetailInfo(){
	fmt.Println("mysql save order detail info")
}

