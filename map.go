package main

import (
	"fmt"
	"errors"
)

func commandMap(config *config) error {
	response, err := config.pokeapiClient.ListLocations(config.Next)
	if err != nil {
		return err
	}

	config.Next = &(response.Next)
	config.Previous = &(response.Previous)

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil

}

func commandMapB(config *config) error {
	if config.Previous == nil {
		return errors.New("you're on the first page")
	}

	response, err := config.pokeapiClient.ListLocations(config.Next)
	if err != nil {
		return err
	}

	config.Next = &(response.Next)
	config.Previous = &(response.Previous)

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}

