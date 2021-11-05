package models

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	ID        int       `gorm:"autoIncrement" json:"id" yaml:"id"`
	Name      string    `gorm:"unique" json:"name" yaml:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
