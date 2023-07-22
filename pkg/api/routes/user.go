package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/handler"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler) {

	// USER SIDE
	router.POST("/signup", userHandler.UserSignUp)
	router.POST("/login", userHandler.LoginHandler)

	router.POST("/send-otp", otpHandler.SendOTP)
	router.POST("/verify-otp", otpHandler.VerifyOTP)

	product := router.Group("/products")
	{
		product.GET("", productHandler.ShowAllProducts)
		product.GET("/page/:page", productHandler.ShowAllProducts)
	}

}
