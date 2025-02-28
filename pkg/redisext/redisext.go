package redisext

import "github.com/redis/go-redis/v9"

type RedisConfig struct {
	Address string `mapstructure:"address"`
}

type RedisConfigGetter interface {
	GetRedisConfig() *RedisConfig
}

func (c *RedisConfig) GetRedisConfig() *RedisConfig {
	return c
}

func ProvideRedisClient(cfgGetter RedisConfigGetter) *redis.Client {
	cfg := cfgGetter.GetRedisConfig()

	return redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: "",
		DB:       0,
	})
}
