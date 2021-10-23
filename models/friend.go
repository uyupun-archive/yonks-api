package models

type Friend struct {
	ID      int `json:"id" yaml:"id"`
	UserAID int `json:"user_a_id" yaml:"user_a_id"`
	UserBID int `json:"user_b_id" yaml:"user_b_id"`
}
