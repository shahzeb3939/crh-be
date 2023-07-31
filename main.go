package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shahzeb3939/crh-be/pkg/handlers"
)

func main() {
	// region := os.Getenv("AWS_REGION")
	// awsSession, err := session.NewSession(&aws.Config{
	// 	Region: aws.String(region),
	// })

	// if err != nil {
	// 	return
	// }

	// ddb := dynamodb.New(awsSession)
	// ddb.Query(&dynamodb.QueryInput{})
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handlers.GetCount()
}
