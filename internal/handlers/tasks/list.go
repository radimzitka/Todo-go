package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/task"
)

func ListHandler(c fiber.Ctx) error {
	list, err := task.List()

	if err != nil {
		if err.Error() == data.AnyErrorReadingDB {
			return response.SendError(c, 500, response.APIError{
				Type:        "DatabaseAccessFailed",
				Msg:         "Access to MDB failed",
				ErrorNumber: 500,
			})
		}
		return response.SendError(c, 500, response.APIError{
			Type:        "InternalServerError",
			Msg:         "",
			ErrorNumber: 500,
		})
	}

	return c.JSON(list)
}
