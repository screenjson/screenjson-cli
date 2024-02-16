package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

func insert_into_mysql(uri, database, table, field, json_data string, additional_data map[string]string) error {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", uri)
	if err != nil {
		color.Red("Failed to connect to MySQL: %s", err)
		return err
	}
	defer db.Close()

	// Check the connection
	if err := db.Ping(); err != nil {
		color.Red("Failed to ping MySQL: %s", err)
		return err
	}

	// Start building the INSERT statement
	columns := []string{field}         // JSON field
	placeholders := []string{"?"}      // Placeholder for the JSON value
	values := []interface{}{json_data} // Values to insert, starting with JSON data

	// Add additional data columns and placeholders
	for col, val := range additional_data {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		values = append(values, val)
	}

	// Construct the query string
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	// Execute the query
	result, err := db.Exec(query, values...)
	if err != nil {
		color.Red("Failed to insert data into MySQL: %s", err)
		return err
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		color.Red("Failed to retrieve last inserted ID: %s", err)
		return err
	}

	color.Green("Data successfully inserted into MySQL with ID: %d", id)
	return nil
}
