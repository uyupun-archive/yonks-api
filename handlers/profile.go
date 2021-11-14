package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
	"github.com/uyupun/yonks-api/utility/database"
)

func GetProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":            1,
		"user_id":       "takashi0602",
		"name":          "高橋たかし",
		"status_id":     3,
		"sns_line":      "@takashi0602",
		"sns_twitter":   "@takashi0602",
		"sns_instagram": "@takashi0602",
		"sns_tiktok":    "@takashi0602",
		"status": models.Status{
			ID:   3,
			Name: "人肌恋しい",
		},
	})
}

func SaveProfile(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: claimsと比較して正当なユーザからのリクエストかどうかを見る

	err = database.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	return c.JSON(http.StatusOK, nil)
}
