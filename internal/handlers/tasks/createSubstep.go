package handlers

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *payloadCreateSubStep) Validate() error {
	if strings.TrimSpace(p.Title) == "" {
		return errors.New("title is required")
	}

	return nil
}

func CreateSubstepHandler(c fiber.Ctx) error {
	userID := c.Locals("userId").(primitive.ObjectID)
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "InvalidID",
			Msg:         "Non-existing ID of task",
			ErrorNumber: 400,
		})
	}

	var payload payloadCreateSubStep
	err = c.Bind().Body(&payload)
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "DataCheckError",
			Msg:         "Error occured when data was readed from Body.",
			ErrorNumber: 400,
		})
	}
	if err = payload.Validate(); err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "TitleNotValid",
			Msg:         "Title for subtask has not valid format.",
			ErrorNumber: 400,
		})
	}

	updatedTask, err := task.CreateSubstep(&data.SubStep{
		Title: payload.Title,
	}, &id, &userID)
	if err != nil {
		if err.Error() == data.AnyErrorInsertingSubtask {
			return response.SendError(c, 500, response.APIError{
				Type:        "DatabaseAccessFailed",
				Msg:         "Error during database access",
				ErrorNumber: 500,
			})
		}
		if err.Error() == data.TaskNotFound {
			return response.SendError(c, 404, response.APIError{
				Type:        "TaskNotFound",
				Msg:         "task not found",
				ErrorNumber: 404,
			})
		}
		return response.SendError(c, 500, response.APIError{
			Type:        "InternalServerError",
			Msg:         "",
			ErrorNumber: 500,
		})
	}

	return c.JSON(updatedTask)
}
