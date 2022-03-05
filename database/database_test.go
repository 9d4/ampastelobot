package database

import (
	"testing"
)

func TestGetSqlFromFile(t *testing.T) {
	file := "create_blynk_table.sql"

	_, err := getSqlFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
}
