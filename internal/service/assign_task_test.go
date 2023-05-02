package service

import (
	"testing"

	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

func TestAssignTask(t *testing.T) {
	// create a new repository and service
	repo := NewRepository()
	service := NewService(repo)

	// create a new task to assign
	task := entity.Task{
		Id:             "123",
		Url:            "https://example.com",
		Method:         "GET",
		ReqBody:        "",
		Status:         "",
		HttpStatusCode: 0,
		RespBody:       "",
		Length:         0,
		Headers:        nil,
	}

	// create a new channel to mock the global queue
	mockQueue := make(chan entity.Task, 1)

	// mock the global queue with the mockQueue channel
	queue = mockQueue

	// call the AssignTask function
	_, err := service.AssignTask(task)

	// check if the error is nil
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// check if the task is in the mockQueue
	select {
	case queuedTask := <-mockQueue:
		if queuedTask != task {
			t.Errorf("unexpected task in queue: got %v, want %v", queuedTask, task)
		}
		if queuedTask.Status != statusNew {
			t.Errorf("unexpected task status: got %v, want %v", queuedTask.Status, statusNew)
		}
	default:
		t.Errorf("no task in queue")
	}
}
