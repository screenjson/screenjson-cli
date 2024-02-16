package main

import (
	"database/sql"
	"fmt"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
)

func insert_into_postgresql(uri, database, table, field, json_data string, additional_data map[string]string) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		color.Red("Failed to connect to PostgreSQL: %s", err)
		return err
	}
	defer db.Close()

	columns := []string{field}         // JSON field
	placeholders := []string{"$1"}     // PostgreSQL uses numbered placeholders
	values := []interface{}{json_data} // Starting with JSON data

	// PostgreSQL specific: Constructing additional columns and placeholders
	i := 1
	for col, val := range additional_data {
		i++
		placeholder := fmt.Sprintf("$%d", i)
		columns = append(columns, col)
		placeholders = append(placeholders, placeholder)
		values = append(values, val)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, placeholders)
	result, err := db.Exec(query, values...)
	if err != nil {
		color.Red("Failed to insert data into PostgreSQL: %s", err)
		return err
	}

	id, err := result.LastInsertId() // PostgreSQL uses RETURNING clause instead
	if err != nil {
		color.Red("Failed to retrieve last inserted ID: %s", err)
		return err
	}

	color.Green("Data successfully inserted into PostgreSQL with ID: %d", id)
	return nil
}
