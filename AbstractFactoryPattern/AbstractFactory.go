package AbstractFactoryPattern

// 订单
type OrderInfo interface {
	GetOrderInfo()
	SaveOrderInfo()
}

// 订单详情
type OrderDetailInfo interface {
	GetOrderDetailInfo()
	SaveOrderDetailInfo()
}

// 抽象工厂
type AbstractFactory interface {
	CreateOrderInfo()OrderInfo
	CreateOrderDetailInfo()OrderDetailInfo
}
