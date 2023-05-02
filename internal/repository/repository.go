package repository

import (
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

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
