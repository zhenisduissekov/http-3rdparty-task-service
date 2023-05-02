package repository

import (
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

const (
	noTaskFoundErrMsg = "no task found"
	uknownTypeErrMsg  = "unknown type error"
)

type RepoCache struct {
	cache *cache.Cache
	cnf   *config.Conf
}

func NewCache(cnf *config.Conf) *RepoCache {
	newCache := cache.New(cnf.Cache.DefaultExpiration, cnf.Cache.CleanupInterval)
	return &RepoCache{
		cnf:   cnf,
		cache: newCache,
	}
}

func (r *RepoCache) Set(key string, value entity.Task) {
	r.cache.Set(key, value, r.cnf.Cache.DefaultExpiration)
}

func (r *RepoCache) Get(key string) (entity.Task, error) {
	val, ok := r.cache.Get(key)
	if !ok {
		return entity.Task{}, errors.New(noTaskFoundErrMsg)
	}
	task, ok := val.(entity.Task)
	if !ok {
		return entity.Task{}, errors.New(uknownTypeErrMsg)
	}
	return task, nil
}
