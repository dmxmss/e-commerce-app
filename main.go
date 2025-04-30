package main

import (
	"github.com/dmxmss/e-commerce-app/config"
	"github.com/dmxmss/e-commerce-app/internal/http"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
	"log"
)

func main() {
	conf := config.GetConfig()
	log.Printf("%d", conf.Auth.Access.Expiration)

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
	accessMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(conf.Auth.JWTSecret),
		TokenLookup: "header:Authorization",
		ContextKey: "user",
	})

	refreshMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(conf.Auth.JWTSecret),
		TokenLookup: "cookie:refresh_token",
		ContextKey: "refresh",
	})

	auth.POST("/signup", s.SignUp)	
	auth.POST("/login", s.LogIn)	
	auth.POST("/refresh", s.RefreshTokens, refreshMiddleware)	

	e.GET("/me", s.GetUserInfo, accessMiddleware)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", conf.App.Address, conf.App.Port)))
}
