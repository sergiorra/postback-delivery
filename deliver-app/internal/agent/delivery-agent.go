package agent

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sergiorra/postback-delivery/deliver-app/internal/models/postback"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/repository/redis"
)

type deliveryAgent struct {
	repo	redis.DeliveryRepo
}

// NewDeliveryAgent initialize Delivery Agent
func NewDeliveryAgent(repo redis.DeliveryRepo) *deliveryAgent {
	return &deliveryAgent{
		repo: repo,
	}
}

func (d *deliveryAgent) Start() {
	d.repo.CheckConnection()

	for {
		message := d.repo.PopMessage()
		go d.process(message)
	}

}

func (d *deliveryAgent) process(message string) {
	postback := postback.Postback{}
	if err := json.Unmarshal([]byte(message), &postback); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(message)
}