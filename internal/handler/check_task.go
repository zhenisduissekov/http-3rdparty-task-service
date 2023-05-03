package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// CheckTask godoc
// @Security BasicAuth
//
//	@Summary		Проверить статус задачи
//	@Description	проверка статуса задачи.
//	@Tags			task
//	@Accept			*/*
//	@Produce		json
//
// @Param id path string true "comment ID"
//
//	@Success		200		object response		"успешный ответ"
//	@Failure		400		object response		"ошибка запроса"
//	@Failure		500		object response		"ошибка сервера"
//	@Router			/api/v1/task/{id} [get]
func (h *Handler) CheckTask(f *fiber.Ctx) error {
	id := f.Params("id") //NOTE: if api/v1/task endpoint is removed then this would enforce the id to be passed in the query
	if id == "" {
		log.Warn().Msgf(NoIdErrMsg)
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: NoIdErrMsg,
		})
	}

	task, err := h.service.Task.Check(id)
	if err != nil {
		log.Error().Err(err).Msg(GetTaskStatusErrMsg)
		return f.Status(fiber.StatusInternalServerError).JSON(&response{
			Status:  StatusError,
			Message: ServerErrMsg,
			Results: err.Error(), //NOTE: would replace this with a template, if it was a real project
		})
	}

	return f.Status(fiber.StatusOK).JSON(&response{
		Status:  StatusSuccess,
		Message: TaskCheckedMsg,
		Results: task,
	})
}
