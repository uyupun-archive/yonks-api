package models

import (
	"time"

	"gorm.io/gorm"
)

type Friend struct {
	ID           int `gorm:"autoIncrement" json:"id" yaml:"id"`
	UserID       int `json:"user_id" yaml:"userId"`
	TargetUserID int `json:"target_user_id" yaml:"targetUserId"`
	User         User
	TargetUser   User
	CreatedAt    time.Time      `json:"created_at" yaml:"createdAt"`
	UpdatedAt    time.Time      `json:"updated_at" yaml:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" yaml:"deletedAt"`
}
