package interfaces

import (
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type ProductUseCase interface {
	ShowAllProducts(page int, count int) ([]models.ProductsBrief, error)
	ShowAllProductsToAdmin(page int, count int) ([]models.ProductsBrief, error)
	AddProduct(product models.ProductsReceiver) (models.ProductResponse, error)
	DeleteProduct(product_id string) error
	GetGenres() ([]domain.Genre, error)
}
