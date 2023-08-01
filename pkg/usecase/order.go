package usecase

import (
	"errors"
	"fmt"

	domain "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	helper "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/helper"
	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"github.com/jinzhu/copier"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	cartsRepository interfaces.CartsRepository
	userRepository  interfaces.UserRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository, cartsRepo interfaces.CartsRepository, userRepo interfaces.UserRepository) services.OrderUseCase {
	return &orderUseCase{
		orderRepository: orderRepo,
		cartsRepository: cartsRepo,
		userRepository:  userRepo,
	}
}

func (o *orderUseCase) OrderItemsFromCarts(orderFromCarts models.OrderFromCarts, userID int) (domain.OrderSuccessResponse, error) {

	var orderBody models.OrderIncoming
	err := copier.Copy(&orderBody, &orderFromCarts)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderBody.UserID = uint(userID)
	cartsExist, err := o.orderRepository.DoesCartsExist(userID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !cartsExist {
		return domain.OrderSuccessResponse{}, errors.New("carts empty can't order")
	}

	addressExist, err := o.orderRepository.AddressExist(orderBody)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !addressExist {
		return domain.OrderSuccessResponse{}, errors.New("address does not exist")
	}

	// get all items a slice of carts
	cartsItems, err := o.cartsRepository.GetAllItemsFromCarts(int(orderBody.UserID))
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	var orderDetails domain.Order
	var orderItemDetails domain.OrderItem

	// add general order details - that is to be added to orders table
	orderDetails = helper.CopyOrderDetails(orderDetails, orderBody)

	// get grand total iterating through each products in carts
	for _, c := range cartsItems {
		orderDetails.GrandTotal += c.TotalPrice
	}

	orderDetails.FinalPrice = orderDetails.GrandTotal
	if orderBody.PaymentID == 2 {
		orderDetails.PaymentStatus = "not paid"
		orderDetails.ShipmentStatus = "pending"
	}

	err = o.orderRepository.CreateOrder(orderDetails)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	for _, c := range cartsItems {
		// for each order save details of products and associated details and use order_id as foreign key ( for each order multiple product will be there)
		orderItemDetails.OrderID = orderDetails.OrderId
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = int(c.Quantity)
		orderItemDetails.TotalPrice = c.TotalPrice

		err := o.orderRepository.AddOrderItems(orderItemDetails, orderDetails.UserID, c.ProductID, c.Quantity)
		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}

	}
	orderSuccessResponse, err := o.orderRepository.GetBriefOrderDetails(orderDetails.OrderId)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	return orderSuccessResponse, nil
}

// get order details
func (o *orderUseCase) GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := o.orderRepository.GetOrderDetails(userID, page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}

	return fullOrderDetails, nil

}

func (o *orderUseCase) CancelOrder(orderID string, userID int) error {

	// check whether the orderID corresponds to the given user (other user with token may try to send orderID as path variables) (have to add this logic to so many places)
	userTest, err := o.orderRepository.UserOrderRelationship(orderID, userID)
	if err != nil {
		return err
	}

	if userTest != userID {
		return errors.New("the order is not done by this user")
	}

	orderProducts, err := o.orderRepository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return err
	}

	shipmentStatus, err := o.orderRepository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}

	if shipmentStatus == "delivered" {
		return errors.New("item already delivered, cannot cancel")
	}

	if shipmentStatus == "pending" || shipmentStatus == "returned" || shipmentStatus == "return" {
		message := fmt.Sprint(shipmentStatus)
		return errors.New("the order is in" + message + ", so no point in cancelling")
	}

	if shipmentStatus == "cancelled" {
		return errors.New("the order is already cancelled, so no point in cancelling")
	}

	err = o.orderRepository.CancelOrder(orderID)
	if err != nil {
		return err
	}

	// update the quantity to products since the order is cancelled
	err = o.orderRepository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return err
	}

	return nil

}

func (o *orderUseCase) CancelOrderFromAdminSide(orderID string) error {

	orderProducts, err := o.orderRepository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return err
	}

	err = o.orderRepository.CancelOrder(orderID)
	if err != nil {
		return err
	}

	// update the quantity to products since the order is cancelled
	err = o.orderRepository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return err
	}

	return nil

}

func (o *orderUseCase) GetAllOrderDetailsForAdmin(page int) ([]models.CombinedOrderDetails, error) {

	orderDetails, err := o.orderRepository.GetOrderDetailsBrief(page)
	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}
	return orderDetails, nil

}

func (o *orderUseCase) ApproveOrder(orderID string) error {

	// check whether the specified orderID exist
	ok, err := o.orderRepository.CheckOrderID(orderID)
	fmt.Println(ok)
	if !ok {
		return err
	}

	// check the shipment status - if the status cancelled, don't approve it
	shipmentStatus, err := o.orderRepository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}

	if shipmentStatus == "cancelled" {

		return errors.New("the order is cancelled, cannot approve it")
	}

	if shipmentStatus == "pending" {

		return errors.New("the order is pending, cannot approve it")
	}

	if shipmentStatus == "processing" {
		fmt.Println("reached here")
		err := o.orderRepository.ApproveOrder(orderID)

		if err != nil {
			return err
		}

		return nil
	}

	// if the shipment status is not processing or cancelled. Then it is defenetely cancelled
	return nil

}

func (o *orderUseCase) OrderDelivered(orderID string) error {

	// check the shipment status - if the status cancelled, don't approve it
	shipmentStatus, err := o.orderRepository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}

	if shipmentStatus == "order placed" {
		shipmentStatus = "delivered"
		return o.orderRepository.UpdateShipmentStatus(shipmentStatus, orderID)
	}

	return errors.New("order not placed or order id does not exist")

}
