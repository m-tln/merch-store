package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	DATABASE_PORT     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
	DATABASE_HOST     string
}

type ServerConfig struct {
	SERVER_PORT string
	SERVER_HOST string
}

type Config struct {
	DBConfig
	ServerConfig
	secretJWT string
}

func NewConfig() (*Config, error) {
	return &Config{
		DBConfig: DBConfig{
			DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
			DATABASE_PORT:     os.Getenv("DATABASE_PORT"),
			DATABASE_USER:     os.Getenv("DATABASE_USER"),
			DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
			DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
		},
		ServerConfig: ServerConfig{
			SERVER_PORT: os.Getenv("SERVER_PORT"),
			SERVER_HOST: os.Getenv("SERVER_HOST"),
		},
		secretJWT: os.Getenv("SECRET_JWT"),
	}, nil
}

func (cfg *Config) GetDSN() (string, error) {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DATABASE_HOST,
		cfg.DATABASE_PORT,
		cfg.DATABASE_USER,
		cfg.DATABASE_NAME,
		cfg.DATABASE_PASSWORD,
	), nil
}

func (cfg *Config) GetSecretJWT() (string, error) {
	return cfg.secretJWT, nil
}

func (cfg *Config) GetPort() string {
	return cfg.SERVER_PORT
}
