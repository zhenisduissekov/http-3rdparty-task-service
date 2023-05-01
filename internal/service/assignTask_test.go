package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
)

func TestService_AssignTask_NoCacheInitialized(t *testing.T) {
	cnf := &config.Conf{}
	s := New(cnf, nil)
	s.Cache = nil
	resp, err := s.AssignTask(AssignTaskReq{})
	assert.Equal(t, resp, "")
	assert.Equal(t, cacheNotInitializedErrMsg, err.Error())
}


func TestService_AssignTask_LimitHit(t *testing.T) {
	s := New(&config.Conf{}, nil)
	for i:=0; i<1000; i++ {
		s.Cache.Set(fmt.Sprintf("key %d", i), i, 0)
	}
	s.Cache.Set("test", "test", 0)
	resp, err := s.AssignTask(AssignTaskReq{})
	assert.Equal(t, resp, "")
	assert.Equal(t, cacheLimitReachedErrMsg, err.Error())
	s.Cache = nil
}

func TestService_AssignTask_GetID(t *testing.T) {
	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", "test", 0)
	resp, err := s.AssignTask(AssignTaskReq{})
	assert.NotEqual(t, resp, "")
	assert.ErrorIs(t, err, nil)
	s.Cache = nil
	close(queue)
}