package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	migrate()
}

func migrate() {
	args := os.Args
	if len(args) < 2 {
		panic("Too few arguments")
	}

	command := args[1]
	if command == "up" {
		err := migrateUp()
		if err != nil {
			panic(err)
		}
	} else if command == "down" {
		err := migrateDown()
		if err != nil {
			panic(err)
		}
	} else {
		panic("Command not match")
	}
}

func migrateUp() error {
	db, err := connectDB()
	if err != nil {
		return err
	}

	for idx, target := range registerMigrationTargets() {
		db.AutoMigrate(target)
		modelName := strings.Split(reflect.TypeOf(target).String(), ".")[1]
		fmt.Printf("%d: Applied %s migration!\n", idx+1, modelName)
	}
	return nil
}

func migrateDown() error {
	db, err := connectDB()
	if err != nil {
		return err
	}

	for idx, target := range registerMigrationTargets() {
		db.Migrator().DropTable(target)
		modelName := strings.Split(reflect.TypeOf(target).String(), ".")[1]
		fmt.Printf("%d: Rollbacked %s migration!\n", idx+1, modelName)
	}
	return nil
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
