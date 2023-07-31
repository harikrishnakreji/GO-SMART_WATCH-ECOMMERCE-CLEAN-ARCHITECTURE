package usecase

import (
	"errors"
	"fmt"

	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type cartUseCase struct {
	cartRepository    interfaces.CartsRepository
	productRepository interfaces.ProductRepository
}

func NewCartsUseCase(repository interfaces.CartsRepository, productRepo interfaces.ProductRepository) services.CartsUseCase {
	return &cartUseCase{
		cartRepository:    repository,
		productRepository: productRepo,
	}
}

func (cr *cartUseCase) AddToCarts(product_id int, userID int) (models.CartsResponse, error) {
	// check if the product existes
	ok, _, err := cr.cartRepository.CheckProduct(product_id)
	if err != nil {
		return models.CartsResponse{}, err
	}

	if !ok {
		return models.CartsResponse{}, errors.New("product not exists")
	}

	quantityOfProductInCarts, err := cr.cartRepository.QuantityOfProductInCarts(userID, product_id)
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
		err := cr.cartRepository.AddItemToCarts(userID, product_id, 1, productPrice)
		if err != nil {
			return models.CartsResponse{}, err
		}
	} else {
		currentTotal, err := cr.cartRepository.TotalPriceForProductInCarts(userID, product_id)
		if err != nil {
			return models.CartsResponse{}, err
		}

		err = cr.cartRepository.UpdateCarts(quantityOfProductInCarts+1, currentTotal+productPrice, userID, product_id)
		if err != nil {
			return models.CartsResponse{}, err
		}
	}

	cartDetails, err := cr.cartRepository.DisplayCarts(userID)
	if err != nil {
		return models.CartsResponse{}, err
	}

	// function to get the grand total price
	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartResponse := models.CartsResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Carts:      cartDetails,
	}

	return cartResponse, nil
}

// remove items cart (if a product of multiple quantity is present - item will be removed one by one)
func (cr *cartUseCase) RemoveFromCarts(product_id int, userID int) (models.CartsResponse, error) {

	// check the product to be removed exist
	productExist, err := cr.cartRepository.ProductExist(product_id, userID)
	if err != nil {
		return models.CartsResponse{}, err
	}
	if !productExist {
		return models.CartsResponse{}, errors.New("the product does not exist in catt")
	}

	var cartDetails struct {
		Quantity   int
		TotalPrice float64
	}

	cartDetails, err = cr.cartRepository.GetQuantityAndTotalPrice(userID, product_id, cartDetails)
	if err != nil {
		return models.CartsResponse{}, err
	}

	cartDetails.Quantity = cartDetails.Quantity - 1
	// after decrementing one quantity if the quantity = 0. delete that item from the cart
	if cartDetails.Quantity == 0 {

		err := cr.cartRepository.RemoveProductFromCarts(userID, product_id)
		if err != nil {
			return models.CartsResponse{}, err
		}
	}

	if cartDetails.Quantity != 0 {

		err := cr.cartRepository.UpdateCartsDetails(cartDetails, userID, product_id)
		if err != nil {
			return models.CartsResponse{}, err
		}

	}

	updatedCarts, err := cr.cartRepository.RemoveFromCarts(userID)
	if err != nil {
		return models.CartsResponse{}, err
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartResponse := models.CartsResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Carts:      updatedCarts,
	}

	return cartResponse, nil
}

func (cr *cartUseCase) DisplayCarts(userID int) (models.CartsResponse, error) {

	displayCarts, err := cr.cartRepository.DisplayCarts(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartResponse := models.CartsResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Carts:      displayCarts,
	}

	return cartResponse, nil
}

func (cr *cartUseCase) EmptyCarts(userID int) (models.CartsResponse, error) {

	cartExist, err := cr.cartRepository.DoesCartsExist(userID)
	if err != nil {
		return models.CartsResponse{}, err
	}

	if !cartExist {
		return models.CartsResponse{}, errors.New("cart already empty")
	}

	err = cr.cartRepository.EmptyCarts(userID)
	if err != nil {
		return models.CartsResponse{}, err
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartsResponse{}, err
	}

	cartResponse := models.CartsResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Carts:      []models.Carts{},
	}

	return cartResponse, nil
}
