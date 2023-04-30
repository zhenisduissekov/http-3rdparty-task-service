package handler

import (
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/logger"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/repository"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/service"
)

type Handler struct {
	repo    *repository.Repository
	service *service.Service
	log     *logger.Logger
}

type Response struct {
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	TechMessage interface{} `json:"techMessage"`
	ActivityId  string      `json:"activityId"`
	Result      interface{} `json:"result"`
}

func New(repo *repository.Repository, service *service.Service, log *logger.Logger) *Handler {
	log.Trace().Msg("handler init")
	return &Handler{
		service: service,
		repo:    repo,
		log:     log,
	}
}
