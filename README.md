# dynamo-mock

Super simple example to show mocking a dynamo queryItems call while unmarshaling the output to lend itself well to marshalling to json.

Prereqs:
- [aws go sdk](https://github.com/aws/aws-sdk-go)
- [mockery](https://github.com/vektra/mockery)
- [testify](https://github.com/stretchr/testify)

In order to generate the dynamo mock:

- Run:
`mockery -path=$GOPATH/src/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface -name=DynamoDBAPI`


This example makes use of testify's `mock.Anything` in the mocked response, which means you'll need to specify the responses/placeholders for each response. [Read more on testify's mock package here](https://github.com/stretchr/testify#mock-package)

If you have any questions, feel free to let me know!