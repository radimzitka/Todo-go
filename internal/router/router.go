package router

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/handlers"
)

func Init() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("App is working")
	})
	app.Get("/tasks", handlers.ListHandler)
	app.Post("/tasks", handlers.CreateHandler)
	app.Delete("/tasks/:id", handlers.DeleteHandler)
	app.Put("/tasks/:id/finish", handlers.FinishHandler)
	app.Post("/tasks/:id/substeps", handlers.CreateSubstepHandler)
	app.Put("/tasks/:id/substeps/:sid/finish", handlers.FinishSubstepHandler)
	app.Delete("/tasks/:id/substeps/:sid", handlers.DeleteSubstepHandler)

	log.Fatalln(app.Listen(":3000"))
}
