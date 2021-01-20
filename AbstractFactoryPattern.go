package main

import "DesignPattern/AbstractFactoryPattern"

func main(){
	var abstractFactory AbstractFactoryPattern.AbstractFactory
	//abstractFactory = &AbstractFactoryPattern.MysqlFactory{}
	abstractFactory = &AbstractFactoryPattern.MongoFactory{}
	order := abstractFactory.CreateOrderInfo()
	orderDetail := abstractFactory.CreateOrderDetailInfo()
	order.GetOrderInfo()
	order.SaveOrderInfo()
	orderDetail.GetOrderDetailInfo()
	orderDetail.SaveOrderDetailInfo()
}
