package main

import (
	"fmt"

	"github.com/sergiorra/postback-delivery/deliver-app/internal/agent"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/repository/redis"
)

func main() {
	fmt.Println("starting...")

	repo := redis.NewRepository("127.0.0.1", "6379")
	agent := agent.NewDeliveryAgent(repo)
	agent.Start()

}
