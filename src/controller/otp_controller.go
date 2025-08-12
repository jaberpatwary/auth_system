package controller

import (
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

type OtpController struct {
	_OtpService service.OtpService
}

func NewOtpController(otpService service.OtpService) *OtpController {
	return &OtpController{

		_OtpService: otpService,
	}
}

// Create otp
func (u *OtpController) CreateOtp(c *fiber.Ctx) error {

	otp, err := u._OtpService.CreateOtp(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(otp)
}
