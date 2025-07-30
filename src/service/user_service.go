package service

import (
	"app/src/model"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(c *fiber.Ctx) (*model.User, error)
	GetAll(c *fiber.Ctx, params *validation.QueryUser) ([]model.User, error)
	GetByUserId(c *fiber.Ctx, id string) (*model.User, error)
	//GetByPhoneNumber(c *fiber.Ctx)
	//Update(c *fiber.Ctx)
	//Delete(c *fiber.Ctx)
}

// Define methods for user service

type userService struct {
	DB *gorm.DB
}

// DB servie init
func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}

// Create
func (s *userService) CreateUser(c *fiber.Ctx) (*model.User, error) {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Get All
func (s *userService) GetAll(c *fiber.Ctx, params *validation.QueryUser) ([]model.User, error) {

	var users []model.User

	offset := (params.Page - 1) * params.Limit

	query := s.DB.WithContext(c.Context()).Order("created_at asc")

	if search := params.Search; search != "" {
		query = query.Where("name LIKE? or phoneNumber LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	result := query.Find(&users).Offset(offset)

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	result = query.Limit(params.Limit).Offset(offset).Find(&users)
	if result.Error != nil {

		return nil, result.Error
	}

	return users, result.Error

}

// GetbyUserId

func (s *userService) GetByUserId(c *fiber.Ctx, id string) (*model.User, error) {
	user := new(model.User)

	result := s.DB.WithContext(c.Context()).First(&user, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}
