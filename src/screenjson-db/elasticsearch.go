package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/fatih/color"
)

func insert_into_elasticsearch(uri, database, table, field, json_data string, additional_data map[string]string) error {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{uri},
	})
	if err != nil {
		color.Red("Failed to create Elasticsearch client: %s", err)
		return err
	}

	var doc map[string]interface{}
	if err := json.Unmarshal([]byte(json_data), &doc); err != nil {
		color.Red("Failed to unmarshal JSON data: %s", err)
		return err
	}
	for k, v := range additional_data {
		doc[k] = v
	}

	docBytes, _ := json.Marshal(doc)
	res, err := es.Index(
		table,                           // Index name
		bytes.NewReader(docBytes),       // Document body
		es.Index.WithContext(context.TODO()),
		es.Index.WithRefresh("true"),
	)
	if err != nil || res.IsError() {
		color.Red("Failed to insert data into Elasticsearch: %s", err)
		return err
	}

	color.Green("Data successfully inserted into Elasticsearch")
	return nil
}