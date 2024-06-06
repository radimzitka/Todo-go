package router

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/auth"
	handlersTask "github.com/radimzitka/zitodo-mongo/internal/handlers/tasks"
	handlersUser "github.com/radimzitka/zitodo-mongo/internal/handlers/users"
	"github.com/radimzitka/zitodo-mongo/internal/response"
)

func Init() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("App is working")
	})
	app.Post("/auth/register", handlersUser.AddHandler)
	app.Post("/auth/login", handlersUser.LoginHandler)
	app.Get("/user/list", handlersUser.ListUserHandler)

	app.Get("/tasks", handlersTask.ListHandler, auth.ValidateJWT("user"))
	app.Post("/tasks", handlersTask.CreateHandler, auth.ValidateJWT("user"))
	app.Delete("/tasks/:id", handlersTask.DeleteHandler, auth.ValidateJWT("user"))
	app.Put("/tasks/:id/finish", handlersTask.FinishHandler, auth.ValidateJWT("user"))
	app.Post("/tasks/:id/substeps", handlersTask.CreateSubstepHandler, auth.ValidateJWT("user"))
	app.Put("/tasks/:id/substeps/:sid/finish", handlersTask.FinishSubstepHandler, auth.ValidateJWT("user"))
	app.Delete("/tasks/:id/substeps/:sid", handlersTask.DeleteSubstepHandler, auth.ValidateJWT("user"))

	//app.Delete("/auth/register/:id", handlersUser.DeleteHandler)

	app.All("*", func(c fiber.Ctx) error {
		return response.SendError(c, 404, response.APIError{
			Type:        "Not found",
			Msg:         "Requested page not found.",
			ErrorNumber: 404,
		})
	})
	log.Fatalln(app.Listen(":3000"))

}
