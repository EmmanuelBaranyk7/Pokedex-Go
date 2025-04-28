package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var body []byte
	var err error

	cachedData, found := c.cache.Get(url)
    if found {
        body = cachedData
	} else {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		c.cache.Add(url, body)
	}

	locationsResp := LocationAreaResponse{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationsResp, nil
}
