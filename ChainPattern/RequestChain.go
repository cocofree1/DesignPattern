package ChainPattern

type RequestChain struct {
	Manager
	NextReq *RequestChain
}

func (r* RequestChain)SetNextChain(nextReq *RequestChain){
	r.NextReq = nextReq
}

func (r* RequestChain)DealChain(name string,money int){
	if r.HavePermission(money) {
		r.DealRequest(name, money)
		return
	}
	if r.NextReq != nil{
		r.NextReq.DealChain(name, money)
	}
}
