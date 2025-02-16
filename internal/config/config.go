package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Port     string
	User     string
	Password string
	Name     string
	Host     string
}

type ServerConfig struct {
	port string
	host string
}

type Config struct {
	DBConfig
	ServerConfig
	secretJWT string
}

func NewConfig() (*Config, error) {
	return &Config{
		DBConfig: DBConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
		},
		ServerConfig: ServerConfig{
			port: os.Getenv("SERVER_PORT"),
			host: os.Getenv("SERVER_HOST"),
		},
		secretJWT: os.Getenv("SECRET_JWT"),
	}, nil
}

func (cfg *Config) GetDSN() (string, error) {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.User,
		cfg.Name,
		cfg.Password,
	), nil
}

func (cfg *Config) GetSecretJWT() (string, error) {
	return cfg.secretJWT, nil
}

func (cfg *Config) GetPort() string {
	return cfg.ServerConfig.port
}

func (cfg *Config) GetHost() string {
	return cfg.ServerConfig.host
}

func (cfg *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", cfg.GetHost(), cfg.GetPort())
}
