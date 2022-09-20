package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"os"
)

type generalConfig struct {
	v *viper.Viper
}

type AllEnvironmentLoadBalancerConfig struct {
	ServiceName   string                `mapstructure:"name"`
	Version       string                `mapstructure:"version"`
	ProdConfig    EnvLoadBalancerConfig `mapstructure:"prod"`
	StagingConfig EnvLoadBalancerConfig `mapstructure:"staging"`
	DevConfig     EnvLoadBalancerConfig `mapstructure:"dev"`
	TestConfig    EnvLoadBalancerConfig `mapstructure:"test"`
}

type EnvLoadBalancerConfig struct {
	AccountConfig []string `mapstructure:"account"`
	PaymentConfig []string `mapstructure:"payment"`
}

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

func CreateLoadBalancerConfigFromAllEnvConfig(name string, version string, envConfig EnvLoadBalancerConfig) loadbalancerConfig {
	return loadbalancerConfig{
		accountServiceAddresses: envConfig.AccountConfig,
		paymentServiceAddresses: envConfig.PaymentConfig,
	}
}

type loadbalancerConfig struct {
	accountServiceAddresses []string
	paymentServiceAddresses []string
}

func (config *loadbalancerConfig) GetAccountServiceAddressesList() []string {
	return config.accountServiceAddresses
}

func (config *loadbalancerConfig) GetPaymentServiceAddressesList() []string {
	return config.paymentServiceAddresses
}
