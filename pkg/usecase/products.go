package usecase

import (
	"fmt"

	domain "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUseCase(repo interfaces.ProductRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
	}
}

func (pr *productUseCase) ShowAllProducts(page int, count int) ([]models.ProductsBrief, error) {
	productsBrief, err := pr.productRepo.ShowAllProducts(page, count)

	if err != nil {
		panic("no products")
	}
	// here memory address of each item in productBrief is taken so that a copy of each instance is not made while updating
	for i := range productsBrief {
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}
	return productsBrief, nil
}

func (pr *productUseCase) ShowAllProductsToAdmin(page int, count int) ([]models.ProductsBrief, error) {

	productsBrief, err := pr.productRepo.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductsBrief{}, err
	}
	fmt.Println(productsBrief)
	// here memory address of each item in productBrief is taken so that a copy of each instance is not made while updating
	for i := range productsBrief {
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}

	return productsBrief, nil
}

func (pr *productUseCase) AddProduct(product models.ProductsReceiver) (models.ProductResponse, error) {
	// this logic is to add the quantity of product if admin try to add duplicate product (have to work on this in the future)
	// alreadyPresent,err := cr.productRepo.CheckIfAlreadyPresent(c,product)

	// if err != nil {
	// 	return err
	// }

	// if alreadyPresent {
	// 	fmt.Println("it came here")
	// 	err := cr.productRepo.UpdateQuantity(c,product)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// }

	productResponse, err := pr.productRepo.AddProduct(product)

	if err != nil {
		return models.ProductResponse{}, err
	}

	return productResponse, nil

}

func (pr *productUseCase) DeleteProduct(product_id string) error {

	err := pr.productRepo.DeleteProduct(product_id)
	if err != nil {
		return err
	}
	return nil

}

func (pr *productUseCase) GetGenres() ([]domain.Genre, error) {

	return pr.productRepo.GetGenres()
}
