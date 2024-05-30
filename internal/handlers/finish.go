package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FinishHandler(c fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	err = task.FinishByID(&id)
	if err.Error() == "task is already finished" {
		return response.SendError(c, 400, response.APIError{
			Type:        "TaskDone",
			Msg:         "Task is already done",
			ErrorNumber: 400,
		})
	}
	if err.Error() == "error occured while deleting subtask" {
		return response.SendError(c, 500, response.APIError{
			Type:        "DatabaseAccessFailed",
			Msg:         "Access to MDB failed",
			ErrorNumber: 500,
		})
	}
	return c.JSON(fiber.StatusOK)
}
