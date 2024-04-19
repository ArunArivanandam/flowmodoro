package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Timer, ...string) error
}

func startRepl(timer *Timer) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Printf("Invalid command")
			continue
		}
		err := command.callback(timer, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display available commands",
			callback:    callbackHelp,
		},
		"start": {
			name:        "start",
			description: "Starts the timer",
			callback:    callbackStart,
		},
		"stop": {
			name:        "stop",
			description: "Stops the timer",
			callback:    callbackStop,
		},
		"exit": {
			name:        "exit",
			description: "Exits the timer",
			callback:    callbackExit,
		},
	}
}

func cleanInput(text string) []string {

	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
