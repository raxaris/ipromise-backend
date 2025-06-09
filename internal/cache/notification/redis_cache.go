package notification

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type redisNotificationCache struct {
	client *redis.Client
}

func NewRedisNotificationCache(client *redis.Client) NotificationCache {
	return &redisNotificationCache{client: client}
}

func (r *redisNotificationCache) SaveOfflineNotification(ctx context.Context, userID string, payload string) error {
	key := fmt.Sprintf("offline_notifications:%s", userID)
	return r.client.RPush(ctx, key, payload).Err()
}

func (r *redisNotificationCache) GetOfflineNotifications(ctx context.Context, userID string) ([]string, error) {
	key := fmt.Sprintf("offline_notifications:%s", userID)
	return r.client.LRange(ctx, key, 0, -1).Result()
}

func (r *redisNotificationCache) ClearOfflineNotifications(ctx context.Context, userID string) error {
	key := fmt.Sprintf("offline_notifications:%s", userID)
	return r.client.Del(ctx, key).Err()
}
