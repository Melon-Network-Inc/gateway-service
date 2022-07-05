package config

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// NOTE: DO NOT CHANGE SPACES TO TABS HERE! - it will break loading
var content = []byte(`
server_defaults: &server_defaults
  auth_token_validity_period: 15m
redis_defaults: &redis_defaults
  password: "pass"
  host: "localhost"
  port: "6379"
  db: 0
  expiration_time: 1m
defaults: &defaults
  <<: *server_defaults
  redis:
    <<: *redis_defaults
development:
  <<: *defaults
production:
  <<: *defaults
  redis:
    <<: *redis_defaults
    host: "cache"
test:
  <<: *defaults
  redis:
    <<: *redis_defaults
    port: "6380"
    expiration_time: 50ms
`)

func TestCache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config")
}
