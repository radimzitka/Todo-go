package handlers

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type payloadCreateTask struct {
	Title    string                 `json:"title"`
	SubSteps []payloadCreateSubStep `json:"substeps"`
}

type payloadCreateSubStep struct {
	Title        string    `json:"title"`
	FinishedTime time.Time `json:"finishedTime"`
	Done         bool      `json:"done"`
}

func (p *payloadCreateTask) ValidateTitle(c fiber.Ctx) error {
	if strings.TrimSpace(p.Title) == "" {
		return errors.New("non-valid title")
	}
	return nil
}

func CreateHandler(c fiber.Ctx) error {
	userID := c.Locals("userId").(primitive.ObjectID)
	var payload payloadCreateTask
	err := c.Bind().Body(&payload)
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "DataCheckError",
			Msg:         "Error occured when data was readed from Body.",
			ErrorNumber: 400,
		})
	}
	if err = payload.ValidateTitle(c); err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "TitleNotValid",
			Msg:         "Title for task has not valid format.",
			ErrorNumber: 400,
		})
	}

	substeps := make([]*data.SubStep, len(payload.SubSteps))
	for i, substep := range payload.SubSteps {
		aid := primitive.NewObjectID()
		substeps[i] = &data.SubStep{
			ID:           &aid,
			Title:        substep.Title,
			Done:         substep.Done,
			FinishedTime: &substep.FinishedTime,
		}
	}

	timeNow := time.Now()
	insertedTask, err := task.Create(&data.Item{
		Title:        payload.Title,
		TimeAdded:    &timeNow,
		SubSteps:     substeps,
		Finished:     false,
		TimeFinished: nil,
		UserID:       &userID,
	})

	// Je toto ok?
	if err != nil {
		if err.Error() == data.ANY_ERROR_INSERTING_TASK {
			return response.SendError(c, 500, response.APIError{
				Type:        "TaskCreateError",
				Msg:         "Error during creating new task",
				ErrorNumber: 500,
			})
		}
		return response.SendError(c, 500, response.APIError{
			Type:        "InternalServerError",
			Msg:         "",
			ErrorNumber: 500,
		})
	}

	return c.JSON(insertedTask)
}
