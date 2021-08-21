package entities

import (
	"time"
)

type Module struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
	Name        string    `json:"name" gorm:"not null"`
	KeyName     string    `json:"keyName" gorm:"not null"`
	Icon        string    `json:"icon" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Accesses    []Access  `json:"accesses"`
}
