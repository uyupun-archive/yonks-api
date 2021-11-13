package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
	"github.com/uyupun/yonks-api/utility/auth"
	"github.com/uyupun/yonks-api/utility/database"
)

func AuthRegister(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	hashed, err := auth.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	user.Password = hashed

	err = database.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	token, err := auth.IssueToken()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
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

	user := models.User{
		UserID: loginInfo.UserID,
	}
	err = database.FindUserByUserID(&user)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	isValidPassword := auth.IsValidPassword(loginInfo.Password, user.Password)
	if !isValidPassword {
		return c.JSON(http.StatusUnauthorized, err)
	}

	token, err := auth.IssueToken()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
