package main

import (
	"database/sql"
	"fmt"

	"github.com/fatih/color"
	_ "github.com/godror/godror"
)

func insert_into_oracle(uri, database, table, field, json_data string, additional_data map[string]string) error {
	db, err := sql.Open("godror", uri)
	if err != nil {
		color.Red("Failed to connect to Oracle: %s", err)
		return err
	}
	defer db.Close()

	columns := []string{field}            // Start with the JSON field
	placeholders := []string{":" + field} // Oracle uses named placeholders
	values := []interface{}{json_data}    // Values to insert, starting with JSON data

	// Oracle specific: Constructing additional columns and placeholders
	i := 1
	for col, val := range additional_data {
		i++
		placeholder := fmt.Sprintf(":%s", col)
		columns = append(columns, col)
		placeholders = append(placeholders, placeholder)
		values = append(values, val)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, placeholders)
	result, err := db.Exec(query, values...)
	if err != nil {
		color.Red("Failed to insert data into Oracle: %s", err)
		return err
	}

	id, err := result.LastInsertId() // Oracle might require a different approach for getting the inserted ID
	if err != nil {
		color.Red("Failed to retrieve last inserted ID: %s", err)
		return err
	}

	color.Green("Data successfully inserted into Oracle with ID: %d", id)
	return nil
}
