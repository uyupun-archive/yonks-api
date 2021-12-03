package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/uyupun/yonks-api/utils/database"
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
	if command != string(MigrateTypeUp) && command != string(MigrateTypeDown) {
		panic("Command not match")
	}
	migrate(MigrateType(command))
}

func migrate(migrateType MigrateType) {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	cnt := 0
	msg := "Applied"
	for _, target := range registerMigrationTargets() {
		isShow := false
		switch migrateType {
		case MigrateTypeUp:
			if db.Migrator().HasTable(target) {
				break
			}
			isShow = true
			db.AutoMigrate(target)
		case MigrateTypeDown:
			msg = "Rollbacked"
			if !db.Migrator().HasTable(target) {
				break
			}
			isShow = true
			db.Migrator().DropTable(target)
		}

		if isShow {
			cnt++
			modelName := strings.Split(reflect.TypeOf(target).String(), ".")[1]
			fmt.Printf("%d: %s %s migration!\n", cnt, msg, modelName)
		}
	}
	if cnt == 0 {
		fmt.Printf("No migration to \"%s\"\n", msg)
	}
}
