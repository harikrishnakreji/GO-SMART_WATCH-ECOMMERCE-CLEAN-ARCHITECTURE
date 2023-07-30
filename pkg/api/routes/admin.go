package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/handler"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/middleware"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, userHandler *handler.UserHandler) {

	router.POST("/adminlogin", adminHandler.LoginHandler)

	router.POST("/create-admin", adminHandler.CreateAdmin)
	router.POST("/add_category", adminHandler.AddCategorys)

	router.Use(middleware.AuthorizationMiddleware)
	{
		router.GET("/dashboard", adminHandler.DashBoard)
		// router.POST("/create-admin", adminHandler.CreateAdmin)

		// categories := router.Group("/categories")
		// {
		// 	categories.GET("", adminHandler.GetCategorys) // change this to get category
		// 	categories.POST("/add_category", adminHandler.AddCategorys)
		// 	categories.DELETE("/delete_category/:id", adminHandler.DeleteCategory)
		// }

		product := router.Group("/products")
		{
			product.GET("", productHandler.SeeAllProductToAdmin)
			product.GET("/:page", productHandler.SeeAllProductToAdmin)
			product.POST("/add-product", productHandler.AddProduct)
			product.DELETE("/delete-product/:id", productHandler.DeleteProduct)

		}

		userDetails := router.Group("/users")
		{
			userDetails.GET("", adminHandler.GetUsers)
			userDetails.GET("/:page", adminHandler.GetUsers)
			userDetails.GET("/block-users/:id", adminHandler.BlockUser)
			userDetails.GET("/unblock-users/:id", adminHandler.UnBlockUser)
		}
	}
}
