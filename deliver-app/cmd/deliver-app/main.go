package main

import (
	"github.com/sergiorra/postback-delivery/deliver-app/internal/agent"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/logger/logfile"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/repository/redis"
)

func main() {
	repo := redis.NewRepository("127.0.0.1", "6379")
	logger := logfile.NewLogger("logs.txt")
	deliveryAgent := agent.NewDeliveryAgent(repo, logger)
	deliveryAgent.Start()
}
