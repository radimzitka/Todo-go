package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *payloadCreateSubStep) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}

	return nil
}

func CreateSubstepHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	var payload payloadCreateSubStep
	err = c.Bind().Body(&payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	if err = payload.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updatedTask, err := task.CreateSubstep(&data.SubStep{
		Title: payload.Title,
	}, &id)

	if err != nil {
		return err
	}

	return c.JSON(updatedTask)
}
