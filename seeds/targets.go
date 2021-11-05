package main

import "github.com/uyupun/yonks-api/models"

type Target struct {
	Model interface{}
	Seed  string
}

func registerInitTargets() []Target {
	return []Target{
		{
			Model: models.Status{},
			Seed:  "statuses.json",
		},
	}
}

func registerMockTargets() []Target {
	return []Target{
		{
			Model: []models.User{},
			Seed:  "users.json",
		},
	}
}
