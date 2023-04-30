package handler

import (
	"github.com/gofiber/fiber/v2"
)

type TaskStatusReq struct {
	Id string `json:"id"`
}

type TaskStatusResp struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}



func (h *Handler) CheckTask(f *fiber.Ctx) error {
	var items TaskStatusReq
	err := f.BodyParser(&items)
	if err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(&response{
			Status:  StatusError,
			Message: BadRequestErrMsg,
			Results: err.Error(),
		})
	}
	
	// todo: validation
	
	
	status, err := h.service.CheckTask(items.Id)
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(&response{
			Status:  StatusError,
			Message: ServerErrMsg,
			Results: err.Error(),
		})
	}
	
	return f.Status(fiber.StatusOK).JSON(&AssignTaskResp{
		Id: status,
	})
}
