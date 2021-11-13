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

func FindUserByUserID(user *models.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	db.Find(&user, "user_id = ?", user.UserID)
	return nil
}
