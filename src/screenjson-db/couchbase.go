package main

import (
	"encoding/json"

	"github.com/couchbase/gocb/v2"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func insert_into_couchbase(uri, database, table, field, json_data string, additional_data map[string]string) error {
	cluster, err := gocb.Connect(uri, gocb.ClusterOptions{})
	if err != nil {
		color.Red("Failed to connect to Couchbase: %s", err)
		return err
	}
	bucket := cluster.Bucket(database)
	collection := bucket.Scope("_default").Collection(table)

	docData := make(map[string]interface{})
	err = json.Unmarshal([]byte(json_data), &docData)
	if err != nil {
		color.Red("Failed to unmarshal JSON data: %s", err)
		return err
	}

	for k, v := range additional_data {
		docData[k] = v
	}

	docId := uuid.New().String()

	_, err = collection.Upsert(docId, docData, nil)
	if err != nil {
		color.Red("Failed to insert data into Couchbase: %s", err)
		return err
	}

	color.Green("Data successfully inserted into Couchbase with docId: %s", docId)
	return nil
}
