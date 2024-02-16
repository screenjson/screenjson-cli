package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/fatih/color"
)

func insert_into_dynamodb(uri, database, table, field, json_data string, additional_data map[string]string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		color.Red("Failed to load AWS configuration: %s", err)
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)

	item := map[string]types.AttributeValue{
		field: &types.AttributeValueMemberS{Value: json_data},
	}

	for k, v := range additional_data {
		item[k] = &types.AttributeValueMemberS{Value: v}
	}

	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      item,
	})
	if err != nil {
		color.Red("Failed to insert data into DynamoDB: %s", err)
		return err
	}

	color.Green("Data successfully inserted into DynamoDB")
	return nil
}
