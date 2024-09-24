package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

type Config struct {
	Host		string
	Port		string
	User		string
	Password	string
	DbName		string
	SSLMode		string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DbName, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}