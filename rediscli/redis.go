package rediscli

import (
	"iptodns/config"

	"github.com/go-redis/redis/v8"
)

// RDB is a redis.Client point
var RDB *redis.Client

// Setup init redis client
func Setup() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Conf.RedisServer,
		Password: config.Conf.RedisPassword,
		DB:       config.Conf.RedisDB,
	})
}
