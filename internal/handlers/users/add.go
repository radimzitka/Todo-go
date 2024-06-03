package users

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/user"
)

type payloadCreateUser struct {
	Username string `json:"title"`
}

func (p *payloadCreateUser) ValidateTitle(c fiber.Ctx) error {
	if strings.TrimSpace(p.Username) == "" {
		return errors.New("non-valid username")
	}
	return nil
}

func AddHandler(c fiber.Ctx) error {
	var payload payloadCreateUser
	err := c.Bind().Body(&payload)
	if err != nil {
		// Je toto spravne odeslani chyby?
		return response.SendError(c, 400, response.APIError{
			Type:        "DataCheckError",
			Msg:         "Error occured when data was readed from Body.",
			ErrorNumber: 400,
		})
	}

	// Proc toto nefunguje?
	if err = payload.ValidateTitle(c); err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "UsernameNotValid",
			Msg:         "Username for task has not valid format.",
			ErrorNumber: 400,
		})
	}

	//substeps := make([]*data.SubStep, len(payload.SubSteps))
	insertedUser, err := user.Add(&data.User{
		Username: payload.Username,
	})

	// Je toto ok?
	if err != nil {
		if err.Error() == data.ANY_ERROR_INSERTING_USER {
			return response.SendError(c, 500, response.APIError{
				Type:        "TaskCreateError",
				Msg:         "Error during creating new user",
				ErrorNumber: 500,
			})
		}
		return response.SendError(c, 500, response.APIError{
			Type:        "InternalServerError",
			Msg:         "",
			ErrorNumber: 500,
		})
	}

	return c.JSON(insertedUser)
}
