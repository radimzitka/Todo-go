package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FinishSubstepHandler(c fiber.Ctx) error {
	tid, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	sid, err := primitive.ObjectIDFromHex(c.Params("sid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid substep ID")
	}

	task, err := task.FinishSubstep(&tid, &sid)
	if err.Error() == "subtask already finished" {
		return response.SendError(c, 400, response.APIError{
			Type:        "SubtaskDone",
			Msg:         "Subtask is already done",
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

	if err.Error() == "error occured during finding subtask" {
		return response.SendError(c, 500, response.APIError{
			Type:        "DatabaseAccessFailed",
			Msg:         "Access to MDB failed",
			ErrorNumber: 500,
		})
	}

	return c.JSON(task)
}
