package database

import "github.com/uyupun/yonks-api/models"

func CreateUser(user models.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	db.Create(&user)

	return nil
}
