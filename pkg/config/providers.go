package config

// LoadBalancerConfigProvider provides load balancer server address.
type LoadBalancerConfigProvider interface {
	GetAccountServiceAddressesList() []string
	GetPaymentServiceAddressesList() []string
}
