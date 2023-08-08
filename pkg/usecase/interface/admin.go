package interfaces

import (
	domain "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	CreateAdmin(admin models.AdminSignUp) (domain.TokenAdmin, error)
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	GetCategorys() ([]domain.Category, error)
	AddCategorys(category models.CategoryUpdate) error
	Delete(category_id string) error
	BlockUser(id string) error
	UnBlockUser(id string) error
	FilteredSalesReport(timePeriod string) (models.SalesReport, error)
	DashBoard() (models.CompleteAdminDashboard, error)
}
