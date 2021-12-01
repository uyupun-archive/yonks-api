package main

import (
	"fmt"
	"time"
)

func main() {
	scheduler()
}

func scheduler() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			matcher()
		}
	}
}

func matcher() {
	fmt.Println("hello, world!")
}
