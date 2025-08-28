package config

import (
	"github.com/spf13/viper"

	"sync"
	"strings"
)

type (
	Config struct {
		App *App
		Database *Database
		Auth *Auth
	}

	App struct {
		Address string	
		Port int
	}

	Database struct {
		Host string
		Name string	
		User string
		Port string
		Password string
	}

	Auth struct {
		JWTSecret string
		SigningMethod string
		Access *Token
		Refresh *Token
	}

	Token struct {
		Expiration int // seconds
	}
)

var (
	once sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func () {
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		viper.SetDefault("app.address", "127.0.0.1")
		viper.SetDefault("app.port", "8080")

		viper.SetDefault("database.name", "postgres")
		viper.SetDefault("database.user", "postgres")
		viper.SetDefault("database.port", "5432")
		viper.SetDefault("database.host", "db")
		viper.SetDefault("database.password", "")

		viper.SetDefault("auth.access.expiration", 60*60)
		viper.SetDefault("auth.refresh.expiration", 60*60*24*7)
		viper.SetDefault("auth.signingmethod", "HS256")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
