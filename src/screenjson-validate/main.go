package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/xeipuuv/gojsonschema"
)

//go:embed schema.json
var schema_file embed.FS

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: screenjson-validate path/to/screenplay.json")
		os.Exit(1)
	}

	json_file_path := os.Args[1]

	json_bytes, err := os.ReadFile(json_file_path)

	if err != nil {
		color.Red("[exception]: Couldn't open/read JSON file: %v", err)
		os.Exit(1)
	}

	var json_data interface{}

	if json.Unmarshal(json_bytes, &json_data) != nil {
		color.Red("[exception]: That file is not valid JSON.")
		os.Exit(1)
	}

	schema_bytes, err := schema_file.ReadFile("schema.json")

	if err != nil {
		color.Red("[exception]: Error reading embedded schema document: %v", err)
		os.Exit(1)
	}

	schema_loader := gojsonschema.NewStringLoader(string(schema_bytes))
	document_loader := gojsonschema.NewBytesLoader(json_bytes)
	result, err := gojsonschema.Validate(schema_loader, document_loader)

	// Check if result is not nil before accessing its methods to avoid panic
	if err != nil || result == nil {
		color.Red("[internal schema exception]: %s", err)
		os.Exit(1)
	}

	if !result.Valid() {

		color.Red("That ScreenJSON document is not formatted properly.")
		color.Red("---------------------------------------------------")

		for _, desc := range result.Errors() {
			color.Red("%s", desc)
		}

	} else {

		color.Green("ScreenJSON document is valid and formatted correctly.")

	}
}
