package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func SaveProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
