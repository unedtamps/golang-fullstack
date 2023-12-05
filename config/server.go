package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Host   string
	Port   string
}

func GetSeverEnv() ServerConfig {
	err := godotenv.Load()
	if err != nil{
		log.Println(err.Error())
	}
	return ServerConfig{
		Host: os.Getenv("SERVER_HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}
}