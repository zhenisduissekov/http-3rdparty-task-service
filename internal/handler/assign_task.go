package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// AssignTask godoc
// @Security BasicAuth
//
//	@Summary		Назначить задачу
//	@Description	назначение задачи.
//	@Tags			task
//
// @Accept  json
// @Produce  json
// @Param	input	body	AssignTaskReq		true "тело запроса"
//
//	@Success		200		object response		"успешный ответ"
//	@Failure		400		object response		"ошибка запроса"
//	@Failure		406		object response		"ошибка валидации"
//	@Failure		500		object response		"ошибка сервера"
//	@Router			/api/v1/task [post]
func (h *Handler) AssignTask(f *fiber.Ctx) error {
	var items AssignTaskReq
	err := f.BodyParser(&items)
	if err != nil {
		log.Err(err).Msg(ParseErrMsg)
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: err.Error(), //NOTE 1: would replace this with a template, if it was a real project, if there is logging then it would be useful to log the error
		})
	}

	convItems, err := items.convert2service()
	if err != nil {
		log.Err(err).Msg(InputParamsValidErrMsg)
		return f.Status(fiber.StatusNotAcceptable).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: err.Error(), //NOTE: I would replace this with a template, if it was a real project
		})
	}

	id, err := h.service.Task.Assign(convItems)
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
		Results: AssignTaskResp{
			Id: id,
		},
	})
}
