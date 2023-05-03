package service

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/repository"
)

var queue = make(chan entity.Task, entity.QueueSize)

type TaskService struct {
	httpClient *http.Client
	repository *repository.Repository
}

func NewTask(repository *repository.Repository) *TaskService {
	return &TaskService{
		repository: repository,
		httpClient: &http.Client{},
	}
}

func (s *TaskService) Assign(items entity.Task) (string, error) {
	items.Status = entity.StatusNew
	s.repository.Set(items.Id, items)
	queue <- items
	return items.Id, nil
}

func (s *TaskService) Check(id string) (entity.Task, error) {
	return s.repository.Get(id)
}

func (s *TaskService) StartQueue() {
	ticker := time.NewTicker(entity.TickPeriod)
	defer ticker.Stop()

	for {
		select {
		case nextTask, ok := <-queue:
			if !ok {
				log.Warn().Msg(entity.ChannelWasClosedMsg)
				return
			}
			log.Info().Msg(entity.TaskReceivedMsg)
			s.processNextTask(nextTask)
		case <-ticker.C:
			log.Debug().Msg(entity.TickMsg)
		}
	}

	return
}

func (s *TaskService) CloseQueue() {
	log.Warn().Msg("closing channel")
	close(queue)
}

func (s *TaskService) processNextTask(items entity.Task) {
	status := entity.StatusDone
	body, headers, statusCode, err := s.repository.MakeRequest(items)
	if err != nil {
		log.Error().Err(err).Msg(entity.FailedToMakeRequestErrMsg)
		status = entity.StatusError
	}

	s.repository.Set(items.Id, entity.Task{
		Id:             items.Id,
		Url:            items.Url,
		Method:         items.Method,
		Status:         status,
		HttpStatusCode: statusCode,
		ReqBody:        items.ReqBody,
		RespBody:       string(body),
		Length:         len(body),
		Headers:        headers,
	})
}
