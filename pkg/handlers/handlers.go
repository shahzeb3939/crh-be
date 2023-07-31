package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/shahzeb3939/crh-be/utils"
)

func GetCount() (events.APIGatewayProxyResponse, error) {

	// return utils.ResponseObject(http.StatusOK, "GetCount is called")

	// cfg, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	fmt.Println("Error loading AWS config:", err)
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
		return events.APIGatewayProxyResponse{}, err
	}

	dbClient := dynamodb.NewFromConfig(cfg)

	result, err := dbClient.ListTables(context.Background(), &dynamodb.ListTablesInput{Limit: aws.Int32(10)})
	if err != nil {
		fmt.Println("Error listing tables:", err)
		return events.APIGatewayProxyResponse{}, err
	}

	tables := make([]string, 0)
	fmt.Println("DynamoDB Tables:")
	for _, tableName := range result.TableNames {
		tables = append(tables, tableName)
	}

	return utils.ResponseObject(http.StatusOK, strings.Join(tables, " "))

	// input := &dynamodb.QueryInput{
	// 	TableName: aws.String(tableName),
	// 	KeyConditions: map[string]*dynamodb.Condition{
	// 		"PK": {
	// 			ComparisonOperator: aws.String("EQ"),
	// 			AttributeValueList: []*dynamodb.AttributeValue{
	// 				{
	// 					S: aws.String("COUNT"), // Replace with the value of the partition key
	// 				},
	// 			},
	// 		},
	// 	},
	// 	ProjectionExpression: aws.String("count"), // Replace "count" with the actual attribute name
	// }

	// result, err := dbClient.Query(input)
	// if err != nil {
	// 	return apiResponse(http.StatusNotFound, errors.New("status not found"))
	// }

	// return apiResponse(http.StatusOK, result)

}
