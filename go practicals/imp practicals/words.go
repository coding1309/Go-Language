package main

import (
	"fmt"
	"strings"
)

func main() {
	dictionary := make(map[string]string)

	for {
		fmt.Println("1. Look up a word")
		fmt.Println("2. Add a word")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			lookupWord(dictionary)
		case 2:
			addWord(dictionary)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}

		fmt.Println()
	}
}

func lookupWord(dictionary map[string]string) {
	var word string
	fmt.Print("Enter the word to look up: ")
	fmt.Scanln(&word)

	definition, found := dictionary[strings.ToLower(word)]
	if found {
		fmt.Printf("Definition of %s: %s\n", word, definition)
	} else {
		fmt.Printf("'%s' not found in the dictionary.\n", word)
	}
}

func addWord(dictionary map[string]string) {
	var word, definition string
	fmt.Print("Enter the word to add: ")
	fmt.Scanln(&word)

	fmt.Print("Enter the definition: ")
	fmt.Scanln(&definition)

	dictionary[strings.ToLower(word)] = definition
	fmt.Printf("'%s' added to the dictionary.\n", word)
}
