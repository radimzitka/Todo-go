package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "InvalidID",
			Msg:         "Non-existing ID of task",
			ErrorNumber: 400,
		})
	}

	err = task.DeleteByID(&id)
	if err != nil {
		if err.Error() == data.TASK_NOT_FOUND {
			return response.SendError(c, 404, response.APIError{
				Type:        "NonExistTask",
				Msg:         "Trying to delete non-exist task.",
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
