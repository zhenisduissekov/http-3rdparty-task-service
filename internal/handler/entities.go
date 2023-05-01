package handler

import (
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/logger"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/service"
)

const (
	StatusSuccess    = "success"
	StatusError      = "error"
	BadRequestErrMsg = "please check your request"
	ServerErrMsg     = "server error, please try again later"
	TaskAssignedMsg  = "task assigned"
	TaskCheckedMsg  	= "task checked"

	ParseErrMsg            = "could not parse query/body"
	InputParamsValidErrMsg = "input parameters validation error"
	TaskAssignErrMsg       = "assign task error"
	GetTaskStatusErrMsg    = "get task status error"
)

type Handler struct {
	service *service.Service
	log     *logger.Logger
}

func New(service *service.Service, log *logger.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Results interface{} `json:"result"`
}

type assignTaskResp struct {
	Id string `json:"id"`
}

type TaskStatusReq struct {
	Id string `json:"id" required:"true"`
}

type TaskStatusResp struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}
