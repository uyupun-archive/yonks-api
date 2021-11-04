package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/yonks-api/models"
)

func GetNotifications(c echo.Context) error {
	return c.JSON(http.StatusOK, []models.Notification{
		{
			ID:      1,
			UserID:  1,
			Content: "かずきちさんとマッチングしました！\nかずきちさんの連絡先:\nLINE: @h_tyokinuhata\nTwitter: @h_tyokinuhata\nInstagram: @h_tyokinuhata\nTikTok: @h_tyokinuhata",
		},
	})
}
