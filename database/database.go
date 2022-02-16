package database

import (
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"github.com/traperwaze/ampastelobot/common"

	"database/sql"
)

const DBNAME string = "session.db"

var DB *sql.DB

func Init() {
	dsn := path.Join(common.Wd(), DBNAME)

	if !dbFileExists(dsn) {
		log.Println("[database] DB File not found, will created automatically")

		// create table after init finish
		defer CreateSessionTable()
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal("[database] Cannot connect to database")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("[database] Database connection failed")
	}

	DB = db
	log.Println("[database] Init Done")
}

func dbFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func CreateSessionTable() bool {
	// make session table
	sql := `CREATE TABLE IF NOT EXISTS session (
id INTEGER PRIMARY KEY AUTOINCREMENT,
user_id INTEGER NOT NULL,
session_id TEXT NOT NULL,
data TEXT,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL
);`

	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Println("Can't create sesison table")
	}

	if _, err := stmt.Exec(); err == nil {
		return true
	}

	return false
}
