package main

import "os"

func main() {
	router := newRouter()
	loadConfig()
	router.Logger.Fatal(router.Start(":" + os.Getenv("APP_PORT")))
}
