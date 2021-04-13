/*
auth_code package is response for issue and check the random auth code.
The code will be stored by Cacher entity for better scalability.
The Cacher interface could be implemented as a Redis or Memcache client for example.
*/
package auth_code

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)


type Cacher interface {
	SetValue(ctx context.Context, key string, value interface{}, ttl time.Duration) (Result, error)
	GetValue(ctx context.Context, key string) (Result, error)
	Remove(ctx context.Context, key string) (Result, error)
	Close() error
}

type Result interface {
	Int() (int, error)
}

type CodeValidator struct {
	CodeLength       int
	ExpirationMinute int
	CacheManager     Cacher
}

func (cv *CodeValidator) Initialize() error {
	if cv.CacheManager == nil {
		return fmt.Errorf("CacheManager should be initialized")
	}
	if cv.CodeLength == 0 {
		cv.CodeLength = 6
	}
	rand.Seed(time.Now().UnixNano())
	return nil
}

func (cv *CodeValidator) GenerateAndStore(clientId string) (int, error) {
	code := fixedLengthCode(cv.CodeLength)
	err := cv.insertCode(clientId, code)
	if err != nil {
		return 0, err
	}
	return code, nil
}

func (cv *CodeValidator) insertCode(key string, code int) error {
	ttl := time.Minute * time.Duration(cv.ExpirationMinute)
	if _, err := cv.CacheManager.SetValue(context.Background(), key, code, ttl); err != nil {
		return err
	}
	return nil
}

func (cv *CodeValidator) CheckCode(clientId string) (int, error) {
	res, err := cv.CacheManager.GetValue(context.Background(), clientId)
	if err != nil {
		return 0, err
	}
	intRes, err := res.Int()
	if err != nil {
		return 0, err
	}
	return intRes, nil
}
