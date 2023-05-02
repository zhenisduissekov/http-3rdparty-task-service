package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
)

func TestService_CheckTask_NoCacheInitialized(t *testing.T) {
	cnf := &config.Conf{}
	s := New(cnf, nil)
	s.Cache = nil
	resp, err := s.CheckTask("test")

	assert.Equal(t, cacheNotInitializedErrMsg, err.Error())
	assert.Equal(t, statusError, resp.Status)
	assert.Equal(t, "test", resp.Id)

}

func TestService_CheckTask_Not_Found(t *testing.T) {
	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", "test", 0)
	resp, err := s.CheckTask("test2")
	assert.Equal(t, resp.Id, "test2")
	assert.Equal(t, resp.Status, statusError)
	assert.Equal(t, notFoundErrMsg, err.Error())

	s.Cache = nil
}

func TestService_CheckTask_Hit_Done(t *testing.T) {
	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", AssignTaskResp{
		Id:     "test",
		Status: statusDone,
	}, 0)
	resp, err := s.CheckTask("test")
	assert.Equal(t, resp.Id, "test")
	assert.Equal(t, resp.Status, statusDone)
	assert.NoError(t, err)
	s.Cache = nil
}

func TestService_CheckTask_Hit_New(t *testing.T) {
	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", AssignTaskReq{
		Status: statusNew,
	}, 0)
	resp, err := s.CheckTask("test")
	assert.Equal(t, resp.Id, "test")
	assert.Equal(t, resp.Status, statusInProcess)
	assert.NoError(t, err)
	s.Cache = nil
}

func TestService_CheckTask_Hit_UknownType(t *testing.T) {
	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", "string", 0)
	resp, err := s.CheckTask("test")
	assert.Equal(t, resp.Id, "test")
	assert.Equal(t, resp.Status, statusError)
	assert.Equal(t, unexpectedTypeErrMsg, err.Error())
	s.Cache = nil
}
