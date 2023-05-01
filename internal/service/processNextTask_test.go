package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
)

func TestService_getItemNotFound(t *testing.T) {
	s := New(&config.Conf{}, nil)
	res, err := s.getItem("test")
	assert.Equal(t, err.Error(), notFoundErrMsg)
	assert.Equal(t, AssignTaskReq{}, res)
}

func TestService_getItemIsNotAssignedReq(t *testing.T) {
	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", "test", 0)
	res, err := s.getItem("test")
	assert.Equal(t, err.Error(), itemIsNotAssignReq)
	assert.Equal(t, AssignTaskReq{}, res)
}

func TestService_getItemNoError(t *testing.T) {
	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", AssignTaskReq{Method: "POST"}, 0)
	res, err := s.getItem("test")
	assert.ErrorIs(t, err, nil)
	assert.Equal(t, AssignTaskReq{Method: "POST"}, res)
}
