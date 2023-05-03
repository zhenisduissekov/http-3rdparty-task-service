package repository

import (
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

type Task struct {
	Id      string
	Method  string
	Url     string
	Headers map[string]string
	ReqBody string
	Status  string
}

func (a *Task) converterToRepository(items entity.Task) {
	a.Id = items.Id
	a.Url = items.Url
	a.Method = items.Method
	a.ReqBody = items.ReqBody
	a.Status = items.Status
	a.Headers = items.Headers
}


type Storage interface {
	Set(key string, value entity.Task)
	Get(key string) (entity.Task, error)
}

type Repository struct {
	Storage
}

func NewRepository(cnf *config.Conf) *Repository {
	return &Repository{
		Storage: NewCache(cnf),
	}
}