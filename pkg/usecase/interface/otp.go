package interfaces

import (
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type OtpUseCase interface {
	VerifyOTP(code models.VerifyData) (models.TokenUsers, error)
	VerifyOTPtoReset(code models.VerifyData) (string, error)
	SendOTP(phone string) error
	SendOTPtoReset(email string) (string, error)
}
