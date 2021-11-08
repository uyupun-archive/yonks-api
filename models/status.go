package models

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	ID        int            `gorm:"autoIncrement" json:"id" yaml:"id"`
	Name      string         `gorm:"unique" json:"name" yaml:"name"`
	CreatedAt time.Time      `json:"created_at" yaml:"createdAt"`
	UpdatedAt time.Time      `json:"updated_at" yaml:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" yaml:"deletedAt"`
}
