package service

import (
	"app/src/model"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthService interface {
	CreateAuth(c *fiber.Ctx) (*model.AuthToken, error)
	GetAll(c *fiber.Ctx, params *validation.QueryAuth) ([]model.AuthToken, error)
	GetByAuthId(c *fiber.Ctx, id string) (*model.AuthToken, error)
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
func (s *authService) CreateAuth(c *fiber.Ctx) (*model.AuthToken, error) {
	var auth model.AuthToken
	if err := c.BodyParser(&auth); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

// Get All
func (s *authService) GetAll(c *fiber.Ctx, params *validation.QueryAuth) ([]model.AuthToken, error) {

	var auth []model.AuthToken

	offset := (params.Page - 1) * params.Limit

	query := s.DB.WithContext(c.Context()).Order("created_at asc")

	if search := params.Search; search != "" {
		query = query.Where("name LIKE? or phoneNumber LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	result := query.Find(&auth).Offset(offset)

	if err := query.Find(&auth).Error; err != nil {
		return nil, err
	}
	result = query.Limit(params.Limit).Offset(offset).Find(&auth)
	if result.Error != nil {

		return nil, result.Error
	}

	return auth, result.Error

}

// GetbyId

func (s *authService) GetByAuthId(c *fiber.Ctx, id string) (*model.AuthToken, error) {
	auth := new(model.AuthToken)

	result := s.DB.WithContext(c.Context()).First(&auth, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return auth, nil
}
