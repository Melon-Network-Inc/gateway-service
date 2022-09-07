package lb

type accountServiceLB interface {
	GetNextAccountServiceAddress() string
}
type paymentServiceLB interface {
	GetNextPaymentServiceAddress() string
}

type Accessor interface {
	accountServiceLB
	paymentServiceLB
}