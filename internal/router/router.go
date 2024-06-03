package router

import (
	"log"

	"github.com/gofiber/fiber/v3"
	handlersTask "github.com/radimzitka/zitodo-mongo/internal/handlers/tasks"
	handlersUser "github.com/radimzitka/zitodo-mongo/internal/handlers/users"
)

func Init() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("App is working")
	})
	app.Get("/tasks", handlersTask.ListHandler)
	app.Post("/tasks", handlersTask.CreateHandler)
	app.Delete("/tasks/:id", handlersTask.DeleteHandler)
	app.Put("/tasks/:id/finish", handlersTask.FinishHandler)
	app.Post("/tasks/:id/substeps", handlersTask.CreateSubstepHandler)
	app.Put("/tasks/:id/substeps/:sid/finish", handlersTask.FinishSubstepHandler)
	app.Delete("/tasks/:id/substeps/:sid", handlersTask.DeleteSubstepHandler)

	app.Post("/users", handlersUser.AddHandler)
	app.Delete("/users/:id", handlersUser.DeleteHandler)
	app.Get("/users", handlersUser.ListHandler)

	log.Fatalln(app.Listen(":3000"))
}
