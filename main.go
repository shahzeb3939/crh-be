package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/shahzeb3939/crh-be/handlers"
)

var ddb *dynamodb.Client
var dynamodbTable string

func init() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Println("Error loading AWS configuration:", err)
		return
	}

	ddb = dynamodb.NewFromConfig(cfg)

	dynamodbTable = os.Getenv("dynamodbTable")

	if dynamodbTable == "" {
		log.Fatal("dynamodbTable environment variable is not set")
	}

	fmt.Println(dynamodbTable)
}

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handlers.GetCount(ddb, dynamodbTable)
}
