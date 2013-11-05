package command

import (
	"fmt"
	"github.com/bcarrell/tvpm/db"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func AddIndexer(c *cli.Context) {
	args := c.Args()
	key := c.String("apiKey")
	// check if everything is supplied properly
	if len(args) == 0 || len(args) > 1 {
		fmt.Println("Wrong number of arguments.\n\n")
		cli.ShowCommandHelp(c, "add-indexer")
		os.Exit(1)
	}
	if len(key) == 0 {
		fmt.Println("Missing an API key.")
		cli.ShowCommandHelp(c, "find-series")
		os.Exit(1)
	}
	url := args[0]
	// add to indexer db
	err := db.InsertIndexer(url, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added %s with key %s to your indexers!\n", url, key)
}
