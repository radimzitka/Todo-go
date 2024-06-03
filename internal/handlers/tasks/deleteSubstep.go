package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteSubstepHandler(c fiber.Ctx) error {
	tid, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "InvalidID",
			Msg:         "Non-existing ID of task",
			ErrorNumber: 400,
		})
	}

	sid, err := primitive.ObjectIDFromHex(c.Params("sid"))
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "InvalidSubtaskID",
			Msg:         "Non-existing ID of subtask",
			ErrorNumber: 400,
		})
	}

	err = task.DeleteSubstep(&tid, &sid)
	if err != nil {
		if err.Error() == data.ANY_ERROR_DELETING_SUBTASK {
			return response.SendError(c, 500, response.APIError{
				Type:        "DatabaseAccessFailed",
				Msg:         "Access to MDB failed",
				ErrorNumber: 500,
			})
		}
		if err.Error() == data.TASK_NOT_FOUND {
			return response.SendError(c, 404, response.APIError{
				Type:        "TaskNotFound",
				Msg:         "Task was not found",
				ErrorNumber: 404,
			})
		}

		if err.Error() == data.SUBTASK_NOT_FOUND {
			return response.SendError(c, 404, response.APIError{
				Type:        "SubtaskNotFound",
				Msg:         "Subtask was not found",
				ErrorNumber: 404,
			})
		}
		return response.SendError(c, 500, response.APIError{
			Type:        "InternalServerError",
			Msg:         "",
			ErrorNumber: 500,
		})
	}
	return c.JSON(fiber.StatusOK)
}
