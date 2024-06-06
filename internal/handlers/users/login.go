package users

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"github.com/radimzitka/zitodo-mongo/internal/app"
	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type payloadLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(c fiber.Ctx) error {
	var payload payloadLoginRequest
	err := c.Bind().Body(&payload)
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type: "DataCheckError",
			Msg:  "Error occured while reading from body",
		})
	}

	var user data.User
	err = db.Coll.Users.FindOne(context.Background(), bson.M{
		"email": payload.Email,
	}).Decode(&user)
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type: "UserNotFound",
			Msg:  "User not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(payload.Password),
	)
	if err != nil {
		return response.SendError(c, 400, response.APIError{
			Type: "PasswordMismatch",
			Msg:  "Password does not match",
		})
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"aud": "user",
	})

	tokenString, err := token.SignedString([]byte(app.State.Cfg.JWTSecret))
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, response.APIError{
			Type: "TokenError",
			Msg:  "Error occured while creating token",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
