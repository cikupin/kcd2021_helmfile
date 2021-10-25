package logic

import (
	"github.com/cikupin/kcd2021_helmfile/internal/bootstrap"
	"github.com/go-gorp/gorp/v3"
	"github.com/go-redis/redis/v8"
)

type Logic struct {
	db     *gorp.DbMap
	cache  *redis.Client
	config bootstrap.ConfigObjects
}

func NewLogic(dbMap *gorp.DbMap, cacheClient *redis.Client, cfg bootstrap.ConfigObjects) *Logic {
	return &Logic{
		db:     dbMap,
		cache:  cacheClient,
		config: cfg,
	}
}
