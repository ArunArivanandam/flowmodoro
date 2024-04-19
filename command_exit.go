package main

import "os"

func callbackExit(timer *Timer, args ...string) error {
	os.Exit(0)
	return nil
}
