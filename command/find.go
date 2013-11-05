package command

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

func outputResults(data []Item) {
	w := tabwriter.NewWriter(os.Stdout, 0, 5, 5, ' ', 0)
	for index, item := range data[:10] {
		fmt.Fprintf(
			w,
			"%d\t%s\n",
			index+1,
			item.Title,
		)
	}
	w.Flush()
	getUserChoice()
}

func getUserChoice() {
	fmt.Print("Enter the file number to send to sabnzbd: ")
	var input float64
	fmt.Scanf("%f", &input)
}

func searchIndexer(query, url, apiKey string, jsonResults chan<- []Item) {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += "api?t=tvsearch&apikey=%s&q=%s&o=json"
	url = fmt.Sprintf(url, apiKey, query)
	data := Payload{}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	dec.Decode(&data)

	jsonResults <- data.getItems()
}

func Find(c *cli.Context) {
	args := c.Args()
	// check if everything is supplied properly
	if len(args) == 0 || len(args) > 1 {
		fmt.Println("Wrong number of arguments.\n\n")
		cli.ShowCommandHelp(c, "find")
		os.Exit(1)
	}
	jsonChan := make(chan []Item)
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
		go searchIndexer(q, url, apiKey, jsonChan)
		numRows += 1
	}
	for i := 0; i < numRows; i++ {
		json := <-jsonChan
		outputResults(json)
	}
	rows.Close()
}
