package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFriends(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func AddFriend(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
