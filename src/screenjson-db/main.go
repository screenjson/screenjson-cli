package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Define a generic interface for database insertion functions
type insert_function func(uri, database, table, field, json_data string, additional_data map[string]string) error

func main() {
	var (
		uri            string
		engine         string
		database       string
		table          string
		field          string
		json_file      string
		additional_str string
	)

	// Parse command-line arguments
	flag.StringVar(&uri, "uri", "", "Database connection URI")
	flag.StringVar(&engine, "engine", "", "Database engine (mysql, postgresql, mongodb, oracle)")
	flag.StringVar(&database, "database", "", "Database name")
	flag.StringVar(&table, "table", "", "Table or collection name")
	flag.StringVar(&field, "field", "", "Field or column name to insert the JSON into")
	flag.StringVar(&json_file, "file", "", "Path to the JSON file containing the data to insert")
	flag.StringVar(&additional_str, "additional", "", "Additional key-value pairs in the format 'key1=value1,key2=value2'")
	flag.Parse()

	// Validate the input
	if uri == "" || engine == "" || database == "" || table == "" || field == "" || json_file == "" {
		color.Red("All arguments are required: -uri, -engine, -database, -table, -field, -file")
		os.Exit(1)
	}

	// Read and parse the JSON file into a string
	json_data, err := read_json_file(json_file)
	if err != nil {
		color.Red("Failed to read or parse the JSON file: %s", err)
		os.Exit(1)
	}

	// Parse additional key-value pairs
	additional_data := parse_additional_key_value_pairs(additional_str)

	// Map of supported databases and their corresponding insert functions
	insert_functions := map[string]insert_function{
		"cassandra":     insert_into_cassandra,
		"couchbase":     insert_into_couchbase,
		"couchdb":       insert_into_couchdb,
		"dynamodb":      insert_into_dynamodb,
		"elasticsearch": insert_into_elasticsearch,
		"mongodb":       insert_into_mongodb,
		"mssql":         insert_into_mssql,
		"mysql":         insert_into_mysql,
		"oracle":        insert_into_oracle,
		"postgresql":    insert_into_postgresql,
		"redis":         insert_into_redis,
		"sqlite":        insert_into_sqlite,
	}

	// Get the appropriate function based on the engine argument
	insert, ok := insert_functions[engine]
	if !ok {
		color.Red("Unsupported engine: %s", engine)
		os.Exit(1)
	}

	// Call the insert function with the JSON data as a string and additional data
	if err := insert(uri, database, table, field, json_data, additional_data); err != nil {
		color.Red("Failed to insert data: %s", err)
		os.Exit(1)
	}

	color.Green("Data successfully inserted")
}

// read_json_file reads the JSON file and returns it as a string
func read_json_file(file_path string) (string, error) {
	bytes, err := ioutil.ReadFile(file_path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(bytes), nil
}

// parse_additional_key_value_pairs converts a string of key=value pairs into a map
func parse_additional_key_value_pairs(kvPairs string) map[string]string {
	data := make(map[string]string)
	pairs := strings.Split(kvPairs, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			data[kv[0]] = kv[1]
		}
	}
	return data
}
