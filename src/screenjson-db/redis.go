package main

import (
	"context"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
)

func insert_into_redis(uri, database, table, field, json_data string, additional_data map[string]string) error {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		color.Red("Failed to parse Redis URI: %s", err)
		return err
	}

	rdb := redis.NewClient(opt)
	ctx := context.Background()

	// Using the field as key and json_data as value for Redis
	if err := rdb.Set(ctx, field, json_data, 0).Err(); err != nil {
		color.Red("Failed to insert data into Redis: %s", err)
		return err
	}

	color.Green("Data successfully inserted into Redis with key: %s", field)
	return nil
}
