package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexnino8/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil

}

// commandMap handles moving Forward in the pokemon world
func commandMap(cfg *config) error {
	// call our engine, passing the next URL -- which is nil the first time
	locationsResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationsURL)
	// print testing
	// fmt.Println("next: ", cfg.nextLocationsURL)
	// fmt.Println("prev: ", cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	// update next and prev urls
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	//  fmt.Println("next: ", cfg.nextLocationsURL)
	// fmt.Println("prev: ", cfg.prevLocationsURL)

	// loop through the results and print the names
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// commandMapb handles moving Backwards in the pokemon world
func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	// call our engine passing the prev url
	locationsResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// fmt.Println("next: ", cfg.nextLocationsURL)
	// fmt.Println("prev: ", cfg.prevLocationsURL)

	// update urls
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// fmt.Println("next: ", cfg.nextLocationsURL)
	// fmt.Println("prev: ", cfg.prevLocationsURL)

	// loop and print
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}
