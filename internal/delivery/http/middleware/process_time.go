package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type localKey struct {
	key string
}

var (
	processTimeKey = &localKey{"process_time"}
)

func ProcessTime(c *fiber.Ctx) (time.Duration, bool) {
	t, ok := c.Locals(processTimeKey).(time.Time)
	if !ok {
		return 0, false
	}
	return time.Since(t), true
}

func NewProcessTimeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(processTimeKey, time.Now())
		return c.Next()
	}
}
