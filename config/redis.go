package config

import (
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Addr string
}

func InitRedisConfig() *RedisConfig {

	redisConfig := RedisConfig{
		Addr: viper.GetString("redis.addr"),
	}
	return &redisConfig
}

func init() {
	viper.SetDefault("redis.addr", "localhost:6379")

	InitError(viper.BindEnv("redis.addr", EnvPrefix+"_REDIS_ADDR"))
}
