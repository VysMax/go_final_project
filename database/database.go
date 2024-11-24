package database

import (
	"database/sql"
	"log"
	"os"
)

// }

func ConnectDB(dbFilePath string) *sql.DB {
	db, err := sql.Open("sqlite", dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateDB(dbName string) *sql.DB {
	newDB, err := os.Create(dbName)
	if err != nil {
		log.Fatal(err)
	}
	newDB.Close()

	db := ConnectDB(dbName)
	table, err := db.Prepare(`CREATE TABLE IF NOT EXISTS scheduler (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date CHAR(8) NOT NULL DEFAULT "",
			title VARCHAR(128) NOT NULL DEFAULT "",
			comment VARCHAR(256) NOT NULL DEFAULT "",
			repeat VARCHAR(128) NOT NULL DEFAULT "")`)
	if err != nil {
		log.Fatal(err)
	}
	table.Exec()

	table, err = db.Prepare("CREATE INDEX scheduler_date ON scheduler (date)")
	if err != nil {
		log.Fatal(err)
	}
	table.Exec()

	return db
}
