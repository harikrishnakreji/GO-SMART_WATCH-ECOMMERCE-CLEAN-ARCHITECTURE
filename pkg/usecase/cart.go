package usecase

import (
	"errors"
	"fmt"

	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type cartUseCase struct {
	cartRepository    interfaces.CartRepository
	productRepository interfaces.ProductRepository
}

func NewCartUseCase(repository interfaces.CartRepository, productRepo interfaces.ProductRepository) services.CartUseCase {
	return &cartUseCase{
		cartRepository:    repository,
		productRepository: productRepo,
	}
}

func (cr *cartUseCase) AddToCart(product_id int, userID int) (models.CartResponse, error) {
	// check if the product existes
	ok, _, err := cr.cartRepository.CheckProduct(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}

	if !ok {
		return models.CartResponse{}, errors.New("product not exists")
	}

	quantityOfProductInCart, err := cr.cartRepository.QuantityOfProductInCart(userID, product_id)
	fmt.Println(quantityOfProductInCart)
	if err != nil {
		return models.CartResponse{}, err
	}

	quantityOfProduct, err := cr.productRepository.GetQuantityFromProductID(product_id)
	fmt.Println(quantityOfProduct)
	if err != nil {
		return models.CartResponse{}, err
	}

	if quantityOfProduct == 0 {
		return models.CartResponse{}, errors.New("product out of stock")
	}

	if quantityOfProduct == quantityOfProductInCart {
		return models.CartResponse{}, errors.New("stock limit exceeded")
	}

	productPrice, err := cr.productRepository.GetPriceOfProductFromID(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}

	if quantityOfProductInCart == 0 {
		err := cr.cartRepository.AddItemToCart(userID, product_id, 1, productPrice)
		if err != nil {
			return models.CartResponse{}, err
		}
	} else {
		currentTotal, err := cr.cartRepository.TotalPriceForProductInCart(userID, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}

		err = cr.cartRepository.UpdateCart(quantityOfProductInCart+1, currentTotal+productPrice, userID, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}

	cartDetails, err := cr.cartRepository.DisplayCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	// function to get the grand total price
	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cartDetails,
	}

	return cartResponse, nil
}

// remove items cart (if a product of multiple quantity is present - item will be removed one by one)
func (cr *cartUseCase) RemoveFromCart(product_id int, userID int) (models.CartResponse, error) {

	// check the product to be removed exist
	productExist, err := cr.cartRepository.ProductExist(product_id, userID)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !productExist {
		return models.CartResponse{}, errors.New("the product does not exist in catt")
	}

	var cartDetails struct {
		Quantity   int
		TotalPrice float64
	}

	cartDetails, err = cr.cartRepository.GetQuantityAndTotalPrice(userID, product_id, cartDetails)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartDetails.Quantity = cartDetails.Quantity - 1
	// after decrementing one quantity if the quantity = 0. delete that item from the cart
	if cartDetails.Quantity == 0 {

		err := cr.cartRepository.RemoveProductFromCart(userID, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}

	if cartDetails.Quantity != 0 {

		err := cr.cartRepository.UpdateCartDetails(cartDetails, userID, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}

	}

	updatedCart, err := cr.cartRepository.RemoveFromCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       updatedCart,
	}

	return cartResponse, nil
}

func (cr *cartUseCase) DisplayCart(userID int) (models.CartResponse, error) {

	displayCart, err := cr.cartRepository.DisplayCart(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       displayCart,
	}

	return cartResponse, nil
}

func (cr *cartUseCase) EmptyCart(userID int) (models.CartResponse, error) {

	cartExist, err := cr.cartRepository.DoesCartExist(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	if !cartExist {
		return models.CartResponse{}, errors.New("cart already empty")
	}

	err = cr.cartRepository.EmptyCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       []models.Carts{},
	}

	return cartResponse, nil
}
