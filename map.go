package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func commandMap(config *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	defer res.Body.Close()

	if err != nil {
		return fmt.Errorf("error getting url response: %w\n", err)
	}

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and body %s\n", res.StatusCode, body)
	}

	if err != nil {
		return fmt.Errorf("error reading res.Body: %w\n", err)
	}

	var response LocationAreaResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("error unmarshaling body: %w\n", err)
	}

	config.Next = response.Next
	if respons
	config.Previous = response.Previous

	fmt.Println(config.Next)
	fmt.Println(config.Previous)

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil

}

func commandMapB(config *config) error {
	if config.Previous == "" {
		fmt.Println("No map command previously selected")
		return nil
	}

	res, err := http.Get(config.Previous)
	defer res.Body.Close()

	if err != nil {
		return fmt.Errorf("error getting url response: %w\n", err)
	}

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and body %s\n", res.StatusCode, body)
	}

	if err != nil {
		return fmt.Errorf("error reading res.Body: %w\n", err)
	}

	var response LocationAreaResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("error unmarshaling body: %w\n", err)
	}

	config.Next = response.Next
	config.Previous = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}

type LocationAreaResponse struct {
    Count    int    `json:"count"`
    Next     string `json:"next"`
    Previous string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

