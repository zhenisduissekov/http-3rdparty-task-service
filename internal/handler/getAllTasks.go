package handler

import (
	"github.com/gofiber/fiber/v2"
)


func (h *Handler) GetAllTasks (f *fiber.Ctx) error {
	allTasks, err := h.service.GetAllTasks()
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
		Results: allTasks,
	})
}
