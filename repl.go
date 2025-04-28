package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"exit": {
			name: 		"exit",
			description: "Exit the pokedex",
			callback:	 commandExit,
		},
	}

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
			switch cleanedInput[0] {
			case "exit":
				if cmd, exists := commands["exit"]; exists {
					err := cmd.callback()
					if err != nil {
						fmt.Printf("Error executing command %s: %v\n", cmd.name, err)
					}
				}
			case "help":
				if cmd, exists := commands["help"]; exists {
					err := cmd.callback()
					if err != nil {
						fmt.Printf("Error executing command %s: %v\n", cmd.name, err)
					}
				}
			default:
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

func (c *map[string]cliCommand) commandHelp() error {
	fmt.Println("Welcome to Pokedex!\nUsage:\n")
	for command, _ := range c {
		fmt.Printf("%s: %s\n", c[command].name, c[command].description)
	}
	return nil
}


func commandExit() error {
	fmt.Println("Closeing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}