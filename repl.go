package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var config config

	for {
		fmt.Print("Pokedex > : ")

		ok := scanner.Scan() 
		if !ok {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %v\n", err)
			}
			break
		}

		input := scanner.Text()
		cleanedInput := cleanInput(input)

		//checking if cleanedInput is empty
		if len(cleanedInput) == 0 {
			fmt.Println("no user input")
		} else {
			if cmd, exists := getCommands()[cleanedInput[0]]; exists {
				err := cmd.callback(&config)
				if err != nil {
					fmt.Printf("Error executing command: %v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next		string
	Previous	string
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand{
    return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: 		"map",
			description: "displays a list of locations",
			callback:	commandMap,
		},

		"mapb": {
			name: 		"mapb",
			description: "displays previous list of locations",
			callback:	commandMapB,
		},
	}
}