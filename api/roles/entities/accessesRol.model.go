package entities

import (
	"time"
)

type AccessesRol struct {
	RolID      int       `gorm:"primaryKey" json:"rolId"`
	AccessID   int       `gorm:"primaryKey" json:"accessId"`
	Permission string    `gorm:"not null" json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time
}
