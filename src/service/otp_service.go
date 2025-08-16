package service

import (
	"app/src/model"
	"app/src/validation"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OtpService interface {
	CreateOtp(c *fiber.Ctx) (*model.OtpToken, error)
	GetAll(c *fiber.Ctx, params *validation.QueryOtp) ([]model.OtpToken, error)
	GetByOtpId(c *fiber.Ctx, id string) (*model.OtpToken, error)
	Update(c *fiber.Ctx, req *validation.UpdateOtp, id string) (*model.OtpToken, error)

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

// Get All
func (s *otpService) GetAll(c *fiber.Ctx, params *validation.QueryOtp) ([]model.OtpToken, error) {

	var otp []model.OtpToken

	offset := (params.Page - 1) * params.Limit

	query := s.DB.WithContext(c.Context()).Order("created_at asc")

	if search := params.Search; search != "" {
		query = query.Where("name LIKE? or phoneNumber LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	result := query.Find(&otp).Offset(offset)

	if err := query.Find(&otp).Error; err != nil {
		return nil, err
	}
	result = query.Limit(params.Limit).Offset(offset).Find(&otp)
	if result.Error != nil {

		return nil, result.Error
	}

	return otp, result.Error

}

// GetbyId

func (s *otpService) GetByOtpId(c *fiber.Ctx, id string) (*model.OtpToken, error) {
	otp := new(model.OtpToken)

	result := s.DB.WithContext(c.Context()).First(&otp, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return otp, nil
}

// Update Otp details
func (s *otpService) Update(c *fiber.Ctx, req *validation.UpdateOtp, id string) (*model.OtpToken, error) {
	if req.OtpCode == "" && req.Purpose == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}
	if req.Purpose == "" {
		hashedPurpose, err := hashedPurpose(req.Purpose)
		if err != nil {
			return nil, err
		}
		req.Purpose = hashedPurpose
	}

	updateBody := &model.OtpToken{

		OtpCode: req.OtpCode,
		Purpose: req.Purpose,
	}
	result := s.DB.WithContext(c.Context()).Where("id = ?", id).Updates(updateBody)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return nil, fiber.NewError(fiber.StatusConflict, "Phone number already exists")
	}

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusConflict, "User Not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	otp, err := s.GetByOtpId(c, id)
	if err != nil {
		return nil, err
	}
	return otp, nil
}
