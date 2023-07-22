package interfaces

import (
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type ProductRepository interface {
	ShowAllProducts(page int, count int) ([]models.ProductsBrief, error)
	AddProduct(product models.ProductsReceiver) (models.ProductResponse, error)
	DeleteProduct(product_id string) error
	// DoesProductExist(productID int) (bool, error)
	GetGenres() ([]domain.Genre, error)
	// GetQuantityFromProductID(id int) (int, error)
	// GetPriceOfProductFromID(productID int) (float64, error)
}
