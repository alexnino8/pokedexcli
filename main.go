package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/alexnino8/pokedexcli/internal/pokeapi"
)

func main() {
	// create the API client with a 5 second timeout
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// load client into our app's brain
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			clean := cleanInput(scanner.Text())

			if len(clean) == 0 {
				continue
			}

			commandName := clean[0]

			availableCommands := getCommands()

			command, exists := availableCommands[commandName]

			if exists {
				var err error
				if len(clean) > 1 {
					err = command.callback(cfg, clean[1])
				} else {
					err = command.callback(cfg, "")
				}

				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Println("Unknown command")
			}

		} else {
			break
		}
	}

	// the loop ended so we check why
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
	}

}
