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
				"user_id":       "@kakimoto_itc",
				"name":          "かっきー",
				"status_id":     1,
				"sns_line":      "@kakimoto_itc",
				"sns_twitter":   "@kakimoto_itc",
				"sns_instagram": "@kakimoto_itc",
				"sns_tiktok":    "@kakimoto_itc",
				"status": models.Status{
					ID:   1,
					Name: "いそがしい",
				},
			},
		},
		{
			"id":        1,
			"user_a_id": 1,
			"user_b_id": 3,
			"user": map[string]interface{}{
				"id":            3,
				"user_id":       "@h_tyokinuhata",
				"name":          "かずきち",
				"status_id":     1,
				"sns_line":      "@h_tyokinuhata",
				"sns_twitter":   "@h_tyokinuhata",
				"sns_instagram": "@h_tyokinuhata",
				"sns_tiktok":    "@h_tyokinuhata",
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
