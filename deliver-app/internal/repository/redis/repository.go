package redis

import (
	"context"

	. "github.com/sergiorra/postback-delivery/deliver-app/internal/repository"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// repository representation of repository into struct
type repository struct {
	client *redis.Client
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
func (r *repository) CheckConnection() error {
	_, err := r.client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// PopMessage pop a message from Redis
func (r *repository) PopMessage() (string, error){
	str, err := r.client.BRPop(ctx, 0, "data").Result()
	if err != nil {
		return "", err
	}
	return str[1], nil
}