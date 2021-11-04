package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int       `gorm:"autoIncrement" json:"id" yaml:"id"`
	UserID       string    `gorm:"unique" json:"user_id" yaml:"user_id"`
	Password     string    `gorm:"notNull" json:"password" yaml:"password"`
	Name         string    `gorm:"notNull" json:"name" yaml:"name"`
	StatusID     int       `gorm:"notNull" json:"status_id" yaml:"status_id"`
	SNSLine      string    `gorm:"default:null" json:"sns_line" yaml:"sns_line"`
	SNSTwitter   string    `gorm:"default:null" json:"sns_twitter" yaml:"sns_twitter"`
	SNSInstagram string    `gorm:"default:null" json:"sns_instagram" yaml:"sns_instagram"`
	SNSTikTok    string    `gorm:"default:null" json:"sns_tiktok" yaml:"sns_tiktok"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt
}
