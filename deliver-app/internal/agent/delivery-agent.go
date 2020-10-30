package agent

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sergiorra/postback-delivery/deliver-app/internal/models/postback"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/repository/redis"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/request"
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
	err := d.repo.CheckConnection()
	if err != nil {
		log.Println(err)
	}

	for {
		message, err := d.repo.PopMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		go d.process(message)
	}

}

func (d *deliveryAgent) process(message string) {
	fmt.Println(message)
	pb := &postback.Postback{}
	if err := json.Unmarshal([]byte(message), pb); err != nil {
		log.Println(err)
		return
	}
	pb.MountURL()

	req := request.NewRequest(pb.Endpoint.Url, pb.Endpoint.Method)

	switch strings.ToLower(pb.Endpoint.Method) {
	case "get":
		res, err := req.Get()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(res)
	case "post":
		body, err := json.Marshal(pb.Data[0])
		if err != nil {
			log.Println(err)
			return
		}
		req.Body = body
		res, err := req.Post()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(res)
	}

}
