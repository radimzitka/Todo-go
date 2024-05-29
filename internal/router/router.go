package router

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/handlers"
)

func Init() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Jede.to")
	})
	app.Post("/tasks", handlers.CreateHandler)
	app.Delete("/tasks/:id", handlers.DeleteHandler)
	app.Get("/tasks", handlers.ListHandler)
	app.Put("/tasks/:id", handlers.FinishHandler)
	app.Put("/tasks/:id/:sid", handlers.FinishSubstepHandler)
	app.Post("/tasks/:id", handlers.CreateSubstepHandler)
	app.Delete("/tasks/:id/:sid", handlers.DeleteSubstepHandler)
	/* app.Patch("/theses/:id/title", thesis.UpdateTitleHandler) */

	log.Fatalln(app.Listen(":3000"))
}
