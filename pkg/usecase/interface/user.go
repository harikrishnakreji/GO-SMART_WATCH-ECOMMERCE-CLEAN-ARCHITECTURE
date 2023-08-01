package interfaces

import (
	"context"

	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)

	AddAddress(address models.AddressInfo, userID int) error
	UpdateAddress(address models.AddressInfo, addressID int, userID int) (models.AddressInfoResponse, error)
	Checkout(userID int) (models.CheckoutDetails, error)
	UserDetails(userID int) (models.UsersProfileDetails, error)
	GetAllAddress(userID int) ([]models.AddressInfoResponse, error)
	UpdateUserDetails(body models.UsersProfileDetails, ctx context.Context) (models.UsersProfileDetails, error)
	UpdatePassword(ctx context.Context, body models.UpdatePassword) error
	ResetPassword(userID int, pass models.ResetPassword) error
}
