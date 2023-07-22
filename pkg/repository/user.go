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
