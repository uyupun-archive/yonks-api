package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uyupun/yonks-api/handlers"
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
	api.POST("/auth/register", handlers.AuthRegister)
	api.POST("/auth/login", handlers.AuthLogin)
	api.POST("/auth/logout", handlers.AuthLogout)
	api.GET("/friends", handlers.GetFriends)
	api.POST("/frineds", handlers.AddFriend)
	api.GET("/profile", handlers.GetProfile)
	api.PATCH("/profile", handlers.SaveProfile)
	api.GET("/notifications", handlers.GetNotifications)
}
