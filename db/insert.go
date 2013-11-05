package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InsertIndexer(url, apiKey string) error {
	db, err := sql.Open("sqlite3", dbInfo.FullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sql := "INSERT INTO Indexer(Url, ApiKey) values('%s', '%s')"
	sql = fmt.Sprintf(sql, url, apiKey)

	_, err = db.Exec(sql)
	return err
}
