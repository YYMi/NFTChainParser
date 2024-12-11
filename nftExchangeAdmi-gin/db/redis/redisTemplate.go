package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"nftExchangeAdmi-gin/config"
	"time"
)

// 全局变量
var (
	ctx           = context.Background()
	clusterClient *redis.ClusterClient
)

// CreateRedisClusterClient 初始化 Redis 集群客户端
func CreateRedisClusterClient(cof config.Redis) {
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cof.Cluster.Nodes,
		Password: cof.Password, // 如果没有密码，可以设置为空字符串
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			logrus.Infof("Connected to Redis: %s", cn.String())
			return nil
		},
	})

	// 测试连接
	pong, err := clusterClient.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("Could not connect to Redis cluster: %v", err)
	} else {
		logrus.Infof("Connected to Redis cluster: %s\n", pong)
	}
}

// Set 设置 key 的值，并设置过期时间
func Set(key string, value interface{}, expiration time.Duration) {
	err := clusterClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		logrus.Errorf("Failed to set key %s: %v", key, err)
	}
}

// Get 获取指定 key 的值
func Get(key string) (string, error) {
	val, err := clusterClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			logrus.Warnf("Key %s does not exist", key)
		} else {
			logrus.Errorf("Failed to get key %s: %v", key, err)
		}
		return "", err
	}
	return val, nil
}

// Delete 删除指定的 key
func Delete(key string) {
	err := clusterClient.Del(ctx, key).Err()
	if err != nil {
		logrus.Errorf("Failed to delete key %s: %v", key, err)
	}
}

// Expire 设置 key 的过期时间
func Expire(key string, expiration time.Duration) {
	err := clusterClient.Expire(ctx, key, expiration).Err()
	if err != nil {
		logrus.Errorf("Failed to set expiration for key %s: %v", key, err)
	}
}

// Increment 增加指定 key 的值
func Increment(key string) (int64, error) {
	val, err := clusterClient.Incr(ctx, key).Result()
	if err != nil {
		logrus.Errorf("Failed to increment key %s: %v", key, err)
		return 0, err
	}
	return val, nil
}

// Decrement 减少指定 key 的值
func Decrement(key string) (int64, error) {
	val, err := clusterClient.Decr(ctx, key).Result()
	if err != nil {
		logrus.Errorf("Failed to decrement key %s: %v", key, err)
		return 0, err
	}
	return val, nil
}

// HSet 设置哈希表字段的值
func HSet(key, field string, value interface{}) {
	err := clusterClient.HSet(ctx, key, field, value).Err()
	if err != nil {
		logrus.Errorf("Failed to set hash field %s in key %s: %v", field, key, err)
	}
}

// HGet 获取哈希表字段的值
func HGet(key, field string) (string, error) {
	val, err := clusterClient.HGet(ctx, key, field).Result()
	if err != nil {
		if err == redis.Nil {
			logrus.Warnf("Field %s does not exist in key %s", field, key)
		} else {
			logrus.Errorf("Failed to get hash field %s in key %s: %v", field, key, err)
		}
		return "", err
	}
	return val, nil
}

// SAdd 向集合添加一个或多个成员
func SAdd(key string, members ...interface{}) {
	err := clusterClient.SAdd(ctx, key, members...).Err()
	if err != nil {
		logrus.Errorf("Failed to add members to set %s: %v", key, err)
	}
}

// SMembers 获取集合中的所有成员
func SMembers(key string) ([]string, error) {
	members, err := clusterClient.SMembers(ctx, key).Result()
	if err != nil {
		logrus.Errorf("Failed to get members of set %s: %v", key, err)
		return nil, err
	}
	return members, nil
}

// ZAdd 向有序集合添加一个或多个成员
func ZAdd(key string, members ...*redis.Z) {
	err := clusterClient.ZAdd(ctx, key, members...).Err()
	if err != nil {
		logrus.Errorf("Failed to add members to sorted set %s: %v", key, err)
	}
}

// ZRange 获取有序集合指定区间内的成员
func ZRange(key string, start, stop int64) ([]string, error) {
	members, err := clusterClient.ZRange(ctx, key, start, stop).Result()
	if err != nil {
		logrus.Errorf("Failed to get range of sorted set %s: %v", key, err)
		return nil, err
	}
	return members, nil
}

// LPush 向列表头部添加一个或多个值
func LPush(key string, values ...interface{}) {
	err := clusterClient.LPush(ctx, key, values...).Err()
	if err != nil {
		logrus.Errorf("Failed to push values to list %s: %v", key, err)
	}
}

// RPush 向列表尾部添加一个或多个值
func RPush(key string, values ...interface{}) {
	err := clusterClient.RPush(ctx, key, values...).Err()
	if err != nil {
		logrus.Errorf("Failed to push values to list %s: %v", key, err)
	}
}

// LPop 移除并返回列表的第一个元素
func LPop(key string) (string, error) {
	val, err := clusterClient.LPop(ctx, key).Result()
	if err != nil {
		logrus.Errorf("Failed to pop value from list %s: %v", key, err)
		return "", err
	}
	return val, nil
}

// RPop 移除并返回列表的最后一个元素
func RPop(key string) (string, error) {
	val, err := clusterClient.RPop(ctx, key).Result()
	if err != nil {
		logrus.Errorf("Failed to pop value from list %s: %v", key, err)
		return "", err
	}
	return val, nil
}

// Exists 检查一个或多个 key 是否存在
func Exists(keys ...string) (int64, error) {
	val, err := clusterClient.Exists(ctx, keys...).Result()
	if err != nil {
		logrus.Errorf("Failed to check existence of keys %v: %v", keys, err)
		return 0, err
	}
	return val, nil
}

// TTL 获取 key 的剩余过期时间
func TTL(key string) (time.Duration, error) {
	ttl, err := clusterClient.TTL(ctx, key).Result()
	if err != nil {
		logrus.Errorf("Failed to get TTL for key %s: %v", key, err)
		return 0, err
	}
	return ttl, nil
}

// FlushAll 清空整个 Redis 服务器数据
func FlushAll() {
	err := clusterClient.FlushAll(ctx).Err()
	if err != nil {
		logrus.Errorf("Failed to flush all data in Redis: %v", err)
	}
}
