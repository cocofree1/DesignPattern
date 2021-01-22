package ChainPattern

type Manager interface {
	HavePermission(money int)bool
	DealRequest(name string, money int)
}
