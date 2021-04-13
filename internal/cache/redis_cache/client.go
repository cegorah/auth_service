package redis_cache

import (
	"context"
	"fmt"
	"github.com/cegorah/auth_service/internal/cache"
	"github.com/cegorah/auth_service/internal/config"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	RedisConfig config.RedisConfig
	rCl         *redis.Client
}

func (r *RedisCache) Initialize() {
	conOpt := redis.Options{
		Addr:      r.RedisConfig.ConnectionString,
		Username:  r.RedisConfig.Username,
		Password:  r.RedisConfig.Password,
		DB:        r.RedisConfig.Database,
		TLSConfig: r.RedisConfig.TlsConfig,
	}
	r.rCl = redis.NewClient(&conOpt)
}
func (r RedisCache) SetValue(ctx context.Context, key string, value interface{}, ttl time.Duration) (cache.Result, error) {
	res := r.rCl.Set(ctx, key, value, ttl)
	if res.Err() != nil {
		return cache.Result{}, res.Err()
	}
	return cache.Result{Val: res.String()}, nil
}
func (r RedisCache) GetValue(ctx context.Context, key string) (cache.Result, error) {
	res := r.rCl.Get(ctx, key)
	if res.Err() == redis.Nil {
		return cache.Result{}, &EmptyResult{fmt.Sprintf("no value for %s found", key)}
	}
	if res.Err() != nil {
		return cache.Result{}, res.Err()
	}
	return cache.Result{Val: res.String()}, nil
}
func (r RedisCache) Remove(ctx context.Context, key string) (cache.Result, error) {
	res := r.rCl.Del(ctx, key)
	if res.Err() != nil {
		return cache.Result{}, res.Err()
	}
	return cache.Result{}, nil
}

func (r RedisCache) Close() error {
	if err := r.rCl.Close(); err != nil {
		return err
	}
	return nil
}
