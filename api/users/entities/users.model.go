package entities

import (
	entities "crud/api/roles/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string       `gorm:"not null"`
	LastName       string       `gorm:"not null"`
	Email          string       `gorm:"not null"`
	Password       string       `gorm:"not null"`
	SecondLastName string       `gorm:"not null"`
	Photo          string       `gorm:"not null"`
	RolId          int          `gorm:"not null"`
	Rol            entities.Rol `gorm:"not null"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
func (u *User) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
