//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api"
	handler "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/api/handler"
	config "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/config"
	db "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/db"
	repository "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository"
	usecase "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP, repository.NewProductRepository, usecase.NewProductUseCase, handler.NewProductHandler, handler.NewOtpHandler, usecase.NewOtpUseCase, repository.NewOtpRepository, repository.NewAdminRepository, usecase.NewAdminUseCase, handler.NewAdminHandler, handler.NewCartsHandler, usecase.NewCartsUseCase, repository.NewCartsRepository)

	return &http.ServerHTTP{}, nil
}
