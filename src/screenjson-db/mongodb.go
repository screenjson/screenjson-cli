package main

import (
	"context"
	"encoding/json"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insert_into_mongodb(uri, database, table, field, json_data string, additional_data map[string]string) error {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		color.Red("Failed to connect to MongoDB: %s", err)
		return err
	}
	defer client.Disconnect(context.TODO())

	// Select the database and collection
	coll := client.Database(database).Collection(table)

	// Unmarshal the JSON data into a map
	var doc map[string]interface{}
	if err := json.Unmarshal([]byte(json_data), &doc); err != nil {
		color.Red("Failed to unmarshal JSON data: %s", err)
		return err
	}

	// Add additional data to the document
	for k, v := range additional_data {
		doc[k] = v
	}

	// Insert the document
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		color.Red("Failed to insert data into MongoDB: %s", err)
		return err
	}

	color.Green("Data successfully inserted into MongoDB with ID: %v", result.InsertedID)
	return nil
}
