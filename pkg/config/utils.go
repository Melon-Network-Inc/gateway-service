package config

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type generalConfig struct {
	v *viper.Viper
}

type AllEnvironmentTokenConfig struct {
	ServiceName       string 			`mapstructure:"name"`
	Version           string 			`mapstructure:"version"`
	ProdConfig		  EnvTokenConfig 	`mapstructure:"prod"`
	StagingConfig	  EnvTokenConfig 	`mapstructure:"staging"`
	DevConfig		  EnvTokenConfig 	`mapstructure:"dev"`
	TestConfig		  EnvTokenConfig 	`mapstructure:"test"`
}

type EnvTokenConfig struct {
	Secret         				string 	`mapstructure:"secret"`
	AuthTokenValidityPeriod 	int64 	`mapstructure:"auth_token_period"`
	RefreshTokenValidityPeriod  int64 	`mapstructure:"refresh_token_period"`
}

func (config *generalConfig) getTokenConfig() *tokenConfig {
	var at AllEnvironmentTokenConfig
	
	if err := config.v.ReadInConfig();  err != nil {
		panic(fmt.Errorf("failed to open config file %w", err))
    }
    config.v.AutomaticEnv()
	
    if err := config.v.Unmarshal(&at); err != nil {
		panic(fmt.Errorf("failed to unmarshal config file %w", err))
    }

	var tc tokenConfig
	switch os.Getenv("TARGET_ENV") {
		case "PROD":
			tc = CreateTokenConfigFromAllEnvConfig(at.ServiceName, at.Version, at.ProdConfig)
		case "STAGING":
			tc = CreateTokenConfigFromAllEnvConfig(at.ServiceName, at.Version, at.StagingConfig)
		case "DEV":
			tc = CreateTokenConfigFromAllEnvConfig(at.ServiceName, at.Version, at.DevConfig)
			spew.Dump(at)
		default:
			tc = CreateTokenConfigFromAllEnvConfig(at.ServiceName, at.Version, at.TestConfig)
			spew.Dump(at)
	}
	return &tc
}

func CreateTokenConfigFromAllEnvConfig(name string, version string, envConfig EnvTokenConfig) tokenConfig {
	return tokenConfig{
		secretKey:                  []byte(envConfig.Secret),
		authTokenValidityPeriod:    time.Hour * time.Duration(envConfig.AuthTokenValidityPeriod),
		refreshTokenValidityPeriod: time.Hour * time.Duration(envConfig.RefreshTokenValidityPeriod),
	}
}

type tokenConfig struct {
	secretKey                  []byte
	authTokenValidityPeriod    time.Duration
	refreshTokenValidityPeriod time.Duration
}

func (config *tokenConfig) GetSecretKey() []byte {
	return config.secretKey
}

func (config *tokenConfig) GetAuthTokenValidityPeriod() time.Duration {
	return config.authTokenValidityPeriod
}

func (config *tokenConfig) GetRefreshTokenValidityPeriod() time.Duration {
	return config.refreshTokenValidityPeriod
}

type AllEnvironmentLoadBalancerConfig struct {
	ServiceName       string 					`mapstructure:"name"`
	Version           string 					`mapstructure:"version"`
	ProdConfig		  EnvLoadBalancernConfig 	`mapstructure:"prod"`
	StagingConfig	  EnvLoadBalancernConfig 	`mapstructure:"staging"`
	DevConfig		  EnvLoadBalancernConfig 	`mapstructure:"dev"`
	TestConfig		  EnvLoadBalancernConfig 	`mapstructure:"test"`
}

type EnvLoadBalancernConfig struct {
	AccountConfig     	[]string `mapstructure:"account"`
	PaymentConfig     	[]string `mapstructure:"payment"`
}

func (config *generalConfig) getLoadBalancerConfig() *loadbalancerConfig {
	var al AllEnvironmentLoadBalancerConfig
	
	if err := config.v.ReadInConfig();  err != nil {
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

func CreateLoadBalancerConfigFromAllEnvConfig(name string, version string, envConfig EnvLoadBalancernConfig) loadbalancerConfig {
	return loadbalancerConfig{
		accountServiceAddresses: 	envConfig.AccountConfig,
		paymentServiceAddresses: 	envConfig.PaymentConfig,
	}
}

type loadbalancerConfig struct {
	accountServiceAddresses 	[]string
	paymentServiceAddresses 	[]string
}

func (config *loadbalancerConfig) GetAccountServiceAddressesList() []string {
	return config.accountServiceAddresses
}

func (config *loadbalancerConfig) GetPaymentServiceAddressesList() []string {
	return config.paymentServiceAddresses
}