package usecase

import (
	"errors"

	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"github.com/razorpay/razorpay-go"
)

type paymentUseCase struct {
	paymentRepository interfaces.PaymentRepository
	orderRepository   interfaces.OrderRepository
	userRepository    interfaces.UserRepository
}

func NewPaymentUseCase(paymentRepo interfaces.PaymentRepository, orderRepo interfaces.OrderRepository, userRepo interfaces.UserRepository) services.PaymentUseCase {
	return &paymentUseCase{
		paymentRepository: paymentRepo,
		orderRepository:   orderRepo,
		userRepository:    userRepo,
	}
}

func (p *paymentUseCase) MakePaymentRazorPay(orderID string, userID int) (models.CombinedOrderDetails, string, error) {

	combinedOrderDetails, err := p.orderRepository.GetOrderDetailsByOrderId(orderID)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}

	client := razorpay.NewClient("rzp_test_AcVxAhyZEXSJ6H", "IXYtGlRdNFkdB5nozcKNbq09")

	data := map[string]interface{}{
		"amount":   int(combinedOrderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}

	razorPayOrderID := body["id"].(string)

	err = p.orderRepository.AddRazorPayDetails(orderID, razorPayOrderID)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}

	return combinedOrderDetails, razorPayOrderID, nil

}

func (p *paymentUseCase) SavePaymentDetails(paymentID string, razorID string, orderID string) error {

	// to check whether the order is already paid
	status, err := p.orderRepository.CheckPaymentStatus(razorID, orderID)
	if err != nil {
		return err
	}

	if status == "not paid" {

		err = p.orderRepository.UpdatePaymentDetails(orderID, paymentID)
		if err != nil {
			return err
		}

		err := p.orderRepository.UpdateShipmentAndPaymentByOrderID("processing", "paid", orderID)
		if err != nil {
			return err
		}

		return nil

	}

	return errors.New("already paid")

}
