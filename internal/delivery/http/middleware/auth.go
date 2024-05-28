package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	auth_token_usecase "github.com/krissukoco/go-project-layout/internal/usecase/auth_token"
)

var (
	authKey = &localKey{"auth"}
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

func NewAuthMiddleware(
	uc auth_token_usecase.Usecase,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		head := c.Get("Authorization")
		if head == "" {
			return ErrUnauthorized
		}
		split := strings.Split(head, " ")
		if len(split) != 2 {
			return ErrUnauthorized
		}
		if strings.ToLower(split[0]) != "Bearer" {
			return ErrUnauthorized
		}

		// Validate token
		userId, err := uc.Validate(c.Context(), split[1])
		if err != nil {
			// TODO: process error
			return ErrUnauthorized
		}

		// Set locals
		c.Locals(authKey, userId)

		return c.Next()
	}
}

func AuthContext(c *fiber.Ctx) (userId int64, err error) {
	userId, ok := c.Locals(authKey).(int64)
	if !ok {
		return 0, errors.New("auth context not provided")
	}
	return userId, nil
}
