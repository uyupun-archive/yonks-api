package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	con, err := connectDB()
	if err != nil {
		panic(err)
	}
	db, err := con.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// マイグレーション
}

func connectDB() (*gorm.DB, error) {
	dsn, err := getDSN()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDSN() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_EXPOSE_PORT")
	name := os.Getenv("DB_NAME")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	parseTime := "true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s", userName, password, host, port, name, parseTime)

	return dsn, nil
}
