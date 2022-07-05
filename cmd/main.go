package main

import (
	"os"

	"github.com/Melon-Network-Inc/gateway-service/cmd/server"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stderr)
}

func main() {
	s := server.New()
	s.Run(":8080")
}
