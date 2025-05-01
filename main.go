package main

import (
	"github.com/dmxmss/e-commerce-app/config"
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/http"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
)

func main() {
	conf := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Name,
		conf.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})	
	if err != nil {
		panic(err)
	}

	db = db.Debug()

	s, err := http.NewEchoServer(conf, db)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	auth := e.Group("/auth")

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("%v %v %v\n", v.Method, v.URI, v.Status)
			return nil
		},
	}))

	accessMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(conf.Auth.JWTSecret),
		TokenLookup: "header:Authorization:Bearer ",
		ContextKey: "user",
		SigningMethod: conf.Auth.SigningMethod,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.Claims)
		},
	})

	refreshMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(conf.Auth.JWTSecret),
		TokenLookup: "cookie:refresh_token",
		ContextKey: "refresh",
		SigningMethod: conf.Auth.SigningMethod,
	})

	auth.POST("/signup", s.SignUp)	
	auth.POST("/login", s.LogIn)	
	auth.POST("/refresh", s.RefreshTokens, refreshMiddleware)	

	e.GET("/me", s.GetUserInfo, accessMiddleware)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", conf.App.Address, conf.App.Port)))
}
