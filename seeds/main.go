package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/uyupun/yonks-api/models"
	"github.com/uyupun/yonks-api/utility/database"
)

func main() {
	bytes, err := ioutil.ReadFile("seeds/mocks/users.json")
	if err != nil {
		panic(err)
	}

	var users []models.User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.Create(&users)

	fmt.Printf("Applied users seed!\n")
}
