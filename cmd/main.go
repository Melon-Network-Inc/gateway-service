package main

import (
	"os"

	"gateway-service/cmd/server"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stderr)
}

func main() {
	s := server.New()
	s.Run(":8080")
}
