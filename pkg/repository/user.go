package repository

import (
	"errors"
	"fmt"

	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserDatabase{DB}
}

// check whether the user is already present in the database . If there recommend to login
func (c *UserDatabase) CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}

// retrieve the user details form the database
func (c *UserDatabase) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {

	var userDetails models.UserSignInResponse

	err := c.DB.Raw(`
		SELECT *
		FROM users where email = ? and blocked = false
		`, user.Email).Scan(&userDetails).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}

	return userDetails, nil

}

func (c *UserDatabase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse
	err := c.DB.Raw(`INSERT INTO users (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id, name, email, phone`, user.Name, user.Email, user.Phone, user.Password).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, nil
}

func (c *UserDatabase) LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userResponse models.UserDetailsResponse
	err := c.DB.Save(&userResponse).Error
	return userResponse, err

}

func (cr *UserDatabase) UserPassword(userID int) (string, error) {

	var userPassword string
	err := cr.DB.Raw("select password from users where id = ?", userID).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}

func (cr *UserDatabase) UserBlockStatus(email string) (bool, error) {

	var isBlocked bool
	err := cr.DB.Raw("select blocked from users where email = ?", email).Scan(&isBlocked).Error
	if err != nil {
		return false, err
	}

	return isBlocked, nil
}

func (cr *UserDatabase) AddAddress(address models.AddressInfo, userID int) error {

	fmt.Println(address)
	err := cr.DB.Exec("insert into addresses (user_id,name,house_name,state,pin,street,city) values (?, ?, ?, ?, ?, ?, ?)", userID, address.Name, address.HouseName, address.State, address.Pin, address.Street, address.City).Error
	if err != nil {
		return err
	}

	return nil

}

func (cr *UserDatabase) UpdateAddress(address models.AddressInfo, addressID int, userID int) (models.AddressInfoResponse, error) {

	err := cr.DB.Exec("update addresses set house_name = ?, state = ?, pin = ?, street = ?, city = ? where id = ? and user_id = ?", address.HouseName, address.State, address.Pin, address.Street, address.City, addressID, userID).Error
	if err != nil {
		return models.AddressInfoResponse{}, err
	}

	var addressResponse models.AddressInfoResponse
	err = cr.DB.Raw("select * from addresses where id = ?", addressID).Scan(&addressResponse).Error
	if err != nil {
		return models.AddressInfoResponse{}, err
	}

	return addressResponse, nil

}

func (cr *UserDatabase) GetAllAddresses(userID int) ([]models.AddressInfoResponse, error) {

	var addressResponse []models.AddressInfoResponse
	err := cr.DB.Raw(`select * from addresses where user_id = $1`, userID).Scan(&addressResponse).Error
	if err != nil {
		return []models.AddressInfoResponse{}, err
	}

	return addressResponse, nil

}

func (cr *UserDatabase) GetAllPaymentOption() ([]models.PaymentDetails, error) {

	var paymentMethods []models.PaymentDetails
	err := cr.DB.Raw("select * from payment_methods").Scan(&paymentMethods).Error
	if err != nil {
		return []models.PaymentDetails{}, err
	}

	return paymentMethods, nil

}

func (cr *UserDatabase) UserDetails(userID int) (models.UsersProfileDetails, error) {

	var userDetails models.UsersProfileDetails
	err := cr.DB.Raw("select users.name,users.email,users.phone from users where id=?", userID).Row().Scan(&userDetails.Name, &userDetails.Email, &userDetails.Phone)
	if err != nil {
		return models.UsersProfileDetails{}, err
	}

	return userDetails, nil
}

func (cr *UserDatabase) UpdateUserEmail(email string, userID int) error {

	err := cr.DB.Exec("update users set email = ? where id = ?", email, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *UserDatabase) UpdateUserPhone(phone string, userID int) error {

	err := cr.DB.Exec("update users set phone = ? where id = ?", phone, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *UserDatabase) UpdateUserName(name string, userID int) error {

	err := cr.DB.Exec("update users set name = ? where id = ?", name, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *UserDatabase) UpdateUserPassword(password string, userID int) error {

	err := cr.DB.Exec("update users set password = ? where id = ?", password, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *UserDatabase) ResetPassword(userID int, password string) error {

	err := cr.DB.Exec("update users set password = ? where id = ?", password, userID).Error
	if err != nil {
		return err
	}

	return nil
}
