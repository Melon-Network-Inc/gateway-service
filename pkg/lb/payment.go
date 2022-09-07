package lb

type paymentLB struct {
	idx		        int
	serviceAddreses []string
}

func (p *paymentLB) GetNextPaymentServiceAddress() string {
	nextIdx := p.idx % len(p.serviceAddreses)
	p.idx++
	return p.serviceAddreses[nextIdx]
}

func newPaymentLB(serviceAddress []string) paymentServiceLB {
	return &paymentLB {
		idx: 0,
		serviceAddreses: serviceAddress,
	}
}