package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host    string
	Port    string
	Pass    string
	User    string
	SSLMode string
	DBName  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s pass=%s user=%s sslmode=%s dbname=%s",
		config.Host, config.Port, config.Pass, config.User, config.SSLMode, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
