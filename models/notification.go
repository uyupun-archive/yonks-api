package models

type Notification struct {
	ID      int    `json:"id" yaml:"id"`
	UserID  int    `json:"user_id" yaml:"user_id"`
	Content string `json:"content" yaml:"content"`
}
