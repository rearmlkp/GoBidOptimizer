package Managers

import (
	"github.com/go-redis/redis"
	"GoBidOptimizer/src/Configurations"
	"time"
)

var client *redis.Client

type RedisManager struct {
}

func GetRedisClient() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     Configurations.GetConfig().RedisHost + ":" + Configurations.GetConfig().RedisPort,
			Password: Configurations.GetConfig().RedisPassword,
			DB:       0,
		})
	}
	return client
}

func AutoSet(key string, value string) error {
	return GetRedisClient().Set(key, value, time.Duration(1)*time.Hour).Err()
}

func AutoHmSet(key string, data map[string]interface{}) error {
	return GetRedisClient().Watch(func(tx *redis.Tx) error {
		if err := tx.HMSet(key, data).Err(); err != nil {
			return err
		} else {
			return tx.Expire(key, time.Duration(1)*time.Hour).Err()
		}
		return nil
	})
}
