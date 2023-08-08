package interfaces

import (
	domain "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type OrderUseCase interface {
	OrderItemsFromCarts(orderBody models.OrderFromCarts, userId int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error)
	CancelOrder(orderID string, userID int) error
	CancelOrderFromAdminSide(orderID string) error
	GetAllOrderDetailsForAdmin(page int) ([]models.CombinedOrderDetails, error)
	ApproveOrder(orderId string) error
	OrderDelivered(orderID string) error
	ReturnOrder(orderID string) error
	RefundOrder(orderID string) error
}
