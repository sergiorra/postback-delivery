package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type repository struct {
	client *redis.Client
}

type DeliveryRepo interface {
	CheckConnection()
	PopMessage() string
}

// NewRepository initialize redis repository
func NewRepository(addr, port string) DeliveryRepo {
	return &repository{
		client: redis.NewClient(&redis.Options{
			Addr: addr + ":" + port,
		}),
	}
}

// CheckConnection check Redis connection
func (r *repository) CheckConnection() {
	_, err := r.client.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err)
	}
}

// PopMessage check Redis connection
func (r *repository) PopMessage() string {
	str, err := r.client.BRPop(ctx, 0, "data").Result()
	if err != nil {
		log.Fatalln(err)
	}
	return str[1]
}