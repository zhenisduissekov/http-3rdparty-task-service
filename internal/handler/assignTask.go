package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/service"
"github.com/zhenisduissekov/http-3rdparty-task-service/internal/validate"
)

// AssignTask godoc
//
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
	var items service.AssignTaskReq

	err := f.BodyParser(&items)
	if err != nil {
		log.Err(err).Msg(ParseErrMsg)
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: err.Error(), //NOTE 1: would replace this with a template, if it was a real project, if there is logging then it would be useful to log the error
		})
	}

	error := validate.Validate(items)
	if error != nil {
		log.Err(err).Msg(InputParamsValidErrMsg)
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: error, //NOTE: would replace this with a template, if it was a real project
		})
	}

	id, err := h.service.AssignTask(items)
	if err != nil {
		log.Err(err).Msgf(TaskAssignErrMsg)
		return f.Status(fiber.StatusInternalServerError).JSON(&response{
			Status:  StatusError,
			Message: ServerErrMsg,
			Results: err.Error(), //NOTE: would replace this with a template, if it was a real project
		})
	}

	return f.Status(fiber.StatusOK).JSON(&response{
		Status:  StatusSuccess,
		Message: TaskAssignedMsg,
		Results: assignTaskResp{
			Id: id,
		},
	})
}
