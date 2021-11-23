package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
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
	friend := new(models.Friend)
	err := c.Bind(&friend)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// DBに保存する処理

	return c.JSON(http.StatusOK, "")
}
