package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
)

func GetProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":            1,
		"user_id":       "yonks",
		"name":          "ヤンクス太郎",
		"status_id":     1,
		"sns_line":      "@yonks",
		"sns_twitter":   "@yonks",
		"sns_instagram": "@yonks",
		"sns_tiktok":    "@yonks",
		"status": models.Status{
			ID:   1,
			Name: "人肌恋しい",
		},
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
