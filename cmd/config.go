package main

import "github.com/joho/godotenv"

func loadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
