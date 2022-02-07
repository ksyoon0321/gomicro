package adapter

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisAdapter struct {
	client *redis.Client
	chanid string
}

func NewRedisAdapter(chanid, addr string) *RedisAdapter {
	return &RedisAdapter{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
		chanid: chanid,
	}
}

func (a *RedisAdapter) Listen(ch chan NotifyData) {
	subs := a.client.Subscribe(context.Background(), a.chanid)

	for {
		data, err := subs.ReceiveMessage(context.Background())
		if err != nil {
			//write log

			continue
		}

		obj, err := ParseRegPayload(data.Payload)
		if err != nil {
			//write log
			continue
		}

		ch <- obj
	}
}
