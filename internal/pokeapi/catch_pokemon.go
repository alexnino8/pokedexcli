package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(name string) (RespCatchPokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	dat, exists := c.cache.Get(url)
	var err error

	if !exists {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespCatchPokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespCatchPokemon{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return RespCatchPokemon{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
		}

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespCatchPokemon{}, err
		}

		c.cache.Add(url, dat)
	}

	catchResp := RespCatchPokemon{}
	err = json.Unmarshal(dat, &catchResp)
	if err != nil {
		return RespCatchPokemon{}, err
	}

	return catchResp, nil
}
