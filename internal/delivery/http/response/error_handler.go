package http_response

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewErrorHandler(debug bool) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		// Response Error
		var respErr *Error
		if errors.As(err, &respErr) {
			return c.Status(respErr.HttpStatus).JSON(respErr)
		}

		// Validation Error
		var valErr validator.ValidationErrors
		if errors.As(err, &valErr) {
			msg := "Validation error :"
			details := make([]ErrorDetail, len(valErr))
			for i, v := range valErr {
				if i == 0 {
					msg += v.Error()
				}
				details[i] = ErrorDetail{Field: v.Field(), Message: v.Error()}
			}
			return c.Status(http.StatusBadRequest).JSON(NewError(http.StatusBadRequest, Code_InvalidRequest, msg, details...))
		}

		// Fiber Error
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return c.Status(fiberErr.Code).JSON(fiberErr)
		}

		log.Printf("UNKNOWN INTERNAL SERVER ERROR : %v", err)
		intErr := NewError(
			http.StatusInternalServerError,
			Code_Internal,
			"Internal server error",
		)
		if debug {
			intErr.Details = append(intErr.Details, ErrorDetail{
				Field:   "INTERNAL_ERROR",
				Message: err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(intErr)
	}
}
