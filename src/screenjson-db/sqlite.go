package main

import (
	"database/sql"
	"fmt"

	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
)

func insert_into_sqlite(uri, database, table, field, json_data string, additional_data map[string]string) error {
	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		color.Red("Failed to connect to SQLite: %s", err)
		return err
	}
	defer db.Close()

	columns := []string{field}         // JSON field
	placeholders := []string{"?"}      // SQLite uses ? for placeholders
	values := []interface{}{json_data} // Starting with JSON data

	// SQLite specific: Constructing additional columns and placeholders
	for col, val := range additional_data {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		values = append(values, val)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, placeholders)
	result, err := db.Exec(query, values...)
	if err != nil {
		color.Red("Failed to insert data into SQLite: %s", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		color.Red("Failed to retrieve last inserted ID: %s", err)
		return err
	}

	color.Green("Data successfully inserted into SQLite with ID: %d", id)
	return nil
}
