package handler

import (
		"github.com/gofiber/fiber/v2"
)

func (h *Handler) CheckTask(f *fiber.Ctx) error {
	id := f.Params("id")

	status, err := h.service.CheckTask(id)
	if err != nil {
		h.log.Error().Err(err).Msg(GetTaskStatusErrMsg)
		return f.Status(fiber.StatusInternalServerError).JSON(&response{
			Status:  StatusError,
			Message: ServerErrMsg,
			Results: err.Error(), //NOTE: would replace this with a template, if it was a real project
		})
	}

	return f.Status(fiber.StatusOK).JSON(&response{
		Status:  StatusSuccess,
		Message: TaskCheckedMsg,
		Results: status,
	})
}
