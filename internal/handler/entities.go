package handler

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/service"
)

const (
	StatusSuccess          = "success"
	StatusError            = "error"
	statusNew              = "new"
	BadRequestErrMsg       = "please check your request"
	NotAcceptableErrMsg    = "not acceptable"
	ServerErrMsg           = "server error, please check your request or try again"
	TaskAssignedMsg        = "task assigned"
	TaskCheckedMsg         = "task checked"
	NoIdErrMsg             = "no id provided"
	ParseErrMsg            = "could not parse query/body"
	InputParamsValidErrMsg = "input parameters validation error"
	TaskAssignErrMsg       = "assign task error"
	GetTaskStatusErrMsg    = "get task status error"
)

type Handler struct {
	service *service.NewService
}

func New(service *service.NewService) *Handler {
	return &Handler{
		service: service,
	}
}

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Results interface{} `json:"result"`
}

type AssignTaskResp struct {
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

type AssignTaskReq struct {
	Method  string            `json:"method" example:"GET"`
	Url     string            `json:"url" example:"http://google.com"`
	Headers map[string]string `json:"headers" example:"'Authentication'': 'Basic bG9naW46cGFzc3dvcmQ='"`
	ReqBody string            `json:"body" example:"some body"`
}

func (a *AssignTaskReq) convert2service() (entity.Task, error) {
	if a.Method == "" || a.Url == "" || (a.Method != "GET" && a.Method != "POST") {
		return entity.Task{}, errors.New(NotAcceptableErrMsg)
	}

	return entity.Task{
		Id:      uuid.New().String(),
		Method:  a.Method,
		Url:     a.Url,
		Headers: a.Headers,
		ReqBody: a.ReqBody,
		Status:  statusNew,
	}, nil
}

type CheckTaskResp struct {
	Method         string            `json:"method" validate:"required,min=3,max=6,alpha,uppercase" example:"GET"`
	Url            string            `json:"url" validate:"required" example:"http://google.com"`
	Headers        map[string]string `json:"headers" validate:"omitempty" example:"\"Authentication\": \"Basic bG9naW46cGFzc3dvcmQ=\""`
	ReqBody        []byte            `json:"body" validate:"omitempty" `
	Status         string            `json:"status" validate:"omitempty" example:"in_process"` //todo: seperate to new structure
	HttpStatusCode int               `json:"http_status_code" validate:"omitempty"`
	Length         int               `json:"length" validate:"omitempty"`
}
