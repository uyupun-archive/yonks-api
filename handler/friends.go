package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
)

func GetFriends(c echo.Context) error {
	return c.JSON(http.StatusOK, []map[string]interface{}{
		{
			"id":        1,
			"user_a_id": 1,
			"user_b_id": 2,
			"user": map[string]interface{}{
				"id":            2,
				"user_id":       "@yonks2",
				"name":          "ヤンクス次郎",
				"status_id":     1,
				"sns_line":      "@yonks2",
				"sns_twitter":   "@yonks2",
				"sns_instagram": "@yonks2",
				"sns_tiktok":    "@yonks2",
				"status": models.Status{
					ID:   1,
					Name: "人肌恋しい",
				},
			},
		},
		{
			"id":        1,
			"user_a_id": 1,
			"user_b_id": 2,
			"user": map[string]interface{}{
				"id":            2,
				"user_id":       "@yonks2",
				"name":          "ヤンクス次郎",
				"status_id":     1,
				"sns_line":      "@yonks2",
				"sns_twitter":   "@yonks2",
				"sns_instagram": "@yonks2",
				"sns_tiktok":    "@yonks2",
				"status": models.Status{
					ID:   1,
					Name: "人肌恋しい",
				},
			},
		},
	})
}

func AddFriend(c echo.Context) error {
	friend := new(models.Friend)
	err := c.Bind(&friend)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// DBに保存する処理

	return c.JSON(http.StatusOK, "")
}
