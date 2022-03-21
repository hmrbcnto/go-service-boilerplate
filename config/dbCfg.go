package config

import (
	"errors"
)

type DBcfg struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
}

// Validates structs without missing fields
func (db *DBcfg) validate() error {
	if db.DB_HOST == "" {
		return errors.New("DB_HOST is not supplied")
	}

	if db.DB_NAME == "" {
		return errors.New("DB_NAME is not supplied")
	}

	if db.DB_USER == "" {
		return errors.New("DB_USERNAME is not supplied")
	}

	if db.DB_PASSWORD == "" {
		return errors.New("DB_PASSWORD is not supplied")
	}

	return nil
}
