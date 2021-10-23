package models

type User struct {
	ID           int    `json:"id" yaml:"id"`
	UserID       string `json:"user_id" yaml:"user_id"`
	Password     string `json:"password" yaml:"password"`
	Name         string `json:"name" yaml:"name"`
	StatusID     int    `json:"status_id" yaml:"status_id"`
	SNSLine      string `json:"sns_line" yaml:"sns_line"`
	SNSTwitter   string `json:"sns_twitter" yaml:"sns_twitter"`
	SNSInstagram string `json:"sns_instagram" yaml:"sns_instagram"`
	SNSTikTok    string `json:"sns_tiktok" yaml:"sns_tiktok"`
}
