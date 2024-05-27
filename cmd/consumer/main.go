package main

import (
	"github.com/third-place/notification-service/internal/kafka"
)

func main() {
	kafka.InitializeAndRunLoop()
}
