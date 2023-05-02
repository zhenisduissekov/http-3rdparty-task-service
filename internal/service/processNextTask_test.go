package service

import (
	"testing"

	"github.com/h2non/gock"
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

func TestService_processNextTask(t *testing.T) {
	gock.New("http://shmoogle.com").
		Get("/").
		Reply(200).
		JSON(``)

	s := New(&config.Conf{}, nil)
	s.Cache.Set("test", AssignTaskReq{Method: "GET", Url: "http://shmoogle.com"}, 0)
	s.processNextTask("test")

	received, found := s.Cache.Get("test")
	expected := AssignTaskResp{Id: "test", Status: "done", HttpStatusCode: 200, Headers: map[string]string{"Content-Type": "application/json"}, Length: 0, Body: ""}
	assert.Equal(t, expected, received)
	assert.True(t, found)
}
