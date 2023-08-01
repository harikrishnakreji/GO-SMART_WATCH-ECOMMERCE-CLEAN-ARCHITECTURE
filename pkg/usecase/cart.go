package usecase

import (
	"errors"
	"fmt"

	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type cartsUseCase struct {
	cartsRepository   interfaces.CartsRepository
	productRepository interfaces.ProductRepository
}

func NewCartsUseCase(repository interfaces.CartsRepository, productRepo interfaces.ProductRepository) services.CartsUseCase {
	return &cartsUseCase{
		cartsRepository:   repository,
		productRepository: productRepo,
	}
}

func (cr *cartsUseCase) AddToCarts(product_id int, userID int) (models.CartsResponse, error) {
	// check if the product existes
	ok, _, err := cr.cartsRepository.CheckProduct(product_id)
	if err != nil {
		return models.CartsResponse{}, err
	}

	if !ok {
		return models.CartsResponse{}, errors.New("product not exists")
	}

	quantityOfProductInCarts, err := cr.cartsRepository.QuantityOfProductInCarts(userID, product_id)
	fmt.Println(quantityOfProductInCarts)
	if err != nil {
		return models.CartsResponse{}, err
	}

	quantityOfProduct, err := cr.productRepository.GetQuantityFromProductID(product_id)
	fmt.Println(quantityOfProduct)
	if err != nil {
		return models.CartsResponse{}, err
	}

	if quantityOfProduct == 0 {
		return models.CartsResponse{}, errors.New("product out of stock")
	}

	if quantityOfProduct == quantityOfProductInCarts {
		return models.CartsResponse{}, errors.New("stock limit exceeded")
	}

	productPrice, err := cr.productRepository.GetPriceOfProductFromID(product_id)
	if err != nil {
		return models.CartsResponse{}, err
	}

	if quantityOfProductInCarts == 0 {
		err := cr.cartsRepository.AddItemToCarts(userID, product_id, 1, productPrice)
		if err != nil {
			return models.CartsResponse{}, err
		}
	} else {
		currentTotal, err := cr.cartsRepository.TotalPriceForProductInCarts(userID, product_id)
		if err != nil {
			return models.CartsResponse{}, err
		}

		err = cr.cartsRepository.UpdateCarts(quantityOfProductInCarts+1, currentTotal+productPrice, userID, product_id)
		if err != nil {
			return models.CartsResponse{}, err
		}
	}

	cartsDetails, err := cr.cartsRepository.DisplayCarts(userID)
	if err != nil {
		return models.CartsResponse{}, err
	}

	// function to get the grand total price
	cartsTotal, err := cr.cartsRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartsResponse := models.CartsResponse{
		UserName:   cartsTotal.UserName,
		TotalPrice: cartsTotal.TotalPrice,
		Carts:      cartsDetails,
	}

	return cartsResponse, nil
}

// remove items carts (if a product of multiple quantity is present - item will be removed one by one)
func (cr *cartsUseCase) RemoveFromCarts(product_id int, userID int) (models.CartsResponse, error) {

	// check the product to be removed exist
	productExist, err := cr.cartsRepository.ProductExist(product_id, userID)
	if err != nil {
		return models.CartsResponse{}, err
	}
	if !productExist {
		return models.CartsResponse{}, errors.New("the product does not exist in catt")
	}

	var cartsDetails struct {
		Quantity   int
		TotalPrice float64
	}

	cartsDetails, err = cr.cartsRepository.GetQuantityAndTotalPrice(userID, product_id, cartsDetails)
	if err != nil {
		return models.CartsResponse{}, err
	}

	cartsDetails.Quantity = cartsDetails.Quantity - 1
	// after decrementing one quantity if the quantity = 0. delete that item from the carts
	if cartsDetails.Quantity == 0 {

		err := cr.cartsRepository.RemoveProductFromCarts(userID, product_id)
		if err != nil {
			return models.CartsResponse{}, err
		}
	}

	if cartsDetails.Quantity != 0 {

		err := cr.cartsRepository.UpdateCartsDetails(cartsDetails, userID, product_id)
		if err != nil {
			return models.CartsResponse{}, err
		}

	}

	updatedCarts, err := cr.cartsRepository.RemoveFromCarts(userID)
	if err != nil {
		return models.CartsResponse{}, err
	}

	cartsTotal, err := cr.cartsRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartsResponse := models.CartsResponse{
		UserName:   cartsTotal.UserName,
		TotalPrice: cartsTotal.TotalPrice,
		Carts:      updatedCarts,
	}

	return cartsResponse, nil
}

func (cr *cartsUseCase) DisplayCarts(userID int) (models.CartsResponse, error) {

	displayCarts, err := cr.cartsRepository.DisplayCarts(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartsTotal, err := cr.cartsRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartsResponse := models.CartsResponse{
		UserName:   cartsTotal.UserName,
		TotalPrice: cartsTotal.TotalPrice,
		Carts:      displayCarts,
	}

	return cartsResponse, nil
}

func (cr *cartsUseCase) EmptyCarts(userID int) (models.CartsResponse, error) {

	cartsExist, err := cr.cartsRepository.DoesCartsExist(userID)
	if err != nil {
		return models.CartsResponse{}, err
	}

	if !cartsExist {
		return models.CartsResponse{}, errors.New("carts already empty")
	}

	err = cr.cartsRepository.EmptyCarts(userID)
	if err != nil {
		return models.CartsResponse{}, err
	}

	cartsTotal, err := cr.cartsRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartsResponse := models.CartsResponse{
		UserName:   cartsTotal.UserName,
		TotalPrice: cartsTotal.TotalPrice,
		Carts:      []models.Carts{},
	}

	return cartsResponse, nil
}
