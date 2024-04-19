package main

import (
	"time"
)

type Timer struct {
	startTime time.Time
}

func main() {
	timer := Timer{}

	startRepl(&timer)
}
