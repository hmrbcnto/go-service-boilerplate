package db

import (
	"github.com/go-service-boilerplate/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&entities.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
