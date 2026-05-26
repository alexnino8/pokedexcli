package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(loc string) (RespExploreArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + loc

	dat, exists := c.cache.Get(url)
	var err error

	if !exists {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespExploreArea{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespExploreArea{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return RespExploreArea{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
		}

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespExploreArea{}, err
		}

		c.cache.Add(url, dat)
	}

	exploreResp := RespExploreArea{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespExploreArea{}, err
	}

	return exploreResp, nil
}
