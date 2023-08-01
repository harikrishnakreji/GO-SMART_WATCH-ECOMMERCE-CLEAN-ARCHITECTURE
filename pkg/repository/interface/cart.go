package interfaces

import "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"

type CartsRepository interface {
	AddItemToCarts(userID int, product_id int, quantity int, productPrice float64) error
	TotalPriceForProductInCarts(userID int, ProductID int) (float64, error)
	UpdateCarts(quantity int, price float64, userID int, product_id int) error
	QuantityOfProductInCarts(userID int, product_id int) (int, error)
	RemoveFromCarts(userID int) ([]models.Carts, error)
	DisplayCarts(userID int) ([]models.Carts, error)
	EmptyCarts(userID int) error
	GetTotalPriceFromCarts(userID int) (float64, error)
	GetQuantityAndTotalPrice(userID int, product_id int, CartsDetails struct {
		Quantity   int
		TotalPrice float64
	}) (struct {
		Quantity   int
		TotalPrice float64
	}, error)
	UpdateCartsDetails(CartsDetails struct {
		Quantity   int
		TotalPrice float64
	}, userID int, productID int) error
	RemoveProductFromCarts(userID int, product_id int) error
	GetTotalPrice(userID int) (models.CartsTotal, error)
	GetAllItemFromCarts(userID int) ([]models.Carts, error)
	CheckProduct(product_id int) (bool, string, error)
	ProductExist(product_id int, userID int) (bool, error)
	DoesCartsExist(userID int) (bool, error)
	GetAllItemsFromCarts(userID int) ([]models.Carts, error)
}
