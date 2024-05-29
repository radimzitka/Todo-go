package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type payloadCreateSubStep struct {
	Title        string    `json:"title"`
	FinishedTime time.Time `json:"finishedTime"`
	Done         bool      `json:"done"`
}

type payloadCreateTask struct {
	Title    string                 `json:"title"`
	SubSteps []payloadCreateSubStep `json:"substeps"`
}

func (p *payloadCreateTask) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}

	return nil
}

func CreateHandler(c fiber.Ctx) error {
	var payload payloadCreateTask
	err := c.Bind().Body(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	if err = payload.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
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
	})

	if err != nil {
		return err
	}

	return c.JSON(insertedTask)
}
