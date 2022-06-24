package gateway

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type CacheGw interface {
	SaveDocumentVersion(documentUid string, version string) error
	GetDocumentVersion(documentUid string) (string, error)
}

type CacheGateway struct {
	redisConn *redis.Client
}

func NewCacheGateway(redisConn *redis.Client) *CacheGateway {
	return &CacheGateway{
		redisConn: redisConn,
	}
}

func (gw *CacheGateway) SaveDocumentVersion(documentUid string, version string) error {
	var ctx = context.Background()

	err := gw.redisConn.Set(ctx, documentUid, version, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (gw *CacheGateway) GetDocumentVersion(documentUid string) (string, error) {
	var ctx = context.Background()

	version, err := gw.redisConn.Get(ctx, documentUid).Result()
	if err != nil {
		return version, err
	}

	return version, nil
}
