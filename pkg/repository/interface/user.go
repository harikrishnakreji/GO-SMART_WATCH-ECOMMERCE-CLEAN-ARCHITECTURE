package interfaces

import (
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error)
	AddAddress(address models.AddressInfo, userID int) error
	UpdateAddress(address models.AddressInfo, addressID int, userID int) (models.AddressInfoResponse, error)
	GetAllAddresses(userID int) ([]models.AddressInfoResponse, error)
	GetAllPaymentOption() ([]models.PaymentDetails, error)
	UserDetails(userID int) (models.UsersProfileDetails, error)
	UpdateUserEmail(email string, userID int) error
	UpdateUserName(name string, userID int) error
	UpdateUserPhone(phone string, userID int) error
	UpdateUserPassword(password string, userID int) error
	UserPassword(userID int) (string, error)
	ResetPassword(userID int, password string) error
}
