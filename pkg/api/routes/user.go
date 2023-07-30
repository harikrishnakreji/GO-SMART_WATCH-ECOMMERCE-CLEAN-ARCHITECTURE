package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/handler"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/middleware"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler, cartHandler *handler.CartHandler) {

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
	router.Use(middleware.AuthMiddleware())

	{
		cart := router.Group("/cart")
		{
			cart.POST("/addtocart/:id", cartHandler.AddToCart)
			cart.DELETE("/removefromcart/:id", cartHandler.RemoveFromCart)
			cart.GET("", cartHandler.DisplayCart)
			cart.DELETE("", cartHandler.EmptyCart)
		}
	}

}
