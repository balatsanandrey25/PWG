package models

import (
	"fmt"
	"log"

	"github.com/go-sqlx/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("we have successfully called DB")
	return db, nil
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load .env file %s", err.Error())
	}
	return err
}

// функция для добавление конфигов
func loadViper() error {
	viper.AddConfigPath("./pkg/configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}

func ConfigurationInstaller() error {
	if err := loadViper(); err != nil {
		return err
	}

	if err := loadEnv(); err != nil {
		return err
	}
	return nil
}
