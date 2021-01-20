package user

import (
	"fmt"

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

// GetUserByEmail is
func (r *PostgreRepository) GetUserByEmail(email string) (data User, err error) {
	err = r.db.Where("email = ?", email).First(&data).Error

	if err != nil {
		return data, fmt.Errorf("get user by email fail: %s", err.Error())
	}
	return
}

// GetUserByName is
func (r *PostgreRepository) GetUserByName(name string) (data User, err error) {
	err = r.db.Where("name = ?", name).First(&data).Error
	if err != nil {
		return data, fmt.Errorf("get user by email fail: %s", err.Error())
	}
	return
}
