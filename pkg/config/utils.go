package config

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

// generalConfig is the general configuration for the gateway service.
type generalConfig struct {
	v *viper.Viper
}

// LoadBalancerConfiguration is the configuration for load balancer.
type AllEnvironmentLoadBalancerConfig struct {
	ServiceName   string                `mapstructure:"name"`
	Version       string                `mapstructure:"version"`
	ProdConfig    EnvLoadBalancerConfig `mapstructure:"prod"`
	StagingConfig EnvLoadBalancerConfig `mapstructure:"staging"`
	DevConfig     EnvLoadBalancerConfig `mapstructure:"dev"`
	TestConfig    EnvLoadBalancerConfig `mapstructure:"test"`
}

// EnvLoadBalancerConfig is the configuration for load balancer in each environment.
type EnvLoadBalancerConfig struct {
	AccountConfig     []string `mapstructure:"account"`
	PaymentConfig     []string `mapstructure:"payment"`
	MaintenanceConfig []string `mapstructure:"maintenance"`
}

// LoadBalancerConfiguration is the configuration for load balancer.
func (config *generalConfig) getLoadBalancerConfig() *loadbalancerConfig {
	var al AllEnvironmentLoadBalancerConfig

	if err := config.v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to open config file %w", err))
	}
	config.v.AutomaticEnv()

	if err := config.v.Unmarshal(&al); err != nil {
		panic(fmt.Errorf("failed to unmarshal config file %w", err))
	}

	var lb loadbalancerConfig
	switch os.Getenv("TARGET_ENV") {
	case "PROD":
		lb = CreateLoadBalancerConfigFromAllEnvConfig(al.ServiceName, al.Version, al.ProdConfig)
	case "STAGING":
		lb = CreateLoadBalancerConfigFromAllEnvConfig(al.ServiceName, al.Version, al.StagingConfig)
	case "DEV":
		lb = CreateLoadBalancerConfigFromAllEnvConfig(al.ServiceName, al.Version, al.DevConfig)
		spew.Dump(al)
	default:
		lb = CreateLoadBalancerConfigFromAllEnvConfig(al.ServiceName, al.Version, al.TestConfig)
		spew.Dump(al)
	}
	return &lb
}

// CreateLoadBalancerConfigFromAllEnvConfig creates a loadbalancerConfig from the AllEnvironmentLoadBalancerConfig
func CreateLoadBalancerConfigFromAllEnvConfig(name string, version string, envConfig EnvLoadBalancerConfig) loadbalancerConfig {
	return loadbalancerConfig{
		accountServiceAddresses:   envConfig.AccountConfig,
		paymentServiceAddresses:   envConfig.PaymentConfig,
		maintenanceServiceAddress: envConfig.PaymentConfig,
	}
}

// loadbalancerConfig is the configuration for the loadbalancer
type loadbalancerConfig struct {
	accountServiceAddresses 	[]string
	paymentServiceAddresses 	[]string
	maintenanceServiceAddress   []string
}

// GetAccountServiceAddressesList returns the list of account service addresses
func (config *loadbalancerConfig) GetAccountServiceAddressesList() []string {
	return config.accountServiceAddresses
}

// GetPaymentServiceAddressesList returns the list of payment service addresses
func (config *loadbalancerConfig) GetPaymentServiceAddressesList() []string {
	return config.paymentServiceAddresses
}

// GetMaintenanceServiceAddressesList returns the list of maintenance service addresses
func (config *loadbalancerConfig) GetMaintenanceServiceAddressesList() []string {
	return config.maintenanceServiceAddress
}