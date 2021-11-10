package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
	"golang.org/x/crypto/bcrypt"
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
	loginInfo := struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}{}

	err := c.Bind(&loginInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	password := []byte(loginInfo.Password)
	hashed, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	// DBのユーザ情報を取得

	// パスワードの検証
	err = bcrypt.CompareHashAndPassword(hashed, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	// トークンの生成

	return c.JSON(http.StatusOK, nil)
}

func AuthLogout(c echo.Context) error {
	return nil
}
