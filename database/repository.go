package database

import (
	"github.com/jinzhu/gorm"
)

// PostgreRepository is
type PostgreRepository struct {
	db *gorm.DB
}

// CreateRepository is
func CreateRepository(database *gorm.DB) PostgreRepository {
	return PostgreRepository{database}
}
