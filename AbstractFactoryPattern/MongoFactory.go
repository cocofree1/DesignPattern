package AbstractFactoryPattern

type MongoFactory struct {}

func (m *MongoFactory) CreateOrderInfo()OrderInfo{
	return &MongoOrderInfo{}
}

func (m *MongoFactory) CreateOrderDetailInfo()OrderDetailInfo{
	return &MongoOrderDetailInfo{}
}
