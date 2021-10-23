package handler

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
			Content: "test",
		},
		{
			ID:      2,
			UserID:  1,
			Content: "test",
		},
	})
}
