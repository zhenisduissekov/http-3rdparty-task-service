package service

import (
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/repository"
)


type Task interface {
	Assign(items entity.Task) (string, error)
	StartQueue()
	CloseQueue()
	Check(id string) (entity.Task, error)
}

type Service struct {
	Task
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Task: NewTask(repository),
	}
}
