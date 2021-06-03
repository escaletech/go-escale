package redisclient

import (
	"context"

	"github.com/escaletech/go-escale/messages"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func Connect(redisAddress string, redisCluster bool, log *logrus.Logger, context context.Context, operation string) redis.UniversalClient {
	redisClient, err := client("redis://"+redisAddress, redisCluster)
	if err != nil {
		log.Fatal(messages.RedisConnectionError(operation, err))
	}
	_, err = redisClient.Ping(context).Result()

	if err != nil {
		log.Fatal(messages.RedisConnectionError(operation, err))
	}

	return redisClient
}

func client(redisURL string, cluster bool) (redis.UniversalClient, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	if cluster {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    []string{opt.Addr},
			Username: opt.Username,
			Password: opt.Password,
		}), nil
	}

	return redis.NewClient(opt), nil
}
