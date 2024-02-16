package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func insert_into_couchdb(uri, database, table, field, json_data string, additional_data map[string]string) error {
	client := &http.Client{}
	doc := make(map[string]interface{})
	err := json.Unmarshal([]byte(json_data), &doc)
	if err != nil {
		color.Red("Failed to unmarshal JSON data: %s", err)
		return err
	}

	for k, v := range additional_data {
		doc[k] = v
	}

	docBytes, err := json.Marshal(doc)
	if err != nil {
		color.Red("Failed to marshal document: %s", err)
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", uri, database), bytes.NewBuffer(docBytes))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		color.Red("Failed to create request: %s", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		color.Red("Failed to execute request: %s", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		color.Red("Failed to insert data into CouchDB: Status %d", resp.StatusCode)
		return fmt.Errorf("failed to insert data into CouchDB: status %d", resp.StatusCode)
	}

	color.Green("Data successfully inserted into CouchDB")
	return nil
}
