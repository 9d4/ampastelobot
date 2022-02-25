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

	}

	// create tables after init finish
	defer func() {
		CreateSessionTable()
		CreateBlynkTokenTable()
	}()

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

func CreateSessionTable() bool {
	// make session table
	sql := `CREATE TABLE IF NOT EXISTS session (
id INTEGER PRIMARY KEY AUTOINCREMENT,
user_id INTEGER NOT NULL,
session_id TEXT NOT NULL,
data TEXT,
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
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

func CreateBlynkTokenTable() bool {
	sql := `CREATE TABLE IF NOT EXISTS blynk_tokens (
id INTEGER PRIMARY KEY AUTOINCREMENT,
user_id INTEGER NOT NULL,
token TEXT,
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Panicln("Can't create blynk_tokens table")
	}

	if _, err := stmt.Exec(); err == nil {
		return true
	}

	return false
}

func dbFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}
