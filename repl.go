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

func commandExit(c *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
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
