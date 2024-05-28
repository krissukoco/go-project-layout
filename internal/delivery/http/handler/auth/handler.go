package auth_handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	http_request "github.com/krissukoco/go-project-layout/internal/delivery/http/request"
	http_response "github.com/krissukoco/go-project-layout/internal/delivery/http/response"
	auth_user_usecase "github.com/krissukoco/go-project-layout/internal/usecase/auth_user"
)

type handler struct {
	authUc auth_user_usecase.Usecase
}

func New(
	authUc auth_user_usecase.Usecase,
) *handler {
	return &handler{authUc}
}

func (h *handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return http_response.NewError(http.StatusUnprocessableEntity, http_response.Code_InvalidRequest, err.Error())
	}
	if err := http_request.Validate(req); err != nil {
		return err
	}

	tokens, err := h.authUc.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, auth_user_usecase.ErrInvalidCredentials) {
			return http_response.NewError(http.StatusBadRequest, http_response.Code_InvalidRequest, err.Error())
		}
		return err
	}

	return http_response.JSON(c, &LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}
