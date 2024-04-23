package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	dbSettings := `host=localhost port=5432 user=admin password=admin dbname=mydb`

	db, err := gorm.Open(postgres.Open(dbSettings), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
