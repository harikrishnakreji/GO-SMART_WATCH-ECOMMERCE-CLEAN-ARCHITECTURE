package interfaces

import (
	"time"

	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	CreateAdmin(admin models.AdminSignUp) (models.AdminDetailsResponse, error)
	CheckAdminAvailability(admin models.AdminSignUp) bool
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	GetCategorys() ([]domain.Category, error)
	AddCategory(category models.CategoryUpdate) error
	Delete(category_id string) error
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error)
	DashboardUserDetails() (models.DashboardUser, error)
	DashBoardProductDetails() (models.DashBoardProduct, error)
	TotalRevenue() (models.DashboardRevenue, error)
	DashBoardOrder() (models.DashboardOrder, error)
	AmountDetails() (models.DashboardAmount, error)
}
