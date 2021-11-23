package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/utility/database"
)

func GetFriends(c echo.Context) error {
	userID := c.Param("user_id")

	friends, err := database.FindFriendsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	return c.JSON(http.StatusOK, friends)
}

func AddFriend(c echo.Context) error {
	addFriendInfo := struct {
		UserID       string `json:"user_id"`
		TargetUserID string `json:"target_user_id"`
	}{}
	err := c.Bind(&addFriendInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = database.CreateFriend(addFriendInfo.UserID, addFriendInfo.TargetUserID)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	return c.JSON(http.StatusOK, nil)
}
