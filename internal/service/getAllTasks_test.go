package service

import (
	"testing"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
)

func TestService_GetAllTasks_NoCacheInitialized(t *testing.T) {
	cnf := &config.Conf{}
	s := New(cnf, nil)
	s.Cache = nil
	resp, err := s.GetAllTasks()
	var expected map[string]cache.Item
	assert.Equal(t, expected, resp)
	assert.Equal(t, cacheNotInitializedErrMsg, err.Error())
}

func TestService_TwoElements(t *testing.T) {
	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", "test", 0)
	s.Cache.Set("test1", "test", 0)

	resp, err := s.GetAllTasks()
	assert.Equal(t, len(resp), 2)

	assert.NoError(t, err)
	s.Cache = nil
}
