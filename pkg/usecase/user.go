package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/helper"
	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo    interfaces.UserRepository
	cartsRepo   interfaces.CartsRepository
	productRepo interfaces.ProductRepository
}

func NewUserUseCase(repo interfaces.UserRepository, cartsRepositiry interfaces.CartsRepository, productRepository interfaces.ProductRepository) services.UserUseCase {
	return &userUseCase{
		userRepo:    repo,
		cartsRepo:   cartsRepositiry,
		productRepo: productRepository,
	}
}

func (u *userUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {

	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := u.userRepo.CheckUserAvailability(user.Email)

	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}

	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	// Hash password since details are validated
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	// add user details to the database
	userData, err := u.userRepo.UserSignUp(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// crete a JWT token string for the user
	tokenString, err := helper.GenerateTokenUsers(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	// copies all the details except the password of the user
	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userData)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil
}

func (u *userUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	// checking if a username exist with this email address
	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, err
	}

	if isBlocked {
		return models.TokenUsers{}, errors.New("user is not authorized to login")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &user_details)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := helper.GenerateTokenUsers(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}

func (u *userUseCase) AddAddress(address models.AddressInfo, userID int) error {

	err := u.userRepo.AddAddress(address, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) UpdateAddress(address models.AddressInfo, addressID int, userID int) (models.AddressInfoResponse, error) {

	return u.userRepo.UpdateAddress(address, addressID, userID)

}

// user checkout section
func (u *userUseCase) Checkout(userID int) (models.CheckoutDetails, error) {

	// list all address added by the user
	allUserAddress, err := u.userRepo.GetAllAddresses(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get available payment options
	paymentDetails, err := u.userRepo.GetAllPaymentOption()
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get all items from users carts
	cartsItems, err := u.cartsRepo.GetAllItemsFromCarts(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get grand total of all the product
	grandTotal, err := u.cartsRepo.GetTotalPrice(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	return models.CheckoutDetails{
		AddressInfoResponse: allUserAddress,
		Payment_Method:      paymentDetails,
		Carts:               cartsItems,
		Grand_Total:         grandTotal.TotalPrice,
		Total_Price:         grandTotal.FinalPrice,
	}, nil
}

func (u *userUseCase) UserDetails(userID int) (models.UsersProfileDetails, error) {

	return u.userRepo.UserDetails(userID)

}

func (u *userUseCase) GetAllAddress(userID int) ([]models.AddressInfoResponse, error) {

	userAddress, err := u.userRepo.GetAllAddresses(userID)

	if err != nil {
		return []models.AddressInfoResponse{}, nil
	}

	return userAddress, nil

}

func (u *userUseCase) UpdateUserDetails(userDetails models.UsersProfileDetails, ctx context.Context) (models.UsersProfileDetails, error) {

	var userID int
	var ok bool
	// sent value through context - just for studying purpose - not required in this case
	if userID, ok = ctx.Value("userID").(int); !ok {
		return models.UsersProfileDetails{}, errors.New("error retreiving user details")
	}

	userExist := u.userRepo.CheckUserAvailability(userDetails.Email)

	// update with email that does not already exist
	if userExist {
		return models.UsersProfileDetails{}, errors.New("user already exist, choose different email")
	}
	// which all field are not empty (which are provided from the front end should be updated)
	if userDetails.Email != "" {
		u.userRepo.UpdateUserEmail(userDetails.Email, userID)
	}

	if userDetails.Name != "" {
		u.userRepo.UpdateUserName(userDetails.Name, userID)
	}

	if userDetails.Phone != "" {
		u.userRepo.UpdateUserPhone(userDetails.Phone, userID)
	}

	return u.userRepo.UserDetails(userID)

}

func (u *userUseCase) UpdatePassword(ctx context.Context, body models.UpdatePassword) error {

	var userID int
	var ok bool
	if userID, ok = ctx.Value("userID").(int); !ok {
		return errors.New("error retrieving user details")
	}

	userPassword, err := u.userRepo.UserPassword(userID)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(body.OldPassword))
	if err != nil {
		return errors.New("password incorrect")
	}
	fmt.Println(body)
	if body.NewPassword != body.ConfirmNewPassword {
		return errors.New("password does not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)
	if err != nil {
		return errors.New("internal server error")
	}

	return u.userRepo.UpdateUserPassword(string(hashedPassword), userID)

}

func (u *userUseCase) ResetPassword(userID int, pass models.ResetPassword) error {

	if pass.Password != pass.CPassword {
		return errors.New("password does not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass.Password), 10)
	if err != nil {
		return errors.New("internal server error")
	}

	err = u.userRepo.ResetPassword(userID, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil

}
