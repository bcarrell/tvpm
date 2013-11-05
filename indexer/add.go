package indexer

import (
	"database/sql"
	"fmt"
	tdb "github.com/bcarrell/tvpm/db"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var (
	dbInfo tdb.DbInfo = tdb.GetDbInfo()
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
		cli.ShowCommandHelp(c, "add-indexer")
		os.Exit(1)
	}
	url := args[0]
	db, err := sql.Open("sqlite3", dbInfo.FullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sql := "INSERT INTO Indexer(Url, ApiKey) values('%s', '%s')"
	sql = fmt.Sprintf(sql, url, key)

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added %s with key %s to your indexers!\n", url, key)
}
