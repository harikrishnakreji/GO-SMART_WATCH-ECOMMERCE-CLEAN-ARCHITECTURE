package interfaces

import (
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type CartsUseCase interface {
	AddToCarts(product_id int, userID int) (models.CartsResponse, error)
	RemoveFromCarts(product_id int, userID int) (models.CartsResponse, error)
	DisplayCarts(userID int) (models.CartsResponse, error)
	EmptyCarts(userID int) (models.CartsResponse, error)
}
