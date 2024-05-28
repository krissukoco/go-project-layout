package http_response

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/krissukoco/go-project-layout/internal/delivery/http/middleware"
)

type Response struct {
	ProcessTime string      `json:"process_time"`
	Data        interface{} `json:"data"`
}

func JSON(c *fiber.Ctx, data interface{}) error {
	resp := &Response{
		Data: data,
	}
	ptime, ok := middleware.ProcessTime(c)
	if ok {
		resp.ProcessTime = fmt.Sprintf("%v", ptime)
	}
	return c.Status(200).JSON(resp)
}

type Error struct {
	HttpStatus int           `json:"http_status"`
	Code       int           `json:"code"`
	Message    string        `json:"message"`
	Details    []ErrorDetail `json:"details"`
}

// Error struct should be of type error so handlers or middlewares can return Error to be catched later by FiberErrorHandler
func (e Error) Error() string {
	return e.Message
}

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func NewError(httpStatus int, code int, message string, details ...ErrorDetail) *Error {
	e := &Error{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    capitalize(message),
		Details:    details,
	}
	if e.Details == nil {
		// Make empty array for clients
		e.Details = make([]ErrorDetail, 0)
	}
	return e
}
