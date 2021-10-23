package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
)

func GetProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, models.User{
		ID:           1,
		UserID:       "yonks",
		Name:         "ヤンクス太郎",
		StatusID:     1,
		SNSLine:      "@yonks",
		SNSTwitter:   "@yonks",
		SNSInstagram: "@yonks",
		SNSTikTok:    "@yonks",
	})
}

func SaveProfile(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// DBに保存する処理

	return c.JSON(http.StatusOK, "")
}
