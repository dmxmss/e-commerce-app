package main

import (
	"github.com/dmxmss/e-commerce-app/config"
	http "github.com/dmxmss/e-commerce-app/internal/http"
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


	s.Start()
}
