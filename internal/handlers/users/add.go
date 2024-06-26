package users

import (
	"errors"
	"strings"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gofiber/fiber/v3"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"github.com/radimzitka/zitodo-mongo/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type payloadCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *payloadCreateUser) ValidateName(c fiber.Ctx) error {
	if len(strings.TrimSpace(p.Name)) < 3 {
		return errors.New("non-valid username")
	}
	return nil
}

func (p *payloadCreateUser) ValidateEmail(c fiber.Ctx) error {
	verifier := emailverifier.NewVerifier()
	ret, err := verifier.Verify(p.Email)
	if err != nil {
		return errors.New("verify email address failed")
	}
	if !ret.Syntax.Valid {
		return errors.New("email address has no valid format")
	}
	return nil
}

func (p *payloadCreateUser) ValidatePassword(c fiber.Ctx) error {
	// overit email
	if len(p.Password) < 10 {
		return errors.New("non-valid password")
	}
	return nil
}

func AddHandler(c fiber.Ctx) error {
	var payload payloadCreateUser
	err := c.Bind().Body(&payload)
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "DataCheckError",
			Msg:         "Error occured when data was readed from Body.",
			ErrorNumber: 400,
		})
	}

	if err = payload.ValidateName(c); err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "NameNotValid",
			Msg:         "Name not valid format (length < 3 ch)",
			ErrorNumber: 400,
		})
	}

	if err = payload.ValidateEmail(c); err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "EmailNotValid",
			Msg:         "Email has not correct format ('user@email.com')",
			ErrorNumber: 400,
		})
	}

	if err = payload.ValidatePassword(c); err != nil {
		return response.SendError(c, 400, response.APIError{
			Type:        "PasswordNotValid",
			Msg:         "Password has not enough length (<10 ch)",
			ErrorNumber: 400,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 15)
	if err != nil {
		return response.SendError(c, 500, response.APIError{
			Type:        "HasPasswordError",
			Msg:         "Error while password hashing",
			ErrorNumber: 500,
		})
	}

	insertedUser, err := user.Add(&data.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
	})

	// Je toto ok?
	if err != nil {
		if err.Error() == data.AnyErrorInsertingUser {
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
