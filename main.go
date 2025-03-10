package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Hello, World!")

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
