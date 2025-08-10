package service

import (
	"app/src/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthService interface {
	CreateAuth(c *fiber.Ctx) (*model.Auth, error)
	//GetAll(c *fiber.Ctx, params *validation.QueryUser) ([]model.User, error)
	//GetByUserId(c *fiber.Ctx, id string) (*model.User, error)
	//GetByPhoneNumber(c *fiber.Ctx)
	//Update(c *fiber.Ctx)
	//Delete(c *fiber.Ctx)
}

// Define methods for user service

type authService struct {
	DB *gorm.DB
}

// DB servie init
func NewAuthService(db *gorm.DB) AuthService {
	return &authService{DB: db}
}

// Create
func (s *authService) CreateAuth(c *fiber.Ctx) (*model.Auth, error) {
	var auth model.Auth
	if err := c.BodyParser(&auth); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
