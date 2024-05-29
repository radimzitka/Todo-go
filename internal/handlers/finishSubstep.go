package handlers

import (
	"github.com/gofiber/fiber/v3"
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

	if err != nil {
		return err
	}

	return c.JSON(task)
}
