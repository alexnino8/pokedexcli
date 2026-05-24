package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
				err := command.callback()
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
