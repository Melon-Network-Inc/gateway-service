package lb

type maintenaceLB struct {
	idx		        int
	serviceAddreses []string
}

func (p *maintenaceLB) GetNextMaintenanceServiceAddress() string {
	nextIdx := p.idx % len(p.serviceAddreses)
	p.idx++
	return p.serviceAddreses[nextIdx]
}

func newMaintenanceLB(serviceAddress []string) maintenanceServiceLB {
	return &maintenaceLB{
		idx: 0,
		serviceAddreses: serviceAddress,
	}
}