package http

import (
	"github.com/dmxmss/e-commerce-app/config"
	"github.com/dmxmss/e-commerce-app/service"
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
	authService service.AuthService
	userService service.UserService
	productService service.ProductService
}

func NewEchoServer(conf *config.Config, db *gorm.DB) (ServerInterface, error) {
	authService, err := service.NewAuthService(conf)
	if err != nil {
		return nil, err
	}

	userService := service.NewUserService(db)

	productService := service.NewProductService(db)

	return &Server{
		authService: authService,
		userService: userService,
		productService: productService,
	}, nil
}
