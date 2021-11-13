package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
	"github.com/uyupun/yonks-api/utility/database"
	"golang.org/x/crypto/bcrypt"
)

func AuthRegister(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// パスワードをbcryptでハッシュ化
	password := []byte(user.Password)
	hashed, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	user.Password = string(hashed)

	// ユーザの作成
	err = database.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	// トークンの生成
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Taro"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
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

	// ユーザの取得
	user := models.User{
		UserID: loginInfo.UserID,
	}
	err = database.FindUserByUserID(&user)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	// パスワードの検証
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	// トークンの生成
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Taro"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func AuthLogout(c echo.Context) error {
	return nil
}
