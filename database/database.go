package database

import (
	"io"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/9d4/ampastelobot/common"
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
)

const (
	sqlResourcePath string = "/resources/sql/"
)

var DBNAME string = "database.db"
var DB *sql.DB

func Init() {
	if common.IsDevelopment() {
		DBNAME = "database-dev.db"
	}

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
	sql, _ := getSqlFromFile("create_session_table.sql")

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
	sql, _ := getSqlFromFile("create_blynk_table.sql")

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

func getSqlFromFile(filename string) (string, error) {
	_, f, _, _ := runtime.Caller(0)
	currentDir := path.Dir(f)
	rootDir := path.Join(currentDir, "../")

	sqlPath := path.Join(rootDir, sqlResourcePath, filename)

	sql, err := os.Open(sqlPath)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(sql)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
