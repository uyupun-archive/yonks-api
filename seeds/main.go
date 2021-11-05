package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/uyupun/yonks-api/utility/database"
)

type SeedType string

const (
	SeedTypeInit SeedType = "init"
	SeedTypeMock SeedType = "mock"
)

func main() {
	seed(SeedTypeInit)
	seed(SeedTypeMock)
}

func seed(seedType SeedType) {
	targets := registerInitTargets()
	if seedType == SeedTypeMock {
		targets = registerMockTargets()
	}

	for idx, target := range targets {
		bytes, err := ioutil.ReadFile("seeds/" + string(seedType) + "/" + target.Seed)
		if err != nil {
			panic(err)
		}

		var records []map[string]interface{}
		err = json.Unmarshal(bytes, &records)
		if err != nil {
			panic(err)
		}

		db, err := database.ConnectDB()
		if err != nil {
			panic(err)
		}

		db.Model(&target.Model).Create(records)

		fmt.Printf("%d: Applied %s seed!\n", idx+1, target.Seed)
	}
}
