package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Series struct {
	Title   string
	Year    int
	Url     string
	Country string
	Network string
	Airday  string `json:"air_day"`
	Airtime string `json:"air_time"`
	Ended   string
}

var Endpoint string = "http://api.trakt.tv/search/shows.json/%s/%s/5"

// Hits Trakt.tv to find a list of series matching a query string
func FindSeries(q string) []Series {
	q = strings.Replace(q, "-", "+", -1)
	url := fmt.Sprintf(Endpoint, TraktKey, q)
	data := []Series{}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	dec.Decode(&data)
	return data
}
