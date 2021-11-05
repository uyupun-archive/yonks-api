package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/uyupun/yonks-api/utility/database"
)

type MigrateType string

const (
	MigrateTypeUp   MigrateType = "up"
	MigrateTypeDown MigrateType = "down"
)

func main() {
	execCmd()
}

func execCmd() {
	args := os.Args
	if len(args) < 2 {
		panic("Too few arguments")
	}

	command := args[1]
	if command == string(MigrateTypeUp) {
		migrate(MigrateTypeUp)
	} else if command == string(MigrateTypeDown) {
		migrate(MigrateTypeDown)
	} else {
		panic("Command not match")
	}
}

func migrate(migrateType MigrateType) {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	for idx, target := range registerMigrationTargets() {
		msg := "Applied"
		if migrateType == MigrateTypeUp {
			db.AutoMigrate(target)
		} else if migrateType == MigrateTypeDown {
			db.Migrator().DropTable(target)
			msg = "Rollbacked"
		} else {
			return
		}

		modelName := strings.Split(reflect.TypeOf(target).String(), ".")[1]
		fmt.Printf("%d: %s %s migration!\n", idx+1, msg, modelName)
	}
}
