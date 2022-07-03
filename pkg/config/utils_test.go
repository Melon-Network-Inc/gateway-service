package config

import (
	"bytes"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = Describe("Utils", func() {
	var (
		config *generalConfig
	)

	BeforeEach(func() {
		v := viper.New()
		v.SetConfigType("yaml")
		v.ReadConfig(bytes.NewBuffer(content))
		config = &generalConfig{v.Sub("development")}
	})

	It("should not return nil when reading content", func() {
		Expect(config).NotTo(BeNil())
	})

	It("should return proper values for Redis config provider", func() {
		var redis RedisConfigProvider = config.getRedisConfig()
		Expect(redis.GetPassword()).To(Equal("pass"))
		Expect(redis.GetHost()).To(Equal("localhost"))
		Expect(redis.GetPort()).To(Equal("6379"))
		Expect(redis.GetExpirationTime()).To(Equal(time.Duration(1) * time.Minute))
		Expect(redis.GetDB()).To(Equal(0))
	})
})
