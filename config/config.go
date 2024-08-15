package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}
	return &Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
}
