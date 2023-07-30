package interfaces

import (
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type CartUseCase interface {
	AddToCart(product_id int, userID int) (models.CartResponse, error)
	RemoveFromCart(product_id int, userID int) (models.CartResponse, error)
	DisplayCart(userID int) (models.CartResponse, error)
	EmptyCart(userID int) (models.CartResponse, error)
}
