package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
)

func AuthRegister(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: DBに保存する処理

	return c.JSON(http.StatusOK, "")
}

func AuthLogin(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: ユーザIDとパスワードを検証する処理
	if user.UserID != "takashi0602" || user.Password != "testtest" {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: トークン等を返す
	return c.JSON(http.StatusOK, "")
}
