package service

import (
	"app/src/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OtpService interface {
	CreateOtp(c *fiber.Ctx) (*model.OtpToken, error)
	//GetAll(c *fiber.Ctx, params *validation.QueryAuth) ([]model.AuthToken, error)
	//GetByAuthId(c *fiber.Ctx, id string) (*model.AuthToken, error)
	//Update(c *fiber.Ctx, req *validation.UpdateAuth2, id string) (*model.AuthToken, error)

	//Update(c *fiber.Ctx)
	//Delete(c *fiber.Ctx)
}

// Define methods for user service

type otpService struct {
	DB *gorm.DB
}

// DB servie init
func NewOtpService(db *gorm.DB) OtpService {
	return &otpService{DB: db}
}

// Create
func (s *otpService) CreateOtp(c *fiber.Ctx) (*model.OtpToken, error) {
	var otp model.OtpToken
	if err := c.BodyParser(&otp); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&otp).Error; err != nil {
		return nil, err
	}

	return &otp, nil
}
