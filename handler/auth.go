package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthRegister(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func AuthLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
