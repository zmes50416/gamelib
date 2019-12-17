package kg

import "net/http"

import "net/url"

import "fmt"

import "io/ioutil"

var APIKey = "Place your api key here"
var api string = "https://kgsearch.googleapis.com/v1/entities:search"

// SearchGame find game info for query.
func SearchGame(query string) ([]byte, error) {
	u, _ := url.Parse(api)
	u.Query().Add("query", query)
	u.Query().Add("types", "VideoGame")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("status is not ok")
	}

	return data, nil
}
