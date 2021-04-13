package redis_cache

import (
	"context"
	"github.com/cegorah/auth_service/internal"
	_ "github.com/cegorah/auth_service/internal"
	"github.com/cegorah/auth_service/internal/config"
	"log"
	"os"
	"testing"
	"time"
)

const (
	key   = "some_key"
	value = "some_value"
	ttl   = time.Second * time.Duration(5)
)

var redisCfg config.RedisConfig
var rdc RedisCache

type MockedRedis struct {
	// TODO
}

func TestMain(m *testing.M) {
	err := setUp()
	if err != nil {
		log.Fatal(err)
	}
	retCode := m.Run()
	err = tearDown()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(retCode)
}

func setUp() error {
	internal.EnvSetup(nil)
	cfg, err := config.FromEnv(os.Getenv("CONFIG_PREFIX"))
	if err != nil {
		return err
	}
	redisCfg = cfg.RedisCfg
	rdc = RedisCache{RedisConfig: redisCfg}
	rdc.Initialize()
	return nil
}

func tearDown() error {
	defer rdc.Close()
	internal.EnvRemove(nil)
	_, err := rdc.Remove(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func TestRedisCache_SetValue(t *testing.T) {
	_, err := rdc.SetValue(context.Background(), key, value, ttl)
	internal.TestOk(t, err)
}

func TestRedisCache_GetNil(t *testing.T) {
	_, err := rdc.GetValue(context.Background(), key)
	internal.TestEqual(t, err, EmptyResult{})
}

func TestRedisCache_Remove(t *testing.T) {
	// TODO
}

func TestRedisCache_SetGet(t *testing.T) {
	// TODO
}
