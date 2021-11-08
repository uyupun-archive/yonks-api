package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int            `gorm:"autoIncrement" json:"id" yaml:"id"`
	UserID       string         `gorm:"unique" json:"user_id" yaml:"userId"`
	Password     string         `gorm:"notNull" json:"password" yaml:"password"`
	Name         string         `gorm:"notNull" json:"name" yaml:"name"`
	StatusID     int            `gorm:"notNull default:1" json:"status_id" yaml:"statusId"`
	SNSLine      string         `gorm:"default:null" json:"sns_line" yaml:"snsLine"`
	SNSTwitter   string         `gorm:"default:null" json:"sns_twitter" yaml:"snsTwitter"`
	SNSInstagram string         `gorm:"default:null" json:"sns_instagram" yaml:"snsInstagram"`
	SNSTikTok    string         `gorm:"default:null" json:"sns_tiktok" yaml:"snsTiktok"`
	CreatedAt    time.Time      `json:"created_at" yaml:"createdAt"`
	UpdatedAt    time.Time      `json:"updated_at" yaml:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" yaml:"deletedAt"`
}
