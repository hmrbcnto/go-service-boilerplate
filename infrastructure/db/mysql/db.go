package db

import (
	"fmt"

	"github.com/go-service-boilerplate/config"
	"github.com/go-service-boilerplate/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient(cfg config.DBcfg) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&entities.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
