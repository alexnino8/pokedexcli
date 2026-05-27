# Pokedex CLI

A command-line Pokedex application built in Go.

## Current Progress

The CLI supports interactive exploration of the Pokémon world via the PokéAPI,
backed by a custom in-memory cache for fast repeat queries. You can now also
attempt to catch Pokémon and add them to your Pokedex.

## Supported Commands

- `help` — Prints a description of how to use the Pokedex and lists available commands.
- `exit` — Gracefully shuts down the Pokedex.
- `map` — Displays the next 20 location areas in the Pokémon world.
- `mapb` — Displays the previous 20 location areas (back-pagination).
- `explore <area-name>` — Lists all Pokémon that can be found in the given location area.
- `catch <pokemon-name>` — Attempts to catch a Pokémon. Success is based on the Pokémon's base experience — the higher it is, the harder the catch.

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
```

## Architecture Highlights

- REPL loop with a command registry that dispatches to handlers and supports arguments.
- PokéAPI client in `internal/pokeapi` for HTTP requests and JSON parsing.
- Custom cache (`internal/pokecache`) with TTL-based eviction via a background reaper goroutine — re-running a query hits the cache instead of the network.
- Caught Pokémon are tracked in a `map[string]Pokemon` on the config struct.

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

## Future Features

- [ ] **Inspect**: View stats, types, height, and weight of caught Pokémon.
- [ ] **Pokedex**: List all the Pokémon you've caught.
