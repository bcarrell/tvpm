package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func AllTablesExist() bool {
	db, err := sql.Open("sqlite3", dbInfo.FullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// TODO: Optimize?
	q :=
		`SELECT EXISTS 
		 (SELECT * FROM sqlite_master 
		  WHERE type='table' AND name='%s');`
	for _, table := range Tables {
		var exists int
		err := db.QueryRow(fmt.Sprintf(q, table)).Scan(&exists)
		if err == sql.ErrNoRows || err != nil || exists == 0 {
			return false
		}
	}
	return true
}
