package dynamoexample

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Dynamoer create a struct
type Dynamoer struct {
	DynamoDB dynamodbiface.DynamoDBAPI
}

func (db *Dynamoer) queryItems(queryInput dynamodb.QueryInput) (map[string][]map[string]string, error) {
	result, err := db.DynamoDB.Query(&queryInput)
	if err != nil {
		fmt.Println("dynamo err:", err)
		return map[string][]map[string]string{}, err
	}

	output, err := processQueryOutput(result)
	if err != nil {
		return map[string][]map[string]string{}, err
	}

	return output, nil
}

func processQueryOutput(queryOutput *dynamodb.QueryOutput) (map[string][]map[string]string, error) {
	log.Println("in processQueryOutput")
	var unmarshaledItems = []map[string]string{}

	err := dynamodbattribute.UnmarshalListOfMaps(queryOutput.Items, &unmarshaledItems)
	if err != nil {
		log.Printf("failed to unmarshal query output items: %v", queryOutput.Items)
		return map[string][]map[string]string{}, fmt.Errorf("failed to unmarshal query output items, %v", err)
	}

	itemsOut := map[string][]map[string]string{"Items": []map[string]string{}}
	for _, item := range unmarshaledItems {
		itemsOut["Items"] = append(itemsOut["Items"], item)
	}

	return itemsOut, nil
}
