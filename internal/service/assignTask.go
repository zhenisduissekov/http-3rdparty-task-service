package service

import (
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

func (s *Service) AssignTask(items AssignTaskReq) (string, error) {
	if s.Cache.ItemCount() > cacheLimit { //NOTE: to control cache size
		s.Cache.DeleteExpired()
		return "", errors.New(cacheLimitReachedErrMsg) //NOTE: just an example of error handling and error message
	}

	id := uuid.New().String()

	items.Status = statusNew
	s.Cache.Set(id, items, cache.DefaultExpiration)
	queue <- id

	return id, nil
}
