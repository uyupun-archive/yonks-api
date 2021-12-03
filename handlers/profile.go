package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
	"github.com/uyupun/yonks-api/utils/database"
)

func GetProfile(c echo.Context) error {
	userID := c.Param("user_id")

	user, err := database.FindUserByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	return c.JSON(http.StatusOK, user)
}

func SaveProfile(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: claimsと比較して正当なユーザからのリクエストかどうかを見る

	err = database.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	return c.JSON(http.StatusOK, nil)
}
