package main

import "github.com/uyupun/yonks-api/models"

func registerMigrationTargets() []interface{} {
	return []interface{}{
		&models.User{},
	}
}
