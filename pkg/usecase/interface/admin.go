package interfaces

import (
	domain "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	CreateAdmin(admin models.AdminSignUp) (domain.TokenAdmin, error)
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	GetGenres() ([]domain.Genre, error)
	AddGenres(genre models.CategoryUpdate) error
	Delete(genre_id string) error
	BlockUser(id string) error
	UnBlockUser(id string) error
	DashBoard() (models.CompleteAdminDashboard, error)
}
