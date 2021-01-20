package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetConection to db
func GetConnection() (db *gorm.DB, err error) {
	connectionString := os.Getenv("GORM_CONNECTION")
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	return
}
