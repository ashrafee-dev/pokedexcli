package main

func main() {
	c := config{}
	c.commands = getCommands()
	repl(&c)
}
