package AbstractFactoryPattern

type MysqlFactory struct {}

func (m *MysqlFactory) CreateOrderInfo()OrderInfo{
	return &MysqlOrderInfo{}
}

func (m *MysqlFactory) CreateOrderDetailInfo()OrderDetailInfo{
	return &MysqlOrderDetailInfo{}
}
