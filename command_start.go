package main

import (
	"fmt"
	"time"
)

func callbackStart(timer *Timer, args ...string) error {
	timer.startTime = time.Now()
	fmt.Println("Timer started.")
	return nil
}
