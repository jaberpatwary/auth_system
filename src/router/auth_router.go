package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(v1 fiber.Router, u service.AuthService) {
	// Initialize the UserController with the UserService
	authController := controller.NewAuthController(u)
	// Define user-related routes
	authGroup := v1.Group("/auths")
	authGroup.Post("/", authController.CreateAuth)
	//userGroup.Get("/", userController.GetAll)
	//userGroup.Get("/:userId", userController.GetByUserId)
}
