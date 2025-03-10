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
	callback    func() error
}

var commands map[string]cliCommand

func main() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    printHelpMessage,
		},
	}
	for {
		// read user input
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		for input.Scan() {
			word := input.Text()
			cleanedWord := cleanInput(word)
			command := cleanedWord[0]

			//check if command is in commands map
			cmd, ok := commands[command]
			if ok {
				err := cmd.callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
			fmt.Print("Pokedex > ")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	// trim input string
	text = strings.TrimSpace(text)
	// convert to lowercase
	text = strings.ToLower(text)
	// replace all double spaces into single space
	text = strings.Replace(text, "  ", " ", -1)
	// split by space
	separated := strings.Split(text, " ")
	return separated
}

func printHelpMessage() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
