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

func FindUserByUserID(userID string) (models.User, error) {
	var user models.User
	db, err := ConnectDB()
	if err != nil {
		return user, err
	}

	db.First(&user, "user_id = ?", userID)
	return user, nil
}

func UpdateUser(user models.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	db.Model(&models.User{}).Where("user_id = ?", user.UserID).Updates(models.User{
		Name:         user.Name,
		StatusID:     user.StatusID,
		SNSLine:      user.SNSLine,
		SNSTwitter:   user.SNSTwitter,
		SNSInstagram: user.SNSInstagram,
		SNSTikTok:    user.SNSTikTok,
	})
	return nil
}

func FindFriendsByUserID(userID string) ([]models.Friend, error) {
	var friends []models.Friend

	user, err := FindUserByUserID(userID)
	if err != nil {
		return friends, err
	}

	db, err := ConnectDB()
	if err != nil {
		return friends, err
	}

	db.Where("user_id = ?", user.ID).Find(&friends)
	return friends, nil
}
