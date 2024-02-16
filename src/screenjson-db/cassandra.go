package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gocql/gocql"
)

func insert_into_cassandra(uri, database, table, field, json_data string, additional_data map[string]string) error {
	cluster := gocql.NewCluster(uri)
	cluster.Keyspace = database
	session, err := cluster.CreateSession()
	if err != nil {
		color.Red("Failed to connect to Cassandra: %s", err)
		return err
	}
	defer session.Close()

	columns := field
	values := []interface{}{json_data}
	placeholders := "?"

	for k, v := range additional_data {
		columns += ", " + k
		values = append(values, v)
		placeholders += ", ?"
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, placeholders)
	if err := session.Query(query, values...).Exec(); err != nil {
		color.Red("Failed to insert data into Cassandra: %s", err)
		return err
	}

	color.Green("Data successfully inserted into Cassandra")
	return nil
}
