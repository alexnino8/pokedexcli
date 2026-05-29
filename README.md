# Pokedex CLI

A command-line Pokedex application built in Go.

## Current Progress

The CLI supports interactive exploration of the Pokémon world via the PokéAPI, backed by a custom in-memory cache for fast repeat queries. You can catch Pokémon, add them to your Pokedex, and inspect their details.

## Supported Commands

- `help` — Prints a description of how to use the Pokedex and lists available commands.
- `exit` — Gracefully shuts down the Pokedex.
- `map` — Displays the next 20 location areas in the Pokémon world.
- `mapb` — Displays the previous 20 location areas (back-pagination).
- `explore <area-name>` — Lists all Pokémon that can be found in the given location area.
- `catch <pokemon-name>` — Attempts to catch a Pokémon. Success is based on the Pokémon's base experience — the higher it is, the harder the catch.
- `inspect <pokemon-name>` — Displays the name, height, weight, stats, and types of a caught Pokémon.
- `pokedex` — Lists the names of all Pokémon you have successfully caught.

## Example

```bash
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados
 - remoraid
 - octillery
 - wingull
 - pelipper
 - shellos
 - gastrodon

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu escaped!
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > pokedex
Your Pokedex:
 - pikachu

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  -hp: 35
  -attack: 55
  -defense: 40
  -special-attack: 50
  -special-defense: 50
  -speed: 90
Types:
  - electric

## Architecture Highlights

- REPL Loop: A robust command registry that dispatches to handlers and supports dynamic arguments.
- PokéAPI Client: A dedicated internal package for managing HTTP requests and JSON decoding.
- Custom Cache: An internal package featuring a thread-safe cache with TTL-based eviction via a background reaper goroutine to optimize network performance.
- State Management: Uses a shared configuration struct to track the Pokedex (a map[string]Pokemon) and pagination state.

## Installation

Ensure you have Go installed on your system.

Clone the repo:

```bash
git clone https://github.com/yourusername/pokedexcli.git
```

Run the application:

```bash
go run .
```

