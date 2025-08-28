package http

import (
	"github.com/dmxmss/e-commerce-app/config"
	"github.com/dmxmss/e-commerce-app/service"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServerInterface interface {
	SignUp(echo.Context) error
	LogIn(echo.Context) error
	RefreshTokens(echo.Context) error
	GetUserInfo(echo.Context) error
	CreateProduct(echo.Context) error
	DeleteProduct(echo.Context) error
	ErrorHandler(error, echo.Context)
}

type Server struct {
	service Service
}

type Service struct {
	auth service.AuthService
	user service.UserService
	product service.ProductService

}

func NewEchoServer(conf *config.Config, db *gorm.DB) (ServerInterface, error) {
	userRepo := repository.NewUserRepository(db)
	authRepo, err := repository.NewAuthRepository(conf.Auth)
	if err != nil {
		return nil, err
	}

	authService, err := service.NewAuthService(conf, authRepo, userRepo)
	if err != nil {
		return nil, err
	}

	userService := service.NewUserService(db, userRepo)
	productService := service.NewProductService(db)

	return &Server{
		Service{
			auth: authService,
			user: userService,
			product: productService,
		},
	}, nil
}
