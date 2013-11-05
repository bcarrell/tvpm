package episode

import (
	"database/sql"
	"encoding/json"
	"fmt"
	tdb "github.com/bcarrell/tvpm/db"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
)

var dbInfo tdb.DbInfo = tdb.GetDbInfo()

type Item struct {
	Title string
}

type Payload struct {
	Channel struct {
		Item []Item
	}
}

func (p *Payload) getItems() []Item {
	return p.Channel.Item
}

// Main function
//
// Responsible for gathering newznab indexers from the database and
// creating goroutines to search them for the query.
func FindEpisode(c *cli.Context) {
	args := c.Args()
	// check if everything is supplied properly
	if len(args) == 0 || len(args) > 1 {
		fmt.Println("Wrong number of arguments.\n\n")
		cli.ShowCommandHelp(c, "find")
		os.Exit(1)
	}
	jsonChannel := make(chan []Item)
	db, err := sql.Open("sqlite3", dbInfo.FullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// grab all indexers so we can hit them up
	rows, err := db.Query("SELECT Url, ApiKey FROM Indexer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	numRows := 0
	for rows.Next() {
		var url, apiKey string
		rows.Scan(&url, &apiKey)
		q := strings.Replace(args[0], "-", "+", -1)
		go searchIndexer(q, url, apiKey, jsonChannel)
		numRows += 1
	}
	output(jsonChannel, numRows)
	rows.Close()
}

// Worker function.
//
// Searches an indexer for the query and sends JSON into a channel of the
// results.
func searchIndexer(query, url, apiKey string, jsonChannel chan<- []Item) {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += "api?t=tvsearch&apikey=%s&q=%s&o=json"
	url = fmt.Sprintf(url, apiKey, query)
	data := Payload{}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	dec.Decode(&data)

	jsonChannel <- data.getItems()
}

// Aggregates the return data from all goroutines and sends to console.
// Blocks waiting for user input on which episode to send to sabnzbd.
func output(jsonChannel <-chan []Item, expected int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 5, 5, ' ', 0)
	for i := 0; i < expected; i++ {
		data := <-jsonChannel
		for index, item := range data {
			fmt.Fprintf(
				w,
				"%d\t%s\n",
				index+1,
				item.Title,
			)
		}
	}
	w.Flush()
	getChoice()
}

// Gets the user's file choice out of the list of returned episodes and sends
// the url to sabnzbd for downloading.
func getChoice() {
	fmt.Print("Enter the file number (1-10) to send to sabnzbd: ")
	var input float64
	fmt.Scanf("%f", &input)
}
