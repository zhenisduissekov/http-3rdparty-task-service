package service

import (
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

func (s *Service) GetAllTasks() (map[string]cache.Item, error) {
	if s.Cache == nil {
		return nil, errors.New(cacheNotInitializedErrMsg)
	}

	return s.Cache.Items(), nil
}
