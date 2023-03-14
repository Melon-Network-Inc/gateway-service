package lb

type accountServiceLB interface {
	GetNextAccountServiceAddress() string
}
type paymentServiceLB interface {
	GetNextPaymentServiceAddress() string
}

type maintenanceServiceLB interface {
	GetNextMaintenanceServiceAddress() string
}

type Accessor interface {
	accountServiceLB
	paymentServiceLB
	maintenanceServiceLB
}