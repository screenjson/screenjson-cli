package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/fatih/color"
)

func insert_into_mssql(uri, database, table, field, json_data string, additional_data map[string]string) error {
	db, err := sql.Open("sqlserver", uri)
	if err != nil {
		color.Red("Failed to connect to MSSQL: %s", err)
		return err
	}
	defer db.Close()

	columns := []string{field}            // JSON field
	placeholders := []string{"@" + field} // SQL Server uses named placeholders
	values := []interface{}{json_data}    // Starting with JSON data

	// SQL Server specific: Constructing additional columns and placeholders
	for col, val := range additional_data {
		placeholder := fmt.Sprintf("@%s", col)
		columns = append(columns, col)
		placeholders = append(placeholders, placeholder)
		values = append(values, sql.Named(col, val))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, placeholders)
	result, err := db.Exec(query, values...)
	if err != nil {
		color.Red("Failed to insert data into MSSQL: %s", err)
		return err
	}

	id, err := result.LastInsertId() // MSSQL might require OUTPUT clause for getting the inserted ID
	if err != nil {
		color.Red("Failed to retrieve last inserted ID: %s", err)
		return err
	}

	color.Green("Data successfully inserted into MSSQL with ID: %d", id)
	return nil
}
