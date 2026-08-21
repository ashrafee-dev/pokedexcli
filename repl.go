package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	type cliCommands struct {
		name        string
		description string
		callback    func() error
	}

	commands := map[string]cliCommands{
		"exit": {
			name:        "exit",
			description: "Closing the Pokedex... Goodbye!",
			callback:    commandExit,
		},
		"help": {
			name:        "exit",
			description: "Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex",
			callback:    commandHelp,
		},
	}

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		commandName := strings.ToLower(strings.Fields(scanner.Text())[0])
		command, exists := commands[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Print("Unknown command")
		}

	}

}
