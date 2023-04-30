package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/service"
"github.com/zhenisduissekov/http-3rdparty-task-service/internal/validate"
)

const (
	StatusSuccess    = "success"
	StatusError      = "error"
	BadRequestErrMsg = "please check your request"
	ServerErrMsg     = "server error, please try again later"
)

type AssignTaskReq struct {
	Method  string            `json:"method" validate:"required,min=3,max=6,alphanum,uppercase" example:"GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD, CONNECT, TRACE"`
	Url     string            `json:"url" validate:"required" example:"http://google.com"`
	Headers map[string]string `json:"headers" validate:"omitempty" example:"\"Authentication\": \"Basic bG9naW46cGFzc3dvcmQ=\""`
	Payload []byte            `json:"payload" validate:"omitempty" example:"{\"name\":\"John\"}"`
}

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Results interface{} `json:"result"`
}

type AssignTaskResp struct {
	Id string `json:"id"`
}

// AssignTask godoc
//	@Summary		Назначить задачу
//	@Description	назначение задачи.
//	@Tags			task
//	@Accept			*/*
//	@Produce		json
//	@Success		200		string		"успешный ответ"
//	@Failure		400		string		"ошибка запроса"
//	@Failure		500		string		"ошибка сервера"
//	@Router			/api/v1/task [get]
func (h *Handler) AssignTask(f *fiber.Ctx) error {
	var items AssignTaskReq

	err := f.BodyParser(&items)
	if err != nil {
		log.Err(err).Msg("could not parse query")
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: err.Error(),
		})
	}

	error := validate.Validate(items)
	if error != nil {
		log.Err(err).Msg("input parameters validation failed")
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: error,
		})

	}
	
	srv := service.New()
	
	id, err := srv.AssignTask(items.Url, items.Method, items.Headers, items.Payload)
	if err != nil {
		log.Err(err).Msgf("could not get release")
		return f.Status(fiber.StatusInternalServerError).JSON(&response{
			Status:  StatusError,
			Message: ServerErrMsg,
			Results: err.Error(),
		})
	}

	return f.Status(fiber.StatusOK).JSON(&response{
		Status:  StatusSuccess,
		Message: "task assigned",
		Results: AssignTaskResp{
			Id: id,
		},
	})
}
