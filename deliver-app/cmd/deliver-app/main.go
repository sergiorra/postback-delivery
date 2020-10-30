package main

import (
	"fmt"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/logger/logfile"

	"github.com/sergiorra/postback-delivery/deliver-app/internal/agent"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/repository/redis"
)

func main() {
	fmt.Println("starting...")

	repo := redis.NewRepository("127.0.0.1", "6379")
	logger := logfile.NewLogger("logs.txt")
	agent := agent.NewDeliveryAgent(repo, logger)
	agent.Start()

}
