package lb

type accountLB struct {
	idx		        int
	serviceAddreses []string
}

func (a *accountLB) GetNextAccountServiceAddress() string {
	nextIdx := a.idx % len(a.serviceAddreses)
	a.idx++
	return a.serviceAddreses[nextIdx]
}

func newAccountLB(serviceAddress []string) accountServiceLB {
	return &accountLB {
		idx: 0,
		serviceAddreses: serviceAddress,
	}
}