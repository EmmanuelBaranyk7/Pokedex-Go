package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

var commands map[string]cliCommand

//prevents initilization cycle
func init() {
    commands = map[string]cliCommand{
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
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

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
			if cmd, exists := commands[cleanedInput[0]]; exists {
				err := cmd.callback()
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
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for command, _ := range commands {
		fmt.Printf("%s: %s\n", commands[command].name, commands[command].description)
	}
	return nil
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}