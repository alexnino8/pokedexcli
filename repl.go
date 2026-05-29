package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/alexnino8/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"explore": {
			name:        "explore",
			description: "Get a list of all the Pokemon in a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw a Pokeball at a pokemon (catch try)",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Print a specific Pokemon's details (if already caught)",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List of all the names of the Pokemon the user has caught",
			callback:    commandPokedex,
		},
	}
}

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex) < 1 {
		return errors.New("you have no Pokemon yet")
	}

	fmt.Println("Your Pokedex:")

	for name := range cfg.pokedex {
		fmt.Println("  -", name)
	}

	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemon := args[0]

	_, exists := cfg.pokedex[pokemon]

	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name:", cfg.pokedex[pokemon].Name)
	fmt.Println("Height:", cfg.pokedex[pokemon].Height)
	fmt.Println("Weight:", cfg.pokedex[pokemon].Weight)
	fmt.Println("Stats:")
	for _, stat := range cfg.pokedex[pokemon].Stats {
		digit := stat.BaseStat
		name := stat.Stat.Name
		fmt.Printf("    -%v: %v\n", name, digit)
	}
	fmt.Println("Types:")
	for _, typ := range cfg.pokedex[pokemon].Types {
		fmt.Println("    -", typ.Type.Name)
	}

	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("Pokemon name or id not provided")
	}

	_, exists := cfg.pokedex[args[0]]
	if exists {
		fmt.Println("you already caught ", args[0])
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])

	catchResp, err := cfg.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}

	catchChance := 100 - (catchResp.BaseExperience)
	if catchChance < 10 {
		catchChance = 10
	}
	roll := rand.Intn(100)

	if roll < catchChance {
		fmt.Println(catchResp.Name, "was caught!")
		cfg.pokedex[catchResp.Name] = catchResp

	} else {
		fmt.Println(catchResp.Name, "escaped!")
	}

	return nil

}

func commandExplore(cfg *config, args []string) error {
	if len(args[0]) < 1 {
		return errors.New("no location passed")
	}
	fmt.Println("Exploring ", args[0])
	exploreResp, err := cfg.pokeapiClient.ExploreArea(args[0])
	if err != nil {
		return err
	}
	for _, pokemon_encounters := range exploreResp.PokemonEncounters {
		fmt.Println(pokemon_encounters.Pokemon.Name)
	}
	return nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
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
func commandMap(cfg *config, args []string) error {
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
func commandMapb(cfg *config, args []string) error {
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
