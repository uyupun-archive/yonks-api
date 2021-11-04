package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/uyupun/yonks-api/utility"
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
		err := migrate(MigrateTypeUp)
		if err != nil {
			panic(err)
		}
	} else if command == string(MigrateTypeDown) {
		err := migrate(MigrateTypeDown)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Command not match")
	}
}

func migrate(migrateType MigrateType) error {
	db, err := utility.ConnectDB()
	if err != nil {
		return err
	}

	for idx, target := range registerMigrationTargets() {
		if migrateType == MigrateTypeUp {
			db.AutoMigrate(target)
		} else if migrateType == MigrateTypeDown {
			db.Migrator().DropTable(target)
		} else {
			return nil
		}
		modelName := strings.Split(reflect.TypeOf(target).String(), ".")[1]
		fmt.Printf("%d: Applied %s migration!\n", idx+1, modelName)
	}
	return nil
}
