package service

import (
	"net/http"
	"testing"
	"time"

	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/repository"
)


func TestNewService_CloseChannel(t *testing.T) {
	t.Cleanup(func() {
		queue = make(chan entity.Task, entity.QueueSize)
	})
	cnf := config.New()
	repo := repository.NewRepository(cnf)
	srv := NewService(repo)
	srv.CloseQueue()
	if val, ok := <-queue; ok {
		t.Errorf("Channel is not closed: %v", val)
	}
}

func TestNewService_AssignTask(t *testing.T) {
	t.Parallel()
	cnf := config.New()
	repo := repository.NewRepository(cnf)
	srv := NewService(repo)
	items := entity.Task{
		Id: "test1",
	}
	id, err := srv.Assign(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if id != items.Id {
		t.Errorf("Unexpected id: %v", id)
	}

}

func TestTaskQueue(t *testing.T) {
	t.Cleanup(func() {
		queue = make(chan entity.Task, entity.QueueSize)
	})
	cnf := config.New()
	repo := repository.NewRepository(cnf)
	srv := NewService(repo)
	go func() {
		srv.StartQueue()
	}()
	queue <- entity.Task{
		Id:      "123",
		Url:     "http://example.com",
		Method:  http.MethodGet,
		ReqBody: "",
	}
	time.Sleep(100 * time.Millisecond)
	if len(queue) != 0 {
		t.Errorf("queue still has items, want empty queue")
	}
	if len(queue) == 0 {
		srv.CloseQueue()
	}
	time.Sleep(100 * time.Millisecond)
}
