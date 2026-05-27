package pokeapi

// RespLocationAreas represents the JSON response from the location-area endpoint
type RespLocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// these are for the explore command which also calls the location-area endpoint

type RespExploreArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// for the catch command -- calls the pokemon endpoint
type RespCatchPokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}
