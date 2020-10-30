package agent

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/sergiorra/postback-delivery/deliver-app/internal/logger/logfile"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/models/postback"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/models/response"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/repository/redis"
	"github.com/sergiorra/postback-delivery/deliver-app/internal/request"
)

type deliveryAgent struct {
	repo	redis.DeliveryRepo
	logger 	logfile.Logger
}

// NewDeliveryAgent initialize Delivery Agent
func NewDeliveryAgent(repo redis.DeliveryRepo, logger logfile.Logger) *deliveryAgent {
	return &deliveryAgent{
		repo: repo,
		logger: logger,
	}
}

func (d *deliveryAgent) Start() {
	err := d.repo.CheckConnection()
	if err != nil {
		log.Println(err)
	}

	f := d.logger.Init()
	defer f.Close()
	log.SetOutput(f)

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
		d.logResponse(res)
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
		d.logResponse(res)
	}

}

func (d *deliveryAgent) logResponse(res *response.Response) {
	log.Println("-----------------------------")
	log.Println("RESPONSE RECEIVED:")
	log.Printf("   Response Code: %v\n", res.Code)
	log.Printf("   Response Body: %v\n", res.Body)
	log.Printf("   Response Time: %v\n", res.Time)
	log.Printf("   Delivery Time: %v\n", res.DeliveryTime)
	log.Println("-----------------------------")
}
