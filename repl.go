package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommands struct {
	name        string
	description string
	callback    func(*config) error
}
type config struct {
	commands map[string]cliCommands
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func repl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		commandName := strings.ToLower(strings.Fields(scanner.Text())[0])
		command, exists := c.commands[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(c)
		if err != nil {
			fmt.Print("Unknown command")
		}

	}

}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"exit": {
			name:        "exit",
			description: "Closing the Pokedex... Goodbye!",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex",
			callback:    commandHelp,
		},
	}

}
