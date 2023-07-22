package http

import (
	"github.com/gin-gonic/gin"

	_ "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/cmd/api/docs"
	handler "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/handler"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, productHandler *handler.ProductHandler, otpHandler *handler.OtpHandler, adminHandler *handler.AdminHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(router.Group("/"), userHandler, otpHandler, productHandler)
	routes.AdminRoutes(router.Group("/admin"), adminHandler, productHandler, userHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
