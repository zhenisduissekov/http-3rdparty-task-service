package service

import (
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/logger"
)

type Service struct {
	log        *logger.Logger
	httpClient *http.Client
	Cache      *cache.Cache
}

func New(cnf *config.Conf, log *logger.Logger) *Service {
	return &Service{
		log: log,
		httpClient: &http.Client{
			Timeout: requestTimeout,
		},
		Cache: cache.New(cnf.Cache.DefaultExpiration, cnf.Cache.CleanupInterval),
	}
}

func (s *Service) TaskQueue() error {
	ticker := time.NewTicker(tickPeriod)
	defer ticker.Stop()

	for {
		select {
		case nextTask, ok := <-queue:
			s.log.Info().Msg(taskReceivedMsg)
			if !ok {
				s.log.Warn().Msg(channelWasClosedMsg)
				return nil
			}

			s.processNextTask(nextTask)
		case <-ticker.C:
			s.log.Trace().Msg(tickMsg)
		}
	}

	return nil
}
