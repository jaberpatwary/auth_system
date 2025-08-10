package controller

import (
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	_AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{

		_AuthService: authService,
	}
}

// Create User
func (u *AuthController) CreateAuth(c *fiber.Ctx) error {

	auth, err := u._AuthService.CreateAuth(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(auth)
}
