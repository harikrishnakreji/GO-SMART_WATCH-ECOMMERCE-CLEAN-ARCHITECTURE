package interfaces

import "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"

type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error)
	FindUserByEmail(email string) (bool, error)
	GetUserPhoneByEmail(email string) (string, error)
}
