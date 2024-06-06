package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FinishSubstepHandler(c fiber.Ctx) error {
	userID := c.Locals("userId").(primitive.ObjectID)
	tid, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	sid, err := primitive.ObjectIDFromHex(c.Params("sid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid substep ID")
	}

	task, err := task.FinishSubstep(&tid, &sid, &userID)
	if err != nil {
		if err.Error() == data.SubtaskFinished {
			return response.SendError(c, 400, response.APIError{
				Type:        "SubtaskDone",
				Msg:         "Subtask is already done",
				ErrorNumber: 400,
			})
		}
		if err.Error() == data.TaskNotFound {
			return response.SendError(c, 404, response.APIError{
				Type:        "TaskNotFound",
				Msg:         "",
				ErrorNumber: 404,
			})
		}
		if err.Error() == data.SubtaskNotFound {
			return response.SendError(c, 404, response.APIError{
				Type:        "SubtaskNotFound",
				Msg:         "",
				ErrorNumber: 404,
			})
		}
		if err.Error() == data.AnyErrorDeletingTask || err.Error() == data.AnyErrorDeletingSubtask {
			return response.SendError(c, 500, response.APIError{
				Type:        "DatabaseAccessFailed",
				Msg:         "Access to MDB failed",
				ErrorNumber: 500,
			})
		}
		return response.SendError(c, 500, response.APIError{
			Type:        "InternalServerError",
			Msg:         "",
			ErrorNumber: 500,
		})
	}

	return c.JSON(task)
}
