package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
	"github.com/uyupun/yonks-api/utility/database"
	"golang.org/x/crypto/bcrypt"
)

type AuthInfo struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

func AuthRegister(c echo.Context) error {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	password := []byte(user.Password)
	hashed, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	user.Password = string(hashed)

	err = database.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	// トークンの生成

	return c.JSON(http.StatusOK, nil)
}

func AuthLogin(c echo.Context) error {
	var loginInfo AuthInfo

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
