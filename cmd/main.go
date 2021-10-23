package main

import "os"

func main() {
	router := newRouter()
	router.Logger.Fatal(router.Start(":" + os.Getenv("APP_PORT")))
}
