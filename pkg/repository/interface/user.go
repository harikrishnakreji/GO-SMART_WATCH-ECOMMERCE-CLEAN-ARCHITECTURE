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
}
