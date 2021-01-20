package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User is
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password" gorm:"size:255"` // set field size to 255
}

// Profile belongs to `User`, `UserID` is the foreign key
type Profile struct {
	gorm.Model
	UserID      int
	User        User
	Address     string
	PhoneNumber string
	UserName    string
}

// CreateTable -
func (t User) CreateTable(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&User{}).Error

	return
}

// CreateTable -
func (t Profile) CreateTable(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&Profile{}).Error

	return
}
