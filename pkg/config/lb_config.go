package config

import "github.com/spf13/viper"

// LoadBalancerConfiguration keeps all config providers
type LoadBalancerConfiguration struct {
	LoadBalancer      LoadBalancerConfigProvider
}

func NewLoadBalancerConfig(path string) *LoadBalancerConfiguration {
	v := viper.New()

	v.SetConfigFile(path)
	v.SetConfigType("yml")

	config := &generalConfig{v}
	return &LoadBalancerConfiguration{
		LoadBalancer:		config.getLoadBalancerConfig(),
	}
}
