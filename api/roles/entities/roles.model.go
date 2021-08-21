package entities

import (
	"time"
)

type Rol struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time     `gorm:"index"`
	Name      string        `gorm:"not null" json:"name"`
	Accesses  []AccessesRol `gorm:"many2many:accesses_rol" json:"accesses"`
}
