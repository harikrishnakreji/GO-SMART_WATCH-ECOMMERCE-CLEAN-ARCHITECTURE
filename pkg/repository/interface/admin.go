package interfaces

import (
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	CreateAdmin(admin models.AdminSignUp) (models.AdminDetailsResponse, error)
	CheckAdminAvailability(admin models.AdminSignUp) bool
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	GetGenres() ([]domain.Genre, error)
	AddGenre(genre models.CategoryUpdate) error
	Delete(genre_id string) error
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	DashboardUserDetails() (models.DashboardUser, error)
	DashBoardProductDetails() (models.DashBoardProduct, error)
}
