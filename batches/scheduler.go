package main

import (
	"time"
)

func main() {
	scheduler()
}

func scheduler() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		execMatcher()
	}
}

func execMatcher() {
	matcher()
}
