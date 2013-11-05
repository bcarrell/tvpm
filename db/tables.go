package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreateTables() {
	statuses := make(chan bool)
	createSeriesSql :=
		`CREATE TABLE IF NOT EXISTS Series
		(Id INTEGER NOT NULL PRIMARY KEY,
		 SeriesName TEXT NOT NULL);`
	createIndexerSql :=
		`CREATE TABLE IF NOT EXISTS Indexer
		(Id INTEGER NOT NULL PRIMARY KEY,
		 Url TEXT NOT NULL UNIQUE, 
		 ApiKey TEXT NOT NULL UNIQUE);`
	createQualitySql := `CREATE TABLE IF NOT EXISTS Quality
		(Id INTEGER NOT NULL PRIMARY KEY,
		 Quality TEXT NOT NULL);`

	go createTable(statuses, createSeriesSql)
	go createTable(statuses, createIndexerSql)
	go createQualityTable(statuses, createQualitySql)

	// block for all table creation goroutines
	for i := 0; i < 3; i++ {
		<-statuses
	}

	createEpisodeSql :=
		`CREATE TABLE IF NOT EXISTS Episode
		(Id INTEGER NOT NULL PRIMARY KEY,
		 SeriesId INTEGER NOT NULL,
		 Title TEXT NOT NULL,
		 Quality INTEGER NOT NULL,
		 FOREIGN KEY(SeriesId) REFERENCES Series(Id),
		 FOREIGN KEY(Quality) REFERENCES Quality(Id));`

	// Needed to wait to create these tables because
	// of foreign keys
	go createTable(statuses, createEpisodeSql)
	for i := 0; i < 1; i++ {
		<-statuses
	}
}

func createTable(done chan bool, sqlCommand string) {
	db, err := sql.Open("sqlite3", dbInfo.FullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(sqlCommand)
	if err != nil {
		log.Fatal(err)
	}
	done <- true
}

func createQualityTable(done chan bool, sqlCommand string) {
	db, err := sql.Open("sqlite3", dbInfo.FullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(sqlCommand)
	if err != nil {
		log.Fatal(err)
	}
	transaction, err := db.Begin()
	for _, qualityValue := range QualityChoices {
		insertion := fmt.Sprintf("INSERT INTO Quality(Quality) VALUES('%s')",
			qualityValue)
		_, err := transaction.Prepare(insertion)
		if err != nil {
			log.Fatal(err)
		}
	}
	transaction.Commit()
	done <- true
}
