package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocationAreas fetches a list of location areas from the PokeAPI
// it attaches to our customer Client struct using (c *Client)
func (c *Client) ListLocationAreas(pageURL *string) (RespLocationAreas, error) {
	url := "https://pokeapi.co/api/v2/location-area/"

	if pageURL != nil {
		url = *pageURL
	}
	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, err
	}

	// send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, err
	}

	// ensure connection colses at the end
	defer resp.Body.Close()

	// check for server errors
	if resp.StatusCode > 299 {
		return RespLocationAreas{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
	}

	// read all the raw bytes of the response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}

	// unmarshal the raw JSON bytes into our GO struct
	locationsResp := RespLocationAreas{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocationAreas{}, err
	}

	return locationsResp, nil

}
