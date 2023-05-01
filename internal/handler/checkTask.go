package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/validate"
)

func (h *Handler) CheckTask(f *fiber.Ctx) error {
	var items TaskStatusReq
	err := f.QueryParser(&items)
	if err != nil {
		h.log.Error().Err(err).Msg(ParseErrMsg)
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: err.Error(), //NOTE: would replace this with a template, if it was a real project
		})
	}

	error := validate.Validate(items)
	if error != nil {
		h.log.Error().Err(err).Msg(InputParamsValidErrMsg)
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: error, //NOTE: would replace this with a template, if it was a real project
		})
	}

	status, err := h.service.CheckTask(items.Id)
	if err != nil {
		h.log.Error().Err(err).Msg(GetTaskStatusErrMsg)
		return f.Status(fiber.StatusInternalServerError).JSON(&response{
			Status:  StatusError,
			Message: ServerErrMsg,
			Results: err.Error(), //NOTE: would replace this with a template, if it was a real project
		})
	}

	return f.Status(fiber.StatusOK).JSON(&response{
		Status:  StatusSuccess, //todo: fill this
		Message: TaskCheckedMsg,
		Results: status,
	})
}
