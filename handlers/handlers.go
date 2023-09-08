package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/shahzeb3939/crh-be/utils"
)

func GetCount(ddb *dynamodb.Client, dynamodbTable string) (events.APIGatewayProxyResponse, error) {

	fmt.Println(dynamodbTable, "2hehe")

	key := map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{
			Value: "COUNT",
		},
	}

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(dynamodbTable),
		Key:       key,
	}

	result, err := ddb.GetItem(context.Background(), getItemInput)
	if err != nil {
		log.Println("Error getting item:", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if len(result.Item) == 0 {
		_, err = ddb.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String(dynamodbTable),
			Item: map[string]types.AttributeValue{
				"PK": &types.AttributeValueMemberS{
					Value: "COUNT",
				},
				"count": &types.AttributeValueMemberN{
					Value: "0",
				},
			},
		})
		if err != nil {
			log.Println("Error creating item:", err)
			return events.APIGatewayProxyResponse{}, err
		}
	}

	updateExpression := "SET #c = #c + :incr"
	expressionAttributeNames := map[string]string{"#c": "count"}
	expressionAttributeValues := map[string]types.AttributeValue{
		":incr": &types.AttributeValueMemberN{
			Value: "1", // Increment the count by 1, change it as needed
		},
	}

	_, err = ddb.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(dynamodbTable),
		Key:                       key,
		UpdateExpression:          &updateExpression,
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
	})
	if err != nil {
		log.Println("Error updating item:", err)
		return events.APIGatewayProxyResponse{}, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(dynamodbTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: "COUNT",
			},
		},
		ProjectionExpression: aws.String("#c"),
		ExpressionAttributeNames: map[string]string{
			"#c": "count",
		},
	}

	result, err = ddb.GetItem(context.Background(), input)
	if err != nil {
		log.Println("Error getting item:", err)
		return events.APIGatewayProxyResponse{}, err
	}

	countValue := 0
	if result.Item["count"] != nil {
		countStr := result.Item["count"].(*types.AttributeValueMemberN).Value
		countValue, err = strconv.Atoi(countStr)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	}

	fmt.Println(dynamodbTable, "3hehe")

	return utils.ResponseObject(http.StatusOK, fmt.Sprint(countValue))
}

func GetTables(ddb *dynamodb.Client) (events.APIGatewayProxyResponse, error) {

	input := &dynamodb.ListTablesInput{
		Limit: aws.Int32(10),
	}

	result, err := ddb.ListTables(context.Background(), input)
	if err != nil {
		log.Println("Error listing tables:", err)
		return events.APIGatewayProxyResponse{}, err
	}

	tables := make([]string, 0)
	tables = append(tables, result.TableNames...)

	return utils.ResponseObject(http.StatusOK, strings.Join(tables, " "))
}
