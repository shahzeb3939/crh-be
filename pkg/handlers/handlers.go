package handlers

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func GetCount(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

	input := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"PK": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("COUNT"), // Replace with the value of the partition key
					},
				},
			},
		},
		ProjectionExpression: aws.String("count"), // Replace "count" with the actual attribute name
	}

	result, err := dbClient.Query(input)
	if err != nil {
		return apiResponse(http.StatusNotFound, errors.New("status not found"))
	}

	return apiResponse(http.StatusOK, result)

}

func Unhandled() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, errors.New("method not allowed"))
}
