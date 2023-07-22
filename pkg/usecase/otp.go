package usecase

import (
	"errors"

	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/config"
	helper "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/helper"
	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"github.com/jinzhu/copier"
)

type otpUseCase struct {
	cfg           config.Config
	otpRepository interfaces.OtpRepository
}

func NewOtpUseCase(cfg config.Config, repo interfaces.OtpRepository) services.OtpUseCase {
	return &otpUseCase{
		cfg:           cfg,
		otpRepository: repo,
	}
}

func (ot *otpUseCase) SendOTP(phone string) error {

	ok := ot.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, ot.cfg.SERVICESSID)
	if err != nil {
		return errors.New("error ocurred while generating OTP")
	}

	return nil

}

func (ot *otpUseCase) VerifyOTP(code models.VerifyData) (models.TokenUsers, error) {

	helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	err := helper.TwilioVerifyOTP(ot.cfg.SERVICESSID, code.Code, code.User.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, errors.New("error while verifying")
	}

	// if user is authenticated using OTP send back user details
	userDetails, err := ot.otpRepository.UserDetailsUsingPhone(code.User.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := helper.GenerateTokenUsers(userDetails)

	var user models.UserDetailsResponse
	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: user,
		Token: tokenString,
	}, nil

}
