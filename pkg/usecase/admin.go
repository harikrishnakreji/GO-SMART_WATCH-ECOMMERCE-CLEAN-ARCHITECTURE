package usecase

import (
	"errors"

	"github.com/go-playground/validator/v10"
	domain "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/helper"
	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	// getting details of the admin based on the email provided
	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

	//  copy all details except password and sent it back to the front end
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil

}

// sign up handler for the admin
func (ad *adminUseCase) CreateAdmin(admin models.AdminSignUp) (domain.TokenAdmin, error) {

	// validator package to check the constraints specified in the struct which is used to retrieve these details
	if err := validator.New().Struct(admin); err != nil {
		return domain.TokenAdmin{}, err
	}

	// check whether the admin already exist in the database -
	userExist := ad.adminRepository.CheckAdminAvailability(admin)
	if userExist {
		return domain.TokenAdmin{}, errors.New("admin already exist, sign in")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return domain.TokenAdmin{}, errors.New("internal server error")
	}
	admin.Password = string(hashedPassword)

	adminDetails, err := ad.adminRepository.CreateAdmin(admin)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetails)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetails,
		Token: tokenString,
	}, nil

}

func (ad *adminUseCase) AddGenres(genre models.CategoryUpdate) error {

	return ad.adminRepository.AddGenre(genre)

}

func (ad *adminUseCase) Delete(genre_id string) error {

	err := ad.adminRepository.Delete(genre_id)
	if err != nil {
		return err
	}
	return nil

}

func (ad *adminUseCase) GetGenres() ([]domain.Genre, error) {

	return ad.adminRepository.GetGenres()

}

func (ad *adminUseCase) GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := ad.adminRepository.GetUsers(page, count)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

// block user
func (ad *adminUseCase) BlockUser(id string) error {

	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}

// unblock user
func (ad *adminUseCase) UnBlockUser(id string) error {

	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}

func (ad *adminUseCase) DashBoard() (models.CompleteAdminDashboard, error) {

	userDetails, err := ad.adminRepository.DashboardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	productDetails, err := ad.adminRepository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	return models.CompleteAdminDashboard{
		DashboardUser:    userDetails,
		DashBoardProduct: productDetails,
	}, nil

}
