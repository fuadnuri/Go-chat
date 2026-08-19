package postgres

import (
	"os"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pql struct {
	DB *gorm.DB
}

func NewDatabase() (*gorm.DB, error) {
	conn_string := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(conn_string), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("coudn't create Coonection: %s", err.Error())
	}
	return db, nil
}
