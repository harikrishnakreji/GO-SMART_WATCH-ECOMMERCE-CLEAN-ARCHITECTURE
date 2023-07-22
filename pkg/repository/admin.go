package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := ad.DB.Raw("select * from admins where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil
}

// check if an admin with specified email already exist
func (ad *adminRepository) CheckAdminAvailability(admin models.AdminSignUp) bool {

	var count int
	if err := ad.DB.Raw("select count(*) from admins where email = ?", admin.Email).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (ad *adminRepository) CreateAdmin(admin models.AdminSignUp) (models.AdminDetailsResponse, error) {

	var adminDetails models.AdminDetailsResponse
	if err := ad.DB.Raw("insert into admins (name,email,password) values (?, ?, ?) RETURNING id, name, email", admin.Name, admin.Email, admin.Password).Scan(&adminDetails).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}

	return adminDetails, nil

}

// Get users details for authenticated admins
func (ad *adminRepository) GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var userDetails []models.UserDetailsAtAdmin

	if err := ad.DB.Raw("select id,name,email,phone,blocked from users limit ? offset ?", count, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

func (ad *adminRepository) GetGenres() ([]domain.Genre, error) {

	var genres []domain.Genre
	if err := ad.DB.Raw("select * from genres").Scan(&genres).Error; err != nil {
		return []domain.Genre{}, err
	}

	return genres, nil

}

// CATEGORY MANAGEMENT
func (ad *adminRepository) AddGenre(genre models.CategoryUpdate) error {

	var count int
	err := ad.DB.Raw("select count(*) from genres where genre_name = ?", genre.Genre).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the genre already exist")
	}

	if err := ad.DB.Exec("insert into genres (genre_name) values (?) ", genre.Genre).Error; err != nil {
		return err
	}
	return nil

}

func (ad *adminRepository) Delete(genre_id string) error {

	id, err := strconv.Atoi(genre_id)
	if err != nil {
		return err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from genres where id = ?", id).Scan(&count).Error; err != nil {
		return err
	}
	if count < 1 {
		return errors.New("genre for given id does not exist")
	}

	query := fmt.Sprintf("delete from genres where id = '%d'", id)
	if err := ad.DB.Exec(query).Error; err != nil {
		return err
	}

	return nil

}

func (ad *adminRepository) GetUserByID(id string) (domain.Users, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.Users{}, err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from users where id = ?", id).Scan(&count).Error; err != nil {
		return domain.Users{}, err
	}
	if count < 1 {
		return domain.Users{}, errors.New("user for the given id does not exist")
	}

	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userDetails domain.Users

	if err := ad.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.Users{}, err
	}

	return userDetails, nil

}

// function which will both block and unblock a user
func (ad *adminRepository) UpdateBlockUserByID(user domain.Users) error {

	err := ad.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		fmt.Println("Error updating user:", err)
		return err
	}

	return nil

}

func (ad *adminRepository) DashboardUserDetails() (models.DashboardUser, error) {

	var userDetails models.DashboardUser
	err := ad.DB.Raw("select count(*) from users").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	err = ad.DB.Raw("select count(*) from users where blocked = true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	return userDetails, nil
}

func (ad *adminRepository) DashBoardProductDetails() (models.DashBoardProduct, error) {

	var productDetails models.DashBoardProduct
	err := ad.DB.Raw("select count(*) from products").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	var productID int
	err = ad.DB.Raw("select product_id from order_items group by product_id order by sum(quantity) desc limit 1").Scan(&productID).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	return productDetails, nil
}
