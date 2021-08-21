package entities

import (
	"time"
)

type Access struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
	Name        string    `gorm:"not null" json:"name"`
	KeyName     string    `gorm:"not null" json:"keyName"`
	Icon        string    `gorm:"not null" json:"icon"`
	Description string    `gorm:"not null" json:"description"`
	ModuleID    int       `gorm:"not null" json:"moduleId"`
	Module      Module    `json:"module"`
}
