package agent

import (
	"fmt"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/repository/redis"
)

type deliveryAgent struct {
	repo	redis.DeliverRepo
}

// NewDeliveryAgent initialize Delivery Agent
func NewDeliveryAgent(repo redis.DeliverRepo) *deliveryAgent {
	return &deliveryAgent{
		repo: repo,
	}
}


func (d *deliveryAgent) Start() {
	d.repo.CheckConnection()
	for {
		message := d.repo.PopMessage()
		fmt.Println(message)
	}
}