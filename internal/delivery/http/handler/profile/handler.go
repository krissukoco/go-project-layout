package profile_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krissukoco/go-project-layout/internal/delivery/http/middleware"
	http_response "github.com/krissukoco/go-project-layout/internal/delivery/http/response"
	user_usecase "github.com/krissukoco/go-project-layout/internal/usecase/user"
)

type handler struct {
	uc user_usecase.Usecase
}

func New(
	uc user_usecase.Usecase,
) *handler {
	return &handler{uc}
}

func (h *handler) GetProfile(c *fiber.Ctx) error {
	userId, err := middleware.AuthContext(c)
	if err != nil {
		return err
	}

	user, err := h.uc.GetProfile(c.Context(), userId)
	if err != nil {
		// TODO: return something else if user not found?
		return err
	}
	return http_response.JSON(c, &Profile{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
}
