package controller

import (
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// Get All
func (u *OtpController) GetAll(c *fiber.Ctx) error {
	query := &validation.QueryOtp{

		Page:   c.QueryInt("Page", 1),
		Limit:  c.QueryInt("Limit", 20),
		Search: c.Query("Search", ""),
	}

	otp, err := u._OtpService.GetAll(c, query)

	if err != nil {

		return err
	}
	return c.Status(fiber.StatusOK).JSON(otp)
}

// Get By OtpId

func (u *OtpController) GetByOtpId(c *fiber.Ctx) error {

	OtpId := c.Params("otpId")

	if _, err := uuid.Parse(OtpId); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	otp, err := u._OtpService.GetByOtpId(c, OtpId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(otp)

}
