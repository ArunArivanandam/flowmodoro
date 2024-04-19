package main

import (
	"fmt"
	"os"
	"time"
)

func callbackStop(timer *Timer, args ...string) error {
	elapsedTime := time.Since(timer.startTime)

	hours := int(elapsedTime.Hours())
	minutes := int(elapsedTime.Minutes()) % 60
	seconds := int(elapsedTime.Seconds()) % 60

	fmt.Printf("Elapsed time: %dh:%dm:%ds\n", hours, minutes, seconds)
	fmt.Println("Get a break, matey!\n   ( ͡° ʖ ͡°)")
	fmt.Print("Go out and don't stare the screen!! \n")

	os.Exit(0)
	return nil
}
