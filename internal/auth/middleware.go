package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"github.com/radimzitka/zitodo-mongo/internal/app"
	"github.com/radimzitka/zitodo-mongo/internal/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateJWT(aud string) fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return response.SendError(c, fiber.StatusUnauthorized, response.APIError{
				Type: "MISSING_TOKEN",
				Msg:  "The token is missing in you Authorization header",
			})
		}
		claims, err := validateTokenWithClaims(token, aud)
		if err != nil {
			return response.SendError(c, fiber.StatusUnauthorized, response.APIError{
				Type: "INVALID_TOKEN",
				Msg:  err.Error(),
			})
		}
		sub := claims["sub"].(string)
		userId, err := primitive.ObjectIDFromHex(sub)
		if err != nil {
			return response.SendError(c, fiber.StatusUnauthorized, response.APIError{
				Type: "INTERNAL_ERROR",
				Msg:  "Error while converting the user ID",
			})
		}
		c.Locals("userId", userId)
		return c.Next()
	}
}

func validateTokenWithClaims(tokenString string, aud string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(app.State.Cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		// test the expiration
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil, errors.New("the token is expired")
		}
		if !claims.VerifyAudience(aud, true) {
			return nil, errors.New("invalid audience")
		}
	} else {
		fmt.Println(err)
	}
	return claims, nil
}
