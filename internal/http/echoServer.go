package http

import (
	"github.com/dmxmss/e-commerce-app/config"
	"github.com/dmxmss/e-commerce-app/service"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"fmt"
)

type ServerInterface interface {
	SignUp(echo.Context) error
	LogIn(echo.Context) error
	RefreshTokens(echo.Context) error
	GetUserInfo(echo.Context) error
	CreateProduct(echo.Context) error
	DeleteProduct(echo.Context) error
	ErrorHandler(error, echo.Context)

	Start()
}

type Server struct {
	echo *echo.Echo
	service Service
	conf *config.Config
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

	echo := echo.New()

	s := Server{
		echo,
		Service{
			auth: authService,
			user: userService,
			product: productService,
		},
		conf,
	}

	s.setUpRouter()

	return &s, nil
}

func (s Server) Start() {
	s.echo.Logger.Fatal(
		s.echo.Start(
			fmt.Sprintf("%s:%d", 
				s.conf.App.Address, 
				s.conf.App.Port,
			),
		),
	)
}

func (s Server) setUpRouter() { // routes, middleware
	s.echo.HTTPErrorHandler = s.ErrorHandler

	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("%v %v\n", v.URI, v.Status)
			return nil
		},
	}))

	accessMiddleware := GetAccessMiddleware(s.conf.Auth.JWTSecret, s.conf.Auth.SigningMethod)
	refreshMiddleware := GetRefreshMiddleware(s.conf.Auth.JWTSecret, s.conf.Auth.SigningMethod)

	api := s.echo.Group("/api")

	auth := s.echo.Group("/auth")
	products := api.Group("/products", accessMiddleware)

	auth.POST("/signup", s.SignUp)	
	auth.POST("/login", s.LogIn)	
	auth.POST("/refresh", s.RefreshTokens, refreshMiddleware)	

	products.POST("/", s.CreateProduct)
	products.GET("/", s.GetUserProducts)
	products.DELETE("/{id}", s.DeleteProduct)

	s.echo.GET("/me", s.GetUserInfo, accessMiddleware)
}
