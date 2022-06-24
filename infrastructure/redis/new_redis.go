package redis

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"thourus-api/config"
)

const serviceName = "redis"

func NewRedisConnection(redisCfg *config.RedisConfig, logger *zap.SugaredLogger) (*redis.Client, error) {

	serviceLogger := logger.Named(serviceName)

	serviceLogger.Info("Establishing the Redis connection...")

	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	serviceLogger.Info("Established the Redis connection successfully.")
	return redisConn, nil
}
