package dynamoexample

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jaysonesmith/dynamo_mock/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQueryItems(t *testing.T) {
	testCases := []struct {
		name       string
		input      dynamodb.QueryInput
		expected   map[string][]map[string]string
		mockOutput *dynamodb.QueryOutput
	}{
		{
			name: "Duration 30 with AdGroup filter",
			input: dynamodb.QueryInput{
				TableName:              aws.String("books"),
				IndexName:              aws.String("title"),
				KeyConditionExpression: aws.String("title = :title"),
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":title": &dynamodb.AttributeValue{S: aws.String("foo")},
				},
			},
			expected: map[string][]map[string]string{"Items": []map[string]string{map[string]string{"title": "foo", "genre": "bar"}}},
			mockOutput: &dynamodb.QueryOutput{Items: []map[string]*dynamodb.AttributeValue{
				map[string]*dynamodb.AttributeValue{
					"title": {S: aws.String("foo")},
					"genre": {S: aws.String("bar")},
				}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDbiface := &mocks.DynamoDBAPI{}
			db := &Dynamoer{DynamoDB: mockDbiface}
			mockDbiface.On("Query", mock.Anything).Return(tc.mockOutput, nil)

			actual, err := db.queryItems(tc.input)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
