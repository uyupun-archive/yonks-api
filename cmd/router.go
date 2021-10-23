package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uyupun/yonks-api/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Yonks API")
	})

	apiVersion := "/api/v0"
	api := e.Group(apiVersion)
	registerApiRoutes(*api)

	return e
}

func registerApiRoutes(api echo.Group) {
	api.POST("/auth/register", handler.AuthRegister)
	api.POST("/auth/login", handler.AuthLogin)
	api.GET("/friends", handler.GetFriends)
	api.POST("/frineds", handler.AddFriend)
	api.GET("/profile", handler.GetProfile)
	api.PATCH("/profile", handler.SaveProfile)
	api.GET("/notifications", handler.GetNotifications)
}
