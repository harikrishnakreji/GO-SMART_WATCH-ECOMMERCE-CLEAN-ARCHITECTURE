package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/response"
)

type OtpHandler struct {
	otpUseCase services.OtpUseCase
}

func NewOtpHandler(useCase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase: useCase,
	}
}

// @Summary  OTP login
// @Description Send OTP to Authenticate user
// @Tags User OTP Login
// @Accept json
// @Produce json
// @Param phone body models.OTPData true "phone number details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /send-otp [post]
func (ot *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData

	if err := c.ShouldBindJSON(&phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}

	err := ot.otpUseCase.SendOTP(phone.PhoneNumber)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Verify OTP
// @Description Verify OTP by passing the OTP in order to authenticate user
// @Tags User OTP Login
// @Accept json
// @Produce json
// @Param phone body models.VerifyData true "Verify OTP Details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /verify-otp [post]
func (ot *OtpHandler) VerifyOTP(c *gin.Context) {

	var code models.VerifyData

	if err := c.ShouldBindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ot.otpUseCase.VerifyOTP(code)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", users, nil)
	c.JSON(http.StatusOK, successRes)

}
