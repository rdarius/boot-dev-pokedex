package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	quit := false

	for !quit {

		// read user input
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		for input.Scan() {
			word := input.Text()
			cleanedWord := cleanInput(word)
			fmt.Printf("Your command was: %s\n", cleanedWord[0])

			if cleanedWord[0] == "quit" {
				quit = true
			}
			break
		}

	}

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
