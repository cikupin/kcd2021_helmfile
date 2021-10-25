package bootstrap

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheOption struct {
	IsEnable        bool          `env:"CACHE_IS_ENABLED"`
	Host            string        `env:"CACHE_HOST"`
	Port            int           `env:"CACHE_PORT"`
	Password        string        `env:"CACHE_PASSWORD"`
	Namespace       int           `env:"CACHE_NAMESPACE"`
	MaxIdle         int           `env:"CACHE_MAX_IDLE_CONNECTION"`
	MaxActive       int           `env:"CACHE_MAX_ACTIVE_CONNECTION"`
	MaxConnLifetime time.Duration `env:"CACHE_CONNECTION_MAX_LIFETIME"`
}

type CacheKey struct {
	FrozenFoodKey  string        `env:"CACHE_KEY_FROZEN_FOOD"`
	FronzenFoodTTL time.Duration `env:"CACHE_TTL_FROZEN_FOOD"`
}

func NewCache(option CacheOption) *redis.Client {
	redisOptions := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", option.Host, option.Port),
		DB:           option.Namespace,
		PoolSize:     option.MaxActive,
		MinIdleConns: option.MaxIdle,
		MaxConnAge:   option.MaxConnLifetime,
	}

	if option.Password != "" {
		redisOptions.Password = option.Password
	}

	if option.IsEnable {
		return redis.NewClient(redisOptions)
	}
	return nil
}
