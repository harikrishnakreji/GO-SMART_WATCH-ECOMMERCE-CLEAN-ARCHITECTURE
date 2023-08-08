package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/handler"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/middleware"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler, cartsHandler *handler.CartsHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) {

	// USER SIDE
	router.POST("/signup", userHandler.UserSignUp)
	router.POST("/login", userHandler.LoginHandler)

	router.POST("/send-otp", otpHandler.SendOTP)
	router.POST("/verify-otp", otpHandler.VerifyOTP)

	forgotPassword := router.Group("/forgot-password")
	{
		forgotPassword.POST("", otpHandler.SendOTPtoReset)
		forgotPassword.POST("/verify-otp", otpHandler.VerifyOTPToReset)

		forgotPassword.Use(middleware.AuthMiddlewareReset())
		forgotPassword.PUT("/reset", userHandler.ResetPassword)

	}
	product := router.Group("/products")
	{
		product.GET("", productHandler.ShowAllProducts)
		product.GET("/page/:page", productHandler.ShowAllProducts)
		product.GET("/:id", productHandler.ShowIndividualProducts)
	}
	router.Use(middleware.AuthMiddleware())

	{
		carts := router.Group("/carts")
		{
			carts.POST("/addtocarts/:id", cartsHandler.AddToCarts)
			carts.DELETE("/removefromcarts/:id", cartsHandler.RemoveFromCarts)
			carts.GET("", cartsHandler.DisplayCarts)
			carts.DELETE("", cartsHandler.EmptyCarts)
			carts.POST("/order", orderHandler.OrderItemsFromCarts)
		}
	}

	address := router.Group("/address")
	{
		address.POST("", userHandler.AddAddress)
		address.PUT("/:id", userHandler.UpdateAddress)
	}

	users := router.Group("/user")
	{

		users.GET("", userHandler.UserDetails)
		users.PUT("", userHandler.UpdateUserDetails)
		users.GET("/address", userHandler.GetAllAddress)
		users.POST("/address", userHandler.AddAddress)
		orders := users.Group("/orders")
		{
			orders.GET("", orderHandler.GetOrderDetails)
			orders.GET("/:page", orderHandler.GetOrderDetails)
		}

		users.PUT("/cancel-order/:id", orderHandler.CancelOrder)
		users.PUT("/update-password", userHandler.UpdatePassword)

		users.GET("/delivered/:order_id", orderHandler.OrderDelivered)
		users.GET("/return/:order_id", orderHandler.ReturnOrder)

		router.GET("/checkout", userHandler.CheckOut)
		router.POST("/order", orderHandler.OrderItemsFromCarts)

		router.GET("/payment/:id", paymentHandler.MakePaymentRazorPay)
		router.GET("/payment-success", paymentHandler.VerifyPayment)

	}

}
