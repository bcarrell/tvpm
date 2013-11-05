package series

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
)

var (
	TraktKey string = os.Getenv("TRAKT_API_KEY")
	Endpoint string = "http://api.trakt.tv/search/shows.json/%s/%s/5"
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

func FindSeries(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 || len(args) > 1 {
		fmt.Println("Wrong number of arguments.\n\n")
		cli.ShowCommandHelp(c, "find-series")
		os.Exit(1)
	}
	seriesQ := strings.Replace(args[0], "-", "+", -1)
	url := fmt.Sprintf(Endpoint, TraktKey, seriesQ)
	data := []Series{}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	dec.Decode(&data)
	w := tabwriter.NewWriter(os.Stdout, 0, 5, 5, ' ', 0)
	for _, item := range data {
		fmt.Fprintf(
			w,
			"%s\t%s\t%d\t%s\n",
			item.Title,
			item.Network,
			item.Year,
			item.Url,
		)
	}
	w.Flush()
}
