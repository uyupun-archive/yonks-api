package main

import (
	"fmt"

	"github.com/uyupun/yonks-api/models"
	"github.com/uyupun/yonks-api/utility/database"
)

func main() {
	users := []models.User{
		{
			UserID:   "test001",
			Password: "testtest",
			Name:     "テスト太郎",
			StatusID: 1,
		},
	}

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.Create(&users)

	fmt.Printf("Applied users seed!\n")
}
