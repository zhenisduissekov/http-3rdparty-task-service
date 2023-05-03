package service

import (
					"time"

	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

func (s *NewService) Assign(items entity.Task) (string, error) {
	items.Status = statusNew
	s.repository.Set(items.Id, items)
	queue <- items
	return items.Id, nil
}

func (s *NewService) Check(id string) (entity.Task, error) {
	return s.repository.Get(id)
}

func (s *NewService) StartQueue() {
	ticker := time.NewTicker(tickPeriod)
	defer ticker.Stop()

	for {
		select {
		case nextTask, ok := <-queue:
			if !ok {
				log.Warn().Msg(channelWasClosedMsg)
				return
			}
			log.Info().Msg(taskReceivedMsg)
			s.processNextTask(nextTask)
		case <-ticker.C:
			log.Debug().Msg(tickMsg)
		}
	}

	return
}

func (s *NewService) CloseQueue() {
	log.Warn().Msg("closing channel")
	close(queue)
}

func (s *NewService) processNextTask(items entity.Task) {
	status := statusDone
	body, headers, statusCode, err := s.repository.MakeRequest(items)
	if err != nil {
		log.Error().Err(err).Msg(failedToMakeRequestErrMsg)
		status = statusError
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

