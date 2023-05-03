package service

import (
	"net/http"
	"time"

	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/repository"
)

const (
	queueSize                 = 100
	tickPeriod                = 1 * time.Second
	statusNew                 = "new"
	statusDone                = "done"
	statusError               = "error"
	taskReceivedMsg           = "task received"
	channelWasClosedMsg       = "channel was closed, exiting task queue"
	tickMsg                   = "tick"
	failedToMakeRequestErrMsg = "failed to make request"
)

var queue = make(chan entity.Task, queueSize)

type Task interface {
	Assign(items entity.Task) (string, error)
	StartQueue()
	CloseQueue()
	Check(id string) (entity.Task, error)
}

type Service struct {
	Task
}

type NewService struct {
	httpClient *http.Client
	repository *repository.Repository
}

func New(repository *repository.Repository, cnf *config.Conf) *Service {
	return &Service{
		Task: &NewService{
			repository: repository,
		},
	}
}
