package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	dat, exists := c.cache.Get(url)
	var err error

	if !exists {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
		}

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, dat)
	}

	catchResp := Pokemon{}
	err = json.Unmarshal(dat, &catchResp)
	if err != nil {
		return Pokemon{}, err
	}

	return catchResp, nil
}
