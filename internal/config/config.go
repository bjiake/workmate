package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/swaggo/swag/example/basic/docs"
	"log"
	configSrv "workmate/internal/config/server"
)

type Config struct {
	Server configSrv.Server
}

func LoadConfig() *Config {
	err := loadDotEnv()
	if err != nil {
		log.Fatal(err)
	}

	srv := configSrv.InitServerConfig()

	return &Config{
		Server: srv,
	}
}

func loadDotEnv() error {
	filePath := fmt.Sprintf(".env")

	err := godotenv.Load(filePath)
	return err
}

func SetSwaggerDefaultInfo(cfg *Config) {
	docs.SwaggerInfo.Title = "IOD App API"
	docs.SwaggerInfo.Description = "API Server for IOD application."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = cfg.Server.Host
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
