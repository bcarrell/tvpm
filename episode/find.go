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
	Link  string
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
	data := make([]Item, 0)
	for i := 0; i < expected; i++ {
		resultFromIndexer := <-jsonChannel
		for _, item := range resultFromIndexer {
			data = append(data, item)
		}
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 5, 5, ' ', 0)
	for index, item := range data {
		fmt.Fprintf(
			w,
			"%d\t»\t%s\n",
			index+1,
			item.Title,
		)
	}
	w.Flush()
	getChoice(data)
}

// Gets the user's file choice out of the list of returned episodes and sends
// the url to sabnzbd for downloading.
func getChoice(data []Item) {
	fmt.Printf("Enter the file number (1-%d) to send to sabnzbd: ", len(data))
	var input int
	fmt.Scanf("%d", &input)
	sendToSabnzbd(data[input-1])
}

// Sends the specified file to be downloaded by Sabnzbd
func sendToSabnzbd(item Item) {
	sabnzbdUrl := os.Getenv("SABNZBD_URL")
	sabnzbdKey := os.Getenv("SABNZBD_API_KEY")
	link := strings.Replace(item.Link, "&", "%26", -1)
	title := item.Title

	// construct url
	url := sabnzbdUrl + "api?mode=addurl&name=" + link + "&nzbname=" + title +
		"&apikey=" + sabnzbdKey

	// off we go!
	http.Get(url)
}
