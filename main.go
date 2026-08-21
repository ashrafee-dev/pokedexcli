package main

func main() {
	c := config{}
	c.commands = make(map[string]cliCommands)
	c.commands["exit"] = cliCommands{
		name:        "exit",
		description: "Closing the Pokedex... Goodbye!",
		callback:    commandExit,
	}
	c.commands["help"] = cliCommands{
		name:        "help",
		description: "Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex",
		callback:    commandHelp,
	}
	repl(&c)
}
