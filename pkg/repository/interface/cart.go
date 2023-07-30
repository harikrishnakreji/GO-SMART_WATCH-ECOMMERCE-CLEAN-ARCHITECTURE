package interfaces

import "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"

type CartRepository interface {
	AddItemToCart(userID int, product_id int, quantity int, productPrice float64) error
	TotalPriceForProductInCart(userID int, ProductID int) (float64, error)
	UpdateCart(quantity int, price float64, userID int, product_id int) error
	QuantityOfProductInCart(userID int, product_id int) (int, error)
	RemoveFromCart(userID int) ([]models.Carts, error)
	DisplayCart(userID int) ([]models.Carts, error)
	EmptyCart(userID int) error
	GetTotalPriceFromCart(userID int) (float64, error)
	GetQuantityAndTotalPrice(userID int, product_id int, CartDetails struct {
		Quantity   int
		TotalPrice float64
	}) (struct {
		Quantity   int
		TotalPrice float64
	}, error)
	UpdateCartDetails(CartDetails struct {
		Quantity   int
		TotalPrice float64
	}, userID int, productID int) error
	RemoveProductFromCart(userID int, product_id int) error
	GetTotalPrice(userID int) (models.CartTotal, error)
	GetAllItemFromCart(userID int) ([]models.Carts, error)
	CheckProduct(product_id int) (bool, string, error)
	ProductExist(product_id int, userID int) (bool, error)
	DoesCartExist(userID int) (bool, error)
}
