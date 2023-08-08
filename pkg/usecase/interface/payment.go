package interfaces

import "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"

type PaymentUseCase interface {
	MakePaymentRazorPay(orderID string, userID int) (models.CombinedOrderDetails, string, error)
	SavePaymentDetails(paymentID string, razorID string, orderID string) error
}
